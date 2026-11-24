# Vremenska postaja

Go modul za upravljanje vremenskih postaj in temperaturnih meritev.

### Funkcionalnosti
- Dodajanje vremenskih postaj in meritev.
- Izpis meritev in statističnega poročila.
- Nastavljanje temperaturnih omejitev.
- Nastavljanje najmanjšega števila meritev za statistiko.

### Zagon

```bash
go run .
```

### Primer z uporabo stikal

```bash
go run . --minTemp=-30 --maxTemp=40 --stMeritev=5
```

### Uporabljeni paket

`github.com/urfave/cli/v3`

### Avtor

MatLes26
