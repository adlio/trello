// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Webhook is the Go representation of a webhook registered in Trello's systems.
// Used when creating, modifying or deleting webhooks.
// https://developers.trello.com/reference/#webhook-object
//
type Webhook struct {
	client      *Client
	ID          string `json:"id,omitempty"`
	IDModel     string `json:"idModel"`
	Description string `json:"description"`
	CallbackURL string `json:"callbackURL"`
	Active      bool   `json:"active"`
}

// BoardWebhookRequest is the object sent by Trello to a Webhook for Board-triggered
// webhooks.
//
type BoardWebhookRequest struct {
	Model  *Board
	Action *Action
}

// ListWebhookRequest is the object sent by Trello to a Webhook for List-triggered
// webhooks.
//
type ListWebhookRequest struct {
	Model  *List
	Action *Action
}

// CardWebhookRequest is the object sent by Trello to a Webhook for Card-triggered
// webhooks.
//
type CardWebhookRequest struct {
	Model  *Card
	Action *Action
}

// CreateWebhook takes a Webhook, POSTs it and returns an error object.
func (c *Client) CreateWebhook(webhook *Webhook) error {
	path := "webhooks"
	args := Arguments{"idModel": webhook.IDModel, "description": webhook.Description, "callbackURL": webhook.CallbackURL}
	err := c.Post(path, args, webhook)
	if err == nil {
		webhook.client = c
	}
	return err
}

// Delete takes a webhook and deletes it
func (w *Webhook) Delete(args Arguments) error {
	path := fmt.Sprintf("webhooks/%s", w.ID)
	return w.client.Delete(path, Arguments{}, w)
}

// GetWebhook takes a webhook id and Arguments, GETs the matching Webhook and returns it or an error.
func (c *Client) GetWebhook(webhookID string, args Arguments) (webhook *Webhook, err error) {
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = c.Get(path, args, &webhook)
	if webhook != nil {
		webhook.client = c
	}
	return
}

// GetWebhooks takes Arguments and returns a list of all Webhooks for the receiver Token or an error.
func (t *Token) GetWebhooks(args Arguments) (webhooks []*Webhook, err error) {
	path := fmt.Sprintf("tokens/%s/webhooks", t.client.Token)
	err = t.client.Get(path, args, &webhooks)
	if err == nil {
		for _, webhook := range webhooks {
			webhook.client = t.client
		}
	}
	return
}

// GetBoardWebhookRequest takes a http.Request and returns the decoded body as BoardWebhookRequest or an error.
func GetBoardWebhookRequest(r *http.Request) (whr *BoardWebhookRequest, err error) {
	if r.Method == "HEAD" {
		return &BoardWebhookRequest{}, nil
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&whr)
	if err != nil {
		err = errors.Wrapf(err, "GetBoardWebhookRequest() failed to decode '%s'.", r.URL)
	}
	return
}

// GetListWebhookRequest takes a http.Request and returns the decoded Body as ListWebhookRequest or an error.
func GetListWebhookRequest(r *http.Request) (whr *ListWebhookRequest, err error) {
	if r.Method == "HEAD" {
		return &ListWebhookRequest{}, nil
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&whr)
	if err != nil {
		err = errors.Wrapf(err, "GetListWebhookRequest() failed to decode '%s'.", r.URL)
	}
	return
}

// GetCardWebhookRequest takes a http.Request and returns the decoded Body as CardWebhookRequest or an error.
func GetCardWebhookRequest(r *http.Request) (whr *CardWebhookRequest, err error) {
	if r.Method == "HEAD" {
		return &CardWebhookRequest{}, nil
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&whr)
	if err != nil {
		err = errors.Wrapf(err, "GetCardWebhookRequest() failed to decode '%s'.", r.URL)
	}
	return
}
