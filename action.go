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
	Board          *Board    `json:"board,omitempty"`
	List           *List     `json:"list,omitempty"`
	Card           *Card     `json:"card,omitempty"`
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
