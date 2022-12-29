package main

import (
	 _ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	 migration "github.com/golang-migrate/migrate/v4/database/mysql"
	 _ "github.com/golang-migrate/migrate/v4/source/file"
	 "github.com/enajera/phone-review/internal/database"
	"github.com/enajera/phone-review/internal/logs"
	s "github.com/enajera/phone-review/internal/server"
	r "github.com/enajera/phone-review/internal/routes"
	
	)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = logs.InitLogger()
	client := database.NewSqlClient("root:admin@tcp(localhost:3306)/phones_review")
	doMigrate(client, "phones_review")

	mux := r.Routes()
	server := s.NewServer(mux)
	server.Run()
}

func doMigrate(client *database.MySqlClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)
	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}
