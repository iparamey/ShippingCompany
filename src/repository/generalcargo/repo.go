package generalcargo

import (
	"database/sql"
	"fmt"
	"db"
)

type Repository interface {
	AddCargoByVN(gc *GeneralCargo) error
	//RetrieveCargoesByVN (vn int) error
	//UpdateCargo(cargo []int, vn int) error
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

func (d *generalCargoRepo) checkCargoByVN(gc *GeneralCargo) error {

	rows, err := d.db.Query("select cargo_id from public.general_cargo where voyage_number = $1", gc.VoyageNumber)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var cargo int
		err := rows.Scan(&cargo)
		if err != nil {
			return err
		}
		if cargo == gc.CargoID {
			return fmt.Errorf("We already have cargo #%d in voyage #%d", gc.CargoID, gc.VoyageNumber)
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func (d *generalCargoRepo) AddCargoByVN(gc *GeneralCargo) error {
	if d.db == nil {
		return db.ConnectionError
	}

	if gc == nil {
		return db.ValueError
	}

	err := d.checkCargoByVN(gc)
	if err != nil {
		return err
	}
	_, err = d.db.Query("insert into public.general_cargo (voyage_number, cargo_id) values ($1, $2)", gc.VoyageNumber, gc.CargoID)

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

//func (d *generalCargoRepo) RetrieveCargoesByVN(vn int) (*[]GeneralCargo, error) {
//	if d.db == nil {
//		return db.ConnectionError
//	}
//
//
//
//	return nil
//}


