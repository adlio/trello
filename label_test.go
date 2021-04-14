// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestGetLabel(t *testing.T) {
	label := testLabel(t)
	if label.Name != "Visited" {
		t.Errorf("Title incorrect. Got '%s'", label.Name)
	}
}

func TestGetLabelsOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("labels", "board-labels-api-example.json").URL
	lists, err := board.GetLabels(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if len(lists) != 3 {
		t.Errorf("Expected 3 labels, got %d", len(lists))
	}
}

func TestCreateLabel(t *testing.T) {
	board := testBoard(t)
	label := Label{Name: "Visited", Color: "green"}
	board.client.BaseURL = mockResponse("labels", "labels-api-example.json").URL
	err := board.CreateLabel(&label)
	if err != nil {
		t.Fatal(err)
	}

	if label.Name != "Visited" {
		t.Errorf("Expected name 'Visited', got '%s'", label.Name)
	}
	if label.Color != "green" {
		t.Errorf("Expected 'green', got '%s'", label.Color)
	}
}

func TestLabelSetClient(t *testing.T) {
	l := Label{}
	client := testClient()
	l.SetClient(client)
	if l.client == nil {
		t.Error("Expected non-nil Label.client")
	}
}

// Utility function to get the standard case Client.GetList() response
//
func testLabel(t *testing.T) *Label {
	c := testClient()
	c.BaseURL = mockResponse("labels", "labels-api-example.json").URL
	label, err := c.GetLabel("4eea4ff", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return label
}
