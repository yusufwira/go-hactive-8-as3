package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type WindWaterData struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

type WindWaterRespon struct {
	ID          int       `json:"id"`
	Wind        int       `json:"wind"`
	Water       int       `json:"water"`
	StatusWind  string    `json:"status_wind"`
	StatusWater string    `json:"status_water"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WindWaterDataRespon struct {
	Data WindWaterRespon
}

func UpdateData() {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			water := rand.Intn(100) + 1
			wind := rand.Intn(100) + 1

			data := WindWaterData{
				Wind:  wind,
				Water: water,
			}
			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("Error encoding JSON: %v", err)
			}
			url := "http://localhost:8080/api/windwater/update"
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				log.Fatalf("Error making POST request: %v", err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Error reading response body: %v", err)
			}

			var response WindWaterDataRespon
			err = json.Unmarshal(body, &response)
			if err != nil {
				log.Fatalf("Error decoding JSON: %v", err)
			}
			dataRes := WindWaterData{
				Water: response.Data.Water,
				Wind:  response.Data.Wind,
			}
			jsonDataRes, _ := json.Marshal(dataRes)
			fmt.Println(string(jsonDataRes))
			fmt.Printf("Status Water : %s\n", response.Data.StatusWater)
			fmt.Printf("Status Wind : %s\n", response.Data.StatusWind)
			// defer resp.Body.Close()

		}
	}
}

func main() {
	go UpdateData()
	select {}
}
