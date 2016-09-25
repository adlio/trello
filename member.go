package trello

import (
	"fmt"
)

type Member struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	FullName   string `json:"fullName"`
	Initials   string `json:"initials"`
	AvatarHash string `json:"avatarHash"`
}

func (b *Board) GetMembers(args Arguments) (members []*Member, err error) {
	path := fmt.Sprintf("boards/%s/members", b.ID)
	err = b.client.Get(path, args, &members)
	return
}

func (c *Card) GetMembers(args Arguments) (members []*Member, err error) {
	path := fmt.Sprintf("cards/%s/members", c.ID)
	err = c.client.Get(path, args, &members)
	return
}
