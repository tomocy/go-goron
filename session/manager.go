package session

// type Manager interface {
// 	Start() Session

// 	serve()
// }

// type manager struct {
// 	cmdCh chan command
// }

// func (m *manager) Start() Session {
// 	go m.serve()
// }

// func (m *manager) serve() {
// 	sessions := make(map[string]interface{})

// 	for {
// 		select {
// 		case cmd := <-m.cmdCh:
// 			switch cmd.t {
// 			case create:
// 				sessionID := "count"
// 			}
// 		}
// 	}
// }

// type Manager interface {
// 	Start()
// 	Stop()
// 	CreateSession() (Id, error)
// 	ReadSession(sid Id) (Session, error)
// 	SaveSession(sid Id, dat map[string]interface{}) error
// 	DeleteSession(sid Id) error
// 	DeleteExpiredSession() error

// 	serve()
// }

// type manager struct {
// 	cmdCh    chan command
// 	stopCh   chan bool
// 	stopGCCh chan bool
// }

// func NewManager() Manager {
// 	return &manager{}
// }

// func (m *manager) Start() {
// 	go m.serve()
// }

// func (m *manager) Stop() {
// 	m.stopCh <- true
// 	m.stopGCCh <- true
// }

// func (m *manager) CreateSession() (Id, error) {
// 	respCh := make(chan response, 1)
// 	defer close(respCh)

// 	cmd := command{create, nil, respCh}
// 	m.cmdCh <- cmd

// 	// wait for response
// 	resp := <-respCh
// 	var id Id
// 	if resp.err != nil {
// 		log.Fatal("Error while creating a session")
// 		return id, resp.err
// 	}

// 	return resp.res[0].(Id), nil
// }

// func (m *manager) ReadSession(id Id) (Session, error) {
// 	respCh := make(chan response, 1)
// 	defer close(respCh)

// 	req := []interface{}{id}
// 	cmd := command{read, req, respCh}
// 	m.cmdCh <- cmd

// 	res := <-respCh
// 	if res.err != nil {
// 		return nil, res.err
// 	}

// 	return res.res[0].(Session), nil
// }
// func (m *manager) SaveSession(sid Id, dat map[string]interface{}) error {
// 	respCh := make(chan response, 1)
// 	defer close(respCh)

// 	req := []interface{}{sid, dat}
// 	cmd := command{save, req, respCh}
// 	m.cmdCh <- cmd

// 	resp := <-respCh
// 	if resp.err != nil {
// 		return resp.err
// 	}

// 	return nil
// }
// func (m *manager) DeleteSession(sid Id) error {
// 	return nil
// }
// func (m *manager) DeleteExpiredSession() error {
// 	return nil
// }

// func (m *manager) serve() {
// 	sessions := make(map[Id]*session)
// 	for {
// 		select {
// 		case cmd := <-m.cmdCh:
// 			switch cmd.cmdType {
// 			case create:
// 				id := Id(uuid.New().String())
// 				sessions[id] = &session{
// 					id:        id,
// 					dat:       make(map[string]interface{}),
// 					expiresAt: time.Now().Add(expiresIn),
// 				}

// 				cmd.respCh <- response{res: []interface{}{sessions[id]}, err: nil}
// 			case read:
// 				sessionId := cmd.req[0].(Id)
// 				session := sessions[sessionId]

// 				if time.Now().After(session.expiresAt) {
// 					cmd.respCh <- response{nil, errors.New("Not found")}
// 					break
// 				}

// 				session.expiresAt = time.Now().Add(expiresIn)

// 				cmd.respCh <- response{[]interface{}{session}, nil}
// 			}
// 		}
// 	}
// }
