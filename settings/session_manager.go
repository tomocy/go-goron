package settings

type sessionManager struct {
	probOfDelete        int
	probOfDeleteDivisor int
}

var SessionManager *sessionManager

func (sm *sessionManager) ProbOfDelete() int {
	return sm.probOfDelete
}

func (sm *sessionManager) ProbOfDeleteDivisor() int {
	return sm.probOfDeleteDivisor
}

func (sm *sessionManager) SetProbOfDelete(pod int) {
	sm.probOfDelete = pod
}

func (sm *sessionManager) SetProbOfDeleteDivisor(podDivisor int) {
	sm.probOfDeleteDivisor = podDivisor
}

func init() {
	SessionManager = &sessionManager{
		probOfDelete:        1,
		probOfDeleteDivisor: 100,
	}
}
