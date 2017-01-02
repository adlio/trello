package trello

import (
	// "fmt"
	"testing"
	// "time"
)

func TestMemberDurationsForTypicalCard(t *testing.T) {
	card := testCard(t)
	card.client.BaseURL = mockResponse("actions", "card-actions-with-member.json").URL
	durations, err := card.GetMemberDurations()
	if err != nil {
		t.Error(err)
	}

	if len(durations) != 1 {
		t.Errorf("Expected 1 MemberDurations{}, got %d instead.", len(durations))
	}
}
