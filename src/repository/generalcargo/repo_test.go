package generalcargo

import (
	"testing"
	"db"
)

var gc = GeneralCargo{
	CargoID: 11,
	VoyageNumber: 3,
}

func TestAddCargoByVN(t *testing.T) {
	gencar := GeneralCargoRepositpory(db.GetDBConnection())
	err := gencar.AddCargoByVN(&gc)

	if err != nil {
		t.Error(err)
	}

	err = gencar.DeleteCargo(&gc)

	if err != nil {
		t.Error(err)
	}
}