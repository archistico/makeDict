package main

import (
	"fmt"
	"strings"
)

func word2Upper(word string) []string {
	
	response := []string {}

	for iMaiuscola := 0; iMaiuscola < len(word); iMaiuscola++ {
		lista := []string {}
		
		for iCarattere := 0; iCarattere < len(word); iCarattere++ {
			if(iCarattere == iMaiuscola) {
				lista = append(lista, strings.ToUpper(string(word[iCarattere])))
			} else {
				lista = append(lista, string(word[iCarattere]))
			}
		}

		response = append(response, strings.Join(lista,""))
	}

	return response
}


func combina(words []string) []string {
	
	response := []string {}

	// aggiungere for che combina un numero di parole da 2 a len(words)

	for i1Parola := 0; i1Parola < len(words); i1Parola++ {
		for i2Parola := 0; i2Parola < len(words); i2Parola++ {
			if i1Parola != i2Parola {
				response = append(response, words[i1Parola]+words[i2Parola])
			}
		}
	}

	return response
}

func main() {
	var words = []string{}
	words = append(words, strings.ToLower("cane"))
	words = append(words, strings.ToLower("gatto"))
	words = append(words, strings.ToLower("topo"))

	// COMBINAZIONI
	if len(words) > 0 {
		words = append(words, combina(words)...)
	}

	// MAIUSCOLI
	var maiuscoli = []string{}
	if len(words) > 0 {
		for _, w := range words {
			maiuscoli = append(maiuscoli, word2Upper(w)...)
		}
	}
	//words = append(words, maiuscoli...)

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

/*
import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "admin")
	fmt.Printf("%x", h.Sum(nil))
}
*/