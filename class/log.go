package SO_Class

import (
	"log"

	"github.com/gin-gonic/gin"
)

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

func (saya lLog) CetakKunci(flag bool, c *gin.Context, v ...any) {
	if flag {
		// log.Println(v...)

		KunciRute, exits := c.Get("KunciRute")
		if exits {
			for i, x := range v {
				s := Fmt.Sprint(x)
				s = Strings.NewReplacer("\r", " ", "\n", " ", "\t", " ").Replace(s)
				v[i] = s
			}
			args := append([]any{KunciRute}, v)
			log.Println(args...)
		} else {
			log.Println(v...)
		}

	}
}

// Exported instance
var Log lLog
