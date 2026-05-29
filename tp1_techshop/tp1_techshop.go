package main

import (
	"errors"
	"fmt"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, produit := range c.Produits {
		if produit.ID == p.ID {
			return errors.New("erreur : ID déjà utilisé")
		}
	}

	c.Produits = append(c.Produits, p)
	return nil
}

func (c *Catalogue) TrouverParID(id int) (Produit, error) {
	for _, produit := range c.Produits {
		if produit.ID == id {
			return produit, nil
		}
	}

	return Produit{}, errors.New("erreur : produit introuvable")
}

func (c *Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultats []Produit

	for _, produit := range c.Produits {
		if strings.EqualFold(produit.Categorie, cat) {
			resultats = append(resultats, produit)
		}
	}

	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nbModifies := 0

	for i := range c.Produits {
		if strings.EqualFold(c.Produits[i].Categorie, categorie) {
			reduction := c.Produits[i].Prix * pct / 100
			c.Produits[i].Prix = c.Produits[i].Prix - reduction
			nbModifies++
		}
	}

	return nbModifies
}

func (c *Catalogue) Vendre(id int, qte int) error {
	for i := range c.Produits {
		if c.Produits[i].ID == id {
			if qte <= 0 {
				return errors.New("erreur : quantité invalide")
			}

			if c.Produits[i].Stock < qte {
				return errors.New("erreur : stock insuffisant")
			}

			c.Produits[i].Stock -= qte
			return nil
		}
	}

	return errors.New("erreur : produit introuvable")
}

func (c Catalogue) Rapport() string {
	nbProduits := len(c.Produits)
	valeurTotale := 0.0

	for _, produit := range c.Produits {
		valeurTotale += produit.Prix * float64(produit.Stock)
	}

	return fmt.Sprintf("Nombre de produits : %d\nValeur totale du stock : %.2f €", nbProduits, valeurTotale)
}

func afficherProduit(p Produit) {
	fmt.Println("-------------------------")
	fmt.Println("ID :", p.ID)
	fmt.Println("Nom :", p.Nom)
	fmt.Println("Marque :", p.Marque)
	fmt.Printf("Prix : %.2f €\n", p.Prix)
	fmt.Println("Stock :", p.Stock)
	fmt.Println("Catégorie :", p.Categorie)
	fmt.Println("Actif :", p.Actif)
}

func main() {
	catalogue := Catalogue{}

	catalogue.AjouterProduit(Produit{ID: 1, Nom: "iPhone 15", Marque: "Apple", Prix: 899.99, Stock: 10, Categorie: "Smartphone", Actif: true})
	catalogue.AjouterProduit(Produit{ID: 2, Nom: "MacBook Air", Marque: "Apple", Prix: 1199.99, Stock: 5, Categorie: "Ordinateur", Actif: true})
	catalogue.AjouterProduit(Produit{ID: 3, Nom: "Galaxy S24", Marque: "Samsung", Prix: 799.99, Stock: 8, Categorie: "Smartphone", Actif: true})
	catalogue.AjouterProduit(Produit{ID: 4, Nom: "ThinkPad X1", Marque: "Lenovo", Prix: 1499.99, Stock: 3, Categorie: "Ordinateur", Actif: true})
	catalogue.AjouterProduit(Produit{ID: 5, Nom: "AirPods Pro", Marque: "Apple", Prix: 249.99, Stock: 20, Categorie: "Accessoire", Actif: true})

	for {
		var choix int

		fmt.Println("\n===== TechShop =====")
		fmt.Println("[1] Ajouter")
		fmt.Println("[2] Chercher")
		fmt.Println("[3] Soldes")
		fmt.Println("[4] Vendre")
		fmt.Println("[5] Rapport")
		fmt.Println("[0] Quitter")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			var p Produit

			fmt.Print("ID : ")
			fmt.Scan(&p.ID)

			fmt.Print("Nom : ")
			fmt.Scan(&p.Nom)

			fmt.Print("Marque : ")
			fmt.Scan(&p.Marque)

			fmt.Print("Prix : ")
			fmt.Scan(&p.Prix)

			fmt.Print("Stock : ")
			fmt.Scan(&p.Stock)

			fmt.Print("Catégorie : ")
			fmt.Scan(&p.Categorie)

			p.Actif = true

			err := catalogue.AjouterProduit(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Produit ajouté avec succès.")
			}

		case 2:
			var id int

			fmt.Print("ID du produit à chercher : ")
			fmt.Scan(&id)

			produit, err := catalogue.TrouverParID(id)
			if err != nil {
				fmt.Println(err)
			} else {
				afficherProduit(produit)
			}

		case 3:
			var categorie string
			var pct float64

			fmt.Print("Catégorie concernée : ")
			fmt.Scan(&categorie)

			fmt.Print("Pourcentage de réduction : ")
			fmt.Scan(&pct)

			nb := catalogue.AppliquerReduction(categorie, pct)
			fmt.Println(nb, "produit(s) modifié(s).")

		case 4:
			var id int
			var qte int

			fmt.Print("ID du produit à vendre : ")
			fmt.Scan(&id)

			fmt.Print("Quantité : ")
			fmt.Scan(&qte)

			err := catalogue.Vendre(id, qte)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Vente effectuée avec succès.")
			}

		case 5:
			fmt.Println("\n===== Rapport =====")
			fmt.Println(catalogue.Rapport())

		case 0:
			fmt.Println("Fin du programme.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
