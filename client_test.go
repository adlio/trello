package trello

func testClient() *Client {
	c := NewClient("user", "pass")
	return c
}
