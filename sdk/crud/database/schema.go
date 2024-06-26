package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type DBConfig struct {
	DSN    string
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort int
}

type Adapter interface {
	Gorm() *gorm.DB
}

type Migrator interface {
	Run() error
	AddModels([]interface{})
	AddMigrations(Migrations)
}

type Migrations interface {
	Models() []interface{}
	GormMigrations() []*gormigrate.Migration
}

type SeedAction struct {
	ID      string
	Execute func(*Database) error
}

type ModelSeeder interface {
	SeedModel(*Database) []SeedAction
}

type ORMSeeder interface {
	Run() error
	AddSeeder(seeders ...ModelSeeder) ORMSeeder
}
