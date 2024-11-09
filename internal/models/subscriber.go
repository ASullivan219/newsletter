package models

import (
	"math/rand"
)

const (
	UNSET_VERIFICATION_CODE = "00000000"
)

// Runes to generate verification code from
var runes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// Randomly generate a new verification code from the runes listed above
func generateNewVerificationCode() string {
	code := make([]rune, 8)
	for i := range code {
		code[i] = runes[rand.Intn(len(runes))]
	}
	return string(code)
}

// Model holding information about a subscriber to the
// newsletter
type SubscriberModel struct {
	Name             string
	Email            string
	Verified         bool
	VerificationCode string
	Relationship     int
}

// Function to create a Brand new subsrriber to the newsletter
func BrandNewSubscriber(
	email string,
	name string,
	relationship int) SubscriberModel {

	return SubscriberModel{
		Name:             name,
		Email:            email,
		Verified:         false,
		VerificationCode: generateNewVerificationCode(),
		Relationship:     relationship}

}
