package encodingOperations

import (
	"bytes"
	"encoding/json"
)

func encode(a []string) (*bytes.Buffer){
	bytesArray, err := json.Marshal(a)
	checkError(err)
	b := bytes.NewBuffer(bytesArray)
	return b
}

func checkError(e error){
	if e != nil {
		panic(e)
	}
}