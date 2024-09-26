package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type carData struct {
	CarsDailyCount        int
	TrafficIntensityIndex float64
}

func main() {
	serverVersion1()
}

func serverVersion1() {
	go func() {
		log.Println("http server starting")

		http.HandleFunc("/example", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(carData{
				CarsDailyCount:        928,
				TrafficIntensityIndex: 3.5,
			})
		})

		http.ListenAndServe(":8090", nil)
	}()

	time.Sleep(1 * time.Second)

	for _ = range 10 {
		time.Sleep(500 * time.Millisecond)

		r, err := http.Get("http://127.0.0.1:8090/example")
		if err != nil {
			fmt.Println(fmt.Errorf("error in Get request: %w", err))

			return
		}

		cd := new(carData)
		err = json.NewDecoder(r.Body).Decode(&cd)
		if err != nil {
			fmt.Println(fmt.Errorf("error in Get request: %w", err))

			return
		}

		fmt.Println("Traffic Intensity:", cd.TrafficIntensityIndex)
		fmt.Println("Cars Daily:", cd.CarsDailyCount)
	}
}
