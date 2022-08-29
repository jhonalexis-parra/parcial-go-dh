package tickets

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jhonalexis-parra/parcial-go-dh/internal/repository"
)

func GetTotalTickets(destination string) (int, error) {
	data, err := repository.Read()
	contador := 0
	if err != nil {
		return 0, err
	}
	for _, travel := range data {
		if travel[3] == destination {
			contador++
		}
	}
	return contador, nil
}

func GetCountByPeriod(time string) (int, error) {
	data, err := repository.Read()
	contadorMadrugada := 0
	contadorManhana := 0
	contadorTarde := 0
	contadorNoche := 0
	if err != nil {
		return 0, err
	}
	for _, travel := range data {
		hora := strings.Split(travel[4], ":")[0]
		intVar, _ := strconv.Atoi(hora)
		if intVar >= 0 && intVar <= 6 {
			contadorMadrugada++
		}
		if intVar > 6 && intVar <= 12 {
			contadorManhana++
		}
		if intVar > 12 && intVar <= 19 {
			contadorTarde++
		}
		if intVar > 19 && intVar <= 23 {
			contadorNoche++
		}
	}

	switch time {
	case "madrugada":
		return contadorMadrugada, nil
	case "mañana":
		return contadorManhana, nil
	case "tarde":
		return contadorTarde, nil
	case "noche":
		return contadorNoche, nil
	default:
		return 0, fmt.Errorf("Does not exist the schedule")
	}
}

func AverageDestination(destination string, total int) (float64, error) {
	data, err := repository.Read()
	contador := 0
	if err != nil {
		return 0, err
	}
	for _, travel := range data {
		if travel[3] == destination {
			contador++
		}
	}
	return (float64(contador) / float64(total)) * 100, nil
}
