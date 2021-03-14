// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

// Member represents a Trello member.
// https://developers.trello.com/reference/#member-object
type Member struct {
	client          *Client
	ID              string   `json:"id"`
	Username        string   `json:"username"`
	FullName        string   `json:"fullName"`
	Initials        string   `json:"initials"`
	AvatarHash      string   `json:"avatarHash"`
	Email           string   `json:"email"`
	IDBoards        []string `json:"idBoards"`
	IDOrganizations []string `json:"idOrganizations"`
}

// GetMember takes a member id and Arguments and returns a Member or an error.
func (c *Client) GetMember(memberID string, extraArgs ...Arguments) (member *Member, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("members/%s", memberID)
	err = c.Get(path, args, &member)
	if err == nil {
		member.client = c
	}
	return
}

// GetMyMember returns Member for the user authenticating the API call
func (c *Client) GetMyMember(args Arguments) (member *Member, err error) {
	path := fmt.Sprintf("members/me")
	err = c.Get(path, args, &member)
	if err == nil {
		member.client = c
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the organization or an error.
func (o *Organization) GetMembers(extraArgs ...Arguments) (members []*Member, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("organizations/%s/members", o.ID)
	err = o.client.Get(path, args, &members)
	for i := range members {
		members[i].client = o.client
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the Board or an error.
func (b *Board) GetMembers(extraArgs ...Arguments) (members []*Member, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("boards/%s/members", b.ID)
	err = b.client.Get(path, args, &members)
	for i := range members {
		members[i].client = b.client
	}
	return
}

// GetMembers takes Arguments and returns a slice of all members of the Card or an error.
func (c *Card) GetMembers(extraArgs ...Arguments) (members []*Member, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("cards/%s/members", c.ID)
	err = c.client.Get(path, args, &members)
	for i := range members {
		members[i].client = c.client
	}
	return
}
