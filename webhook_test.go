// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestCreateWebhook(t *testing.T) {
	client := testClient()
	client.BaseURL = mockResponse("webhooks", "webhook-create.json").URL
	wh := Webhook{IDModel: "test", Description: "Webhook name", CallbackURL: "http://example.com/test"}
	err := client.CreateWebhook(&wh)
	if err != nil {
		t.Error(err)
	}

	if wh.ID != "57f1c02b618bc5da74ad3874" {
		t.Errorf("Unexecpedted resultant Webhook ID: '%s'.", wh.ID)
	}

	if wh.Active != true {
		t.Error("Expected resulting webhook to be active.")
	}

	if wh.Description == "webhook name" {
		t.Errorf("Webhook description should have been retrieved from the server. Got '%s'.", wh.Description)
	}
}

func TestGetWebhook(t *testing.T) {
	client := testClient()
	client.BaseURL = mockResponse("webhooks", "webhook.json").URL
	webhook, err := client.GetWebhook("webhookID", Defaults())

	if err != nil {
		t.Error(err)
	}

	if webhook.ID != "57f1c02b618bc5da74ad3874" {
		t.Errorf("Incorrect webhook.ID: '%s'.", webhook.ID)
	}

	if webhook.Active != true {
		t.Error("Webhook should be active.")
	}

	if webhook.CallbackURL != "http://example.com/uvbhswuv" {
		t.Errorf("Webhook has incorrect callback URL: '%s'.", webhook.CallbackURL)
	}

}

func TestGetWebhooks(t *testing.T) {
	token := testToken(t)
	token.client.BaseURL = mockResponse("webhooks", "webhooks.json").URL

	webhooks, err := token.GetWebhooks(Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(webhooks) != 2 {
		t.Errorf("Expected 2 webhooks. Got %d", len(webhooks))
	}
}
