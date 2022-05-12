package dao

import (
	"fmt"

	"api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("root:%s@tcp(%s:3306)/blog?charset=utf8&parseTime=True&loc=Local", config.MYSQL_PASSWORD, config.MYSQL_HOST),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	return db, err
}
