package examples

import (
	"fmt"
)

func MapsEx1() {

	x := make(map[string]int)
	x["Key"] = 1
	fmt.Println(x["Key"])
}

func MapsEx2() {
	elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["O"])

	fmt.Println(elements)

}

func MapsEx3() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "Gas",
		},
	}

	if el, ok := elements["H"]; ok {

		fmt.Println(el["name"], el["state"])
	}

}
