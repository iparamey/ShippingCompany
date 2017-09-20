package cont

import (
	"cargo"
	"generalcargo"
	"db"
	"shipinfo"
	"github.com/mux"
)

type AppContext struct {
	DB 	*db.DB
	Repo 	*Repositories
	Router  *mux.Router
}

type Repositories struct {
	CargoRepository 	cargo.Repository
	GeneralCargoRepository  generalcargo.Repository
	ShipInfoRepository 	shipinfo.Repository
}

func (r *Repositories) Initialize(db *db.DB) error {
	r.CargoRepository = cargo.CargoRepository(db.GetConnection())
	r.GeneralCargoRepository = generalcargo.GeneralCargoRepositpory(db.GetConnection())
	r.ShipInfoRepository = shipinfo.ShipScheduleRepository(db.GetConnection())
	return nil
}

