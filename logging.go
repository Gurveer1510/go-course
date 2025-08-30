package main

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message")

	log.SetPrefix("INFO:")
	log.Println("This is an info message")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("This is a log message with date")
	infoLogger.Println("This has a prefix")

}

var infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
