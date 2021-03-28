package trello

import "testing"

func TestAttachmentSetClient(t *testing.T) {
	a := Attachment{}
	client := testClient()
	a.SetClient(client)
	if a.client == nil {
		t.Error("Expected non-nil Attachment.client")
	}
}
