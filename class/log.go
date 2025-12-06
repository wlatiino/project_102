package SO_Class

import "log"

// Define a struct with a Print method
type lLog struct{}

func (saya lLog) Println(flag bool, v ...any) {
	if flag {
		log.Println(v...)
	}
}

func (saya lLog) Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}

// Exported instance
var Log lLog
