package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func dump(w http.ResponseWriter, req *http.Request) {
	log.Println("request received")
	_, _ = fmt.Fprintf(w, "HTTP Request\n")
	_, _ = fmt.Fprintf(w, "URL: %s\n", req.URL.String())
	_, _ = fmt.Fprintf(w, "-- Headers --\n")
	for name, headers := range req.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "\t%v: %v\n", name, h)
		}
	}
	_, _ = fmt.Fprintf(w, "-- Body --\n")
	d, _ := ioutil.ReadAll(req.Body)
	_, _ = fmt.Fprintf(w, "%s\n", string(d))
}

func main() {
	var (
		port = flag.Int("p", 8090, "Port to host the webapp on")
	)
	flag.Parse()

	http.HandleFunc("/", dump)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil)
	if err != nil {
		panic(err)
	}
}
