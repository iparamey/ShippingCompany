package shippinginformation

import (
	"time"
)

type ShippingInfo struct {
	VoyageNumber      int
	StartingPoint     int
	FinalDestination  int
	StartDate         time.Time
	EndDate        	  time.Time
	Ship		  int
}

