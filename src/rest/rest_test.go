package rest

import (
	"testing"
	"github.com/mux"
	"log"
	"generalcargo"
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"math/rand"
	"time"
)

var cargo = generalcargo.GeneralCargo{
	CargoID: 3,
	VoyageNumber: 8,
}

var channel int

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func startServer() {
	router := mux.NewRouter()
	ctx := GetContext()
	router.HandleFunc("/", appHandler{ctx, GetScheduleHandler}.ServerHTTP).Methods("GET")
	router.HandleFunc("/si", appHandler{ctx, CreateShippingInfoHandler}.ServerHTTP).Methods("POST")
	router.HandleFunc("/cargo", appHandler{ctx, AddCargoHandler}.ServerHTTP).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func startHttpServer(ch int) {
	router := mux.NewRouter()
	ctx := GetContext()
	router.HandleFunc("/", appHandler{ctx, GetScheduleHandler}.ServerHTTP).Methods("GET")
	router.HandleFunc("/cargo", appHandler{ctx, AddCargoHandler}.ServerHTTP).Methods("POST")

	addr := fmt.Sprintf(":%d", ch)
	srv := &http.Server{Addr: addr, Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	srv.Shutdown(nil)
}

func TestAddCargoHandler(t *testing.T) {
	channel = randInt(1000, 65000)
	startHttpServer(channel)
	fmt.Println("Test Add Cargo Handler")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(cargo)

	addr := fmt.Sprintf("http://localhost:%d/cargo", channel)
	_, err := http.Post(addr, "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal("Something gone wrong: " + err.Error())
	}

}

func TestGetScheduleHandler(t *testing.T) {
	channel = randInt(1000, 65000)
	startHttpServer(channel)

	addr := fmt.Sprintf("http://localhost:%d/", channel)

	_, err := http.Get(addr)

	if err != nil {
		log.Fatal("Wrong request: " + err.Error())
	}
}

