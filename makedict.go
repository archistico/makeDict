package main

import (
	"fmt"
	"strings"
)

func word2Upper(word string) {
	
	for iMaiuscola := 0; iMaiuscola < len(word); iMaiuscola++ {
		lista := []string {}
		
		for iCarattere := 0; iCarattere < len(word); iCarattere++ {
			if(iCarattere == iMaiuscola) {
				lista = append(lista, strings.ToUpper(string(word[iCarattere])))
			} else {
				lista = append(lista, string(word[iCarattere]))
			}
		}

		fmt.Println("Parola maiuscola:", strings.Join(lista,""))
	}
}

func main() {
	var words = []string{}
	words = append(words, "emi")
	//words = append(words, "rollandin")

	var maiuscoli = []string{}
	if len(words) > 0 {
		for _, w := range words {
			word2Upper(w)
		}
	}
	words = append(words, maiuscoli...)


	// VISUALIZZA 

	if len(words) > 0 {
		for _, w := range words {
			fmt.Println("Word:", w)
		}
	} else {
		fmt.Println("Words list is empty")
	}

	fmt.Println("-------- END ---------")
}
