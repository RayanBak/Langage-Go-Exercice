package main

import "fmt"

func main() {
	var poids float64 = 65
	var taille float64 = 1.79

	const nom = "Rayan"

	const IMCMaigreur = 18.5
	const IMCNormal = 25.0
	const IMCSurpoids = 30.0

	imc := poids / (taille * taille)

	fmt.Println("Calcul de l'IMC de", nom)
	fmt.Printf("IMC : %.2f\n", imc)

	if imc < IMCMaigreur {
		fmt.Println("Catégorie : Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Catégorie : Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Catégorie : Surpoids")
	} else {
		fmt.Println("Catégorie : Obésité")
	}
}
