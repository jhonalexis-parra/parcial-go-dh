package main

import (
	"fmt"

	"github.com/jhonalexis-parra/parcial-go-dh/internal/tickets"
)

func main() {

	// Primer Punto
	valor, _ := tickets.GetTotalTickets("Russia")
	fmt.Println(valor)
	valor, _ = tickets.GetTotalTickets("Jamaica")
	fmt.Println(valor)

	// Segundo Punto
	valor, _ = tickets.GetCountByPeriod("madrugada")
	fmt.Println(valor)
	valor, _ = tickets.GetCountByPeriod("mañana")
	fmt.Println(valor)
	valor, _ = tickets.GetCountByPeriod("tarde")
	fmt.Println(valor)
	valor, _ = tickets.GetCountByPeriod("noche")
	fmt.Println(valor)

	// Tercer punto
	// Primer Punto
	valorPromedio, _ := tickets.AverageDestination("Russia", 1000)
	fmt.Println(fmt.Sprintf("el valor es de %f %%", valorPromedio))
	valorPromedio, _ = tickets.AverageDestination("Jamaica", 1000)
	fmt.Println(fmt.Sprintf("el valor es de %f %%", valorPromedio))

}
