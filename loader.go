package migrator

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// GenerateMigrationsForModels scans models folder and generates migrations
func GenerateMigrationsForModels() {
	LoadConfig() // load .env config

	err := filepath.Walk(ModelsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			// Importing structs dynamically is tricky in Go.
			// We'll assume users register their models manually.
			fmt.Println("Found model file:", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error scanning models folder:", err)
	}
}

// Helper to register a model and generate migration
func RegisterModel(model interface{}) {
	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Struct {
		fmt.Println("Not a struct, skipping")
		return
	}

	tableName := StructToTableName(t.Name())
	if tableTag, ok := t.Field(0).Tag.Lookup("db_table"); ok {
		tableName = tableTag
	}

	sql := GenerateCreateTableSQL(model, tableName)
	filename := GenerateMigrationFileName("create_" + tableName + "_table")

	err := WriteMigrationFile(filename, sql)
	if err != nil {
		fmt.Println("Error writing migration for", tableName, ":", err)
		return
	}
	fmt.Println("Migration created for table:", tableName, "->", filename)
}
