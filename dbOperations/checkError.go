package dbOperations

func checkError(err error) { //checks for errors
	if err != nil {
		panic(err)
	}
}
