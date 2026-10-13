package fonction

import "strings"

func An(tableau []string) []string {
    voyelles := map[string]bool{
        "a": true,
		"e": true,
		"i": true,
        "o": true,
		"u": true,
		"h": true,
    }

    for i := 0; i < len(tableau)-1; i++ {
        if tableau[i] == "a" || tableau[i] == "A" {
            premiere := strings.ToLower(tableau[i+1][:1])
            if voyelles[premiere] {
                tableau[i] += "n"
            }
        }
    }
    return tableau
}