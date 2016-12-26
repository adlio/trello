// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Card struct {
	client *Client

	// Key metadata
	ID               string     `json:"id"`
	IDShort          int        `json:"idShort"`
	Name             string     `json:"name"`
	Pos              float64    `json:"pos"`
	Email            string     `json:"email"`
	ShortLink        string     `json:"shortLink"`
	ShortUrl         string     `json:"shortUrl"`
	Url              string     `json:"url"`
	Desc             string     `json:"desc"`
	Due              *time.Time `json:"due"`
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

	// Actions
	Actions []*Action `json:"actions,omitempty"`

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

func (c *Card) CreatedAt() time.Time {
	t, err := IDToTime(c.ID)
	if err != nil {
		return time.Time{}
	} else {
		return t
	}
}

func (c *Client) CreateCard(card *Card, extraArgs Arguments) error {
	path := "cards"
	args := Arguments{
		"name":      card.Name,
		"desc":      card.Desc,
		"pos":       strconv.FormatFloat(card.Pos, 'g', -1, 64),
		"idList":    card.IDList,
		"idMembers": strings.Join(card.IDMembers, ","),
	}
	if card.Due != nil {
		args["due"] = card.Due.Format(time.RFC3339)
	}
	// Allow overriding the creation position with 'top' or 'botttom'
	if pos, ok := extraArgs["pos"]; ok {
		args["pos"] = pos
	}
	err := c.Post(path, args, &card)
	if err == nil {
		card.client = c
	}
	return err
}

func (l *List) AddCard(card *Card, extraArgs Arguments) error {
	path := fmt.Sprintf("lists/%s/cards", l.ID)
	args := Arguments{
		"name":      card.Name,
		"desc":      card.Desc,
		"idMembers": strings.Join(card.IDMembers, ","),
	}
	if card.Due != nil {
		args["due"] = card.Due.Format(time.RFC3339)
	}
	// Allow overwriting the creation position with 'top' or 'bottom'
	if pos, ok := extraArgs["pos"]; ok {
		args["pos"] = pos
	}
	err := l.client.Post(path, args, &card)
	if err == nil {
		card.client = l.client
	} else {
		err = errors.Wrapf(err, "Error adding card to list %s", l.ID)
	}
	return err
}

// Try these Arguments
//
// 	Arguments["keepFromSource"] = "all"
//  Arguments["keepFromSource"] = "none"
// 	Arguments["keepFromSource"] = "attachments,checklists,comments"
//
func (c *Card) CopyToList(listID string, args Arguments) (*Card, error) {
	path := "cards"
	args["idList"] = listID
	args["idCardSource"] = c.ID
	newCard := Card{}
	err := c.client.Post(path, args, &newCard)
	if err == nil {
		newCard.client = c.client
	} else {
		err = errors.Wrapf(err, "Error copying card '%s' to list '%s'.", c.ID, listID)
	}
	return &newCard, err
}

// If this Card was created from a copy of another Card, this func retrieves
// the originating Card. Returns an error only when a low-level failure occurred.
// If this Card has no parent, nil, nil is returned.
//
func (c *Card) GetParentCard(args Arguments) (*Card, error) {
	actions, err := c.GetActions(Arguments{"filter": "copyCard"})
	if err != nil {
		err = errors.Wrapf(err, "ParentCard() failed to GetActions() for card '%s'", c.ID)
		return nil, err
	}
	if len(actions) == 0 {
		return nil, nil
	}
	for _, action := range actions {
		if action.Data.CardSource.ID != c.ID {
			card, err := c.client.GetCard(action.Data.CardSource.ID, args)
			return card, err
		}
	}
	return nil, nil
}

func (c *Card) GetAncestorCards(args Arguments) (ancestors []*Card, err error) {
	parent := c
	for parent != nil {
		if parent != c {
			ancestors = append(ancestors, parent)
		}
		parent, err = parent.GetParentCard(args)
		if err != nil {
			return
		}
	}
	return
}

func (b *Board) ContainsCopyOfCard(cardID string, args Arguments) (bool, error) {
	args["filter"] = "copyCard"
	actions, err := b.GetActions(args)
	if err != nil {
		err := errors.Wrapf(err, "GetCards() failed inside ContainsCopyOf() for board '%s' and card '%s'.", b.ID, cardID)
		return false, err
	}
	for _, action := range actions {
		if action.Data != nil && action.Data.CardSource != nil && action.Data.CardSource.ID == cardID {
			return true, nil
		}
	}
	return false, nil
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
