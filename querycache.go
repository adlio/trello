// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

// queryCache struct to hold maps
// Type map[IDType]*Type
type queryCache struct {
	Boards map[string]*Board
	Lists  map[string]*List
}

// GetBoard returns *List if found, else nil
func (qc *queryCache) GetBoard(id string) (board *Board, found bool) {

	board, found = qc.Boards[id]

	return
}

// SetBoard adds board by id
func (qc *queryCache) SetBoard(id string, board *Board) {
	qc.Boards[id] = board
	return
}

// RemoveBoard adds board by id
func (qc *queryCache) RemoveBoard(id string) {
	delete(qc.Boards, id)
	return
}

// GetList returns *List if found, else nil
func (qc *queryCache) GetList(id string) (list *List, found bool) {
	list, found = qc.Lists[id]
	return
}

// SetList adds board by id
func (qc *queryCache) SetList(id string, list *List) {
	qc.Lists[id] = list
	return
}

// RemoveList adds board by id
func (qc *queryCache) RemoveList(id string) {
	delete(qc.Lists, id)
	return
}
