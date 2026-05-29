package main

import "fmt"

// Struct Personne
type Personne struct {
	Prenom string `json:"prenom,omitempty"`
	Nom    string `json:"nom,omitempty"`
	Age    int    `json:"age,omitempty"`
	Email  string `json:"email,omitempty"`
}

// Méthode NomComplet
func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

// Méthode Presentation
func (p Personne) Presentation() string {
	return fmt.Sprintf("Je m'appelle %s, j'ai %d ans et mon email est %s.",
		p.NomComplet(), p.Age, p.Email)
}

// Struct Adresse
type Adresse struct {
	Rue        string `json:"rue,omitempty"`
	Ville      string `json:"ville,omitempty"`
	CodePostal string `json:"code_postal,omitempty"`
}

// Méthode Format
func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

// Struct Employe avec embedding
type Employe struct {
	Personne
	Adresse
	Poste   string  `json:"poste,omitempty"`
	Salaire float64 `json:"salaire,omitempty"`
}

// Méthode FicheEmploye
func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"Employé : %s\nAge : %d\nEmail : %s\nAdresse : %s\nPoste : %s\nSalaire : %.2f €",
		e.NomComplet(),
		e.Age,
		e.Email,
		e.Adresse.Format(),
		e.Poste,
		e.Salaire,
	)
}

// Méthode AugmenterSalaire avec pointeur
func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire + (e.Salaire * pct / 100)
}

// Struct Etudiant avec embedding
type Etudiant struct {
	Personne
	Promo   string  `json:"promo,omitempty"`
	Moyenne float64 `json:"moyenne,omitempty"`
}

// Méthode MentionObtenue
func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "Très bien"
	case e.Moyenne >= 14:
		return "Bien"
	case e.Moyenne >= 12:
		return "Assez bien"
	case e.Moyenne >= 10:
		return "Passable"
	default:
		return "Non admis"
	}
}

// Méthode FicheEtudiant
func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"Étudiant : %s\nAge : %d\nEmail : %s\nPromo : %s\nMoyenne : %.2f\nMention : %s",
		e.NomComplet(),
		e.Age,
		e.Email,
		e.Promo,
		e.Moyenne,
		e.MentionObtenue(),
	)
}

func main() {
	// Création de 2 employés
	employe1 := Employe{
		Personne: Personne{
			Prenom: "Rayan",
			Nom:    "Bakhouche",
			Age:    24,
			Email:  "rayan@example.com",
		},
		Adresse: Adresse{
			Rue:        "10 rue de Lyon",
			Ville:      "Lyon",
			CodePostal: "69000",
		},
		Poste:   "Développeur web",
		Salaire: 2500,
	}

	employe2 := Employe{
		Personne: Personne{
			Prenom: "Sarah",
			Nom:    "Martin",
			Age:    29,
			Email:  "sarah@example.com",
		},
		Adresse: Adresse{
			Rue:        "5 avenue Victor Hugo",
			Ville:      "Paris",
			CodePostal: "75000",
		},
		Poste:   "Cheffe de projet",
		Salaire: 3200,
	}

	// Création de 2 étudiants
	etudiant1 := Etudiant{
		Personne: Personne{
			Prenom: "Lucas",
			Nom:    "Durand",
			Age:    21,
			Email:  "lucas@example.com",
		},
		Promo:   "M2 Informatique",
		Moyenne: 15.5,
	}

	etudiant2 := Etudiant{
		Personne: Personne{
			Prenom: "Emma",
			Nom:    "Bernard",
			Age:    22,
			Email:  "emma@example.com",
		},
		Promo:   "M2 Informatique",
		Moyenne: 17.2,
	}

	// Augmentation du salaire du premier employé
	employe1.AugmenterSalaire(10)

	// Affichage des fiches
	fmt.Println("----- FICHE EMPLOYÉ 1 -----")
	fmt.Println(employe1.FicheEmploye())

	fmt.Println("\n----- FICHE EMPLOYÉ 2 -----")
	fmt.Println(employe2.FicheEmploye())

	fmt.Println("\n----- FICHE ÉTUDIANT 1 -----")
	fmt.Println(etudiant1.FicheEtudiant())

	fmt.Println("\n----- FICHE ÉTUDIANT 2 -----")
	fmt.Println(etudiant2.FicheEtudiant())
}
