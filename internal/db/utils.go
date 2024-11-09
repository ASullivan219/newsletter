package db

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/asullivan219/newsletter/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

var ERROR_SUBSCRIBER_NOT_FOUND = errors.New("Subscriber not found")

// Get The subscriber from the database and return
func getSubscriber(db *sql.DB, email string) (models.SubscriberModel, error) {
	subscriber := models.SubscriberModel{}

	row := db.QueryRow(
		"SELECT * FROM subscribers WHERE email = ?;",
		email,
	)

	err := row.Scan(
		&subscriber.Email,
		&subscriber.Name,
		&subscriber.Verified,
		&subscriber.VerificationCode,
		&subscriber.Relationship)

	if err != nil {
		slog.Error(
			"Retrieving user from database",
			"exists", false,
			"email", email,
			"error", err.Error())
		return models.SubscriberModel{}, ERROR_SUBSCRIBER_NOT_FOUND
	}

	slog.Info(
		"Retrieving user from database",
		"exists", true,
		"email", email,
	)

	return subscriber, nil
}

// Return true if the user exists in the database already
func subscriberExists(db *sql.DB, email string) bool {
	_, err := getSubscriber(db, email)
	if err != nil {
		return false
	}
	return true
}

// Generate and return a new 8 character string to use as a
// Verification code

// Put a new subscriber return an Error if the insert fails
func createSubscriber(db *sql.DB, subscriber models.SubscriberModel) error {
	_, err := db.Exec(
		`INSERT INTO subscribers VALUES(
			?, ?, ?, ?, ?);`,
		subscriber.Email,
		subscriber.Name,
		subscriber.Verified,
		subscriber.VerificationCode,
		subscriber.Relationship)

	if err != nil {
		slog.Error(
			"error creating subscriber",
			"email", subscriber.Email,
			"error", err,
		)
		return err
	}

	slog.Info("New User input in database",
		"email", subscriber.Email,
		"name", subscriber.Name,
	)

	return nil
}

func upsertSubscriber(
	db *sql.DB,
	email string,
	name string,
	verified bool,
	relationship int) error {

	_, err := db.Exec(
		`INSERT INTO subscribers(email, name, verified, relationship)
			VALUES(?, ?, ?, ?)
			ON CONFLICT(email) DO UPDATE SET
			email = excluded.email,
			name = excluded.name,
			verified = excluded.verified,
			relationship = excluded.relationship;
		`, email, name, verified, relationship)

	if err != nil {
		slog.Error("errr upserting subscriber",
			"email", email,
			"error", err,
		)
		return err
	}

	return nil
}
