package fonction

import (
	"strconv"
	"strings"
)

func Low(tableau []string) []string {
	for i := 0; i < len(tableau); i++ {
		balise := tableau[i]
		nombre := 0

		if balise == "(low)" {
			nombre = 1
		} else if len(balise) > 6 && balise[:5] == "(low," && balise[len(balise)-1] == ')' {
			nombre, _ = strconv.Atoi(strings.TrimSpace(balise[5 : len(balise)-1]))
		} else {
			continue
		}

		for j := i - nombre; j < i; j++ {
			if j >= 0 {
				tableau[j] = strings.ToLower(tableau[j])
			}
		}
		tableau = append(tableau[:i], tableau[i+1:]...)
		i--
	}
	return tableau
}
