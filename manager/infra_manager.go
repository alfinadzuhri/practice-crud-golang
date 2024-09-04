package manager

import (
	"fmt"

	"github.com/alfinadzuhri/practice-crud-golang/config"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}

	if err := conn.openConn(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (m *infraManager) openConn() error {

	dsn := fmt.Sprintf("host=%s, port=%s, user=%s, password%s, dbname=%s sslmode=disable",
		m.cfg.DbConfig.Host,
		m.cfg.DbConfig.Port,
		m.cfg.DbConfig.User,
		m.cfg.DbConfig.Password,
		m.cfg.DbConfig.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	m.db = db

	return nil
}

func (m *infraManager) Conn() *gorm.DB {
	return m.db
}
