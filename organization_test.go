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
	if len(organization.PowerUps) != 1 {
		t.Errorf("Expected PowerUps to have length of 1 but was %d", len(organization.PowerUps))
	}
	if organization.PowerUps[0] != 42 {
		t.Errorf("Expected first PowerUp to be %d but was %d", 42, organization.PowerUps[0])
	}
	if len(organization.Products) != 1 {
		t.Errorf("Expected Products to have length of 1 but was %d", len(organization.Products))
	}
	if organization.Products[0] != 110 {
		t.Errorf("Expected first Product to be %d but was %d", 110, organization.Products[0])
	}
}

func TestGetBoardsInOrganization(t *testing.T) {
	organization := testOrganization(t)
	if organization.DisplayName != "Culture Foundry" {
		t.Errorf("Expected name 'Culture Foundry'. Got '%s'.", organization.DisplayName)
	}

	client := testClient()
	client.BaseURL = mockResponse("organizations", "571ab6ad9dc91c597d6e9f90", "boards", "boards.json").URL

	boards, err := client.GetBoardsInOrganization(organization.ID)
	if err != nil {
		t.Fatalf("Expected boards in organization to be returned. Got error: %v", err)
	}
	if boards == nil {
		t.Fatalf("Expected boards slice to contain boards in the test organization with ID %s. Slice was nil.", organization.ID)
	}
}

func TestOrganizationSetClient(t *testing.T) {
	o := Organization{}
	client := testClient()
	o.SetClient(client)
	if o.client == nil {
		t.Error("Expected non-nil Organization.client")
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
