package tlog

import "fmt"

func GetWantedHad(msg string, wanted, had interface{}) string {
	return fmt.Sprintf("%s\n\nwanted: %#v\n   had: %#v\n", msg, wanted, had)
}
