package db

import (
	"Parser/utils"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// ConnectDB function connects to the database and inits migration.
//   - Returns *gorm.DB
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv(utils.DB)), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("[Connection Error]", err)
	}

	migrateErr := db.AutoMigrate(&Block{}, &Wallet{}, &Transaction{})
	if migrateErr != nil {
		log.Fatal("[Migration Error]", migrateErr)
	}

	return db
}
