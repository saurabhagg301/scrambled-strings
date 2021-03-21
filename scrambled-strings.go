package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

// global variables
var (
	charToInteger = map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12,
		"m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23,
		"x": 24, "y": 25, "z": 26}
	pathToDictFile  = flag.String("dictionary", "dict/dict.txt", "path to dictionary file")
	pathToInputFile = flag.String("input", "input/input.txt", "path to input file")
	output          = []result{}
)

func init() {
	// If the file doesn't exist, create it
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

}

type result struct {
	LineNum int `json:"lineNum"`
	Count   int `json:"count"`
}

func main() {
	// parse input flag
	flag.Parse()

	// validate dictionary file
	errValidate := validateDictionaryFile(*pathToDictFile)
	if errValidate != nil {
		fmt.Println("error:", errValidate)
		return
	}

	// populate dict slice with dictionary word's signature as per below:
	// 2 dimentional slice of n rows and 5 columns
	// 1st column = starting char integer of dictionary word
	// 2nd column = ending character interger of dictionary word,
	// 3rd column = number of characters in between
	// 4th column = sum of chars(integer against character as per charToInteger map) between starting and ending char
	// 5th column = counter value to be incremented by (this is to handle same signature words in dictionary file)
	// same signature word means words having same starting, last characters, and middle characters but is different order. For example, 'axpaj' and 'apxaj'
	dict := [][5]int{}

	fDict, err := os.Open(*pathToDictFile)
	defer fDict.Close()

	check(err)

	scanner := bufio.NewScanner(fDict)
	var dictStrArr []string

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		temp := [5]int{}
		sum := 0
		word := scanner.Text()
		dictStrArr = append(dictStrArr, word)
		for k, v := range word {
			num := charToInteger[string(v)]
			if k == 0 {
				temp[0] = num
			} else if k == len(word)-1 {
				temp[1] = num
			} else {
				sum += num
			}
		}
		temp[2] = len(word) - 2 // number of character between starting and ending character
		temp[3] = sum           // sum of integer equivalent of middle letters

		// check if a same signature word already exists in the dict slice
		var flagSameSignatureWordExists bool
		for k, v := range dict {
			if v[0] == temp[0] && v[1] == temp[1] && v[2] == temp[2] && v[3] == temp[3] {
				dict[k][4] += 1
				flagSameSignatureWordExists = true
				break
			}
		}
		if !flagSameSignatureWordExists {
			temp[4] = 1
			dict = append(dict, temp)
		}
	}

	// create map of dictionary
	// key: integer corresponding to starting character/letter
	// value: all words(its signature) starting with this particular character
	dictMap := map[int][][5]int{}
	for _, v := range dict {
		dictMap[v[0]] = append(dictMap[v[0]], v)

	}

	printDictMap(dictMap) // this will print dictionary map to log file
	// --------------x------------------------------x---------------------------

	f, err := os.Open(*pathToInputFile)
	defer f.Close()

	check(err)

	scanner = bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	lineNum := 0
	for scanner.Scan() {
		// To handle - The scrambled or original form of the dictionary word may appear multiple times but we only count it once since we only need to know whether it shows up at least once.
		wordOccured := [][5]int{}
		counter := 0
		line := scanner.Text()
		lineNum++
		for k, v := range line {
			num, exist := charToInteger[string(v)]
			if !exist {
				fmt.Printf("INFO: invalid character '%s' found at Line# %d - only lowercase characters[a-z] allowed\n", string(v), lineNum)
			} else {
				// here arr is the slice of words(its signature) available in dictionary, corresponding to the current letter in the input line
				arr, charExistsInDictionary := dictMap[num]
				if charExistsInDictionary {
					// here v2 is the signature of one of the word(from dictionary) corresponding to the current letter in the input line
					// check the last letter and number of letter in between from the  signature of dictionary word and see if there is a letter = last letter at the distance = number of letters in between
					// if yes, check it the sum of letters in between in input line = sum of letters in between in word (signature)
					for _, v2 := range arr {
						if len(line[k:]) > v2[2]+1 {
							if v2[1] == charToInteger[string(line[k+v2[2]+1])] {
								sum := 0
								word := string(line[k])
								for i := 1; i <= v2[2]; i++ {
									sum += charToInteger[string(line[k+i])]
									word += string(line[k+i])
								}
								word += string(line[k+v2[2]+1])
								exists := checkElementExistsInSlice(v2, wordOccured)

								if !exists && sum == v2[3] {
									counter += v2[4]
									wordOccured = append(wordOccured, v2)
									log.Printf("INFO: Line# %d, Word matched: %s, Counter Value incremented by: %d\n", lineNum, word, v2[4])
								}
							}
						}
					}
				}
			}
		}

		fmt.Printf("Case #%d: %d\n", lineNum, counter)
		log.Printf("Case #%d: %d\n", lineNum, counter)
		output = append(output, result{lineNum, counter})

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// checkElementExistsInStringSlice to check if element s exists in input string array
func checkElementExistsInStringSlice(s string, arr []string) bool {
	var res bool
	for _, v := range arr {
		if v == s {
			res = true
			break
		}
	}

	return res
}

// checkElementExistsInStringSlice to check if element s exists in input string array
func checkElementExistsInSlice(s [5]int, arr [][5]int) bool {
	var res bool
	for _, v := range arr {
		if v == s {
			res = true
			break
		}
	}

	return res
}

// validateDictionaryFile to validate dictionary file
func validateDictionaryFile(pathToFile string) error {
	fDict, err := os.Open(pathToFile)
	defer fDict.Close()

	check(err)

	scanner := bufio.NewScanner(fDict)
	dictStrArr := []string{}

	scanner.Split(bufio.ScanLines)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		word := scanner.Text()
		if len(word) < 2 || len(word) > 105 {
			// return error
			return errors.New(fmt.Sprintf("Error in Line #%d. Word in the dictionary should be between 2 and 105 letters long(inclusive)", lineNum))
		}
		for _, v := range word {
			if v < 97 || v > 122 {
				// 97 - rune equivalent for a
				// 122 - rune equivalent for z
				// return error
				return errors.New(fmt.Sprintf("Error in Line #%d. Word in the dictionary should contain only lowercase characters [a-z]", lineNum))
			}
		}

		exists := checkElementExistsInStringSlice(word, dictStrArr)
		if exists {
			// return error
			return errors.New(fmt.Sprintf("Error in Line #%d. No two words in the dictionary should be same", lineNum))
		} else {
			dictStrArr = append(dictStrArr, word)
		}

	}
	sum := 0
	for _, word := range dictStrArr {
		sum += len(word)

	}
	if sum > 105 {
		// return error
		return errors.New("The sum of lengths of all words in the dictionary should not exceed 105")
	}

	// return success
	return nil
}

// printDictMap to print dictionary
func printDictMap(dictMap map[int][][5]int) {
	log.Println("INFO: Dictionary Map: ")
	for k, v := range dictMap {
		var c string
		for char, integerValueOfChar := range charToInteger {
			if integerValueOfChar == k {
				c = char
				break
			}
		}
		log.Printf("%s:%+v\n", c, v)
	}
}
