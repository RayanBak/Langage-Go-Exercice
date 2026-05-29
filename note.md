# Langage-Go-note

Go = langage simple, rapide et performant.
L’idée est de faire un compromis entre performance, simplicité et rapidité de dev.

Go est souvent utilisé pour :
microservices
API REST
CLI
DevOps
cloud / infra

Rappel sur les types

Go utilise des types classiques :
int
float64
string
bool

Rien de spécial ici, c’est globalement déjà maîtrisé.

Variables

Déclaration classique :
var age int = 24

Déclaration courte :
nom := "Reda"

Double déclaration :
x, y := 10, 30

On peut aussi déclarer plusieurs valeurs de types différents :
nom, age := "Reda", 24

La déclaration courte `:=` est très utilisée dans les fonctions.

Constantes

const = valeur fixe qui ne change pas.

Exemple :
const TVA = 20

iota

iota permet de générer automatiquement des valeurs dans un bloc const.

Exemple :
const (
Debutant = iota
Intermediaire
Avance
)

Résultat :
Debutant = 0
Intermediaire = 1
Avance = 2

Array / tableau

Un tableau = array.

C’est comme dans les autres langages.
Sa taille est fixe.

Exemple :
notes := [3]int{12, 15, 18}

Slice

Un slice = tableau de taille modifiable.

C’est plus souple qu’un array.

Exemple :
notes := []int{12, 15, 18}
notes = append(notes, 20)

Donc :
array = taille fixe
slice = taille dynamique

Map

Une map = dictionnaire comme en Python.

C’est un système clé / valeur.

Exemple :
notes := map[string]int{
"Reda": 15,
"Rayan": 16,
}

Pour vérifier si une clé existe :
valeur, ok := notes["Reda"]

Fonctions

En Go, une fonction peut retourner plusieurs valeurs.

Souvent :
résultat + erreur

Exemple :
resultat, err := diviser(10, 2)

Il n’y a pas de surcharge de fonction en Go.
Donc on ne peut pas créer plusieurs fonctions avec le même nom mais des paramètres différents.

Fonction variadique

Une fonction variadique accepte un nombre variable d’arguments.

Exemple :
func addition(nombres ...int) {
// plusieurs nombres possibles
}

Erreurs

En Go, on gère souvent les erreurs avec `error`.

Si err != nil, il y a une erreur.

Exemple :
if err != nil {
fmt.Println("Erreur :", err)
}

nil = absence de valeur ou absence d’erreur.

for

En Go, il n’y a qu’un seul mot-clé pour les boucles : for.

Il remplace :
while
do while
foreach

Switch

Le switch permet de tester plusieurs cas.

En Go, pas besoin de break.
Le switch s’arrête tout seul.

fallthrough

Pour forcer l’exécution du cas suivant, Go utilise le mot-clé `fallthrough`.

Il permet de continuer dans le case suivant même si sa condition n’est pas vérifiée.

On l’utilise rarement.

Pointeurs

Un pointeur pointe vers l’emplacement d’une donnée en mémoire.

& récupère l’adresse.

* récupère la valeur.

Exemple :
x := 10
p := &x

Les pointeurs servent surtout à modifier directement une valeur ou une struct sans faire de copie.

Struct

Une struct permet de regrouper plusieurs données.

Exemple :
type Personne struct {
Nom string
Age int
}

OOP en Go

Go permet une approche orientée objet, mais pas comme Java ou PHP.

Il n’y a pas de classes classiques.
On utilise :
structs
méthodes
embedding

Méthodes

On peut attacher une méthode à une struct.

Exemple :
func (p Personne) Presentation() string {
return p.Nom
}

Si on veut modifier la struct, on utilise un pointeur :

func (p *Personne) ModifierNom(nom string) {
p.Nom = nom
}

Pas d’héritage

Il n’y a pas d’héritage en Go.

À la place, Go utilise la composition avec l’embedding.

Exemple :
type Employe struct {
Personne
Poste string
}

Visibilité

La visibilité est déterminée par la casse, pas par des mots-clés public/private.

Majuscule = exporté / accessible depuis un autre package.
Minuscule = non exporté / privé au package.

Exemple :
Produit → accessible
produit → non accessible depuis un autre package

defer

defer permet d’exécuter une instruction à la fin d’une fonction.

Utile pour :
fermer un fichier
libérer une ressource
nettoyer après une opération

Si plusieurs defer sont utilisés, le dernier est exécuté en premier.

Principe :
LIFO = Last In, First Out

Packages utiles

fmt :
affichage et formatage

strings :
traitement des textes

sort :
tri des slices / tableaux

math :
calculs mathématiques

JSON

Marshal :
struct → JSON

Unmarshal :
JSON → struct

Les tags permettent de contrôler le JSON :

Nom string `json:"nom,omitempty"`

omitempty = ne pas afficher si la valeur est vide.
json:"-" = ignorer le champ.

À retenir

Go repose surtout sur :
types
variables
constantes
iota
arrays
slices
maps
fonctions
erreurs
structs
méthodes
pointeurs
embedding
defer

Pas de surcharge de fonction.
Pas d’héritage classique.
La visibilité dépend de la casse.
Les slices sont dynamiques.
Les maps fonctionnent comme des dictionnaires.
