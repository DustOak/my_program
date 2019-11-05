package logger

import (
	"log"
	"os"
)

var InfoLog *log.Logger
var ErrLog *log.Logger

func init() {
	infoFile, err := os.OpenFile("INFOFILE", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	errFile, err := os.OpenFile("ERROR", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	InfoLog = log.New(infoFile, "INFORMATION:", log.LstdFlags|log.Lshortfile)
	ErrLog = log.New(errFile, "INFORMATION:", log.LstdFlags|log.Lshortfile)
}
