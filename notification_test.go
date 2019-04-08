package trello

import "testing"

func TestGetMyNotifications(t *testing.T) {
	c := testClient()

	c.BaseURL = mockResponse("notifications", "member-notifications-example.json").URL
	notifications, err := c.GetMyNotifications(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if len(notifications) != 2 {
		t.Errorf("Expected 2 notifications. Got %d", len(notifications))
	}

	if notifications[0].Data.Board.Name != "Board Name" {
		t.Errorf("Name of first notification incorrect. Got: '%s'", notifications[0].Data.Board.Name)
	}

	if notifications[1].Data.Board.Name != "Board Name 2" {
		t.Errorf("Name of second notification incorrect. Got: '%s'", notifications[1].Data.Board.Name)
	}
}
