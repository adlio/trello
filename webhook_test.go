// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestCreateWebhook(t *testing.T) {
	client := testClient()
	server := mockResponse("webhooks", "webhook-create.json")
	defer server.Close()
	client.BaseURL = server.URL
	wh := Webhook{IDModel: "test", Description: "Webhook name", CallbackURL: "http://example.com/test"}
	err := client.CreateWebhook(&wh)
	if err != nil {
		t.Error(err)
	}

	if wh.ID != "57f1c02b618bc5da74ad3874" {
		t.Errorf("Unexpected resultant Webhook ID: '%s'.", wh.ID)
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
	server := mockResponse("webhooks", "webhook.json")
	defer server.Close()
	client.BaseURL = server.URL
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
	server := mockResponse("webhooks", "webhooks.json")
	defer server.Close()
	token.client.BaseURL = server.URL

	webhooks, err := token.GetWebhooks(Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(webhooks) != 2 {
		t.Errorf("Expected 2 webhooks. Got %d", len(webhooks))
	}
}

func TestDeleteWebhook(t *testing.T) {
	c := testClient()
	server := mockResponse("webhooks", "deleted.json")
	defer server.Close()
	c.BaseURL = server.URL

	webhook := Webhook{
		ID:          "57f1c02b618bc5da74ad3874",
		Description: "Test Web Hook",
		IDModel:     "57f039fbc0f98772398d289d",
		CallbackURL: "http://example.com/uvbhswuv",
		Active:      true,
	}
	webhook.client = c

	err := webhook.Delete(Defaults())
	if err != nil {
		t.Error(err)
	}
}
