package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.WriteHeader(viper.GetInt("status-code"))
	fmt.Fprint(w, string(b))

	if viper.GetBool("log") {
		log.Print(string(b))
	}
}

func main() {
	pflag.Int("port", 8080, "port to listen on")
	pflag.Int("status-code", 200, "status code to return")
	pflag.Bool("log", false, "specifies to log requests to stdout")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()

	port := viper.GetInt("port")

	http.HandleFunc("/", handler)

	log.Printf("Listening on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
