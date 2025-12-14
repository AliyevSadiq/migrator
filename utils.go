package migrator

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func GenerateTimestamp() string {
	return time.Now().Format("20060102150405")
}

// GenerateMigrationFileName returns full path in migrations folder
func GenerateMigrationFileName(tableName string) string {
	return filepath.Join(MigrationsFolder, fmt.Sprintf("%s_create_%s_table.sql", GenerateTimestamp(), tableName))
}

// Convert "UserProfile" to "user_profile"
func StructToTableName(name string) string {
	var re = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(name, "${1}_${2}")
	return strings.ToLower(snake)
}

// Check if a migration file for table exists
func MigrationExists(tableName string) bool {
	files, err := os.ReadDir(MigrationsFolder)
	if err != nil {
		return false
	}

	for _, f := range files {
		if !f.IsDir() && strings.Contains(f.Name(), "_create_"+tableName+"_table") {
			return true
		}
	}
	return false
}
