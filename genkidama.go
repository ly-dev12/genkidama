package main

import (
	"fmt"
)

type Genkidama struct {
	personaje *string
	energy    *int
	cenergy   *int
	complete  bool
}

func (genk *Genkidama) getEnergy() {
	// fmt.Println(*genk.personaje)
	fmt.Printf("Acomulated Energy: %d / %d\n", *genk.cenergy, *genk.energy)

	if *genk.cenergy > *genk.energy {
		*&genk.complete = true
	}
}

func main() {

	// Vareiables declaration
	var charName string
	var energy int

	fmt.Printf("Character's name: ")
	fmt.Scanln(&charName)

	fmt.Printf("Enter max energy: ")
	fmt.Scanln(&energy)

	//fmt.Println(&charName)
	var cenergy = 0
	var genkidama Genkidama = Genkidama{personaje: &charName, energy: &energy, cenergy: &cenergy, complete: false}
	//fmt.Println("genkmada memory: ", genkidama)
	fmt.Println("------------------ DATA ------------------")
	fmt.Println("Character Name: ", *genkidama.personaje)
	fmt.Println("Energy required to form Genkidama: ", *genkidama.energy)

	fmt.Println("Forming Genkidama...")
	//fmt.Println(genkidama.personaje)
	var genergy int
	for genkidama.complete == false {
		fmt.Println("Give me energy: ")
		fmt.Scanln(&genergy)
		cenergy = cenergy + genergy
		(&genkidama).getEnergy()
		//fmt.Println("gerkidama memory2: ", &genkidama)
	}
	//fmt.Println(genkidama.complete)
	if genkidama.complete == true {
		fmt.Println("We have managed to gather the necessary energy...")

		fmt.Println("Time to throw it away, bye majin buu")
		fmt.Println("AHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH!!!....")

		fmt.Printf("%s Wins!", *genkidama.personaje)
	}
}
