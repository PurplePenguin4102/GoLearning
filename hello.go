package main

import "fmt"

type streng string

type WingCommander struct {
	name string
	callsign string
}

const (
	StarSystem = "vega"
	Enemy = "Kilrathi"
	Ship = iota // this marks as changeable
)

//enum
const (
	A = iota
	B
	C
)

func PrintsGreeting (wc WingCommander) {
	fmt.Println(wc.name, wc.callsign, Ship)
}

func main() {
	
	// declare variable, name symbol, declare type
	var message string
	message = "Welcome to the Vega System wing commander. Here are your ships\n"
	var ShipIda, ShipIdb, ShipIdc int = 1, 2, 3
	var myStreng = "\n I will now divine your name"

	//var altset string = "another way to declare + initialise"
	
	// strongly typed, this is an error line
	//message = true

	// type inferencing
	//var infer = "type is inferred to string"

	// shorthand
	//infer2 := "now with := we have declared and initialised"

	var commander = WingCommander{name: "Alex Stevens", callsign: "Maverick"}

	// pointer initialisation
	var bling *string = &message
	fmt.Println(message, ShipIda, ShipIdb, ShipIdc, bling, myStreng)
	PrintsGreeting(commander)
}