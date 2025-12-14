package migrator

import (
	"fmt"
	"os"
	"reflect"
)

// CLI function to run migration generator from command line
func RunCLI(structInstance interface{}, tableName string) {
	sql := GenerateCreateTableSQL(structInstance, tableName)
	filename := GenerateMigrationFileName("create_" + tableName + "_table")

	err := WriteMigrationFile(filename, sql)
	if err != nil {
		fmt.Println("Error writing migration:", err)
		os.Exit(1)
	}
	fmt.Println("Migration file created:", filename)
}

// Helper to check if input is a struct
func IsStruct(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.Struct
}
