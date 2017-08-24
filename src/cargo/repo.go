package cargo

import (
	"db"
	"database/sql"
)

type Repository interface {
	CreateCargo (cargo *Cargo) error
	UpdateCargo (cargo *Cargo) error
	DeleteCargo (id int) error
}

func CargoRepository(database *sql.DB) Repository {
	return &cargoRepo{
		db: database,
	}
}

type cargoRepo struct {
	db *sql.DB
}

func (cr *cargoRepo) CreateCargo (cargo *Cargo) error {
	if cargo == nil {
		return db.ValueError
	}

	if cr.db == nil {
		return db.ConnectionError
	}

	return cr.db.QueryRow("insert into public.cargo (name) values ($1) returning id", cargo.Name).Scan(&cargo.ID)

}

func (cr *cargoRepo) UpdateCargo(cargo *Cargo) error {
	if cargo == nil {
		return db.ValueError
	}

	if cr.db == nil {
		return db.ConnectionError
	}

	_, err := cr.db.Exec("update public.cargo SET name = $1 where id = $2", cargo.Name, cargo.ID)
	if err != nil  {
		return err
	}
	return nil
}

func (cr *cargoRepo) DeleteCargo (id int) error {
	if cr.db == nil {
		return db.ConnectionError
	}

	if id <= 0 {
		return db.ValueError
	}

	_, err := cr.db.Exec("delete from public.cargo where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}