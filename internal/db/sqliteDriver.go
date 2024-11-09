package db

import (
	"database/sql"
	"log/slog"

	"github.com/asullivan219/newsletter/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDriver struct {
	I_database
	Db *sql.DB
}

// Open and initialize the database with the subscriber table
func NewDb(filePath string) *SqliteDriver {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil
	}
	sqlite := &SqliteDriver{
		Db: db,
	}

	if err = sqlite.InitializeTables(); err != nil {
		slog.Error("Error initializing database tables",
			"error", err.Error(),
		)
		return nil
	}

	return sqlite
}

// Create the subscriber table if it does not already exist
func (db *SqliteDriver) InitializeTables() error {
	const schema string = `
		CREATE TABLE IF NOT EXISTS subscribers (
		email VARCHAR(255) NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		verified BOOLEAN,
		verfification_code CHARACTER(10) 
	)`
	if _, err := db.Db.Exec(schema); err != nil {
		return err
	}
	return nil
}

// We wont really need to up
func (db *SqliteDriver) CreateSubscriber(email string, name string) error {
	newSubscriber := models.SubscriberModel{
		Email:    email,
		Name:     name,
		Verified: false,
	}

	return createSubscriber(db.Db, newSubscriber)
}

func (db *SqliteDriver) GetSubscriber(email string) (models.SubscriberModel, error) {
	subscriber, err := getSubscriber(db.Db, email)
	if err != nil {
		return models.SubscriberModel{}, nil
	}

	return subscriber, nil

}

func (db *SqliteDriver) VerifySubscriber(model models.SubscriberModel) (models.SubscriberModel, error) {

	if model.Verified {
		slog.Error("Subscriber already Verified",
			"email", model.Email,
			"verification", model.Verified,
		)
		return model, nil
	}

	model.Verified = true
	err := db.PutSubscriber(model)
	if err != nil {
		slog.Error("Error verifying user",
			"email", model.Email,
			"error", err,
		)
		return model, err
	}

	return model, nil
}

func (db *SqliteDriver) PutSubscriber(model models.SubscriberModel) error {
	return upsertSubscriber(db.Db, model.Email, model.Name, model.Verified)
}

func (db *SqliteDriver) DropSubscribers() error {
	_, err := db.Db.Exec("DROP TABLE IF EXISTS subscribers;")

	if err != nil {
		slog.Error("Could not drop table",
			"error", err,
		)
		return err
	}

	return nil
}
