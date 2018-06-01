package settings

type sessionManager struct {
	ProbOfDelete        int
	ProbOfDeleteDivisor int
}

var SessionManager *sessionManager

func (sm *sessionManager) SetProbOfDelete(pod int) {
	sm.ProbOfDelete = pod
}

func (sm *sessionManager) SetProbOfDeleteDivisor(podDivisor int) {
	sm.ProbOfDeleteDivisor = podDivisor
}

func init() {
	SessionManager = &sessionManager{
		ProbOfDelete:        1,
		ProbOfDeleteDivisor: 100,
	}
}
