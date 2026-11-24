// Package vreme omogoča upravljanje vremenskih postaj in njihovih
// temperaturnih meritev. Vsebuje funkcije za dodajanje postaj,
// dodajanje meritev, izpis meritev in pripravo statističnih poročil.
package vreme

import (
	"fmt"
	"math"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) error {
	if oznaka == "" || ime == "" || kraj == "" {
		return fmt.Errorf("oznaka, ime in kraj ne smejo biti prazni")
	}

	if postaje == nil {
		return fmt.Errorf("slovar postaj ni inicializiran")
	}

	if _, obstaja := postaje[oznaka]; obstaja {
		return fmt.Errorf("postaja %s že obstaja", oznaka)
	}

	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}

	return nil
}

func DodajMeritev(
	postaje map[string]Postaja,
	oznaka string,
	temperatura float64,
	minTemp float64,
	maxTemp float64,
) error {
	postaja, obstaja := postaje[oznaka]

	if !obstaja {
		return fmt.Errorf("postaja %s ne obstaja", oznaka)
	}

	if math.IsNaN(temperatura) || math.IsInf(temperatura, 0) {
		return fmt.Errorf("temperatura mora biti končno število")
	}

	if minTemp > maxTemp {
		return fmt.Errorf("najnižja temperatura ne sme biti večja od najvišje")
	}

	if temperatura < minTemp || temperatura > maxTemp {
		return fmt.Errorf(
			"temperatura %.2f °C je izven dovoljenega območja",
			temperatura,
		)
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja

	fmt.Println("Meritev dodana")
	return nil
}

func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE")

	for oznaka, postaja := range postaje {
		fmt.Printf(
			"%s - %s (%s): %v\n",
			oznaka,
			postaja.ime,
			postaja.kraj,
			postaja.meritve,
		)
	}
}

func statistika(
	postaje map[string]Postaja,
	oznaka string,
	stMeritev int,
) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]

	if !obstaja || stMeritev < 1 {
		return 0, 0, 0, false
	}

	meritve := postaja.meritve

	if len(meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = meritve[0]
	najvisja = meritve[0]

	for _, temperatura := range meritve {
		if temperatura < najnizja {
			najnizja = temperatura
		}

		if temperatura > najvisja {
			najvisja = temperatura
		}

		povprecna += temperatura
	}

	povprecna /= float64(len(meritve))

	return najnizja, najvisja, povprecna, true
}

func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	fmt.Println("POROČILO")

	for oznaka, postaja := range postaje {
		min, max, povprecje, ok := statistika(
			postaje,
			oznaka,
			stMeritev,
		)

		if !ok {
			fmt.Printf("%s: Premalo meritev\n", postaja.ime)
			continue
		}

		var opis string

		switch {
		case povprecje >= 25:
			opis = "Vroče"
		case povprecje >= 15:
			opis = "Toplo"
		case povprecje >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf(
			"%s: min %.2f °C, max %.2f °C, povprečje %.2f °C -> %s\n",
			postaja.ime,
			min,
			max,
			povprecje,
			opis,
		)
	}
}

