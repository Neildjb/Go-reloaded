package fonction

import "strconv"

func Bin(tableau []string) []string {
	for mot := range tableau { 
		if mot > 0 && mot < len(tableau) && tableau[mot] == "(bin)" {
			tableau = append(tableau[:mot], tableau[mot+1:]...)
			valeur, _ := strconv.ParseInt(tableau[mot-1], 2, 64)
			tableau[mot-1] = strconv.Itoa(int(valeur))
		}
	}
	return tableau
}
