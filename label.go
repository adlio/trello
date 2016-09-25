package trello

type Label struct {
	ID      string `json:"id"`
	IDBoard string `json:"id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Uses    int    `json:"uses"`
}
