package database

import (
	"embed"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
	"youras/infra/config"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS
var (
	dialect = "pgx"
	path    = "migrations"
)

func OpenPg(cfg *config.DbConfig, logger gormlogger.Interface) (*gorm.DB, error) {
	pconfig := postgres.Config{
		DSN:                  cfg.Dsn(),
		PreferSimpleProtocol: true,
	}
	dialector := postgres.New(pconfig)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	sqlDB.SetMaxOpenConns(cfg.MaxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLife) * time.Second)
	return db, nil
}

func Migrate(db *gorm.DB, table string) error {
	goose.SetTableName(table)
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	goose.SetBaseFS(embedMigrations)
	if err = goose.SetDialect(dialect); err != nil {
		return err
	}
	if err = goose.Up(sqlDb, path, goose.WithAllowMissing()); err != nil {
		return err
	}
	return nil
}
