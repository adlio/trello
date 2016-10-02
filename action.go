// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

type Action struct {
	ID              string      `json:"id"`
	IDMemberCreator string      `json:"idMemberCreator"`
	Type            string      `json:"type"`
	Date            time.Time   `json:"date"`
	Data            *ActionData `json:"data,omitempty"`
	MemberCreator   *Member     `json:"memberCreator,omitempty"`
}

type ActionData struct {
	Text           string    `json:"text,omitempty"`
	List           *List     `json:"list,omitempty"`
	Card           *Card     `json:"card,omitempty"`
	Board          *Board    `json:"board,omitempty"`
	Old            *Card     `json:"old,omitempty"`
	ListBefore     *List     `json:"listBefore,omitempty"`
	ListAfter      *List     `json:"listAfter,omitempty"`
	DateLastEdited time.Time `json:"dateLastEdited"`

	CheckItem *CheckItem `json:"checkItem"`
	Checklist *Checklist `json:"checklist"`
}

func (b *Board) GetActions(args Arguments) (actions []*Action, err error) {
	path := fmt.Sprintf("boards/%s/actions", b.ID)
	err = b.client.Get(path, args, &actions)
	return
}

func (l *List) GetActions(args Arguments) (actions []*Action, err error) {
	path := fmt.Sprintf("lists/%s/actions", l.ID)
	err = l.client.Get(path, args, &actions)
	return
}

func (c *Card) GetActions(args Arguments) (actions []*Action, err error) {
	path := fmt.Sprintf("cards/%s/actions", c.ID)
	err = c.client.Get(path, args, &actions)
	return
}

// GetListChangeActions retrieves a slice of Actions which resulted in changes
// to the card's active List. This includes the createCard and copyCard action (which
// place the card in its first list, and the updateCard:closed action, which remove it
// from its last list.
//
// This function is just an alias for:
//   card.GetActions(Arguments{"filter": "createCard,copyCard,updateCard:idList,updateCard:closed", "limit": "1000"})
//
func (c *Card) GetListChangeActions() (actions []*Action, err error) {
	path := fmt.Sprintf("cards/%s/actions", c.ID)

	args := Arguments{
		"filter": "createCard,copyCard,updateCard:idList,updateCard:closed",
		"limit":  "1000",
	}
	err = c.client.Get(path, args, &actions)
	return
}

// ListAfterAction calculates which List the card ended up in after this action
// completed. Returns nil when the action resulted in the card being archived (in
// which case we consider it to not be in a list anymore), or when the action isn't
// related to a list at all (in which case this is a nonsensical question to ask).
//
func ListAfterAction(a *Action) *List {
	switch a.Type {
	case "copyCard":
		return a.Data.List
	case "createCard":
		return a.Data.List
	case "updateCard":
		if a.Data.Card != nil && a.Data.Card.Closed {
			return nil
		} else if a.Data.Old != nil && a.Data.Old.Closed {
			return a.Data.List
		}
		if a.Data.ListAfter != nil {
			return a.Data.ListAfter
		}
	}
	return nil
}
