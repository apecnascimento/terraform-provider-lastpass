package api

import (
	"context"
)

// Get secret by name
func (c *Client) GetByName(name string) (Secret, error) {
	var secret Secret
	client, err := c.login()
	if err != nil {
		return secret, err
	}

	accounts, _ := client.Accounts(context.Background())

	for _, account := range accounts {
		if account.Name == name {
			secret.fill(account)
		}
	}
	return secret, nil
}
