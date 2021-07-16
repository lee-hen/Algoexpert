package airport_connections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var Airports = []string{
	"BGI",
	"CDG",
	"DEL",
	"DOH",
	"DSM",
	"EWR",
	"EYW",
	"HND",
	"ICN",
	"JFK",
	"LGA",
	"LHR",
	"ORD",
	"SAN",
	"SFO",
	"SIN",
	"TLV",
	"BUD",
}

var StartingAirport = "LGA"

func TestCase1(t *testing.T) {
	routes := [][]string{
		{"DSM", "ORD"},
		{"ORD", "BGI"},
		{"BGI", "LGA"},
		{"SIN", "CDG"},
		{"CDG", "SIN"},
		{"CDG", "BUD"},
		{"DEL", "DOH"},
		{"DEL", "CDG"},
		{"TLV", "DEL"},
		{"EWR", "HND"},
		{"HND", "ICN"},
		{"HND", "JFK"},
		{"ICN", "JFK"},
		{"JFK", "LGA"},
		{"EYW", "LHR"},
		{"LHR", "SFO"},
		{"SFO", "SAN"},
		{"SFO", "DSM"},
		{"SAN", "EYW"},
	}
	output := AirportConnections(Airports, routes, StartingAirport)
	expected := 3
	require.Equal(t, expected, output)
}
