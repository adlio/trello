// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// IDToTime is a convenience function. It takes a Trello ID string and
// extracts the encoded create time as time.Time or an error.
func IDToTime(id string) (t time.Time, err error) {
	if id == "" {
		return time.Time{}, nil
	}
	// The first 8 characters in the object ID are a Unix timestamp
	ts, err := strconv.ParseUint(id[:8], 16, 64)
	if err != nil {
		err = errors.Wrapf(err, "ID '%s' failed to convert to timestamp.", id)
	} else {
		t = time.Unix(int64(ts), 0)
	}
	return
}
