
package main

import (
	"fmt"
)

type Planet struct {
	planetName        string
	numberOfMoons         int
}

func main() {
	Mercurio := Planet{
		name:        "Mercurio",
		moons:         0,
	}
	Venus := Planet{
		name:        "Venus",
		moons:         0,
	}
	Tierra := Planet{
		name:        "Tierra",
		moons:         1,
	}
	Marte := Planet{
		name:        "Marte",
		moons:         2,
	}
	Jupiter := Planet{
		name:        "Jupiter",
		moons:         63,
	}
	Saturno := Planet{
		name:        "Saturno",
		moons:         62,
	}
	Urano := Planet{
		name:        "Urano",
		moons:         27,
	}
	Neptuno := Planet{
		name:        "Neptuno",
		moons:         13,
	}

	PlanetArray:= [8]Planet{Mercurio,Venus,Tierra,Marte,Jupiter,Saturno,Urano,Neptuno}

	if PlanetArray.moons == 0 {
		fmt.Println("Mercurio y Venus tienen Venus.moons lunas")
	} else if Planet.moons >=1 & <=62 {
		fmt.Println("Tierra tiene Tierra.moons luna, Marte tiene Marte.moons lunas y Jupiter tiene Jupiter.moons lunas")
	} else{
		fmt.Println("Saturno tiene Saturno.moons lunas, Urano tiene Urano.moons lunas y Neptuno tiene Neptuno.moons lunas")
		fmt.Println(planetName, numberOfMoons)
}
}
