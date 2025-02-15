package models

import (
	"C/Users/anura/OneDrive/Documents/GitHub/MyProject/db"
	"C/Users/anura/OneDrive/Documents/GitHub/MyProject/utils"
	"errors"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `"binding":"required"`
	Password string `"binding":"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	newpass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, newpass)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	fmt.Print(u.ID, u.Email)
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id,  password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPassword(retrievedPassword, u.Password)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
