package shippinginformation

import (
	"db"
	"database/sql"
	"fmt"
)

type Repository interface {
	CreateShippingInfo(si *ShippingInfo) error
	//RetrieveSIByNumber(id int) error
	//RetrieveAllSchedule() error
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
func (d *shippingRepo) checkIfVNCreated(voyageNumber int) error {

	rows, err := d.db.Query("select * from public.ship_schedule where voyage_number = $1", voyageNumber)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return fmt.Errorf("We already have info with voyage number %d", voyageNumber)
	}

	return nil
}

func (d *shippingRepo) сheckDates(si *ShippingInfo) error {

	if si.StartDate.Before(si.EndDate) == false {
		return fmt.Errorf("Start Date(%s) should to be before End Date(%s).", si.StartDate, si.EndDate)
	}

	return nil
}

func (d *shippingRepo) сheckPorts(si *ShippingInfo) error {

	rows, err := d.db.Query("select * from public.port where id = $1", si.StartingPoint)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() == false {
		return fmt.Errorf("There is no port with id %d in public.port table", si.StartingPoint)
	}

	rows, err = d.db.Query("select * from public.port where id = $1", si.FinalDestination)
	if err != nil {
		return err
	}

	if rows.Next() == false {
		return fmt.Errorf("There is no port with id %d in our database", si.FinalDestination)
	}

	if si.StartingPoint == si.FinalDestination {
		return fmt.Errorf("Starting Point(port #%d) shouldn't be equal Final Destination(port #%d).", si.StartingPoint, si.FinalDestination)
	}

	return nil
}

func (d *shippingRepo) CreateShippingInfo (si *ShippingInfo) error {
	if d.db == nil {
		return db.ConnectionError
	}

	if si == nil {
		return db.ValueError
	}

	err := d.checkIfVNCreated(si.VoyageNumber)

	if err != nil {
		return err
	}

	err = d.сheckDates(si)

	if err != nil {
		return err
	}

	err = d.сheckPorts(si)
	if err != nil {
		return err
	}

	_, err = d.db.Query("insert into public.ship_schedule (voyage_number, starting_point, final_destination, start_date, end_date, ship) values ($1, $2, $3, $4, $5, $6)",
	si.VoyageNumber, si.StartingPoint, si.FinalDestination, si.StartDate, si.EndDate, si.Ship)

	return err
}

func (d *shippingRepo) DeleteShippingInfo(voyageNumber int) error {
	if d.db == nil {
		return db.ConnectionError
	}


	if d.checkIfVNCreated(voyageNumber) == nil {
		return db.ValueError
	}

	_, err := d.db.Exec("delete from public.ship_schedule where voyage_number = $1", voyageNumber)

	return err
}

func (d *shippingRepo) getAllInfoByVN(voyageNumber int) error {
	if d.db != nil {
		return db.ConnectionError
	}

	if d.checkIfVNCreated(voyageNumber) == nil {
		return db.ValueError
	}

	return nil
}

