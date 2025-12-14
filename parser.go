package migrator

import (
	"reflect"
	"strings"
)

// parseTag converts a struct field and tag into SQL column definition
func parseTag(fieldName string, fieldType reflect.Type, tag string) string {
	parts := strings.Split(tag, ",")
	colName := parts[0]
	colType := goTypeToSQLType(fieldType, parts)
	constraints := []string{}

	for _, p := range parts[1:] {
		if p == "primary" {
			constraints = append(constraints, "PRIMARY KEY")
		} else if p == "auto_increment" {
			constraints = append(constraints, "AUTOINCREMENT")
		} else if strings.HasPrefix(p, "default:") {
			val := strings.TrimPrefix(p, "default:")
			constraints = append(constraints, "DEFAULT "+val)
		} else if p == "unique" {
			constraints = append(constraints, "UNIQUE")
		} else if p == "notnull" {
			constraints = append(constraints, "NOT NULL")
		}
	}

	return strings.TrimSpace(colName + " " + colType + " " + strings.Join(constraints, " "))
}

// goTypeToSQLType maps Go types to SQL types
func goTypeToSQLType(t reflect.Type, tagParts []string) string {
	switch t.Kind() {
	case reflect.Int, reflect.Int64:
		return "INTEGER"
	case reflect.String:
		size := "255"
		for _, p := range tagParts {
			if strings.HasPrefix(p, "size:") {
				size = strings.TrimPrefix(p, "size:")
			}
		}
		return "VARCHAR(" + size + ")"
	case reflect.Float64:
		return "FLOAT"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Struct:
		if t.Name() == "Time" {
			return "TIMESTAMP"
		}
	}
	return "TEXT"
}
