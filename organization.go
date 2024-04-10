// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

// Organization represents a Trello organization or team, i.e. a collection of members and boards.
// https://developers.trello.com/reference/#organizations
type Organization struct {
	client      *Client
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Desc        string `json:"desc"`
	URL         string `json:"url"`
	Website     string `json:"website"`
	Products    []int  `json:"products"`
	PowerUps    []int  `json:"powerUps"`
}

// GetOrganization takes an organization id and Arguments and either
// GETs returns an Organization, or an error.
func (c *Client) GetOrganization(orgID string, extraArgs ...Arguments) (organization *Organization, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("organizations/%s", orgID)
	err = c.Get(path, args, &organization)
	if organization != nil {
		organization.SetClient(c)
	}
	return
}

// GetBoardsInOrganization takes an organization id and Arguments and either GET returns
// a slice of boards within that organization, or an error.
func (c *Client) GetBoardsInOrganization(orgID string, extraArgs ...Arguments) (boards []*Board, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("organizations/%s/boards", orgID)
	err = c.Get(path, args, &boards)
	if err != nil {
		return nil, err
	}
	return
}

// SetClient can be used to override this Organization's internal connection
// to the Trello API. Normally, this is set automatically after API calls.
func (o *Organization) SetClient(newClient *Client) {
	o.client = newClient
}
