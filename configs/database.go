package configs

import (
	"fmt"

	"github.com/winartodev/leaderboard/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (c *Config) LoadDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return db, err
	}

	err = db.AutoMigrate(
		entity.Leaderboard{},
		entity.PointLog{},
		entity.User{},
	)
	if err != nil {
		return db, err
	}

	return db, err
}
