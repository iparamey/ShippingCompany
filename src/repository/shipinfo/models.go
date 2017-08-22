package shipinfo

import "time"

type ShippingInfo struct {
	VoyageNumber      int		`json:"voyagenumber"`
	StartingPoint     int		`json:"startingpoint"`
	FinalDestination  int		`json:"endpoint"`
	StartDate         time.Time	`json:"startdate"`
	EndDate        	  time.Time	`json:"enddate"`
	Ship		  int		`json:"ship"`
}

type Schedule struct {
	VoyageNumber       int 		`json:"voyagenumber"`
	Ship 		   string	`json:"ship"`
	StartingPoint      string	`json:"startingpoint"`
	FinalDestination   string	`json:"endpoint"`
	StartDate 	   string	`json:"startdate"`
	EndDate		   string	`json:"enddate"`
	Cargo 		   []string	`json:"cargo"`
}


