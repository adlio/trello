// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

// SearchResult represents a search result as collections of various
// types returned by a search, e.g. Cards or Boards.
type SearchResult struct {
	Options SearchOptions `json:"options"`
	Actions []*Action     `json:"actions,omitempty"`
	Cards   []*Card       `json:"cards,omitempty"`
	Boards  []*Board      `json:"boards,omitempty"`
	Members []*Member     `json:"members,omitempty"`
}

// SearchOptions contains options for search requests.
type SearchOptions struct {
	Terms      []SearchTerm     `json:"terms"`
	Modifiers  []SearchModifier `json:"modifiers,omitempty"`
	ModelTypes []string         `json:"modelTypes,omitempty"`
	Partial    bool             `json:"partial"`
}

// SearchModifier is wrapper for a search string.
type SearchModifier struct {
	Text string `json:"text"`
}

// SearchTerm is a string that may be negated in a search query.
type SearchTerm struct {
	Text    string `json:"text"`
	Negated bool   `json:"negated,omitempty"`
}

// SearchCards takes a query string and Arguments and returns a slice of Cards or an error.
func (c *Client) SearchCards(query string, extraArgs ...Arguments) (cards []*Card, err error) {
	args := Arguments{
		"query":      query,
		"modelTypes": "cards",
	}
	args.flatten(extraArgs)
	res := SearchResult{}
	err = c.Get("search", args, &res)
	cards = res.Cards
	for _, card := range cards {
		card.client = c
	}
	return
}

// SearchBoards takes a query string and Arguments and returns a slice of Boards or an error.
func (c *Client) SearchBoards(query string, extraArgs ...Arguments) (boards []*Board, err error) {
	args := Arguments{
		"query":      query,
		"modelTypes": "boards",
	}
	args.flatten(extraArgs)
	res := SearchResult{}
	err = c.Get("search", args, &res)
	boards = res.Boards
	for _, board := range boards {
		board.client = c
	}
	return
}

// SearchMembers takes a query string and Arguments and returns a slice of Members or an error.
func (c *Client) SearchMembers(query string, extraArgs ...Arguments) (members []*Member, err error) {
	args := Arguments{
		"query": query,
	}
	args.flatten(extraArgs)
	err = c.Get("search/members", args, &members)
	return
}
