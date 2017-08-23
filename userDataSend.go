package serverOps

import (
	"fmt"
	"net/http"
)

func main(w http.ResponseWriter, r *http.Request) {
	i, e := w.Write([]byte("True"))
	fmt.Println(i, " ", e)
}
