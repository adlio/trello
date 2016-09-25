package trello

import (
	"net/url"
)

type Arguments map[string]string

var Defaults = make(Arguments)

func (args Arguments) ToURLValues() url.Values {
	v := url.Values{}
	for key, value := range args {
		v.Set(key, value)
	}
	return v
}
