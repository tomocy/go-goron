package session

type Session interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	// Delete(key string) error
	// ID() string
}

type session struct {
}

func New() Session {
	return &session{}
}

var sessions map[string]interface{}

func (s *session) Set(key string, value interface{}) {
	sessions[key] = value
}

func (s *session) Get(key string) interface{} {
	if v, ok := sessions[key]; ok {
		return v
	}

	return nil
}

func init() {
	sessions = make(map[string]interface{})
}

// type session struct {
// 	id        Id
// 	dat       map[string]interface{}
// 	expiresAt time.Time
// }

// type Id string

// const expiresIn time.Duration = (3 * time.Minute)

// func (s *session) Set(key string, value interface{}) error {
// 	return nil
// }

// func (s *session) Get(key string) (interface{}, error) {
// 	return nil, nil
// }

// func (s *session) Delete(key string) error {
// 	return nil
// }

// func (s *session) ID() Id {
// 	return s.id
// }
