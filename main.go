//go:generate statik -src "reaper"

package main

import (
	"flag"
	"fmt"
	_ "github.com/karnauskas/reaperd/statik"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

var (
	flagBind  = flag.String("bind", ":8080", "listen address")
	flagGrimd = flag.String("grimd", "http://10.0.0.2:8080", "grimd api")
)

func main() {
	flag.Parse()

	statikFs, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(statikFs))
	http.HandleFunc("/js/config.js", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("var apiURL = '%s/'\n", *flagGrimd))
	})
	http.ListenAndServe(*flagBind, nil)
}
