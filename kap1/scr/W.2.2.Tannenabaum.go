package main

import (
	"fmt"
	"strings"
)

/*
Schreibe ein Programm, das das folgende Muster in der Konsole ausgibt:
*
***
*****
*******
*********
***********
*************
***


*/

func main() {
	// Definiere die Höhe des Baums
	height := 6

	// Erzeuge die oberen Teile des Tannenbaums
	for i := 1; i <= height; i++ {
		// Berechne die Anzahl der Leerzeichen und Sterne
		spaces := strings.Repeat(" ", height-i)
		stars := strings.Repeat("*", 2*i-1)
		// Gib die Zeile aus
		fmt.Println(spaces + stars)
	}

	// Erzeuge den Stamm des Tannenbaums
	stemHeight := 2
	stemWidth := 3
	stemSpaces := strings.Repeat(" ", height-stemWidth/2-1)
	for i := 0; i < stemHeight; i++ {
		stem := strings.Repeat("*", stemWidth)
		// Gib die Zeile des Stamms aus
		fmt.Println(stemSpaces + stem)
	}
}
