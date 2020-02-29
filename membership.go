package trello

// Membership represents a Trello membership.
// https://developers.trello.com/reference#memberships-nested-resource
type Membership struct {
	ID          string `json:"id"`
	MemberID    string `json:"idMember"`
	Type        string `json:"memberType"`
	Unconfirmed bool   `json:"unconfirmed"`
	Deactivated bool   `json:"deactivated"`
}
