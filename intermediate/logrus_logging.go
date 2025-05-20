package intermediate

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()

	// set log level
	log.SetLevel(logrus.InfoLevel)

	// set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// logging examples
	log.Info("This is a info message.")
	log.Warn("This is a warn message.")
	log.Error("This is a error message.")

	log.WithFields(logrus.Fields{
		"username":"Utku Tasguluk",
		"method":"GET",
	}).Info("User logged in.")
}