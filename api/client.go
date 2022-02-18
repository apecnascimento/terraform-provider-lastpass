package api

import (
	"context"

	"github.com/ansd/lastpass-go"
	"github.com/jltorresm/otpgo"
)

// Secret describes a Lastpass Secret object.
type Secret struct {
	Group           string `json:"group"`
	ID              string `json:"id"`
	LastModifiedGmt string `json:"last_modified_gmt"`
	LastTouch       string `json:"last_touch"`
	Name            string `json:"name"`
	Note            string `json:"note"`
	Password        string `json:"password"`
	Share           string `json:"share"`
	URL             string `json:"url"`
	Username        string `json:"username"`
}

func (secret *Secret) fill(account *lastpass.Account) {
	secret.Group = account.Group
	secret.ID = account.ID
	secret.LastModifiedGmt = account.LastModifiedGMT
	secret.LastTouch = account.LastTouch
	secret.Name = account.Name
	secret.Note = account.Notes
	secret.Password = account.Password
	secret.Share = account.Share
	secret.URL = account.URL
	secret.Username = account.Username

}

// Client is a Lastpass client wrapper.
type Client struct {
	Username string
	Password string
	Totp     string
}

func (c *Client) login() (*lastpass.Client, error) {
	var lastpassClient *lastpass.Client

	if c.Totp != "" {

		totp := otpgo.TOTP{
			Key: c.Totp,
		}
		token, error := totp.Generate()
		if error != nil {
			return lastpassClient, error
		}

		client, error := lastpass.NewClient(context.Background(), c.Username, c.Password, lastpass.WithOneTimePassword(token), lastpass.WithTrust())
		if error != nil {
			return lastpassClient, error
		}
		lastpassClient = client
	} else {

		client, error := lastpass.NewClient(context.Background(), c.Username, c.Password, lastpass.WithTrust())
		if error != nil {
			return lastpassClient, error
		}
		lastpassClient = client

	}

	return lastpassClient, nil
}
