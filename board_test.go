// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"testing"
	"time"
)

func TestBoardCreatedAt(t *testing.T) {
	b := Board{ID: "4d5ea62fd76aa1136000000c"}
	ts := b.CreatedAt()
	if ts.IsZero() {
		t.Error("Time shouldn't be zero.")
	}
	if ts.Unix() != 1298048559 {
		t.Errorf("Incorrect CreatedAt() time: '%s'.", ts.Format(time.RFC3339))
	}
}

func TestGetBoard(t *testing.T) {
	board := testBoard(t)
	if board.Name != "Trello Public API" {
		t.Errorf("Incorrect board name '%s'", board.Name)
	}

	if board.LabelNames.Green != "Participate!" {
		t.Errorf("Expected Green label 'Participate!'. Got '%s'", board.LabelNames.Green)
	}
}

func TestGetBoards(t *testing.T) {
	c := testClient()

	c.BaseURL = mockResponse("members", "api-example.json").URL
	member, err := c.GetMember("4ee7df1", Defaults())
	if err != nil {
		t.Fatal(err)
	}

	c.BaseURL = mockResponse("boards", "member-boards-example.json").URL
	boards, err := member.GetBoards(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if len(boards) != 2 {
		t.Errorf("Expected 2 boards. Got %d", len(boards))
	}

	if boards[0].Name != "Example Board" {
		t.Errorf("Name of first board incorrect. Got: '%s'", boards[0].Name)
	}

	if boards[1].Name != "Public Board" {
		t.Errorf("Name of second board incorrect. Got: '%s'", boards[1].Name)
	}

}

func TestGetUnauthorizedBoard(t *testing.T) {
	c := testClient()
	c.BaseURL = mockErrorResponse(401).URL

	_, err := c.GetBoard("boardid", Defaults())
	if err == nil {
		t.Error("GetBoard() should have failed")
	}
}

func testBoard(t *testing.T) *Board {
	c := testClient()
	boardResponse := mockResponse("boards", "cI66RoQS.json")
	c.BaseURL = boardResponse.URL
	board, err := c.GetBoard("cIRoQS", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return board
}
