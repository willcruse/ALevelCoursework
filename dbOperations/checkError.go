package dbOperations

import (
	"log"
)

func checkError(err error) { //checks for errors
	if err != nil {
		log.Println("Error was thrown: ", err)
	}
}
