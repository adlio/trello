package trello

import (
	"fmt"
	"testing"
	"time"
)

func TestListDurationsForTypicalCard(t *testing.T) {
	card := testCard(t)
	card.client.BaseURL = mockResponse("actions", "card-actions-typical.json").URL
	durations, err := card.GetListDurations()
	if err != nil {
		t.Error(err)
	}
	d1 := durations[0]
	if d1.ListName != "Backlog" {
		t.Errorf("Incorrect list name '%s'.", d1.ListName)
	}

	if len(durations) != 7 {
		t.Errorf("Incorrect ListDuration count %d", len(durations))
	}

	if d1.Duration.Seconds() > 30 {
		t.Errorf("Expected less than 30 seconds in Backlog. Was %.2f", d1.Duration.Seconds())
	}

	if d1.FirstEntered.Format(time.Kitchen) != "10:44PM" {
		t.Errorf("Incorrect FirstEntered time: '%s'.", d1.FirstEntered.Format(time.Kitchen))
	}
}

func TestListDurationsForDoneCard(t *testing.T) {
	card := testCard(t)
	card.client.BaseURL = mockResponse("actions", "card-actions-done.json").URL
	durations, err := card.GetListDurations()
	if err != nil {
		t.Error(err)
	}

	d7 := durations[6]

	if len(durations) != 7 {
		for _, d := range durations {
			fmt.Println(d.ListName, d.Duration)
		}
		t.Errorf("Expected 7 durations. Got %d.", len(durations))
	}

	if d7.ListName != "Done" {
		for _, d := range durations {
			fmt.Println(d.ListName, d.Duration, d.FirstEntered)
		}
		t.Errorf("Incorrect list name '%s'.", d7.ListName)
	}

	if d7.Duration.Minutes() < 40 {
		t.Errorf("Expected card to report in Done longer than 60 minutes. Was %.2f", d7.Duration.Minutes())
	}
}

func TestListDurationsForReworkedCard(t *testing.T) {
	card := testCard(t)
	card.client.BaseURL = mockResponse("actions", "card-actions-rework.json").URL
	durations, err := card.GetListDurations()
	if err != nil {
		t.Error(err)
	}

	d4 := durations[3]
	if d4.ListName != "Doing" {
		t.Errorf("Incorrect list name '%s'.", d4.ListName)
	}

	if d4.Duration.Minutes() != 480 {
		t.Errorf("Expected duration in Doing to be 480 minutes. Was %.2f", d4.Duration.Minutes())
	}

	if d4.FirstEntered.Format(time.Kitchen) != "12:03AM" {
		t.Errorf("Incorrect FirstEntered date in Doing: '%s'.", d4.FirstEntered.Format(time.Kitchen))
	}

	if d4.TimesInList != 2 {
		t.Error("Reworked card went through Doing twice, but TimesInList doesn't indicate that.")
	}

	d6 := durations[5]
	if d6.ListName != "QA" {
		t.Errorf("Incorrect list name: '%s'.", d6.ListName)
	}

	if d6.TimesInList != 2 {
		t.Error("Reworked card went through QA twice, but TimesInList doesn't indicate that.")
	}

	d1 := durations[0]
	if d1.TimesInList != 1 {
		t.Errorf("Reworked card only went in Backlog once, but TimesInList says %d", d1.TimesInList)
	}
}

func TestListDurationsForRevivedCard(t *testing.T) {
	card := testCard(t)
	card.client.BaseURL = mockResponse("actions", "card-actions-revived.json").URL
	durations, err := card.GetListDurations()
	if err != nil {
		t.Error(err)
	}

	d2 := durations[1]
	if d2.ListName != "Approved Work" {
		t.Errorf("Incorrect ListName '%s'", d2.ListName)
	}
	if d2.Duration.Minutes() != 62 {
		t.Errorf("Expected 62.0 minutes in Approved Work, got %.2f", d2.Duration.Minutes())
	}
}
