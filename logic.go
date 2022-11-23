package main

import (
	"encoding/json"
	"github.com/samber/lo"
	"sort"
	"sync"
)

func bixiLogic(city string, nbResult int, lat, lon float64) []StationAround {
	var bixiStationInformation BixiStationInformation
	var bixiStationStatus BixiStationStatus

	wg := sync.WaitGroup{}
	wg.Add(2)

	go fetchBixiStationInformation(&wg, &bixiStationInformation)
	go fetchBixiStationStatus(&wg, &bixiStationStatus)

	wg.Wait()

	var stationArounds []StationAround

	stationArounds = lo.Map(bixiStationInformation.Data.StationInformation, func(station StationInformation, i int) StationAround {
		stationCoordinates := Coordinates{station.Lat, station.Lon}
		return StationAround{
			Id:          station.StationID,
			Name:        station.Name,
			Coordinates: stationCoordinates,
			Distance:    int(distance(lat, lon, station.Lat, station.Lon, "K") * 1000)}
	})

	sort.SliceStable(stationArounds, func(i, j int) bool {
		return stationArounds[i].Distance < stationArounds[j].Distance
	})

	nearestStations := lo.Slice(stationArounds, 0, nbResult)

	nearestStationCompleted := lo.Map(nearestStations, func(stationAround StationAround, i int) StationAround {
		stationStatus := lo.Filter[StationStatus](bixiStationStatus.Data.Stations, func(station StationStatus, index int) bool {
			return station.StationId == stationArounds[i].Id
		})[0]
		stationAround.Bikes = stationStatus.NumBikesAvailable
		stationAround.ElectricBikes = stationStatus.NumEbikesAvailable
		stationAround.Docks = stationStatus.NumDocksAvailable
		return stationAround
	})

	return nearestStationCompleted
}

func fetchBixiStationStatus(wg *sync.WaitGroup, stationStatus *BixiStationStatus) {
	bixiStationStateRes := fetchApi("https://gbfs.velobixi.com/gbfs/fr/station_status.json")
	json.Unmarshal(bixiStationStateRes, stationStatus)
	wg.Done()
}

func fetchBixiStationInformation(wg *sync.WaitGroup, stationInformation *BixiStationInformation) {
	bixiStationInformationRes := fetchApi("https://gbfs.velobixi.com/gbfs/fr/station_information.json")
	json.Unmarshal(bixiStationInformationRes, stationInformation)
	wg.Done()
}
