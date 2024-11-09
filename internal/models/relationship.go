package models

import "errors"

const (
	CEW_CREW = iota
	FRIEND
	FAMILY
	LORIE
	OTHER
)

func IntToRelationshipString(relationship int) (string, error) {
	switch relationship {
	case 0:
		return "CEW Crew", nil
	case 1:
		return "Friend", nil
	case 2:
		return "Family", nil
	case 3:
		return "Lorie", nil
	case 4:
		return "Other", nil
	default:
		return "", errors.New("undefined relationship")
	}
}
