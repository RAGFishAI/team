package migration

import (
	"github.com/RAGFishAI/team/migration/mysql"
	"github.com/RAGFishAI/team/migration/postgres"
	"gorm.io/gorm"
)

func AutoMigration(DB *gorm.DB, dbtype any) {

	if dbtype == "postgres" {

		postgres.MigrationTables(DB) //auto migrate table

	} else if dbtype == "mysql" {

		mysql.MigrationTables(DB) //auto migrate table
	}

}
