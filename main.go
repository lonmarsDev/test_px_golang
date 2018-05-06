package main

import (
	"fmt"
)

// PricePoints maintains a list of outstanding order at a speciﬁc price
type PricePoints struct {
	Price          float64
	OrderBookEntry []OrderBookEntry
}

// OrderBookEntry is a representation Order information
type OrderBookEntry struct {
	Size  string
	Price int
	Type  string
}

// TypeMatching function of filter data by Type
func TypeMatching(arrPricePoints []PricePoints, Type string) []OrderBookEntry {
	var ArrOrderBookEntry []OrderBookEntry
	for _, v := range arrPricePoints {
		for _, vi := range v.OrderBookEntry {
			if vi.Type == Type {
				ArrOrderBookEntry = append(ArrOrderBookEntry, vi)
			}

		}
	}
	return ArrOrderBookEntry
}

var arrPricePoints []PricePoints

func main() {

	//Example Record
	var ArrOrderb1 = []OrderBookEntry{
		OrderBookEntry{
			Size:  "Small",
			Price: 200,
			Type:  "BID",
		},
		OrderBookEntry{
			Size:  "Large",
			Price: 300,
			Type:  "ASK",
		},
	}

	var ArrOrderb2 = []OrderBookEntry{
		OrderBookEntry{
			Size:  "Small",
			Price: 400,
			Type:  "BID",
		},
		OrderBookEntry{
			Size:  "Large",
			Price: 1100,
			Type:  "BID",
		},
	}

	var arrPricePoints = []PricePoints{
		PricePoints{ // Now each entry must be an address to a Message struct
			Price:          500,
			OrderBookEntry: ArrOrderb1,
		},
		PricePoints{ // Now each entry must be an address to a Message struct
			Price:          3000,
			OrderBookEntry: ArrOrderb2,
		},
	}
	//End of example record

	//Example call filter by type of BID
	var bidItem = TypeMatching(arrPricePoints, "BID")
	fmt.Print(bidItem)

}
