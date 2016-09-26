// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

type Card struct {
	client *Client

	// Key metadata
	ID               string     `json:"id"`
	IDShort          int        `json:"idShort"`
	Name             string     `json:"name"`
	Pos              float32    `json:"pos"`
	Email            string     `json:"email"`
	ShortLink        string     `json:"shortLink"`
	ShortUrl         string     `json:"shortUrl"`
	Url              string     `json:"url"`
	Desc             string     `json:"desc"`
	Due              string     `json:"due"`
	Closed           bool       `json:"closed"`
	Subscribed       bool       `json:"subscribed"`
	DateLastActivity *time.Time `json:"dateLastActivity"`

	// Board
	Board   *Board
	IDBoard string `json:"idBoard"`

	// List
	List   *List
	IDList string `json:"idList"`

	// Badges
	Badges struct {
		Votes              int        `json:"votes"`
		ViewingMemberVoted bool       `json:"viewingMemberVoted"`
		Subscribed         bool       `json:"subscribed"`
		Fogbugz            string     `json:"fogbugz,omitempty"`
		CheckItems         int        `json:"checkItems"`
		CheckItemsChecked  int        `json:"checkItemsChecked"`
		Comments           int        `json:"comments"`
		Attachments        int        `json:"attachments"`
		Description        bool       `json:"description"`
		Due                *time.Time `json:"due,omitempty"`
	} `json:"badges"`

	// Checklists
	IDCheckLists    []string          `json:"idCheckLists"`
	Checklists      []*Checklist      `json:"checklists,omitempty"`
	CheckItemStates []*CheckItemState `json:"checkItemStates,omitempty"`

	// Members
	IDMembers      []string  `json:"idMembers,omitempty"`
	IDMembersVoted []string  `json:"idMembersVoted,omitempty"`
	Members        []*Member `json:"members,omitempty"`

	// Attachments
	IDAttachmentCover     string        `json:"idAttachmentCover"`
	ManualCoverAttachment bool          `json:"manualCoverAttachment"`
	Attachments           []*Attachment `json:attachments,omitempty"`

	// Labels
	Labels []*Label `json:"labels,omitempty"`
}

func (c *Client) GetCard(cardID string, args Arguments) (card *Card, err error) {
	path := fmt.Sprintf("cards/%s", cardID)
	err = c.Get(path, args, &card)
	if card != nil {
		card.client = c
	}
	return
}

/**
 * Retrieves all Cards on a Board
 */
func (b *Board) GetCards(args Arguments) (cards []*Card, err error) {
	path := fmt.Sprintf("boards/%s/cards", b.ID)
	err = b.client.Get(path, args, &cards)
	for i := range cards {
		cards[i].client = b.client
	}
	return
}

/**
 * Retrieves all Cards in a List
 */
func (l *List) GetCards(args Arguments) (cards []*Card, err error) {
	path := fmt.Sprintf("lists/%s/cards", l.ID)
	err = l.client.Get(path, args, &cards)
	for i := range cards {
		cards[i].client = l.client
	}
	return
}
