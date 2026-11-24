// Program vremenska-postaja omogoča upravljanje vremenskih postaj,
// dodajanje temperaturnih meritev in izpis statističnega poročila.
// Uporablja knjižnico github.com/urfave/cli/v3 za ukazni vmesnik.
// Omogoča nastavljanje najnižje in najvišje temperature ter
// najmanjšega števila meritev za izračun statistike.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MatLes26/vremenska-postaja/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "vremenska-postaja",
		Usage: "Program za upravljanje vremenskih postaj",

		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50,
				Usage: "Najnižja dovoljena temperatura",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "Najvišja dovoljena temperatura",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "Najmanjše število meritev za statistiko",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			if minTemp > maxTemp {
				return fmt.Errorf("minTemp ne sme biti večji od maxTemp")
			}

			if stMeritev < 1 {
				return fmt.Errorf("stMeritev mora biti vsaj 1")
			}

			slovarPostaj := make(map[string]vreme.Postaja)

			err := vreme.DodajPostajo(
				slovarPostaj, "LJ01", "Center", "Ljubljana",
			)
			if err != nil {
				return err
			}

			err = vreme.DodajPostajo(
				slovarPostaj, "KP01", "Hrib", "Kranj",
			)
			if err != nil {
				return err
			}

			err = vreme.DodajPostajo(
				slovarPostaj, "MB01", "Tabor", "Maribor",
			)
			if err != nil {
				return err
			}

			err = vreme.DodajPostajo(
				slovarPostaj, "LJ02", "Rudnik", "Ljubljana",
			)
			if err != nil {
				return err
			}

			err = vreme.DodajPostajo(
				slovarPostaj, "MS01", "Predmestje", "Murska Sobota",
			)
			if err != nil {
				return err
			}

			meritve := map[string][]float64{
				"LJ01": {10, 25, -4, 8, 3},
				"KP01": {33, 2, -10, 11, -19},
				"MB01": {-20, 4, 13, 14, -2},
				"LJ02": {33, -19, 6, -11, 0},
			}

			for oznaka, seznam := range meritve {
				for _, temperatura := range seznam {
					err := vreme.DodajMeritev(
						slovarPostaj,
						oznaka,
						temperatura,
						minTemp,
						maxTemp,
					)

					if err != nil {
						fmt.Println(err)
					}
				}
			}

			fmt.Println()

			err = vreme.DodajMeritev(
				slovarPostaj,
				"LJ01",
				23,
				minTemp,
				maxTemp,
			)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println()
			vreme.IzpisMeritev(slovarPostaj)

			fmt.Println()
			vreme.IzpisPorocila(slovarPostaj, stMeritev)

			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

