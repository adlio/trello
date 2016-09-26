package trello

import (
	"fmt"
)

type List struct {
	client  *Client
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	IDBoard string  `json:"idBoard,omitempty"`
	Closed  bool    `json:"closed"`
	Pos     float32 `json:"pos,omitempty"`
	Board   *Board  `json:"board,omitempty"`
	Cards   []Card  `json:"cards,omitempty"`
}

func (c *Client) GetList(listID string, args Arguments) (list *List, err error) {
	path := fmt.Sprintf("lists/%s", listID)
	err = c.Get(path, args, &list)
	if list != nil {
		list.client = c
	}
	return
}

func (b *Board) GetLists(args Arguments) (lists []*List, err error) {
	path := fmt.Sprintf("boards/%s/lists", b.ID)
	err = b.client.Get(path, args, &lists)
	for i := range lists {
		lists[i].client = b.client
	}
	return
}
