package rest

import (
	"testing"
	"github.com/mux"
	"log"
	"net/http"
	"fmt"
	"net/url"
)

func TestCreateShippingInfoHandler(t *testing.T) {
	router := mux.NewRouter()
	ctx := GetContext()
	fmt.Println("1")
	router.HandleFunc("/si", appHandler{ctx, CreateShippingInfoHandler}.ServerHTTP).Methods("POST")
	fmt.Println("1")
	router.HandleFunc("/cargo", appHandler{ctx, AddCargoHandler}.ServerHTTP).Methods("POST")
	fmt.Println("1")

	form := url.Values{}
	form.Add("voyagenumber", "8")
	form.Add("cargoid", "3")


	_, err := http.PostForm("https://localhost:8080/cargo", form)
	if err != nil {
		fmt.Println(err.Error())
	}
	//defer resp.Body.Close()

	log.Fatal(http.ListenAndServe(":8080", router))

}

