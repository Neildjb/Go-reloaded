package main

import (
	"fmt"
	"os"
	"strings"
	f "travail/fonction"
)

func main() {
	fichier, _ := os.ReadFile(os.Args[1]) //récupération du simple
	contenu := string(fichier)

	tableau := f.Decoupe(contenu)
	fmt.Print(tableau)
	tableau = fonctions(tableau)

	Stockage := strings.Join(tableau, " ") //envoie le fichier
	os.WriteFile("result.txt", []byte(Stockage), 0644)
}

func fonctions(tableau []string) []string {
	tableau = f.Hex(tableau)
	tableau = f.Bin(tableau)
	tableau = f.Up(tableau)
	tableau = f.Low(tableau)
	tableau = f.Cap(tableau)
	tableau = f.An(tableau)
	tableau = f.Apostrophe(tableau)
	tableau = f.Ponctuation(tableau)
	return tableau
}
