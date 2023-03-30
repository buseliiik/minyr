package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/buseliiik/minyr/yr"
)

func main() {
	// Opprett en scanner for standard input (brukerens tastaturinput)
	scanner := bufio.NewScanner(os.Stdin)

	// Gjenta programmet inntil brukeren velger å avslutte
	for {
		// Skriv ut menyen
		fmt.Println("Velg en av disse alternativene ved å skrive:(Convert), (Average), (Exit)")

		// Vent på brukerens input
		if !scanner.Scan() {
			break // Avslutt programmet hvis det ikke er mer input
		}

		// Lagre brukerens input som en streng
		input := scanner.Text()

		// Behandle brukerens valg
		switch input {
		case "exit":
			fmt.Println("Avsluttet programmet.")
			return // Avslutt programmet hvis brukeren velger å avslutte
		case "convert":
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit")
			yr.ConvertTemperature() // Kjør funksjonen som konverterer temperatur
		case "average":
			fmt.Println("Gjennomsnitt kalkulator")

			// Fortsett å beregne gjennomsnitt inntil brukeren velger å gå tilbake til hovedmenyen
			for {
				yr.AverageTemperature() // Kjør funksjonen som beregner gjennomsnitt

				// Spør brukeren om de vil gå tilbake til hovedmenyen
				fmt.Println("Tilbake til hovedmeny? (j),(n)")
				if !scanner.Scan() {
					break // Avslutt programmet hvis det ikke er mer input
				}
				input2 := scanner.Text()
				if input2 == "j" {
					break // Gå tilbake til hovedmenyen hvis brukeren velger "j"
				} else if input2 == "n" {
					continue // Fortsett å beregne gjennomsnitt hvis brukeren velger "n"
				}
			}
		default:
			fmt.Println("Ugyldig valg.") // Gi en feilmelding hvis brukeren taster inn en ugyldig input
		}
	}
}

