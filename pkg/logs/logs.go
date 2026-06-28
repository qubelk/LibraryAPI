package logs

import (
	"log"
	"os"
)

func Init() {
	os.MkdirAll("../logs", 0755)
}

func LogError(mes string) {
	f, err := os.Create("logs.log")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	l := log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Println(mes)
}
