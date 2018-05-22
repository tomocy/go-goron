package settings

type sessionManager struct {
	ProbOfDelete        int
	ProbOfDeleteDivisor int
}

var SessionManager *sessionManager
