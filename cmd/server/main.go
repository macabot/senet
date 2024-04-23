package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"regexp"
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

var (
	gzipPattern   = regexp.MustCompile(`\bgzip\b`)
	brotliPattern = regexp.MustCompile(`\bbrotli\b`)
)

func (a *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	tryBrotli := false
	tryGzip := false
	if len(r.Header["Accept-Encoding"]) > 0 {
		tryBrotli = brotliPattern.MatchString(r.Header["Accept-Encoding"][0])
		tryGzip = gzipPattern.MatchString(r.Header["Accept-Encoding"][0])
	}
	htmlDir := HTMLDir{
		Dir:       a.Dir,
		TryBrotli: tryBrotli,
		TryGzip:   tryGzip,
	}
	http.FileServer(htmlDir).ServeHTTP(w, r)
}

type HTMLDir struct {
	Dir       http.Dir
	TryBrotli bool
	TryGzip   bool
}

func (d HTMLDir) Open(name string) (http.File, error) {
	var tryNames []string
	if d.TryBrotli {
		tryNames = append(tryNames, name+".br")
	}
	if d.TryGzip {
		tryNames = append(tryNames, name+".gz")
	}
	tryNames = append(tryNames, name)

	if d.TryBrotli {
		tryNames = append(tryNames, name+".html.br")
	}
	if d.TryGzip {
		tryNames = append(tryNames, name+".html.gz")
	}
	tryNames = append(tryNames, name+".html")

	var f http.File
	var err error
	for _, tryName := range tryNames {
		f, err = d.Dir.Open(tryName)
		if err == nil || !os.IsNotExist(err) {
			break
		}
	}
	return f, err
}
