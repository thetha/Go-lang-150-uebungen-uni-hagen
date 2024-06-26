package main

import (
	"fmt"
)

/*
Aufgabenstellung
Schreibe ein Programm, das die bei den Testfällen stehenden Rechenaufgaben berechnet,
die Ergebnisse in einer Variablen speichert und die Werte der Variablen in der Konsole
ausgibt.
Testfälle
1 + 7 - 9 + 43 = 42
12 % 2 = 0
43 - 9 + 7 + 1 = 42
13 % 5 = 3
4 * 3 + 1 = 13
12345 / 10 = 1234
9 * (2 + 1) = 27
1234 / 10 = 123
4 * 3 / 6 = 2
123 / 10 = 12
(3 - 7) * (7 + 4) = -44
12 / 10 = 1
3 / 4 = 0
1 / 10 = 0
3.0 / 4.0 = 0.75
12345 % 10 = 5
4 / 3 + 1 * 7 = 8
1234 % 10 = 4
4.0 / 3.0 + 1 * 7 = 8.333333
123 % 10 = 3
2 * 2 * 2 * 2 * 2 * 2 * 2 = 128
12 % 10 = 2
42 / 7 / 3 = 2
1 % 10 = 1
*/

func main() {
	a := 1 + 7 - 9 + 43
	fmt.Println(a)

	b := 12 % 2
	fmt.Println(b)

	c := 43 - 9 + 7 + 1
	fmt.Println(c)

	d := 13 % 5
	fmt.Println(d)

}
