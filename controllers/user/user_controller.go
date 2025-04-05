package user

import (
	"database/sql"
	"errors"
	"time"
	"api.finance.com/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRegister struct {
	Name 		string `json:"name" binding:"required"`
	Email 		string `json:"email" binding:"email,required"`
	Password	string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email		string `json:"email" binding:"email,required"`
	Password 	string `json:"password" binding:"required"`
}

type UserDatabaseFormat struct {
	UserRegister
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	IsActive 	bool `json:"is_active"`
	Id 			string `json:"id"`
}

type JWTTokens struct {
	Auth_token 		string `json:"auth_token" binding:"required"`
	Refresh_token  	string  `json:"refresh_token" binding:"required"`
}

func CreateUser(db *sql.DB, userData *UserRegister) error {
	uuid := uuid.New();
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost)
	if err != nil {
		return errors.New("internal server error")
	}

	_, err = db.Exec(`INSERT INTO "User" (id, name, email, password) VALUES ($1, $2, $3, $4)`, uuid, userData.Name, userData.Email, hashedPass)
	if err != nil {
		return errors.New("err")
	}

	return nil
}

func LoginUser(db *sql.DB, userData *UserLogin) (*JWTTokens, error) {
	// Query to get user
	rows, err := db.Query(`SELECT password, name, id FROM "User" WHERE email = $1`, userData.Email)
	if err != nil {
		return nil, errors.New("query error")
	}

	// Close the connection to the db
	defer rows.Close()

	// Verify if the user exists
	if !rows.Next() {
		return nil, errors.New("user not found")
	}

	// Process the data
	var hashedPassword, name, id string
	err = rows.Scan(&hashedPassword, &name, &id)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userData.Password)); err != nil {
		return nil, errors.New("invalid data")
	}

	//Create the token and add the fields to it
	jwtConfig := config.GetJWTConfig()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = name
	claims["email"] = userData.Email
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iss"] = jwtConfig.Iss
	
	authToken, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return nil, errors.New("internal server error")
	}

	token = jwt.New(jwt.SigningMethodHS256)
	claims = token.Claims.(jwt.MapClaims)

	claims["id"] = userData.Email
	claims["exp"] = time.Now().Add(time.Hour * 672).Unix()
	claims["iss"] = jwtConfig.Iss
	refreshToken, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return nil, errors.New("internal server error")
	}

	tokens := &JWTTokens{
		Auth_token: authToken,
		Refresh_token: refreshToken,
	}

	return tokens, nil
}