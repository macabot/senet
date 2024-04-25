package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	address := flag.String("a", ":8001", "Address to serve on.")
	directory := flag.String("d", ".", "The directory of static files to host.")
	flag.Parse()

	log.Printf("Serving %s on %s\n", *directory, *address)
	log.Fatal(http.ListenAndServe(*address, &Handler{*directory}))
}

type Handler struct {
	PublicDir string
}

var (
	brotliPattern = regexp.MustCompile(`\bbr\b`)
	gzipPattern   = regexp.MustCompile(`\bgzip\b`)
)

func decodeAcceptEncoding(header http.Header) (tryBrotli bool, tryGzip bool) {
	tryBrotli = false
	tryGzip = false
	if len(header["Accept-Encoding"]) > 0 {
		tryBrotli = brotliPattern.MatchString(header["Accept-Encoding"][0])
		tryGzip = gzipPattern.MatchString(header["Accept-Encoding"][0])
	}
	return
}

type PathEncoding struct {
	Path            string
	ContentEncoding string
}

// See https://docs.gitlab.com/ee/user/project/pages/introduction.html#resolving-ambiguous-urls
// It does not serve files without an extension. A file without an extension is assumed to correspond to an HTML file.
func (a *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	path := r.URL.Path
	if path == "" {
		path = "/"
	}

	if path[len(path)-1:] == "/" {
		path += "index.html"
	}

	if fileInfo, err := os.Stat(path); err == nil && fileInfo.IsDir() {
		http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
		return
	}

	if filepath.Ext(path) == "" {
		path += ".html"
	}

	path = filepath.Join(a.PublicDir, path)

	tryBrotli, tryGzip := decodeAcceptEncoding(r.Header)

	var pathEncodings []PathEncoding
	if tryBrotli {
		pathEncodings = append(pathEncodings, PathEncoding{
			Path:            path + ".br",
			ContentEncoding: "br",
		})
	}
	if tryGzip {
		pathEncodings = append(pathEncodings, PathEncoding{
			Path:            path + ".gz",
			ContentEncoding: "gzip",
		})
	}
	pathEncodings = append(pathEncodings, PathEncoding{
		Path:            path,
		ContentEncoding: "",
	})

	var pathEncoding PathEncoding
	var fileInfo fs.FileInfo
	var err error
	for _, pathEncoding = range pathEncodings {
		if fileInfo, err = os.Stat(pathEncoding.Path); !errors.Is(err, fs.ErrNotExist) && !errors.Is(err, fs.ErrPermission) {
			break
		}
	}

	if errors.Is(err, fs.ErrNotExist) {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	} else if errors.Is(err, fs.ErrPermission) {
		http.Error(w, "no permission to read file", http.StatusForbidden)
		return
	} else if err != nil {
		err := fmt.Errorf("failed to get file into: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	f, err := os.Open(pathEncoding.Path)
	if err != nil {
		err := fmt.Errorf("failed to open file: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer f.Close()

	contentType := mime.TypeByExtension(filepath.Ext(path))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	if pathEncoding.ContentEncoding != "" {
		w.Header().Set("Content-Encoding", pathEncoding.ContentEncoding)
	}

	http.ServeContent(w, r, pathEncoding.Path, fileInfo.ModTime(), f)
}
