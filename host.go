package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type logger struct {
	h http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	l.h.ServeHTTP(w, r)
}

func main() {
	port := flag.Int("p", 8001, "HTTP port number")
	flag.Parse()

	args := []string{
		"daemon",
		"--base-path=./src",
		"--listen=localhost.localdomain",
		"--reuseaddr",
		"--export-all",
	}
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go func() {
		log.Println("git", strings.Join(args, " "))
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fs := http.FileServer(http.Dir("www"))
	http.Handle("/", &logger{fs})
	addr := "localhost.localdomain:" + strconv.Itoa(*port)
	log.Println("listening at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
