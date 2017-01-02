package trello

import (
	"testing"
	"time"
)

func TestIDToTimeWithBlankString(t *testing.T) {
	t1, err := IDToTime("")

	if err != nil {
		t.Error(err)
	}

	if !t1.IsZero() {
		t.Error("IDToTime should return a zero-value time for empty strings.")
	}
}

func TestIDToStringWithKnownTime(t *testing.T) {
	t1, err := IDToTime("5865ba842057a239b32d38c0")
	if err != nil {
		t.Error(err)
	}

	// Convert to UTC for consistent comparisons regardless of which
	// time zone the test computer is in.
	t1 = t1.UTC()

	if t1.Format(time.RFC3339) != "2016-12-30T01:38:12Z" {
		t.Errorf("Expected '2016-12-30T01:38:12Z', got '%s'", t1.Format(time.RFC3339))
	}

}

func TestIDToStringWithInvalidTime(t *testing.T) {
	_, err := IDToTime("tooshort")
	if err == nil {
		t.Error("ID 'tooshort' should produce an error.")
	}
}
