package repository

import (
	"database/sql"
	"log"

	"github.com/miltsm/hubung-service/internal/types"
	"github.com/miltsm/hubung-service/internal/utils"
)

type Repository interface {
	GetUser(email string) *types.UserEntity
	CreateNewUser(email string, username string) (*types.UserEntity, error)
}

type repository struct {
	db *sql.DB
}

func New() Repository {
	db := establishDB()
	return &repository{db: db}
}

func (r *repository) GetUser(email string) *types.UserEntity {
	var user *types.UserEntity
	//TODO: SQL injection prevention
	row := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email)
	row.Scan(&user)
	return user
}

func (r *repository) CreateNewUser(email string, username string) (*types.UserEntity, error) {
	uid := utils.GenerateUUID()
	stmt, err := r.db.Prepare("INSERT INTO users (userid,email,username) VALUES ($1,$2,$3)")
	if err != nil {
		log.Printf("Prepare user insert failure (Unique ID:%v)\n%v", email, err)
		return nil, &types.CreateNewUserError{}
	}
	_, err = stmt.Exec(uid, email, username)
	if err != nil {
		log.Printf("Unable to create user (Unique ID:%v)\n%v", email, err)
		return nil, &types.CreateNewUserError{}
	}
	user := types.UserEntity{
		UserID:   uid,
		Email:    email,
		Username: username,
	}
	return &user, nil
}
