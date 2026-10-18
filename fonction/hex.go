package fonction

import "strconv"

func Hex(tableau []string) []string {
	for mot := range tableau {
		if mot > 0 && mot < len(tableau) && tableau[mot] == "(hex)" {
			tableau = append(tableau[:mot], tableau[mot+1:]...)
			valeur, _ := strconv.ParseInt(tableau[mot-1], 16, 64)
			tableau[mot-1] = strconv.Itoa(int(valeur))
		}
	}
	return tableau
}
