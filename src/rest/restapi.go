package rest

import (
	"net/http"
	"cont"
	"github.com/golang/go/src/pkg/log"
	"encoding/json"
	"github.com/mux"
	"db"
	"shipinfo"
	"generalcargo"
	"github.com/golang/go/src/pkg/strconv"
	"fmt"
)

type appHandler struct {
	*cont.AppContext
	H func(*cont.AppContext, http.ResponseWriter, *http.Request) (int, error)
}

func (ah appHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := ah.H(ah.AppContext, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

func GetScheduleHandler(a *cont.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	schedule, err := a.Repo.ShipInfoRepository.RetrieveAllSchedule()
	if err != nil {
		log.Println(err.Error())
	}
	json.NewEncoder(w).Encode(schedule)
	return 200, nil
}

func CreateShippingInfoHandler(a *cont.AppContext, wr http.ResponseWriter, req *http.Request) (int, error) {
	fmt.Println("go")
	si := shipinfo.ShippingInfo{}
	decoder := json.NewDecoder(req.Body)
	fmt.Println(decoder)
	err := decoder.Decode(&si)
	if err != nil {
		log.Println(err.Error())
		return 400, err
	}
	fmt.Println(si)

	err = a.Repo.ShipInfoRepository.CreateShippingInfo(&si)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Success")
	return 200, nil
}

func AddCargoHandler (a *cont.AppContext, wr http.ResponseWriter, req *http.Request) (int, error){
	gc := generalcargo.GeneralCargo{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&gc)

	if err != nil {
		log.Println(err.Error())
		return 400, err
	}

	err, _ = checkCargoByVN(&gc)

	if err != nil {
		log.Printf("The cargo #%d was already created for voyage #%d", gc.CargoID, gc.VoyageNumber)
		return 400, err
	}

	err = a.Repo.GeneralCargoRepository.AddCargoByVN(&gc)

	if err != nil {
		log.Println(err.Error())
		return 400, err
	}
	log.Println("Success")

	return 200, nil
}

func DeleteShippingInfo(a *cont.AppContext, wr http.ResponseWriter, req *http.Request) (int, error) {
	params := mux.Vars(req)
	number, _ := strconv.Atoi(params["vn"])
	err := shipinfo.ShipScheduleRepository(db.GetDBConnection()).DeleteShippingInfo(number)
	if err != nil {
		log.Println(err.Error())
		return 400, err
	}
	log.Println("Success")
	return 200, nil
}

func checkCargoByVN(gc *generalcargo.GeneralCargo) (err error, check bool) {
	rows, err := db.GetDBConnection().Query("select cargo_id from public.general_cargo where voyage_number = $1", gc.VoyageNumber)
	if err != nil {
		log.Println(err.Error())
		return err, check
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


//func StartRestApi() {
//	router := mux.NewRouter()
//	config := &db.Config{}
//	config.Default()
//	context := &cont.AppContext{
//		DB:    &db.DB{},
//		Repo:  &cont.Repositories{},
//	}
//	context.DB.Initialize(config)
//	context.Repo.Initialize(context.DB)
//
//	router.HandleFunc("/schedule", appHandler{context, GetScheduleHandler}.ServerHTTP).Methods("GET")
//	router.HandleFunc("/si", appHandler{context, CreateShippingInfoHandler}.ServerHTTP).Methods("POST")
//	router.HandleFunc("/cargo", appHandler{context, AddCargoHandler}.ServerHTTP).Methods("POST")
//	router.HandleFunc("/si/{vn}", appHandler{context, DeleteShippingInfo}.ServerHTTP).Methods("DELETE")
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//
//}

