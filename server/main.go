package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Quote struct {
	PickupPostcode   string `json:"pickupPostcode"`
	DeliveryPostcode string `json:"deliveryPostcode"`
	Vehicle          string `json:"vehicle,omitempty"`
	Price            uint64 `json:"price"`
}

func (q Quote) getMarkupPercentage(vehicle string) float64 {
	// case sensitive inputs? this is an assumption
	// I'm making.
	vehicle = strings.ToLower(vehicle)

	// given that there are only a few vehicles in the spec
	// I'm using a switch here.
	// other ways to do this could be having a map
	// that contains vehicles => markup percentages
	switch vehicle {
	case "bicycle":
		return 1.10
	case "motorbike":
		return 1.15
	case "parcel_car":
		return 1.20
	case "small_van":
		return 1.30
	case "large_van":
		return 1.40
	default:
		// don't know the vehicle? no markup price.
		return 1.0
	}
}

func (q Quote) calculateDeliveryCost() (uint64, error) {
	// strip the spaces from the pickup and delivery postcodes
	// so it plays nice with strconv
	pickup := strings.Replace(q.PickupPostcode, " ", "", -1)
	delivery := strings.Replace(q.DeliveryPostcode, " ", "", -1)

	pickupBase36, err := strconv.ParseInt(pickup, 36, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	deliveryBase36, err := strconv.ParseInt(delivery, 36, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	diff := (pickupBase36 - deliveryBase36) / 100000000
	final := float64(diff) * q.getMarkupPercentage(q.Vehicle)
	return uint64(math.Abs(final)), nil
}

func main() {
	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var quote Quote
		err := json.NewDecoder(r.Body).Decode(&quote)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		quote.Price, err = quote.calculateDeliveryCost()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("done", quote)

		encoder := json.NewEncoder(w)
		encoder.Encode(&quote)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe("127.0.0.1:8080", nil)
}
