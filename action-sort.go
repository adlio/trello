package trello

type ByActionDate []*Action

func (actions ByActionDate) Len() int {
	return len(actions)
}

func (actions ByActionDate) Less(i, j int) bool {
	return actions[i].Date.Before(actions[j].Date)
}

func (actions ByActionDate) Swap(i, j int) {
	actions[i], actions[j] = actions[j], actions[i]
}
