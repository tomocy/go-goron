package session

type cmdType int

const (
	unknown cmdType = iota
	create
	load
	save
	destroy
	destroyExpired
)

type command struct {
	cmdType cmdType
	req     []interface{}
	respCh  chan response
}
