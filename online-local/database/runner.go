package database

import (
	"database/sql"
	"github.com/sportorg/online/database/migrations"
)

var MIGRATIONS = []func(db *sql.DB){
	migrations.Init,
}


func RunMigrations(db *sql.DB) {
	// TODO: index = migration version
	for _, migration := range MIGRATIONS {
		migration(db)
	}
}
