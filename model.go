package main

type BixiStationInformation struct {
	LastUpdated int `json:"last_updated,omitempty"`
	TTL         int `json:"ttl,omitempty"`
	Data        struct {
		StationInformation []StationInformation `json:"stations,omitempty"`
	} `json:"data,omitempty"`
}

type StationInformation struct {
	StationID string  `json:"station_id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Lat       float64 `json:"lat,omitempty"`
	Lon       float64 `json:"lon,omitempty"`
}

type BixiStationStatus struct {
	LastUpdated int `json:"last_updated"`
	Ttl         int `json:"ttl"`
	Data        struct {
		Stations []StationStatus `json:"stations"`
	} `json:"data"`
}

type StationStatus struct {
	StationId          string `json:"station_id"`
	NumBikesAvailable  int    `json:"num_bikes_available"`
	NumEbikesAvailable int    `json:"num_ebikes_available"`
	NumBikesDisabled   int    `json:"num_bikes_disabled"`
	NumDocksAvailable  int    `json:"num_docks_available"`
	LastReported       int    `json:"last_reported"`
}

type StationAround struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	Coordinates   Coordinates `json:"coordinates"`
	Distance      int         `json:"distanceFromUser"`
	Bikes         int         `json:"bikes"`
	ElectricBikes int         `json:"ebikes"`
	Docks         int         `json:"docks"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
