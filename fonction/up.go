package fonction

import "strings"

func Up(tableau []string) []string {
	for mot := range tableau { // partie recherche
		if mot > 0 && mot < len(tableau) && tableau[mot] == "(up, )" {
			tableau = append(tableau[:mot], tableau[mot+1:]...)
			
		}
	}
}
