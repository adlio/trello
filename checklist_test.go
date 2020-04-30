package trello

import (
	"testing"
)

func TestCreateChecklist(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("checklists", "checklist-create.json").URL

	card := Card{
		Name:    "Test Card Create",
		IDList:  "57f03a06b5ff33a63c8be316",
		ID:      "5c41028905a859019e323bc9",
		IDBoard: "5c41027ca9c378795b5a5036",
	}

	checklistName := "hello"

	cl, err := c.CreateChecklist(&card, checklistName, Arguments{"pos": "33"})
	if err != nil {
		t.Error(err)
	}

	if card.Checklists[0].ID != "5cc064cb72fbdb774ff22bac" {
		t.Errorf("Expected checklist to pick up an ID. Instead got '%s'.", card.ID)
	}

	if cl.ID != "5cc064cb72fbdb774ff22bac" {
		t.Errorf("Expected checklist to pick up an ID. Instead got '%s'.", cl.ID)
	}

	if cl.IDCard != card.ID {
		t.Errorf("Expected checklist to pick up card ID. Instead got '%s'.", cl.IDCard)
	}

	if cl.IDBoard != card.IDBoard {
		t.Errorf("Expected checklist to pick up board ID. Instead got '%s'.", cl.IDBoard)
	}

	if card.IDBoard != "5c41027ca9c378795b5a5036" {
		t.Errorf("Expected card to keep its IDBOard. Instead got '%s'.", card.IDBoard)
	}

	if cl.Name != checklistName {
		t.Errorf("Expected checklist name to be set. Instead got '%s'.", cl.Name)
	}

	if cl.client == nil {
		t.Errorf("Expected checklist to pick up a client. Instead got nil.")
	}
	if cl.Pos != 33 {
		t.Errorf("Expected the returned checklist to pick up a position. Instead got '%v'.", cl.Pos)
	}
}

func TestCreateCheckItem(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("checklists", "checkitem-create.json").URL

	cl := Checklist{
		Name:   "SomeName",
		ID:     "5cc05fc2a44eed7872662d1b",
		client: c,
	}

	checkItemName := "hello2"

	item, err := cl.CreateCheckItem(checkItemName, Arguments{"pos": "35", "checked": "true"})
	if err != nil {
		t.Error(err)
	}

	if item.IDChecklist != "5cc05fc2a44eed7872662d1b" {
		t.Errorf("Expected checkitem to pick up checklist ID. Instead got '%s'.", item.IDChecklist)
	}
	if item.ID != "5cc05fddf0d64d1c89e2a3b5" {
		t.Errorf("Expected checkitem to pick up an ID. Instead got '%s'.", item.ID)
	}
	if len(cl.CheckItems) != 1 {
		t.Errorf("Expected checklist to pick up the created checkitem. Instead got '%v'.", len(cl.CheckItems))
	}

	if cl.CheckItems[0] != *item {
		t.Errorf("Expected the returned item and the checkitem inside the checklist to be equal.\n got: %#v\nwant: %#v", cl.CheckItems[0], *item)
	}
	if item.Pos != 35 {
		t.Errorf("Expected the returned item to pick up a position. Instead got '%v'.", item.Pos)
	}
	if item.State != "complete" {
		t.Errorf("Expected checked to be set. Instead got '%s'.", item.State)
	}
}

func TestGetChecklist(t *testing.T) {
	checklist := testChecklist(t)
	if checklist.Name != "Example checklist" {
		t.Errorf("Name incorrect. Got '%s'", checklist.Name)
	}

	if checklist.Pos != 1 {
		t.Errorf("Pos incorrect. Got '%0.2f'", checklist.Pos)
	}

	if len(checklist.CheckItems) != 1 {
		t.Errorf("len(checklist.CheckItems) incorrect. Got '%0.2f'", checklist.Pos)
	}

	if checklist.CheckItems[0].Name != "Example checkItem" {
		t.Errorf("CheckItem Name incorrect. Got '%s'", checklist.CheckItems[0].Name)
	}

	if checklist.CheckItems[0].State != "complete" {
		t.Errorf("CheckItem State incorrect. Got '%s'", checklist.CheckItems[0].State)
	}

	if checklist.CheckItems[0].Pos != 2 {
		t.Errorf("CheckItem Pos incorrect. Got '%0.2f'", checklist.CheckItems[0].Pos)
	}
}

// Utility function to get a simple response from Client.GetChecklist()
//
func testChecklist(t *testing.T) *Checklist {
	c := testClient()
	c.BaseURL = mockResponse("checklists", "checklist-api-example.json").URL
	checklist, err := c.GetChecklist("4eea503", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return checklist
}
