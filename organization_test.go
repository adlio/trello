// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestGetOrganization(t *testing.T) {
	organization := testOrganization(t)
	if organization.DisplayName != "Culture Foundry" {
		t.Errorf("Expected name 'Culture Foundry'. Got '%s'.", organization.DisplayName)
	}
}

func testOrganization(t *testing.T) *Organization {
	client := testClient()
	client.BaseURL = mockResponse("organizations", "culturefoundry.json").URL
	organization, err := client.GetOrganization("culturefoundry", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return organization
}
