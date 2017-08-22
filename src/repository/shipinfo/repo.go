package shipinfo

import (
	"db"
	"database/sql"
)

type Repository interface {
	CreateShippingInfo(si *ShippingInfo) error
	//RetrieveSIByNumber(id int) error
	RetrieveAllSchedule() (*[]Schedule, error)
	//UpdateShippingInfo(si *ShippingInfo) error
	DeleteShippingInfo(voyageNumber int) error
}

func ShipScheduleRepository (database *sql.DB) Repository {
	return &shippingRepo{
		db: database,
	}
}

type shippingRepo struct {
	db *sql.DB
}

//Надо переделать
//func (d *shippingRepo) checkIfVNCreated(voyageNumber int) error {
//
//	rows, err := d.db.Query("select * from public.ship_schedule where voyage_number = $1", voyageNumber)
//	if err != nil {
//		return err
//	}
//	defer rows.Close()
//
//	if rows.Next() {
//		return fmt.Errorf("We already have info with voyage number %d", voyageNumber)
//	}
//
//	return nil
//}

//func (d *shippingRepo) сheckDates(si *ShippingInfo) error {
//
//	if si.StartDate.Before(si.EndDate) == false {
//		return fmt.Errorf("Start Date(%s) should to be before End Date(%s).", si.StartDate, si.EndDate)
//	}
//
//	return nil
//}

//func (d *shippingRepo) сheckPorts(si *ShippingInfo) error {
//
//	rows, err := d.db.Query("select * from public.port where id = $1", si.StartingPoint)
//	if err != nil {
//		return err
//	}
//	defer rows.Close()
//
//	if rows.Next() == false {
//		return fmt.Errorf("There is no port with id %d in public.port table", si.StartingPoint)
//	}
//
//	rows, err = d.db.Query("select * from public.port where id = $1", si.FinalDestination)
//	if err != nil {
//		return err
//	}
//
//	if rows.Next() == false {
//		return fmt.Errorf("There is no port with id %d in our database", si.FinalDestination)
//	}
//
//	if si.StartingPoint == si.FinalDestination {
//		return fmt.Errorf("Starting Point(port #%d) shouldn't be equal Final Destination(port #%d).", si.StartingPoint, si.FinalDestination)
//	}
//
//	return nil
//}

func (d *shippingRepo) CreateShippingInfo (si *ShippingInfo) error {
	if d.db == nil {
		return db.ConnectionError
	}

	if si == nil {
		return db.ValueError
	}

	_, err := d.db.Query("insert into public.ship_schedule (voyage_number, starting_point, final_destination, start_date, end_date, ship) values ($1, $2, $3, $4, $5, $6)",
	si.VoyageNumber, si.StartingPoint, si.FinalDestination, si.StartDate, si.EndDate, si.Ship)
	return err
}

func (d *shippingRepo) DeleteShippingInfo(voyageNumber int) error {
	if d.db == nil {
		return db.ConnectionError
	}

	_, err := d.db.Exec("delete from public.ship_schedule where voyage_number = $1", voyageNumber)

	return err
}

func (d *shippingRepo) RetrieveAllSchedule() (*[]Schedule, error) {
	if d.db == nil {
		return nil, db.ConnectionError
	}

	rows, err := d.db.Query("select ship_schedule.voyage_number, ship.name as ship_name, p1.name as starting_point, " +
		"p2.name as final_destination, ship_schedule.start_date, ship_schedule.end_date from public.ship_schedule " +
		"join ship on ship.id = ship_schedule.ship join port as p1 on p1.id = ship_schedule.starting_point " +
		"join port as p2 on p2.id = ship_schedule.final_destination where ship_schedule.voyage_number > 0")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedule = []Schedule{}
	for rows.Next() {
		v := Schedule{}
		err := rows.Scan(&v.VoyageNumber, &v.Ship, &v.StartingPoint, &v.FinalDestination, &v.StartDate, &v.EndDate)

		if err != nil {
			return nil, err
		}

		cargos, err := d.db.Query("select name from general_cargo join cargo on cargo.id = general_cargo.cargo_id where voyage_number = $1", &v.VoyageNumber)
		if err != nil {
			return nil, err
		}
		for cargos.Next() {
			var temp string
			err = cargos.Scan(&temp)
			v.Cargo = append(v.Cargo, temp)
		}

			schedule = append(schedule, v)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &schedule, err
}

