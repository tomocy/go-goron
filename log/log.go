package log

import (
	"fmt"
	"os"
)

func Debug(str string) {
	name := "storage/logs/today.log"
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, str)
}
