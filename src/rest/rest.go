package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"repository/shipinfo"
	"db"
	"repository/generalcargo"
	"log"
	"github.com/golang/go/src/pkg/strconv"
)

func GetScheduleHandler(wr http.ResponseWriter, req *http.Request) {
	var schedule, err = shipinfo.ShipScheduleRepository(db.GetDBConnection()).RetrieveAllSchedule()
	if err != nil {
		log.Println(err.Error())

	}
	json.NewEncoder(wr).Encode(schedule)

}

func CreateShippingInfoHandler(wr http.ResponseWriter, req *http.Request) {
	si := shipinfo.ShippingInfo{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&si)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = shipinfo.ShipScheduleRepository(db.GetDBConnection()).CreateShippingInfo(&si)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Success")
}

func AddCargoByVN(wr http.ResponseWriter, req *http.Request) {
	gc := generalcargo.GeneralCargo{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&gc)

	if err != nil {
		log.Println(err.Error())
		return
	}

	err, _ = checkCargoByVN(gc)

	if err != nil {
		log.Printf("The cargo #%d in voyage #%d", gc.CargoID, gc.VoyageNumber)
		return
	}

	err = generalcargo.GeneralCargoRepositpory(db.GetDBConnection()).AddCargoByVN(&gc)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Success")
}

func DeleteShippingInfo(wr http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	number, _ := strconv.Atoi(params["vn"])
	err := shipinfo.ShipScheduleRepository(db.GetDBConnection()).DeleteShippingInfo(number)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Success")
}

func checkCargoByVN(gc *generalcargo.GeneralCargo) (err, check bool) {
	rows, err := db.GetDBConnection().Query("select cargo_id from public.general_cargo where voyage_number = $1", gc.VoyageNumber)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cargo int
		err := rows.Scan(&cargo)
		if err != nil {
			return err, check
		}
		if cargo == gc.CargoID {
			check = true
			return nil, check
		}
	}

	err = rows.Err()
	if err != nil {
		return err, check
	}

	return nil, check

}

func StartRestService() {
	router := mux.NewRouter()
	router.HandleFunc("/schedule", GetScheduleHandler).Methods("GET")
	router.HandleFunc("/si", CreateShippingInfoHandler).Methods("POST")
	router.HandleFunc("/gc", AddCargoByVN).Methods("POST")
	router.HandleFunc("/si/{vn}", DeleteShippingInfo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
