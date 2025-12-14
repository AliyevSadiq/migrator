package migrator

import (
	"fmt"
	"reflect"
	"strings"
)

// GenerateCreateTableSQL generates a SQL CREATE TABLE statement from a struct
func GenerateCreateTableSQL(s interface{}, tableName string) string {
	t := reflect.TypeOf(s)
	sql := fmt.Sprintf("CREATE TABLE %s (\n", tableName)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			continue
		}
		columnDef := parseTag(field.Name, field.Type, tag)
		sql += "    " + columnDef + ",\n"
	}
	sql = strings.TrimSuffix(sql, ",\n") + "\n);\n"
	return sql
}
