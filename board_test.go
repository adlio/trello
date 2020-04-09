// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
	"time"
)

func TestCreateBoard(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("boards", "AkFGHS12.json").URL

	board := Board{
		Name: "Test Board Create",
	}

	err := c.CreateBoard(&board, Defaults())
	if err != nil {
		t.Error(err)
	}

	if board.ID != "5c602cf77061a8169a69deb5" {
		t.Errorf("Expected board to pick up an ID. Instead got '%s'.", board.ID)
	}
}

func TestDeleteBoard(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("boards", "deleted.json").URL

	board := Board{
		ID:   "5c602cf77061a8169a69deb5",
		Name: "Test Board Create",
	}
	board.client = c

	err := board.Delete(Defaults())
	if err != nil {
		t.Error(err)
	}
}

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

func TestGetBoardWithListsAndActions(t *testing.T) {
	board := testBoardWithListsAndActions(t)
	if board.Name != "Public Trello Boards" {
		t.Errorf("Incorrect board name '%s'", board.Name)
	}

	if len(board.Lists) != 4 {
		t.Errorf("Expected %d lists. Got %d", 4, len(board.Lists))
	}

	if len(board.Actions) != 43 {
		t.Errorf("Expected %d actions. Got %d", 4, len(board.Actions))
	}
}

func TestGetBoards(t *testing.T) {
	c := testClient()

	c.BaseURL = mockResponse("members", "api-example.json").URL
	member, err := c.GetMember("4ee7df1", Defaults())
	if err != nil {
		t.Fatal(err)
	}

	args := Defaults()
	args["lists"] = "all"

	c.BaseURL = mockResponse("boards", "member-boards-example.json").URL
	boards, err := member.GetBoards(args)
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

	if len(boards[1].Lists) != 1 {
		t.Error("Lists not sideloaded:", boards[0].Lists)
	}

	if boards[1].client != boards[1].Lists[0].client {
		t.Error("Client not passed to list")
	}
}

func TestGetMyBoards(t *testing.T) {
	c := testClient()

	c.BaseURL = mockResponse("boards", "member-boards-example.json").URL
	boards, err := c.GetMyBoards(Defaults())
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

func TestBoardUpdate(t *testing.T) {
	expected := map[string]map[string]string{
		"created": map[string]string{
			"id":          "5d2ccd3015468d3df508f10d",
			"name":        "test-board-for-update",
			"description": "Some description",
			"cardAging":   "regular",
		},
		"updated": map[string]string{
			"id":          "5d2ccd3015468d3df508f10d",
			"name":        "test-board-for-update plus",
			"description": "Some other description",
			"cardAging":   "pirate",
		},
	}

	board := Board{
		ID:   expected["created"]["id"],
		Name: expected["created"]["name"],
		Desc: expected["created"]["description"],
	}
	board.Prefs.CardAging = "regular"

	client := testClient()
	board.client = client
	boardResponse := mockResponse("boards", "5d2ccd3015468d3df508f10d", "create.json")
	client.BaseURL = boardResponse.URL

	err := client.CreateBoard(&board, Defaults())
	if err != nil {
		t.Error(err)
	}
	if board.ID != expected["created"]["id"] {
		t.Errorf("Expected board to pick up ID. Instead got '%s'.", board.ID)
	}
	if board.Name != expected["created"]["name"] {
		t.Errorf("Expected board name. Instead got '%s'.", board.Name)
	}
	if board.Desc != expected["created"]["description"] {
		t.Errorf("Expected board description. Instead got '%s'.", board.Desc)
	}
	if board.Prefs.CardAging != expected["created"]["cardAging"] {
		t.Errorf("Expected board's card aging. Instead got '%s'.", board.Prefs.CardAging)
	}

	board.Name = expected["updated"]["name"]
	board.Desc = expected["updated"]["description"]
	board.Prefs.CardAging = expected["updated"]["cardAging"]

	boardResponse = mockResponse("boards", "5d2ccd3015468d3df508f10d", "update.json")
	client.BaseURL = boardResponse.URL

	err = board.Update(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if board.ID != expected["updated"]["id"] {
		t.Errorf("Expected board to pick up ID. Instead got '%s'.", board.ID)
	}
	if board.Name != expected["updated"]["name"] {
		t.Errorf("Expected board name. Instead got '%s'.", board.Name)
	}
	if board.Desc != expected["updated"]["description"] {
		t.Errorf("Expected board description. Instead got '%s'.", board.Desc)
	}
	if board.Prefs.CardAging != expected["updated"]["cardAging"] {
		t.Errorf("Expected board's card aging. Instead got '%s'.", board.Prefs.CardAging)
	}

	return
}

func testBoardWithListsAndActions(t *testing.T) *Board {
	c := testClient()
	boardResponse := mockResponse("boards", "rq2mYJNn.json")
	c.BaseURL = boardResponse.URL
	board, err := c.GetBoard("rq2mYJNn", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return board
}

func TestBoardAddMember(t *testing.T) {
	board := Board{
		ID:   "5d2ccd3015468d3df508f10d",
		Name: "Test Board Create",
	}

	client := testClient()
	board.client = client

	boardResponse := mockResponse("boards/5d2ccd3015468d3df508f10d", "added_members.json")
	client.BaseURL = boardResponse.URL

	member := Member{Email: "test@test.com"}

	response, err := board.AddMember(&member, Arguments{"type": "fake"})
	if err != nil {
		t.Error(err)
	}

	if response.ID != "5d2ccd3015468d3df508f10d" {
		t.Errorf("Name of first board incorrect. Got: '%s'", response.ID)
	}

	if len(response.Members) != 2 {
		t.Errorf("Expected 2 members, got %d", len(response.Members))
	}

	if response.Members[1].Username != "user98198126" {
		t.Errorf("Username of invited member incorrect, got %s", response.Members[1].Username)
	}

	if response.Members[1].FullName != "user" {
		t.Errorf("Full name of invited member incorrect, got %s", response.Members[1].FullName)
	}

	if len(response.Memberships) != 2 {
		t.Errorf("Expected 2 memberships, got %d", len(response.Memberships))
	}

	if response.Memberships[1].Type != "normal" {
		t.Errorf("Type of membership incorrect, got %v", response.Memberships[1].Type)
	}

	if response.Memberships[1].Unconfirmed != true {
		t.Errorf("Status membership incorrect, got %v", response.Memberships[1].Unconfirmed)
	}
}
