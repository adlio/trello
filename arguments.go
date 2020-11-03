// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"net/url"
)

// Arguments are used for passing URL parameters to the client for making API calls.
type Arguments map[string]string

// Defaults is a constructor for default Arguments.
func Defaults() Arguments {
	return make(Arguments)
}

// ToURLValues returns the argument's URL value representation.
func (args Arguments) ToURLValues() url.Values {
	v := url.Values{}
	for key, value := range args {
		v.Set(key, value)
	}
	return v
}

// flattenArguments will return a Arguments by merging a slice of Arguments,
// where each successive slice can override fields in the previous.
func flattenArguments(extraArgs []Arguments) (args Arguments) {
	args = make(Arguments)
	args.flatten(extraArgs)
	return
}

// flatten will merge a slice of Arguments into the current one
// where each successive slice can override fields in the previous.
func (args Arguments) flatten(extraArgs []Arguments) {
	for _, extraArg := range extraArgs {
		for key, val := range extraArg {
			args[key] = val
		}
	}
}
