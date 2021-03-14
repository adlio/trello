// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import "fmt"

// Label represents a Trello label.
// Labels are defined per board, and can be applied to the cards on that board.
// https://developers.trello.com/reference/#label-object
type Label struct {
	ID      string `json:"id"`
	IDBoard string `json:"idBoard"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Uses    int    `json:"uses"`
}

// GetLabel takes a label id and Arguments and returns the matching label (per Trello member)
// or an error.
func (c *Client) GetLabel(labelID string, extraArgs ...Arguments) (label *Label, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("labels/%s", labelID)
	err = c.Get(path, args, &label)
	return
}

// GetLabels takes Arguments and returns a slice containing all labels of the receiver board or an error.
func (b *Board) GetLabels(extraArgs ...Arguments) (labels []*Label, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("boards/%s/labels", b.ID)
	err = b.client.Get(path, args, &labels)
	return
}
