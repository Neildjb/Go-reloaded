package fonction

import "strings"

func Low(tableau []string) []string {
	for i := 0; i < len(tableau); i++ {
		if tableau[i] != "(low)" {
			continue
		}
		if i > 0 {
			tableau[i-1] = strings.ToLower(tableau[i-1])
		}
		tableau = append(tableau[:i], tableau[i+1:]...)
		i--
	}
	return tableau
}
