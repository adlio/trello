// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"time"
)

// Board represents a Trello Board.
// https://developers.trello.com/reference/#boardsid
type Board struct {
	client         *Client
	ID             string `json:"id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Closed         bool   `json:"closed"`
	IDOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	Starred        bool   `json:"starred"`
	URL            string `json:"url"`
	ShortURL       string `json:"shortUrl"`
	Prefs          struct {
		PermissionLevel       string            `json:"permissionLevel"`
		Voting                string            `json:"voting"`
		Comments              string            `json:"comments"`
		Invitations           string            `json:"invitations"`
		SelfJoin              bool              `json:"selfjoin"`
		CardCovers            bool              `json:"cardCovers"`
		CardAging             string            `json:"cardAging"`
		CalendarFeedEnabled   bool              `json:"calendarFeedEnabled"`
		Background            string            `json:"background"`
		BackgroundColor       string            `json:"backgroundColor"`
		BackgroundImage       string            `json:"backgroundImage"`
		BackgroundImageScaled []BackgroundImage `json:"backgroundImageScaled"`
		BackgroundTile        bool              `json:"backgroundTile"`
		BackgroundBrightness  string            `json:"backgroundBrightness"`
		CanBePublic           bool              `json:"canBePublic"`
		CanBeOrg              bool              `json:"canBeOrg"`
		CanBePrivate          bool              `json:"canBePrivate"`
		CanInvite             bool              `json:"canInvite"`
	} `json:"prefs"`
	Subscribed bool `json:"subscribed"`
	LabelNames struct {
		Black  string `json:"black,omitempty"`
		Blue   string `json:"blue,omitempty"`
		Green  string `json:"green,omitempty"`
		Lime   string `json:"lime,omitempty"`
		Orange string `json:"orange,omitempty"`
		Pink   string `json:"pink,omitempty"`
		Purple string `json:"purple,omitempty"`
		Red    string `json:"red,omitempty"`
		Sky    string `json:"sky,omitempty"`
		Yellow string `json:"yellow,omitempty"`
	} `json:"labelNames"`
	Lists        []*List      `json:"lists"`
	Actions      []*Action    `json:"actions"`
	Organization Organization `json:"organization"`
}

// NewBoard is a constructor that sets the default values
// for Prefs.SelfJoin and Prefs.CardCovers also set by the API.
func NewBoard(name string) Board {
	b := Board{Name: name}

	// default values in line with API POST
	b.Prefs.SelfJoin = true
	b.Prefs.CardCovers = true

	return b
}

// BackgroundImage is a nested resource of Board.
type BackgroundImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// SetClient can be used to override this Board's internal connection to the
// Trello API. Normally, this is set automatically after calls to GetBoard()
// from the Client. This method exists for special cases where
// functions which need a Client need to be called on Board structs which
// weren't created from a Client in the first place.
func (b *Board) SetClient(newClient *Client) {
	b.client = newClient
}

// CreatedAt returns a board's created-at attribute as time.Time.
func (b *Board) CreatedAt() time.Time {
	t, _ := IDToTime(b.ID)
	return t
}

// CreateBoard creates a board remote.
// Attribute currently supported as extra argument: defaultLists, powerUps.
// Attributes currently known to be unsupported: idBoardSource, keepFromSource.
//
// API Docs: https://developers.trello.com/reference/#boardsid
func (c *Client) CreateBoard(board *Board, extraArgs ...Arguments) error {
	path := "boards"
	args := Arguments{
		"desc":             board.Desc,
		"name":             board.Name,
		"prefs_selfJoin":   fmt.Sprintf("%t", board.Prefs.SelfJoin),
		"prefs_cardCovers": fmt.Sprintf("%t", board.Prefs.CardCovers),
		"idOrganization":   board.IDOrganization,
	}

	if board.Prefs.Voting != "" {
		args["prefs_voting"] = board.Prefs.Voting
	}
	if board.Prefs.PermissionLevel != "" {
		args["prefs_permissionLevel"] = board.Prefs.PermissionLevel
	}
	if board.Prefs.Comments != "" {
		args["prefs_comments"] = board.Prefs.Comments
	}
	if board.Prefs.Invitations != "" {
		args["prefs_invitations"] = board.Prefs.Invitations
	}
	if board.Prefs.Background != "" {
		args["prefs_background"] = board.Prefs.Background
	}
	if board.Prefs.CardAging != "" {
		args["prefs_cardAging"] = board.Prefs.CardAging
	}

	args.flatten(extraArgs)

	err := c.Post(path, args, &board)
	if err == nil {
		board.client = c
	}
	return err
}

// Update PUTs the supported board attributes remote and updates
// the struct from the returned values.
func (b *Board) Update(extraArgs ...Arguments) error {
	args := flattenArguments(extraArgs)
	return b.client.PutBoard(b, args)
}

// Delete makes a DELETE call for the receiver Board.
func (b *Board) Delete(extraArgs ...Arguments) error {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("boards/%s", b.ID)
	return b.client.Delete(path, args, b)
}

// AddedMembersResponse represents a response after adding a new member.
type AddedMembersResponse struct {
	ID          string        `json:"id"`
	Members     []*Member     `json:"members"`
	Memberships []*Membership `json:"memberships"`
}

// AddMember adds a new member to the board.
// https://developers.trello.com/reference#boardsidlabelnamesmembers
func (b *Board) AddMember(member *Member, extraArgs ...Arguments) (response *AddedMembersResponse, err error) {
	args := Arguments{
		"email": member.Email,
	}

	args.flatten(extraArgs)

	path := fmt.Sprintf("boards/%s/members", b.ID)
	err = b.client.Put(path, args, &response)
	return
}

// GetBoard retrieves a Trello board by its ID.
func (c *Client) GetBoard(boardID string, extraArgs ...Arguments) (board *Board, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("boards/%s", boardID)
	err = c.Get(path, args, &board)
	if board != nil {
		board.client = c
	}
	return
}

// GetMyBoards returns a slice of all boards associated with the credentials set on the client.
func (c *Client) GetMyBoards(extraArgs ...Arguments) (boards []*Board, err error) {
	args := flattenArguments(extraArgs)
	path := "members/me/boards"
	err = c.Get(path, args, &boards)
	for i := range boards {
		boards[i].client = c
	}
	return
}

// GetBoards returns a slice of all public boards of the receiver Member.
func (m *Member) GetBoards(extraArgs ...Arguments) (boards []*Board, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("members/%s/boards", m.ID)
	err = m.client.Get(path, args, &boards)
	for i := range boards {
		boards[i].client = m.client

		for j := range boards[i].Lists {
			boards[i].Lists[j].client = m.client
		}
	}
	return
}

// PutBoard PUTs a board remote. Extra arguments are currently unsupported.
//
// API Docs: https://developers.trello.com/reference#idnext
func (c *Client) PutBoard(board *Board, extraArgs ...Arguments) error {
	path := fmt.Sprintf("boards/%s", board.ID)
	args := Arguments{
		"desc":             board.Desc,
		"name":             board.Name,
		"prefs/selfJoin":   fmt.Sprintf("%t", board.Prefs.SelfJoin),
		"prefs/cardCovers": fmt.Sprintf("%t", board.Prefs.CardCovers),
		"idOrganization":   board.IDOrganization,
	}

	if board.Prefs.Voting != "" {
		args["prefs/voting"] = board.Prefs.Voting
	}
	if board.Prefs.PermissionLevel != "" {
		args["prefs/permissionLevel"] = board.Prefs.PermissionLevel
	}
	if board.Prefs.Comments != "" {
		args["prefs/comments"] = board.Prefs.Comments
	}
	if board.Prefs.Invitations != "" {
		args["prefs/invitations"] = board.Prefs.Invitations
	}
	if board.Prefs.Background != "" {
		args["prefs/background"] = board.Prefs.Background
	}
	if board.Prefs.CardAging != "" {
		args["prefs/cardAging"] = board.Prefs.CardAging
	}

	args.flatten(extraArgs)

	err := c.Put(path, args, &board)
	if err == nil {
		board.client = c
	}
	return err
}
