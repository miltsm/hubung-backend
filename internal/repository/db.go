package repository

import (
	"fmt"
	"io/fs"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PGUser="POSTGRES_USER"
	PGPassword="POSTGRES_PASSWORD"
	PGDb="POSTGRES_DB"
	PGPort="POSTGRES_PORT"
	PGHost="POSTGRES_HOST"
	PGPasswordPath="POSTGRES_PASSWORD_FILE"
)

func establishDB() (*gorm.DB) {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to connect to db!")
	// }
	
	POSTGRES_USER, POSTGRES_DB, POSTGRES_PORT, POSTGRES_HOST, POSTGRES_PASSWORD_FILE := os.Getenv(PGUser), os.Getenv(PGDb), os.Getenv(PGPort), os.Getenv(PGHost), os.Getenv(PGPasswordPath)

	//get password
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rootFs := os.DirFS(root)
	pwByte, err := fs.ReadFile(rootFs, POSTGRES_PASSWORD_FILE)
	if err != nil {
		panic(err)
	}

	pswd := string(pwByte)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable Timezone=Asia/Kuala_Lumpur", POSTGRES_HOST, POSTGRES_USER, pswd, POSTGRES_DB, POSTGRES_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db!")
	}
	
	fmt.Println("DB connected!")

	return db
}