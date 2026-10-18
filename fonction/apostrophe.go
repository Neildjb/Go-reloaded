package fonction

func Apostrophe(tableau []string) []string {
	ouvert := false

	for i := 0; i < len(tableau); i++ {
		if tableau[i] != "'" {
			continue
		}

		if ouvert == false {
			tableau[i+1] = "'" + tableau[i+1]
			ouvert = true
		} else {
			tableau[i-1] = tableau[i-1] + "'"
			ouvert = false
		}

		tableau = append(tableau[:i], tableau[i+1:]...)
		i--
	}
	return tableau
}
