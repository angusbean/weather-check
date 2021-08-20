package helpers

import (
	"log"
	"time"
)

//timeTrack returns the time since first calling the application
func TimeTrack(startTime time.Time, name string) {
	elapsed := time.Since(startTime)
	log.Printf("%s took %s", name, elapsed)
}
