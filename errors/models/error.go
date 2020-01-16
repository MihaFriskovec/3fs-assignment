package customerrors

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
