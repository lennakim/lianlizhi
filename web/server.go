package main

import (
	"fmt"
	"net/http"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")

	if err := InitConfig("config.yml"); err != nil {
		panic(err)
	}

	runmode := viper.GetString("runmode")
	fmt.Printf("%s", runmode)

	http.ListenAndServe("0.0.0.0:2930", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}
