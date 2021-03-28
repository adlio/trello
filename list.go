// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

// List represents Trello lists.
// https://developers.trello.com/reference/#list-object
type List struct {
	client     *Client
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	IDBoard    string  `json:"idBoard,omitempty"`
	Closed     bool    `json:"closed"`
	Pos        float32 `json:"pos,omitempty"`
	Subscribed bool    `json:"subscribed"`
	Board      *Board  `json:"board,omitempty"`
	Cards      []*Card `json:"cards,omitempty"`
}

// SetClient can be used to override this List's internal connection to the
// Trello API. Normally, this is set automatically after calls to GetList()
// from the Client. This method exists for special cases where
// functions which need a Client need to be called on List structs which
// weren't created from a Client in the first place.
func (l *List) SetClient(newClient *Client) {
	l.client = newClient
}

// CreatedAt returns the time.Time from the list's id.
func (l *List) CreatedAt() time.Time {
	t, _ := IDToTime(l.ID)
	return t
}

// GetList takes a list's id and Arguments and returns the matching list.
func (c *Client) GetList(listID string, extraArgs ...Arguments) (list *List, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("lists/%s", listID)
	err = c.Get(path, args, &list)
	if list != nil {
		list.client = c
		for i := range list.Cards {
			list.Cards[i].client = c
		}
	}
	return
}

// GetLists takes Arguments and returns the lists of the receiver Board.
func (b *Board) GetLists(extraArgs ...Arguments) (lists []*List, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("boards/%s/lists", b.ID)
	err = b.client.Get(path, args, &lists)
	for i := range lists {
		lists[i].client = b.client
		for j := range lists[i].Cards {
			lists[i].Cards[j].client = b.client
		}
	}
	return
}

// CreateList creates a list.
// Attribute currently supported as extra argument: pos.
// Attributes currently known to be unsupported: idListSource.
//
// API Docs: https://developers.trello.com/reference/#lists-1
func (c *Client) CreateList(onBoard *Board, name string, extraArgs ...Arguments) (list *List, err error) {
	path := "lists"
	args := Arguments{
		"name":    name,
		"pos":     "top",
		"idBoard": onBoard.ID,
	}

	args.flatten(extraArgs)

	list = &List{}
	err = c.Post(path, args, &list)
	if err == nil {
		list.client = c
	}
	return
}

// CreateList creates a list.
// Attribute currently supported as extra argument: pos.
// Attributes currently known to be unsupported: idListSource.
//
// API Docs: https://developers.trello.com/reference/#lists-1
func (b *Board) CreateList(name string, extraArgs ...Arguments) (list *List, err error) {
	args := flattenArguments(extraArgs)
	return b.client.CreateList(b, name, args)
}

// Update UPDATEs the list's attributes.
// API Docs: https://developers.trello.com/reference/#listsid-1
func (l *List) Update(extraArgs ...Arguments) error {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("lists/%s", l.ID)
	return l.client.Put(path, args, l)
}
