package rest

import (
	"testing"
	"time"
	"github.com/mux"
	"github.com/golang/go/src/pkg/log"
	"net/http"
	"net/url"
	"github.com/golang/go/src/pkg/bytes"
	"fmt"
)

func TestCreateShippingInfoHandler(t *testing.T) {
	router := mux.NewRouter()
	ctx := GetContext()
	router.HandleFunc("/si", appHandler{ctx, CreateShippingInfoHandler}.ServerHTTP).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

	form := url.Values{}

	form.Add("voyagenumber", "5")
	form.Add("startingpoint", "5")
	form.Add("endpoint", "3")
	form.Add("startdate", time.Now().String())
	form.Add("startdate", time.Now().Add(time.Hour).String())
	form.Add("ship", "2")

	req, err := http.NewRequest("POST", "localhost:8080/si", bytes.NewBufferString(form.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err != nil {
		t.Errorf("New shipping info wasn't created: %s", err)
	}

	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Print(resp.Status)


}

