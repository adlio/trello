// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

type Webhook struct {
	client      *Client
	ID          string `json:"id,omitempty"`
	IDModel     string `json:"idModel"`
	Description string `json:"description"`
	CallbackURL string `json:"callbackURL"`
	Active      bool   `json:"active"`
}

func (c *Client) CreateWebhook(webhook *Webhook) error {
	path := "webhooks"
	args := Arguments{"idModel": webhook.IDModel, "description": webhook.Description, "callbackURL": webhook.CallbackURL}
	err := c.Post(path, args, webhook)
	if err == nil {
		webhook.client = c
	}
	return err
}

func (c *Client) GetWebhook(webhookID string, args Arguments) (webhook *Webhook, err error) {
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = c.Get(path, args, &webhook)
	if webhook != nil {
		webhook.client = c
	}
	return
}

func (t *Token) GetWebhooks(args Arguments) (webhooks []*Webhook, err error) {
	path := fmt.Sprintf("tokens/%s/webhooks", t.ID)
	err = t.client.Get(path, args, &webhooks)
	return
}
