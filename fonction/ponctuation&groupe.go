package fonction

func Ponctuation(tableau []string) []string {
	signes := ".,!?:;"

	for i := 1; i < len(tableau); i++ {
		mot := tableau[i]

		nombre := 0
		for nombre < len(mot) {
			trouve := false
			for j := 0; j < len(signes); j++ {
				if mot[nombre] == signes[j] {
					trouve = true
				}
			}
			if !trouve {
				break
			}
			nombre++
		}
		if nombre == 0 {
			continue
		}

		tableau[i-1] += mot[:nombre]
		tableau[i] = mot[nombre:]

		if tableau[i] == "" {
			tableau = append(tableau[:i], tableau[i+1:]...)
			i--
		}
	}
	return tableau
}
