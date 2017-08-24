package generalcargo

import (
	"database/sql"
	"db"
)

type Repository interface {
	AddCargoByVN(gc *GeneralCargo) error
	DeleteCargo(gc *GeneralCargo) error
}

func GeneralCargoRepositpory(database *sql.DB) Repository {
	return &generalCargoRepo{
		db: database,
	}
}

type generalCargoRepo struct {
	db *sql.DB
}

func (d *generalCargoRepo) AddCargoByVN(gc *GeneralCargo) error {
	if d.db == nil {
		return db.ConnectionError
	}

	if gc == nil {
		return db.ValueError
	}

	_, err := d.db.Query("insert into public.general_cargo (voyage_number, cargo_id) values ($1, $2)", gc.VoyageNumber, gc.CargoID)

	return err
}

func (d *generalCargoRepo) DeleteCargo(gc *GeneralCargo) error {
	if d.db == nil {
		return db.ConnectionError
	}

	if gc == nil {
		return db.ValueError
	}

	_, err := d.db.Exec("delete from public.general_cargo where voyage_number = $1 and cargo_id = $2", gc.VoyageNumber, gc.CargoID)

	return err
}



