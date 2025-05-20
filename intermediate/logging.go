package intermediate

import (
	"log"
	"os"
)

func main() {

	log.Println("This is log message.")

	// setting prefix before log message
	log.SetPrefix("INFO: ")
	log.Println("This is an info message.")

	// Log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("This is a log message with only date, time.")

	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an error message.")

	// you can write debug message to the file in this example showed.
	file, err := os.OpenFile("app.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	infoLogger1  := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger1  := log.New(file, "WARN: ", log.Ldate|log.Ltime)
	errorLogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger  := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)


	debugLogger.Println("this is a debug message.")
	warnLogger1.Println("this is a warning message.")
	infoLogger1.Println("this is a info message.")
	errorLogger1.Println("this is a error message.")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
