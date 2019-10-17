package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func combina(words []string) []string {

	response := []string{}

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
					if i1Parola != i2Parola && i2Parola != i3Parola && i1Parola != i3Parola {
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
						if i1Parola != i2Parola && i2Parola != i3Parola && i1Parola != i3Parola && i3Parola != i4Parola && i1Parola != i4Parola && i2Parola != i4Parola {
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

	// 6 parole
	if len(words) >= 6 {
		for i1Parola := 0; i1Parola < len(words); i1Parola++ {
			for i2Parola := 0; i2Parola < len(words); i2Parola++ {
				for i3Parola := 0; i3Parola < len(words); i3Parola++ {
					for i4Parola := 0; i4Parola < len(words); i4Parola++ {
						for i5Parola := 0; i5Parola < len(words); i5Parola++ {
							for i6Parola := 0; i6Parola < len(words); i6Parola++ {
								if i1Parola != i2Parola &&
									i2Parola != i3Parola &&
									i1Parola != i3Parola &&
									i3Parola != i4Parola &&
									i1Parola != i4Parola &&
									i2Parola != i4Parola &&
									i1Parola != i5Parola &&
									i2Parola != i5Parola &&
									i3Parola != i5Parola &&
									i4Parola != i5Parola &&
									i1Parola != i6Parola &&
									i2Parola != i6Parola &&
									i3Parola != i6Parola &&
									i4Parola != i6Parola &&
									i5Parola != i6Parola {
									response = append(response, words[i1Parola]+words[i2Parola]+words[i3Parola]+words[i4Parola]+words[i5Parola]+words[i6Parola])
								}
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

	response := []string{}

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, strings.Title(w))
		}
	}

	return response
}

func sost(w string, m map[string]string) string {
	var response string

	if len(w) > 0 {
		for _, c := range w {

			for cOrig, cSost := range m {
				if cOrig == string(c) {
					c = rune(cSost[0])
				}
			}

			response += string(c)
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
			case 'e', 'E':
				c = '3'
			case 'i', 'I':
				c = '1'
			case 'o', 'O':
				c = '0'
			case 'a', 'A':
				c = '4'
			case 's', 'S':
				c = '5'
			case 'g', 'G':
				c = '6'
			case 'b', 'B':
				c = '8'
			case 'z', 'Z':
				c = '2'
			case 't', 'T':
				c = '7'
			}
			response += string(c)
		}
	}

	return response
}

func leets(words []string) []string {

	response := []string{}

	// LEET A
	var mapA = make(map[string]string)
	mapA["a"] = "4"
	mapA["A"] = "4"

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, sost(w, mapA))
		}
	}

	// LEET A2
	var mapA2 = make(map[string]string)
	mapA2["a"] = "@"
	mapA2["A"] = "@"

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, sost(w, mapA2))
		}
	}

	// LEET E
	var mapE = make(map[string]string)
	mapE["e"] = "3"
	mapE["E"] = "3"

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, sost(w, mapE))
		}
	}

	// LEET I
	var mapI = make(map[string]string)
	mapI["i"] = "1"
	mapI["I"] = "1"

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, sost(w, mapI))
		}
	}

	// LEET I2
	var mapI2 = make(map[string]string)
	mapI2["i"] = "!"
	mapI2["I"] = "!"

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, sost(w, mapI2))
		}
	}

	// LEET O
	var mapO = make(map[string]string)
	mapO["o"] = "0"
	mapO["O"] = "0"

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, sost(w, mapO))
		}
	}

	// LEET U
	var mapU = make(map[string]string)
	mapU["u"] = "v"
	mapU["U"] = "V"

	if len(words) > 0 {
		for _, w := range words {
			response = append(response, sost(w, mapU))
		}
	}



	// LEET TOTALE
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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-------------------------------------------\n")
	fmt.Println(" __  __       _        _____  _      _   ")
	fmt.Println("|  \\/  |     | |      |  __ \\(_)    | |  ")
	fmt.Println("| \\  / | __ _| | _____| |  | |_  ___| |_ ")
	fmt.Println("| |\\/| |/ _` | |/ / _ \\ |  | | |/ __| __|")
	fmt.Println("| |  | | (_| |   <  __/ |__| | | (__| |_ ")
	fmt.Println("|_|  |_|\\__,_|_|\\_\\___|_____/|_|\\___|\\__|")
	fmt.Print("\n")
	fmt.Print("-------------------------------------------\n")
	fmt.Print("| Password permutation and hash MD5, SHA1 |\n")
	fmt.Print("-------------------------------------------\n")

	for c := 0; c < 6; c++ {
		fmt.Print("Insert word (empty to calc)-> ")
		text, _ := reader.ReadString('\n')

		// convert
		text = strings.Replace(text, "\r\n", "", -1)
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("", text) == 0 {
			break
		} else {
			words = append(words, strings.ToLower(text))
		}
	}
	fmt.Println("")

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

	// PREPARA FILE
	fmt.Print("Output filename (default 'output.csv')-> ")
	filename, _ := reader.ReadString('\n')
	filename = strings.Replace(filename, "\r\n", "", -1)
	filename = strings.Replace(filename, "\n", "", -1)

	if strings.Compare("", filename) == 0 {
		filename = "output.csv"
	}
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	// VISUALIZZA
	if len(passwords) > 0 {
		for _, p := range passwords {

			h := md5.New()
			hs := sha1.New()

			io.WriteString(h, p)
			io.WriteString(hs, p)

			str := fmt.Sprintf("%s, %x, %x\n", p, h.Sum(nil), hs.Sum(nil))

			_, err := w.WriteString(str)
			check(err)
		}
	} else {
		fmt.Println("Words list is empty")
	}

	w.Flush()

	fmt.Println("\nBy Archistico - Emilie Rollandin")
}
