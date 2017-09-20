package rest

import (
	"cont"
	"db"
	"time"
	"math/rand"
	"github.com/mux"
	"fmt"
	"net/http"
	"log"
)

func getContext() *cont.AppContext {
	config := &db.Config{}
	config.Default()

	context := &cont.AppContext{
		DB:    &db.DB{},
		Repo:  &cont.Repositories{},
	}

	context.DB.Initialize(config)
	context.Repo.Initialize(context.DB)
	return context
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func startHttpServer(ch int) {
	router := mux.NewRouter()
	ctx := getContext()
	router.HandleFunc("/", appHandler{ctx, GetScheduleHandler}.ServerHTTP).Methods("GET")
	router.HandleFunc("/cargo", appHandler{ctx, AddCargoHandler}.ServerHTTP).Methods("POST")
	router.HandleFunc("/si", appHandler{ctx, CreateShippingInfoHandler}.ServerHTTP).Methods("POST")

	host := fmt.Sprintf(":%d", ch)
	srv := &http.Server{Addr: host, Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	srv.Shutdown(nil)
}