package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/mail"
	"net/url"
)

type subscriberModel struct {
	Name     string
	Email    string
	Verified bool
	Id       string
}

type I_database interface {
	GetSubscriber(string) (subscriberModel, error)
	PutSubscriber(string, string) error
}

type Supabase struct {
	I_database
	ApiUrl     string
	ServiceKey string
}

func (s *Supabase) GetSubscriber(email string) (subscriberModel, error) {
	url, err := url.Parse(s.ApiUrl)
	if err != nil {
		return subscriberModel{}, errors.New("bad url")
	}
	queryParams := url.Query()
	queryParams.Set("Email", fmt.Sprintf("eq%s", email))
	url.RawQuery = queryParams.Encode()
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return subscriberModel{}, errors.New("Malformed request")
	}
	req.Header.Add("apiKey", s.ServiceKey)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return subscriberModel{}, errors.New("error making request")
	}
	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return subscriberModel{}, errors.New("error reading response")
	}
	fmt.Println(bytes.NewBuffer(respBytes))
	return subscriberModel{}, nil
}

func (s *Supabase) PutSubscriber(email string, name string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		slog.Error(
			"Error invaid email",
			slog.String("email", email),
		)
		return errors.New("Invalid Email")
	}
	s.GetSubscriber(email)
	subscriber := subscriberModel{Name: name, Email: email, Verified: false, Id: "XXXXY"}
	subscriberJson, err := json.Marshal(subscriber)
	fmt.Println(bytes.NewBuffer(subscriberJson))
	body := []byte(subscriberJson)
	r, _ := http.NewRequest("POST", s.ApiUrl, bytes.NewBuffer(body))
	r.Header.Add("apiKey", s.ServiceKey)
	r.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(r)
	if err != nil {
		return errors.New("Invalid Request")
	}
	responseText, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(bytes.NewBuffer(responseText))
	return nil
}
