package user

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name 		string `json:"name" binding:"required"`
	Email 		string `json:"email" binding:"email,required"`
	Password	string `json:"password" binding:"required"`
}

type UserDatabaseFormat struct {
	User
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	IsActive 	bool `json:"is_active"`
	Id 			string `json:"id"`
}

func CreateUser(db *sql.DB, userData *User) error {
	uuid := uuid.New();
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost)
	if err != nil {
		return errors.New("enternal server error")
	}

	_, err = db.Exec(`INSERT INTO "User" (id, name, email, password) VALUES ($1, $2, $3, $4)`, uuid, userData.Name, userData.Email, hashedPass)
	if err != nil {
		fmt.Println(err)
		return errors.New("err")
	}

	return nil
}
