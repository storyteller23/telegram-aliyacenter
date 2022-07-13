package controller

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/storyteller23/telegram-aliyacenter/pkg/config"
	"github.com/storyteller23/telegram-aliyacenter/pkg/models"
)

func ConnectDB(dbName string) *sql.DB {
	cfg := config.NewConfig()
	db, err := sql.Open("sqlite3", cfg.DBName)
	if err != nil {
		log.Panic(err)
	}
	return db
}

func CreateTable(dbName string) error {
	db := ConnectDB(dbName)
	defer db.Close()
	sql_table := `
		CREATE TABLE IF NOT EXISTS users(
		chat_id INTEGER NOT NULL UNIQUE PRIMARY KEY,
		name TEXT,
		phone TEXT,
		email TEXT,
		specialist TEXT,
		date TEXT,
		time TEXT,
		flag TEXT
	)`
	_, err := db.Exec(sql_table)
	return err
}

func AddChatID(dbName string, chatId int64) error {
	db := ConnectDB(dbName)
	defer db.Close()
	var id int64
	err := db.QueryRow(`SELECT chat_id FROM users WHERE chat_id = ?`, chatId).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := db.Exec(`
		INSERT INTO users (chat_id, flag)
		VALUES(?, 'info')
		`, chatId)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func GetByID(dbName string, chatId int64, name string) (string, error) {
	db := ConnectDB(dbName)
	defer db.Close()
	var data string
	err := db.QueryRow(`SELECT `+name+` FROM users WHERE chat_id = ?`, chatId).Scan(&data)
	if err != nil {
		return "", err
	}
	return data, nil
}

func GetRequestData(dbName string, chatId int64) (models.User, error) {
	db := ConnectDB(dbName)
	defer db.Close()
	user := models.User{}
	err := db.QueryRow(`SELECT name, phone, email, specialist, date, time FROM users WHERE chat_id = ?`, chatId).Scan(&user.Name, &user.Phone, &user.Email, &user.Specialist, &user.Date, &user.Time)
	if err != nil {
		return user, nil
	}
	return user, nil
}

func SetFlagByID(dbName string, chatId int64, newFlag string) error {
	db := ConnectDB(dbName)
	defer db.Close()
	_, err := db.Exec(`
	UPDATE users
	SET flag = ?
	WHERE chat_id = ?;
	`, newFlag, chatId)
	return err
}

func SetNameByID(dbName string, chatId int64, data string) error {
	db := ConnectDB(dbName)
	defer db.Close()
	_, err := db.Exec(`
	UPDATE users
	SET name = ?
	WHERE chat_id = ?;
	`, data, chatId)
	return err
}

func SetPhoneByID(dbName string, chatId int64, data string) error {
	db := ConnectDB(dbName)
	defer db.Close()
	_, err := db.Exec(`
	UPDATE users
	SET phone = ?
	WHERE chat_id = ?;
	`, data, chatId)
	return err
}

func SetEmailByID(dbName string, chatId int64, data string) error {
	db := ConnectDB(dbName)
	defer db.Close()
	_, err := db.Exec(`
	UPDATE users
	SET email = ?
	WHERE chat_id = ?;
	`, data, chatId)
	return err
}

func SetSpecialistByID(dbName string, chatId int64, data string) error {
	db := ConnectDB(dbName)
	defer db.Close()
	_, err := db.Exec(`
	UPDATE users
	SET specialist = ?
	WHERE chat_id = ?;
	`, data, chatId)
	return err
}

func SetDateByID(dbName string, chatId int64, data string) error {
	db := ConnectDB(dbName)
	defer db.Close()
	_, err := db.Exec(`
	UPDATE users
	SET date = ?
	WHERE chat_id = ?;
	`, data, chatId)
	return err
}

func SetTimeByID(dbName string, chatId int64, data string) error {
	db := ConnectDB(dbName)
	defer db.Close()
	_, err := db.Exec(`
	UPDATE users
	SET time = ?
	WHERE chat_id = ?;
	`, data, chatId)
	return err
}
