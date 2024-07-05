package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PGUser="POSTGRES_USER"
	PGPassword="POSTGRES_PASSWORD"
	PGDb="POSTGRES_DB"
	PGPort="POSTGRES_PORT"
)

func establishDB() (*gorm.DB) {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect to db!")
	// }
	
	POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT := os.Getenv(PGUser), os.Getenv(PGPassword), os.Getenv(PGDb), os.Getenv(PGPort) 

	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%v sslmode=disable Timezone=Asia/Kuala_Lumpur", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db!")
	}
	
	fmt.Println("DB connected!")

	return db
}