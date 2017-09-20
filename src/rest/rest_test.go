package rest

import (
	"testing"
	"log"
	"generalcargo"
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"shipinfo"
	"time"
)

var cargo = generalcargo.GeneralCargo{
	CargoID: 3,
	VoyageNumber: 8,
}

var si = shipinfo.ShippingInfo{
	VoyageNumber: 9,
	StartingPoint: 4,
	FinalDestination: 6,
	StartDate: time.Now().Add(time.Hour),
	EndDate: time.Now().Add(120*time.Hour),
	Ship: 5,

}

var channel int

func TestAddCargoHandler(t *testing.T) {
	channel = randInt(1000, 65000)
	startHttpServer(channel)
	fmt.Println("Test Add Cargo Handler")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(cargo)

	url := fmt.Sprintf("http://localhost:%d/cargo", channel)
	_, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal("Something gone wrong: " + err.Error())
	}

}

func TestCreateShippingInfoHandler(t *testing.T) {
	channel = randInt(1000, 65000)
	startHttpServer(channel)

	fmt.Println("Test Create Shipping Info Handler")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(si)

	url := fmt.Sprintf("http://localhost:%d/si", channel)
	_, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal("Something gone wrong: " + err.Error())
	}

}

func TestGetScheduleHandler(t *testing.T) {
	channel = randInt(1000, 65000)
	startHttpServer(channel)

	url := fmt.Sprintf("http://localhost:%d/", channel)

	_, err := http.Get(url)

	if err != nil {
		log.Fatal("Wrong request: " + err.Error())
	}
}

