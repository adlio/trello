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
}

func (c *Client) GetWebhook(webhookID string, args Arguments) (webhook *Webhook, err error) {
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = c.Get(path, args, &webhook)
	if webhook != nil {
		webhook.client = c
	}
	return
}
