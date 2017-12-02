package trello

import (
	"testing"
)

func TestContainsCardCreation(t *testing.T) {
	ac := make(ActionCollection, 4)
	ac[0] = &Action{Type: "commentCard"}
	ac[1] = &Action{Type: "updateCard"}
	ac[2] = &Action{Type: "convertToCardFromCheckItem"}
	ac[3] = &Action{Type: "createBoard"}
	if !ac.ContainsCardCreation() {
		t.Errorf("This ActionCollection contains a card creation action, but the method returned false.")
	}
}

func TestFilterToCardCreationActions(t *testing.T) {
	ac := make(ActionCollection, 6)
	ac[0] = &Action{Type: "commentCard"}
	ac[1] = &Action{Type: "createCard"}
	ac[2] = &Action{Type: "createBoard"}
	ac[3] = &Action{Type: "emailCard"}
	ac[4] = &Action{Type: "copyCard"}
	ac[5] = &Action{Type: "convertToCardFromCheckItem"}
	ccs := ac.FilterToCardCreationActions()

	if len(ccs) != 4 {
		t.Errorf("Expected 4 cards, got %d", len(ccs))
	}

	if ccs[0].Type != "createCard" {
		t.Error("Order was not preserved.")
	}
}

func TestFilterToListChangeActions(t *testing.T) {
	ac := make(ActionCollection, 5)
	ac[0] = &Action{Type: "updateCard"} // An update that didn't change the list
	ac[1] = &Action{Type: "updateCard", Data: &ActionData{ListAfter: &List{ID: "testID", Name: "List 2"}}}
	ac[2] = &Action{Type: "updateCard", Data: &ActionData{Card: &ActionDataCard{Closed: true}}} // Card was archived
	ac[3] = &Action{Type: "updateCard", Data: &ActionData{Old: &ActionDataCard{Closed: true}}}  // Card was unarchived
	ac[4] = &Action{Type: "commentCard"}
	lcas := ac.FilterToListChangeActions()

	if len(lcas) != 3 {
		t.Errorf("Expected 3, got %d", len(lcas))
	}
}

func TestFilterToCardMembershipChangeActions(t *testing.T) {
	ac := make(ActionCollection, 5)
	ac[0] = &Action{Type: "addMemberToCard"}
	ac[1] = &Action{Type: "removeMemberFromCard"}
	ac[2] = &Action{Type: "updateCard", Data: &ActionData{Old: &ActionDataCard{Closed: true}}}
	ac[3] = &Action{Type: "updateCard", Data: &ActionData{Card: &ActionDataCard{Closed: true}}}
	ac[4] = &Action{Type: "commentCard"}
	mcas := ac.FilterToCardMembershipChangeActions()

	if len(mcas) != 4 {
		t.Errorf("Expected 1, got %d", len(mcas))
	}
}
