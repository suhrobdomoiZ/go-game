package structures

type Player struct {
	Room        string
	HasBackpack bool
	Inventory   []string
}

type Room struct {
	Name        string
	Description string
	Items       map[string][]string
	Ways        []string
	Message     string
}
