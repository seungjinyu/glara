package errorHandler

import "log"

// configError checks whether the configuration error is caused
func ConfigError(err error) {
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("The configuration file is now successfully imported")
	}
}
