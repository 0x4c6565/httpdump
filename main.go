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
	fmt.Fprint(w, string(b))
}

func main() {
	pflag.Int("port", 8080, "port to listen on")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.BindEnv("port")
	port := viper.GetInt("port")

	http.HandleFunc("/", handler)

	log.Printf("Listening on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
