// Copyright © 2018 Miguel Ángel Ajo
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestGetCustomField(t *testing.T) {
	customField := testCustomField(t)
	if customField.Name != "Priority" {
		t.Errorf("Name incorrect. Got '%s'", customField.Name)
	}
}

func TestGetCustomFieldsOnBoard(t *testing.T) {
	customFields := testBoardCustomFields(t)

	if len(customFields) != 2 {
		t.Errorf("Expected 2 custom fields, got %d", len(customFields))
	}

	for _, cf := range customFields {
		if len(cf.Options) != 2 {
			t.Errorf("Expected 2 options on custom field %s, got %d", cf.Name, len(cf.Options))
		}
	}

}

func testBoardCustomFields(t *testing.T) []*CustomField {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("boards", "4ed7e27fe6abb2517a21383d", "customFields.json").URL
	customFields, err := board.GetCustomFields(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return customFields
}

func testCustomField(t *testing.T) *CustomField {
	c := testClient()
	c.BaseURL = mockResponse("customFields", "api-example.json").URL
	customField, err := c.GetCustomField("5a98670bd6afbd6de1c8c360", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return customField
}
