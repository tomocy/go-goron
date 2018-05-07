package session

type store struct {
	dat   map[string]interface{}
	token string
}
