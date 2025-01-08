package opti

import (
	"fmt"
	"strings"
)

// Fonction pour enlever les caractères indésirables (_ , -, *)
func Enlever_Caracteres_Dates(TabStr []string, str2 string) ([]string, []string) {
	var resultat []string
	var resultat2 []string
	// Parcours de chaque élément du tableau
	for _, str := range TabStr {
		// Remplacer les caractères "_" et "-" et "*" par rien (en les enlevant)
		str_Traiter := strings.ReplaceAll(str, "_", "")
		str_Traiter = strings.ReplaceAll(str_Traiter, "-", "")
		str_Traiter = strings.ReplaceAll(str_Traiter, "*", "")
		resultat = append(resultat, str_Traiter)
	}
	// Traitement du deuxième string (str2)
	str_result := strings.ReplaceAll(str2, "_", "")
	str_result = strings.ReplaceAll(str_result, "-", "")
	str_result = strings.ReplaceAll(str_result, "*", "")
	resultat2 = append(resultat2, str_result)

	// Retourner les deux résultats
	return resultat, resultat2
}

// Fonction pour formater les dates sous la forme "jj/mm/aaaa"
func Triage_Date(dates []string) []string {
	var result []string

	// Pour chaque date dans le tableau
	for _, date := range dates {
		// Vérifier si la date est bien de 8 chiffres (elle sera utilisée après Enlever_Caracteres)
		if len(date) == 8 {
			jour := date[:2]  // Les deux premiers caractères (jour)
			mois := date[2:4] // Les deux suivants (mois)
			annee := date[4:] // Les quatre derniers (année)

			// Formater la date sous forme "jj/mm/aaaa"
			formattedDate := fmt.Sprintf("%s/%s/%s", jour, mois, annee)

			// Ajouter la date formatée au tableau résultat
			result = append(result, formattedDate)
		} else {
			// Si la date n'est pas correcte, ajouter la date brute au tableau
			result = append(result, date)
		}
	}

	// Retourner les dates formatées
	return result
}
