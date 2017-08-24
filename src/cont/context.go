package cont

import (
	"cargo"
	"generalcargo"
	"db"
	"shipinfo"
)

type AppContext struct {
	DB 	*db.DB
	//CFG 	*db.Config
	Repo 	*Repositories
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

