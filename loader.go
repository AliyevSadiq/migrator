package migrator

import (
	"fmt"
	"reflect"
)

// RegisterModel registers a struct and generates migration if not exists
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

	// Skip if migration already exists
	if MigrationExists(tableName) {
		fmt.Println("Migration already exists for table:", tableName)
		return
	}

	sql := GenerateCreateTableSQL(model, tableName)
	filename := GenerateMigrationFileName(tableName)

	err := WriteMigrationFile(filename, sql)
	if err != nil {
		fmt.Println("Error writing migration for", tableName, ":", err)
		return
	}

	fmt.Println("Migration created for table:", tableName, "->", filename)
}
