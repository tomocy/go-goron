package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/tomocy/goron/session"
)

type file struct {
	path string
	mu   sync.Mutex
}

func New() *file {
	return &file{path: "storage/sessions"}
}

func (f *file) InitSession(sessionID string) session.Session {
	f.mu.Lock()
	defer f.mu.Unlock()

	if _, err := os.Create(f.path + "/" + sessionID); err != nil {
		panic(err)
	}

	dat := make(map[string]string)
	return session.New(sessionID, dat)
}

func (f *file) GetSession(sessionID string) (session.Session, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, err := os.Open(f.path + "/" + sessionID)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dat := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), ":")
		if len(ss) < 2 {
			dat[ss[0]] = ""
		} else {
			dat[ss[0]] = ss[1]
		}
	}

	return session.New(sessionID, dat), nil
}

func (f *file) SetSession(session session.Session) {
	file, err := os.OpenFile(f.path+"/"+session.ID(), os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for k, v := range session.Data() {
		file.WriteString(fmt.Sprintf("%s:%s", k, v))
	}
}
