package models

// Model holding information about a subscriber to the
// newsletter
type SubscriberModel struct {
	Name             string
	Email            string
	Verified         bool
	VerificationCode string
}
