package file

func getSessionFilePath(sessionID string) string {
	b := make([]byte, 0, 10)
	b = append(b, "storage/sessions/"...)
	b = append(b, sessionID...)

	return string(b)
}
