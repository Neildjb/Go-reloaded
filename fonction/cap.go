package fonction

import (
	"strconv"
	"strings"
)

func Cap(tableau []string) []string {
	for i := 0; i < len(tableau); i++ {
		nombre := 0 // mots avants

		if tableau[i] == "(cap)" {
			nombre = 1
		} else if len(tableau[i]) > 6 && tableau[i][:5] == "(cap," {
			nombre, _ = strconv.Atoi(strings.Trim(tableau[i][5:], " )"))
		} else {
			continue // pas une balise
		}

		// maj 1er lettre
		for j := max(i-nombre, 0); j < i; j++ {
			mot := tableau[j]
			if mot != "" && mot[0] >= 'a' && mot[0] <= 'z' {
				tableau[j] = string(mot[0]-32) + mot[1:]
			}
		}

		tableau = append(tableau[:i], tableau[i+1:]...) // retire
		i--
	}
	return tableau
}
