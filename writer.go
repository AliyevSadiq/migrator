package migrator

import (
	"os"
)

// WriteMigrationFile writes the SQL string to a file
func WriteMigrationFile(filename, sql string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(sql)
	return err
}
