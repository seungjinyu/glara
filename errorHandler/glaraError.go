package errorHandler

import "log"

// configError checks whether the configuration error is caused
func PrintError(err error) {
	if err != nil {
		log.Println(err)
	}
}
