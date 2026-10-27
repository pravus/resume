package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"resume/internal/html"
	"resume/internal/model"
	"resume/internal/yaml"
)

const (
	data     = `data/resume.yaml`
	template = `html/content.html`
)

type flags struct {
	http   *string
	prefix *string
}

func main() {
	flags := flags{
		http:   flag.String(`http`, `:8080`, `sets the http bind address`),
		prefix: flag.String(`prefix`, ``, `sets the template prefix`),
	}
	flag.Parse()

	// helpers
	_prefix := func(name string) string {
		if *flags.prefix != `` {
			return *flags.prefix + `/` + name
		}
		return name
	}

	// load
	var resume *model.Resume
	{
		loader := yaml.NewLoader(_prefix(data))
		if value, err := loader.Load(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		} else {
			resume = value
		}
	}

	// writer
	writer := html.NewWriter(_prefix(template))

	// server
	server := http.Server{
		Addr: *flags.http,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet && r.Method != http.MethodHead {
				w.Header().Set("Allow", "GET, HEAD")
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if r.Method == http.MethodHead {
				w.WriteHeader(http.StatusOK)
				return
			}
			if err := writer.Write(w, resume); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				// Headers may already be committed if the client disconnects during write.
			}
		}),
	}

	// main
	fmt.Println(`http.up ` + *flags.http)
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
