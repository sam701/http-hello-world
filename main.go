package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	serverId = createId(10)
	launchId = flag.String("launchId", "none", "Launch ID")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", serve)
	log.Println("Starting service on port 9000")
	http.ListenAndServe(":9000", nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ServerId:", serverId)
	fmt.Fprintln(w, "Timestamp:", time.Now().Format(time.RFC3339))
	fmt.Fprintln(w, "LaunchID:", *launchId)
}

func createId(length int) string {
	buf := make([]byte, length)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}
