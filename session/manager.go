package session

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type Manager interface {
	Start()
	Stop()
	CreateSession() (Session, error)
	ReadSession(sid string) (Session, error)
	SaveSession(sid string, dat interface{})
	DeleteSession(sid string) error
	DeleteExpiredSession() error

	serve()
}

type manager struct {
	cmdCh    chan command
	stopCh   chan bool
	stopGCCh chan bool
}

func (m *manager) Start() {
	go m.serve()
}

func (m *manager) Stop() {
	m.stopCh <- true
	m.stopGCCh <- true
}

func (m *manager) CreateSession() (Session, error) {
	respCh := make(chan response, 1)
	defer close(respCh)

	cmd := command{create, nil, respCh}
	m.cmdCh <- cmd

	// wait for response
	resp := <-respCh
	if resp.err != nil {
		log.Fatal("Error while creating a session")
		return nil, resp.err
	}

	return resp.session, nil
}

func (m *manager) serve() {
	for {
		select {
		case cmd := <-m.cmdCh:
			switch cmd.cmdType {
			case create:
				session := &session{
					id:        uuid.New().String(),
					dat:       make(map[string]interface{}),
					expiresAt: time.Now().Add(expiresIn),
				}

				cmd.respCh <- response{session: session, err: nil}
			}
		}
	}
}
