package fonction

import (
	"strconv"
	"strings"
)

func Up(tableau []string) []string {
	for i := 0; i < len(tableau); i++ {
		balise := tableau[i]
		nombre := 0

		if balise == "(up)" {
			nombre = 1
		} else if len(balise) > 5 && balise[:4] == "(up," && balise[len(balise)-1] == ')' {
			nombre, _ = strconv.Atoi(strings.TrimSpace(balise[4 : len(balise)-1]))
		} else {
			continue
		}

		for j := i - nombre; j < i; j++ {
			if j >= 0 {
				tableau[j] = strings.ToUpper(tableau[j])
			}
		}
		tableau = append(tableau[:i], tableau[i+1:]...)
		i--
	}
	return tableau
}
