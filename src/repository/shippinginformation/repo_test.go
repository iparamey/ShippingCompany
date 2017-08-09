package shippinginformation

import (
	"testing"
	"db"
	"time"
)

var si = ShippingInfo {
	VoyageNumber: 4,
	StartingPoint: 1,
	FinalDestination: 2,
	StartDate: time.Date(2017, 1, 10, 0, 0 ,0, 0, time.UTC),
	EndDate: time.Date(2017, 1, 12, 0, 0 ,0, 0, time.UTC),
	Ship: 3,
}
func TestShippingRepo(t *testing.T) {
	sr := ShipScheduleRepository(db.GetDBConnection())
	err := sr.CreateShippingInfo(&si)
	if err != nil {
		t.Error(err)
	}

	err = sr.DeleteShippingInfo(4)
	if err != nil {
		t.Error(err)
	}
}
