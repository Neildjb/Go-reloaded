package fonction

func Ponctuation(tableau []string) []string {
	ponc := []string{".", ",", "!", "?", ":", ";"}

	for i := 0; i < len(tableau); i++ {
		mot := tableau[i]
		trouve := false

		for _, elt := range ponc {
			if mot == elt {
				trouve = true
			}
		}

		if trouve && i > 0 {
			motAvant := tableau[i-1]
			motAvant = motAvant + mot
			tableau[i-1] = motAvant

			tableau = append(tableau[:i], tableau[i+1:]...)
			i--
		}
	}
	return tableau
}
