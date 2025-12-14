package migrator

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func GenerateTimestamp() string {
	return time.Now().Format("20060102150405")
}

func GenerateMigrationFileName(name string) string {
	return fmt.Sprintf("%s_%s.sql", GenerateTimestamp(), name)
}

// Convert "UserProfile" to "user_profile"
func StructToTableName(name string) string {
	var re = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(name, "${1}_${2}")
	return strings.ToLower(snake)
}
