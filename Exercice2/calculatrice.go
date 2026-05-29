package main

import (
	"errors"
	"fmt"
)

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro impossible")
		}
		return a / b, nil
	default:
		return 0, errors.New("opération inconnue")
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	default:
		return nil
	}
}

func main() {
	for {
		var a, b float64
		var op string

		fmt.Println("Entrez deux nombres et une opération (+, -, *, /) ou 'quit' pour quitter :")
		fmt.Scan(&a, &b, &op)

		if op == "quit" {
			fmt.Println("Fin du programme.")
			break
		}

		resultat, err := operer(a, b, op)

		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Résultat :", resultat)
		}
	}
}
