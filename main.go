package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

var (
	serverId = createId(10)
)

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":9000", nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ServerId:", serverId)
	fmt.Fprintln(w, "Timestamp:", time.Now().Format(time.RFC3339))
}

func createId(length int) string {
	buf := make([]byte, length)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}
