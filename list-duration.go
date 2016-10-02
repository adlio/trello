package trello

import (
	"sort"
	"time"

	"github.com/pkg/errors"
)

type ListDuration struct {
	ListID       string
	ListName     string
	Duration     time.Duration
	FirstEntered time.Time
	TimesInList  int
}

func (l *ListDuration) AddDuration(d time.Duration) {
	l.Duration = l.Duration + d
	l.TimesInList++
}

func (c *Card) GetListDurations() (durations []*ListDuration, err error) {

	// Get all actions which affected the Card's List
	actions, err := c.GetListChangeActions()
	if err != nil {
		err = errors.Wrap(err, "GetListChangeActions() call failed.")
		return
	}
	sort.Sort(ByActionDate(actions))

	var prevTime time.Time
	var prevList *List

	durs := make(map[string]*ListDuration)
	for _, action := range actions {
		if prevList != nil {
			duration := action.Date.Sub(prevTime)
			_, durExists := durs[prevList.ID]
			if !durExists {
				durs[prevList.ID] = &ListDuration{ListID: prevList.ID, ListName: prevList.Name, Duration: duration, TimesInList: 1, FirstEntered: prevTime}
			} else {
				durs[prevList.ID].AddDuration(duration)
			}
		}
		prevList = ListAfterAction(action)
		prevTime = action.Date
	}

	if prevList != nil {
		duration := time.Now().Sub(prevTime)
		_, durExists := durs[prevList.ID]
		if !durExists {
			durs[prevList.ID] = &ListDuration{ListID: prevList.ID, ListName: prevList.Name, Duration: duration, TimesInList: 1, FirstEntered: prevTime}
		} else {
			durs[prevList.ID].AddDuration(duration)
		}
	}

	durations = make([]*ListDuration, 0, len(durs))
	for _, ld := range durs {
		durations = append(durations, ld)
	}
	sort.Sort(ByFirstEntered(durations))

	return durations, nil
}
