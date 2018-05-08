package session

type Session interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	// Delete(key string) error
	ID() string
}

type session struct {
	id  string
	dat map[string]interface{}
}

func New(id string, dat map[string]interface{}) Session {
	return &session{id: id, dat: dat}
}

func (s *session) Set(key string, value interface{}) {
	s.dat[key] = value
}

func (s *session) Get(key string) interface{} {
	v, ok := s.dat[key]
	if !ok {
		return nil
	}

	return v
}

func (s *session) ID() string {
	return s.id
}
