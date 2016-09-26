package trello

import (
	"testing"
)

func TestGetMembersOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("members", "board-members-api-example.json").URL
	members, err := board.GetMembers(Defaults)
	if err != nil {
		t.Fatal(err)
	}

	if len(members) != 3 {
		t.Errorf("Expected 3 member, got %d", len(members))
	}
}

func TestGetMembersOnCard(t *testing.T) {
	card := testCard(t)
	card.client.BaseURL = mockResponse("members", "card-members-api-example.json").URL
	members, err := card.GetMembers(Defaults)
	if err != nil {
		t.Fatal(err)
	}

	if len(members) != 1 {
		t.Errorf("Expected 1 member, got %d", len(members))
	}
}
