package db

import "github.com/asullivan219/newsletter/internal/models"

// Interface describing what the db layer needs to be able
// to accomplish
type I_database interface {
	InitializeTables() error
	GetSubscriber(string) (models.SubscriberModel, error)
	PutSubscriber(models.SubscriberModel) error
	CreateSubscriber(string, string, int) error
	VerifySubscriber(models.SubscriberModel) (models.SubscriberModel, error)
	DropSubscribers() error
}
