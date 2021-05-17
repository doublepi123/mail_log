package util

import "time"

func PauseForRun() {
	for {
		time.Sleep(time.Second)
	}
}
