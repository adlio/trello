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
		list.setClient(c)
		// list.Board, err = c.GetBoard(list.IDBoard, Defaults()) // Set Parent
	}
	return
}

// GetLists takes Arguments and returns the lists of the receiver Board.
func (b *Board) GetLists(extraArgs ...Arguments) (lists []*List, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("boards/%s/lists", b.ID)
	err = b.client.Get(path, args, &lists)
	for _, list := range lists {
		list.setClient(b.client)
		list.Board = b // Set Parent
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
		list.setClient(c)
		list.Board = onBoard
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

// setClient on List and sub-objects
func (l *List) setClient(client *Client) {
	l.client = client
	for _, card := range l.Cards {
		card.setClient(client)
		card.List = l // Set Parent
	}
}
