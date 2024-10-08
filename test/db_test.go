package test

import (
	//"log/slog"
	"database/sql"
	"log/slog"
	"testing"

	"github.com/asullivan219/newsletter/internal/db"
	"github.com/asullivan219/newsletter/internal/models"
)

// Clean up function to clear the database between runs
func dropSubscribersTable(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS subscribers;")

	if err != nil {
		slog.Error("Could not drop table",
			"error", err,
		)
	}
}

// Sanity test to ensure object is not null on creation
func TestCreateSqlDb(t *testing.T) {
	sqliteDb := db.NewDb("../resources/subscribers.db")
	if sqliteDb == nil {
		t.Fatalf("Error initializiing new sqlite db")
	}
}

// Make sure the subscriber table is initialized on startup
func TestSchemaExists(t *testing.T) {
	sqliteDb := db.NewDb("../resources/subscribers.db")
	row := sqliteDb.Db.QueryRow("SELECT name FROM sqlite_master WHERE type='table';")
	var name string
	row.Scan(&name)

	if name != "subscribers" {
		t.Fatal("Table with name subscribers not found in db")
	}
}

// Test inputting new subscriber and retrieving them from the DB
func TestInputNewSubscriber(t *testing.T) {
	sqliteDb := db.NewDb("../resources/subscribers.db")
	defer dropSubscribersTable(sqliteDb.Db)

	err := sqliteDb.CreateSubscriber(
		"a.sullivan219@gmail.com",
		"Alex Sullivan",
	)

	if err != nil {
		t.Fatal("Failed creating new user")
	}

	subscriber, err := sqliteDb.GetSubscriber("a.sullivan219@gmail.com")

	if err != nil {
		t.Fatal(err.Error())
	}

	if subscriber.Name != "Alex Sullivan" {
		t.Fatalf("got name: %s expected %s", subscriber.Name, "Alex Sullivan")
	}

	if subscriber.Verified {
		t.Fatalf("Expected Unverified User, got %t", subscriber.Verified)
	}

}

// Test that the user can be set to verified
func TestVerifyUser(t *testing.T) {
	sqliteDb := db.NewDb("../resources/subscribers.db")
	defer dropSubscribersTable(sqliteDb.Db)

	newSubcriber := models.SubscriberModel{
		Email:            "a.sullivan219@gmail.com",
		Name:             "Alex Sullivan",
		Verified:         false,
		VerificationCode: "XXXXXXXX",
	}

	sqliteDb.CreateSubscriber(
		newSubcriber.Email,
		newSubcriber.Name,
	)

	newSubcriber, err := sqliteDb.VerifySubscriber(newSubcriber)
	db_sub, _ := sqliteDb.GetSubscriber(newSubcriber.Email)

	if err != nil {
		t.Fatal("Error Verifying User", err.Error())
	}

	if !newSubcriber.Verified {
		t.Fatal("expected new subscriber to be verified")
	}

	if db_sub.Email != newSubcriber.Email {
		t.Fatalf("Email incorrected, expected %s got %s",
			newSubcriber.Email, db_sub.Email,
		)
	}

	if db_sub.Name != newSubcriber.Name {
		t.Fatalf("Email incorrected, expected %s got %s",
			newSubcriber.Name, db_sub.Name,
		)
	}
}

func TestGenerateCode(t *testing.T) {
	sqliteDb := db.NewDb("../resources/subscribers.db")
	defer dropSubscribersTable(sqliteDb.Db)

	email := "a.sullivan@gmail.com"
	name := "Alex"

	sqliteDb.CreateSubscriber(email, name)
	subscriber, err := sqliteDb.GetSubscriber(email)

	if err != nil {
		t.Fatalf("Error getting subscriber with email: %s", email)
	}

	if subscriber.VerificationCode == "" {
		t.Fatal("Verification code empty")
	}

	t.Log(subscriber.VerificationCode)

}
