package app

import (
	"fmt"
	"time"
)

// App function
func App() {

	for i := 0; i < 10; i++{
		// Print text
		fmt.Println(i, " App called successfully!")
	
		// Sleep duration
		time.Sleep(time.Second)
	}
}