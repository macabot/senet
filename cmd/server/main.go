package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	address := flag.String("a", ":8001", "Address to serve on.")
	directory := flag.String("d", ".", "The directory of static files to host.")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on %s\n", *directory, *address)
	log.Fatal(http.ListenAndServe(*address, &Handler{http.Dir(*directory)}))
}

type Handler struct {
	Dir http.Dir
}

func (a *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	http.FileServer(HTMLDir{a.Dir}).ServeHTTP(w, r)
}

type HTMLDir struct {
	Dir http.Dir
}

func (d HTMLDir) Open(name string) (http.File, error) {
	// Try name as supplied
	f, err := d.Dir.Open(name)
	if os.IsNotExist(err) {
		// If not found, try with .html
		if f, err := d.Dir.Open(name + ".html"); err == nil {
			return f, nil
		}
	}
	return f, err
}
