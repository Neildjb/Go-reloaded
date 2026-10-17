package fonction

import "strings"

func Decoupe(phrase string) []string {
	mots := strings.Fields(phrase)
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		for mot[0] == '(' && mot[len(mot)-1] != ')' && i+1 < len(mots) {
			i++
			mot += " " + mots[i]
		}

		resultat = append(resultat, mot)
	}
	return resultat
}
