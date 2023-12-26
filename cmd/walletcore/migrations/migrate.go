// main.go
package main

import (
    "database/sql"
    "fmt"
    "log"
	"os"
	"path/filepath"

    _ "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
    driverName     = "mysql"
    dataSourceName = "root:root@tcp(mysql:3306)/wallet?charset=utf8&parseTime=True&loc=Local"
)

func main() {
    db, err := sql.Open(driverName, dataSourceName)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    driver, err := mysql.WithInstance(db, &mysql.Config{})
    if err != nil {
        log.Fatal(err)
    }

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migrationPath := filepath.Join(dir, "internal/database/migrations")

    m, err := migrate.NewWithDatabaseInstance("file://" + migrationPath, driverName, driver)
    if err != nil {
        log.Fatal(err)
    }

    // Aplica todas as migrações
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }

    fmt.Println("Migrations applied successfully!")
}