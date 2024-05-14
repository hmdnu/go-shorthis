package common

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
