// geolocation.go
package geolocation

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Geolocation struct {
    IP          string  `json:"ip"`
    Country     string  `json:"country_name"`
    Region      string  `json:"region_name"`
    City        string  `json:"city"`
    Latitude    float64 `json:"latitude"`
    Longitude   float64 `json:"longitude"`
    Timezone    string  `json:"time_zone"`
}

func GetGeolocation(ip string) (*Geolocation, error) {
    url := fmt.Sprintf("https://freegeoip.app/json/%s", ip)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var geo Geolocation
    if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
        return nil, err
    }

    return &geo, nil
}
