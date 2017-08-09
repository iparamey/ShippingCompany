package cargo

import (
	"testing"
	"db"
	"fmt"
)

var cargo = Cargo {
	Name: "Fructs",
}

func TestDatabaseCargo(t *testing.T) {

	cr := CargoRepository(db.GetDBConnection())
	err := cr.CreateCargo(&cargo)


	if err != nil {
		t.Errorf("Can't add new cargo: %+v", err)
	}

	cargo.Name = "Multifructs"
	err = cr.UpdateCargo(&cargo)
	if err != nil {
		t.Errorf("Can't update cargo: %+v", err)
	}

	err = cr.DeleteCargo(cargo.ID)
	if err != nil {
		t.Errorf("Can't delete cargo: %+v", err)
	}
}

func TestDatabaseCargoBroken(t *testing.T) {
	cr := CargoRepository(db.GetDBConnection())
	err := cr.CreateCargo(nil)

	if err == nil {
		t.Error("It shouldn't allow to create a cargo with nil arguments.")
	}

	err = cr.UpdateCargo(nil)

	if err == nil {
		t.Error("It shouldn't allow to update a cargo with nil arguments.")
	}

	err = cr.DeleteCargo(0)

	if err == nil {
		t.Error("It shouldn't work with id less or equal 0.")
	}
}



