// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/ipgeolocation/geolocation"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ipgeolocation <ip>")
		os.Exit(1)
	}

	ip := os.Args[1]
	geo, err := geolocation.GetGeolocation(ip)
	if err != nil {
		fmt.Printf("Failed to get geolocation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("IP: %s\nCountry: %s\nRegion: %s\nCity: %s\nLatitude: %f\nLongitude: %f\nTimezone: %s\n",
		geo.IP, geo.Country, geo.Region, geo.City, geo.Latitude, geo.Longitude, geo.Timezone)
}
