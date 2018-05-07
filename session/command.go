package session

type cmd int

const (
	unknown cmd = iota
	create
	load
	save
	destroy
	destroyExpired
)

type command struct {
	cmd   cmd
	req   []interface{}
	resCh chan response
}
