package api

import (
	"fmt"
	"strings"
	"unicode"
)

// ----------------------------------------------------------------------------------------------------------------------

// Fonction pour enlever les caractères indésirables (_ , -)
func EnleverCaracteresLocations(TabStr []string) []string {
	var resultat []string

	// Parcours de chaque élément du tableau
	for _, str := range TabStr {
		// Remplacer les caractères "_" par des espaces
		str = strings.ReplaceAll(str, "_", " ")
		// Remplacer les caractères "-" par ", "
		str = strings.ReplaceAll(str, "-", " • ")
		// Ajouter l'élément modifié au résultat
		resultat = append(resultat, str)
	}

	// Retourner les lieux nettoyés
	return resultat
}

// ----------------------------------------------------------------------------------------------------------------------

// MettreMajuscule : Applique les deux transformations (après espace et virgule) sur chaque chaîne.
func MettreMajuscule(tableauStr []string) []string {
	var resultat []string

	// Parcours de chaque chaîne dans le tableau
	for _, chaine := range tableauStr {
		// Appliquer la transformation après un espace
		chaine = MettreMajusculeApresEspace(chaine)

		// Ajouter la chaîne transformée à la liste des résultats
		resultat = append(resultat, chaine)
	}

	return resultat
}

// ----------------------------------------------------------------------------------------------------------------------

// MettreMajusculeApresEspace : Met en majuscule la lettre après chaque espace.
func MettreMajusculeApresEspace(chaine string) string {
	var nouvelleChaine string

	// Parcours de chaque caractère de la chaîne
	for i, caractere := range chaine {
		// Si c'est le premier caractère ou après un espace, mettre la lettre en majuscule
		if i == 0 || chaine[i-1] == ' ' {
			// Ajouter la lettre en majuscule si c'est une lettre
			if unicode.IsLetter(caractere) {
				nouvelleChaine += string(unicode.ToUpper(caractere))
			} else {
				nouvelleChaine += string(caractere)
			}
		} else {
			// Sinon ajouter la lettre sans modification
			nouvelleChaine += string(caractere)
		}
	}

	return nouvelleChaine
}

func Stocker_Lieux_Formates() {
	// Vérifier s'il y a des lieux à traiter dans IndexLocations
	if len(IndexLocations.Locations) == 0 {
		// Pas de lieux disponibles dans l'API
		return
	}

	// Initialiser la map pour stocker les lieux par ID
	VarFonctions.LOCATIONS_FINALES = make(map[int][]string)

	// Pour chaque entrée dans IndexLocations, on traite les lieux
	for _, location := range IndexLocations.Locations {
		// Appliquer EnleverCaracteresLocations pour enlever les caractères indésirables
		TabLieuxModifier := EnleverCaracteresLocations(location.Locations)

		// Appliquer MettreMajuscule pour mettre en majuscule les lieux
		LieuxAvecMajuscule := MettreMajuscule(TabLieuxModifier)

		// Ajouter les lieux formatés dans la map sous la clé de l'ID correspondant
		VarFonctions.LOCATIONS_FINALES[location.Id] = LieuxAvecMajuscule
	}
}

// ----------------------------------------------------------------------------------------------------------------------

// Afficher_Lieux_sans_maj : Affiche les lieux formatés pour chaque index dans IndexLocations.
func Afficher_Lieux_sans_maj() {
	for i, location := range IndexLocations.Locations {
		/* Pour afficher uniquement le nombre d'éléments souhaités */
		if i >= ElementAfficher {
			break
		}

		// Affiche l'ID
		fmt.Println("\n", "ID: ", location.Id)

		// Affiche les lieux dans Locations
		fmt.Println("Locations :")
		// Appel de la fonction pour nettoyer et formater les lieux
		formattedLocations := EnleverCaracteresLocations(location.Locations)

		// Affichage des lieux formatés
		for j, lieu := range formattedLocations {
			fmt.Printf("%d  .  %s\n", j, lieu)
		}
	}
}

// ----------------------------------------------------------------------------------------------------------------------

// Afficher_Lieux_avec_maj : Affiche les lieux formatés avec mise en majuscule pour chaque index dans IndexLocations.
func Afficher_Lieux_avec_maj() {
	for i, location := range IndexLocations.Locations {
		/* Pour afficher uniquement le nombre d'éléments souhaités */
		if i >= ElementAfficher {
			break
		}

		// Affiche l'ID
		fmt.Println("\n", "ID: ", location.Id)

		// Affiche les lieux dans Locations
		fmt.Println("Locations :")
		// Appel de la fonction pour nettoyer et formater les lieux
		formattedLocations := EnleverCaracteresLocations(location.Locations)

		// Appel de la fonction pour mettre en majuscule les lieux
		lieuxAvecMaj := MettreMajuscule(formattedLocations)

		// Affichage des lieux formatés avec majuscule
		for j, lieu := range lieuxAvecMaj {
			fmt.Printf("%d  .  %s\n", j, lieu)
		}
	}
}
