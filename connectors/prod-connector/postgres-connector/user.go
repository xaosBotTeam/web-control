package postgres_connector

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func shaHashing(input string) string {
	plainText := []byte("edN0Cq8a3S54456S" + input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

type User struct {
	ID       int
	Username string
}

type AbstractUserStorage interface {
	AuthUser(username, password string) (bool, int)
	GetById(ID int) string
	ChangePassword(ID int, password string)
}

func AuthUser(username, password string) (int, bool) {
	query := fmt.Sprintf("select id from web_control.user_password where username='%s' and password_hash = '%s'", username, shaHashing(password))
	rows := db.QueryRow(query)

	var id int

	err := rows.Scan(&id)
	if err != nil {
		return 0, false
	}

	return id, true
}

func GetById(ID int) (string, bool) {
	query := fmt.Sprintf("select username from web_control.user_password where id='%d'", ID)
	rows := db.QueryRow(query)

	var name string

	err := rows.Scan(&name)
	if err != nil {
		return "", false
	}

	return name, false
}

func ChangePassword(ID int, password string) bool {
	query := fmt.Sprintf("update web_control.user_password set password_hash='%s' where id='%d'", shaHashing(password), ID)
	_, err := db.Exec(query)

	if err != nil {
		return false
	}

	return true
}
