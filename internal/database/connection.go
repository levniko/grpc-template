package database

import (
	"fmt"
	"grpc-template/internal/config"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IConnection interface {
	GetMaster() *gorm.DB
	GetSlave() *gorm.DB
}

func NewDBManager(cfg *config.Config) (*DBManager, error) {
	master, err := connect(cfg.PostgreSQL.Master)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to master: %w", err)
	}

	var slaves []*gorm.DB
	for _, slaveCfg := range cfg.PostgreSQL.Slaves {
		slave, err := connect(slaveCfg)
		if err != nil {
			log.Printf("Failed to connect to slave: %v", err)
			continue
		}
		slaves = append(slaves, slave)
	}

	return &DBManager{
		Master: master,
		Slaves: slaves,
	}, nil
}

func connect(cfg config.PostgreSQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}

// GetMaster always returns master conn
func (m *DBManager) GetMaster() *gorm.DB {
	return m.Master
}

// GetSlave, return an available slave conn in slave pool randomly
// If does not exist, return master
func (m *DBManager) GetSlave() *gorm.DB {
	if len(m.Slaves) == 0 {
		return m.GetMaster()
	}
	index := rand.Intn(len(m.Slaves))
	return m.Slaves[index]
}
