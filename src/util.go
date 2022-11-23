package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

func fetchApi(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func getDistanceInMeters(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	radlat1 := float64(lat1 * math.Pi / 180)
	radlat2 := float64(lat2 * math.Pi / 180)

	radtheta := float64((lng1 - lng2) * math.Pi / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515 * 1609.344

	return dist
}
