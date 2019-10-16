package main

import (
	"fmt"
	"strings"
	"crypto/md5"
	"io"
)

func combina(words []string) []string {
	
	response := []string {}

	// aggiungere for che combina un numero di parole da 2 a len(words)

	if len(words) >= 2 {
		for i1Parola := 0; i1Parola < len(words); i1Parola++ {
			for i2Parola := 0; i2Parola < len(words); i2Parola++ {
				if i1Parola != i2Parola {
					response = append(response, words[i1Parola]+words[i2Parola])
				}
			}
		}
	}

	// 3 parole
	if len(words) >= 3 {
		for i1Parola := 0; i1Parola < len(words); i1Parola++ {
			for i2Parola := 0; i2Parola < len(words); i2Parola++ {
				for i3Parola := 0; i3Parola < len(words); i3Parola++ {
					if i1Parola != i2Parola && i2Parola != i3Parola && i1Parola != i3Parola  {
						response = append(response, words[i1Parola]+words[i2Parola]+words[i3Parola])
					}
				}
			}
		}
	}

	// 4 parole
	if len(words) >= 4 {
		for i1Parola := 0; i1Parola < len(words); i1Parola++ {
			for i2Parola := 0; i2Parola < len(words); i2Parola++ {
				for i3Parola := 0; i3Parola < len(words); i3Parola++ {
					for i4Parola := 0; i4Parola < len(words); i4Parola++ {
						if i1Parola != i2Parola && i2Parola != i3Parola && i1Parola != i3Parola && i3Parola != i4Parola && i1Parola != i4Parola && i2Parola != i4Parola  {
							response = append(response, words[i1Parola]+words[i2Parola]+words[i3Parola]+words[i4Parola])
						}
					}
				}
			}
		}
	}

	// 5 parole
	if len(words) >= 5 {
		for i1Parola := 0; i1Parola < len(words); i1Parola++ {
			for i2Parola := 0; i2Parola < len(words); i2Parola++ {
				for i3Parola := 0; i3Parola < len(words); i3Parola++ {
					for i4Parola := 0; i4Parola < len(words); i4Parola++ {
						for i5Parola := 0; i5Parola < len(words); i5Parola++ {
							if i1Parola != i2Parola &&
								i2Parola != i3Parola &&
								i1Parola != i3Parola &&
								i3Parola != i4Parola &&
								i1Parola != i4Parola &&
								i2Parola != i4Parola &&
								i1Parola != i5Parola &&
								i2Parola != i5Parola &&
								i3Parola != i5Parola &&
								i4Parola != i5Parola {
								response = append(response, words[i1Parola]+words[i2Parola]+words[i3Parola]+words[i4Parola]+words[i5Parola])
							}
						}
					}
				}
			}
		}
	}

	return response
}

func word2Title(words []string) []string {

	response := []string {}

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, strings.Title(w))
		}
	}

	return response
}

func leet(w string) string {
	var response string

	/*
	qwertyuiopasdfghjklzxcvbnm
	qw3r7yu10p45df6hjklzxcv8nm
	*/

	if len(w) > 0 {
		for _, c := range w {
			switch c {
				case 'e', 'E': c = '3'
				case 'i', 'I': c = '1'
				case 'o', 'O': c = '0'
				case 'a', 'A': c = '4'
				case 's', 'S': c = '5'
				case 'g', 'G': c = '6'
				case 'b', 'B': c = '8'
				case 'z', 'Z': c = '2'
				case 't', 'T': c = '7'
			}
			response += string(c)
		}
	}

	return response
}

func leets(words []string) []string {

	response := []string {}

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, leet(w))
		}
	}

	return response
}


func main() {
	var words = []string{}
	var passwords = []string{}

	words = append(words, strings.ToLower("my"))
	words = append(words, strings.ToLower("secret"))
	words = append(words, strings.ToLower("number"))
	words = append(words, strings.ToLower("is"))
	words = append(words, strings.ToLower("42"))

	// Singoli
	passwords = append(passwords, words...)

	// COMBINAZIONI
	if len(words) > 0 {
		passwords = append(passwords, combina(words)...)
	}

	// TITLE MAIUSCOLA
	if len(words) > 0 {
		var wordsTitle = []string{}
		wordsTitle = append(wordsTitle, word2Title(words)...)
		passwords = append(passwords, wordsTitle...)
		passwords = append(passwords, combina(wordsTitle)...)
	}

	// LEET BASIC
	if len(words) > 0 {
		passwords = append(passwords, leets(passwords)...)
	}

	// VISUALIZZA

	if len(passwords) > 0 {
		for _, p := range passwords {

			h := md5.New()
			io.WriteString(h, p)

			fmt.Printf("%s, %x\n", p, h.Sum(nil))
		}
	} else {
		fmt.Println("Words list is empty")
	}
}

