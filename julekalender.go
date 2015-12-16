package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	fmt.Println("Start")

	//	oppgave1()
	//	oppgave2()
	//	oppgave3()
	//	oppgave4()
	//	oppgave5()
	//	oppgave6()
	//	oppgave7()
	//	oppgave8()
	//	oppgave9()
	oppgave10()
	// oppgave13()
	elapsed := time.Since(start)
	log.Printf("run took %s", elapsed)
}

func oppgave1() {
	fmt.Println("Oppgave1")
	content, err := ioutil.ReadFile("oppgave1")
	if err != nil {
		//Do something
	}
	numberOfMatches := 0
	lines := strings.Split(string(content), "\n")
	for _, element := range lines {
		matched, _ := regexp.MatchString("^[a-z]{0,3}[0-9]{2,8}[A-Z]{3,}$", element)
		if matched {
			numberOfMatches = numberOfMatches + 1
		}
	}

	fmt.Println("Number of Matches: " + strconv.Itoa(numberOfMatches))
}

func oppgave2() {
	fmt.Println("Oppgave2")
	content, _ := ioutil.ReadFile("oppgave2")
	lines := strings.Split(string(content), "\n")
	var bestProfit float64
	for index, element := range lines {
		var currentPrice, _ = strconv.ParseFloat(element, 64)
		for i := index; i < len(lines); i++ {
			secondPrice, _ := strconv.ParseFloat(lines[i], 64)
			currentProfit := secondPrice - currentPrice
			if currentProfit > bestProfit {
				bestProfit = currentProfit
			}
		}
	}
	bestProfitStr := strconv.FormatFloat(bestProfit, 'f', 4, 64)
	fmt.Println("Best profit:" + bestProfitStr)
}

func findBest(lines []string, f func(f float64)) {

	var leftSide []string
	for i := 0; i < len(lines)/2; i++ {
		leftSide = append(leftSide, lines[i])
	}
	if len(leftSide) < 4 {
		return
	} else {
		f(findBestProfit(leftSide))
		findBest(leftSide, f)
	}
	var rightSide []string
	for i := (len(lines) / 2); i < len(lines); i++ {
		rightSide = append(rightSide, lines[i])
	}
	if len(rightSide) < 4 {
		return
	} else {
		f(findBestProfit(rightSide))
		findBest(rightSide, f)
	}

	return
}

func oppgave3() {
	fmt.Println("Oppgave3")
	currentDay := 0
	numberOfOccurrences := 0
	numberOfDays := 365
	numberOfYears := 2015
	for year := 0; year < numberOfYears; year++ {
		if year%4 == 0 && year%100 == 1 || year%400 == 0 {
			numberOfDays = 366
		} else {
			numberOfDays = 265
		}
		for day := 0; day < numberOfDays; day++ {
			if currentDay != 6 {
				currentDay = currentDay + 1
			} else {
				currentDay = 0
			}
			if currentDay == 0 && day == 256 {
				numberOfOccurrences = numberOfOccurrences + 1
			}
		}
	}
	fmt.Println("Number of occurrances " + strconv.Itoa(numberOfOccurrences))
}

func oppgave4() {
	fmt.Println("Oppgave3")
	content, _ := ioutil.ReadFile("oppgave4")
	m := make(map[string]int)
	for _, element := range content {
		if !strings.Contains(string(element), "\n") {
			m[string(element)] = m[string(element)] + 1
		}
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", strconv.Itoa(m[key]))
	}
}

func oppgave5() {
	content, _ := ioutil.ReadFile("oppgave5")
	lines := strings.Split(string(content), "\n")
	m := make(map[string]int)
	for index, element := range lines {
		for i := index; i < len(lines); i++ {
			if !strings.EqualFold(element, lines[i]) {
				if len(element) == len(lines[i]) {
					word1 := sorted(element)
					word2 := sorted(lines[i])
					match := strings.EqualFold(word1, word2)
					if match {
						m[element] = 1
						m[lines[i]] = 1
					}
				}
			}
		}
	}
	fmt.Println("Number of matches: " + strconv.Itoa(len(m)))
}

func oppgave6() {
	n := 13
	matches := 0
	hit := func() {
		matches = matches + 1
	}
	findParans(0, n, 0, 0, hit)
	fmt.Println("Number of matches: " + strconv.Itoa(matches))
}

func oppgave7() {
	sum := 0
	for i := 0; i < 1000; i++ {
		reverse, _ := strconv.Atoi(Reverse(strconv.Itoa(i)))
		if i%7 == 0 && reverse%7 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("Number of matches: " + strconv.Itoa(sum))
}

func oppgave8() {
	numberOfHits := 0
	for i := 0; i < 1000; i++ {
		reverse, _ := strconv.Atoi(Reverse(strconv.Itoa(i)))
		if reverse != i {
			if big.NewInt(int64(i)).ProbablyPrime(1) && big.NewInt(int64(reverse)).ProbablyPrime(1) {
				numberOfHits++
			}
		}
	}
	fmt.Println("Number of matches: " + strconv.Itoa(numberOfHits))
}

func oppgave9() {
	columntitle := make([]rune, 0)
	num := 142453146368
	for num > 0 {
		num-- // 1 => a, not 0 => a
		remainder := num % 26
		digit := (remainder + 97)
		columntitle = append(columntitle, rune(digit))
		num = (num - remainder) / 26
	}
	fmt.Println(Reverse(string(columntitle)))
}

func oppgave10() {
	fmt.Println("Oppgave10")
	content, _ := ioutil.ReadFile("oppgave2")
	lines := strings.Split(string(content), "\n")
	profits := make([]float64, 0)
	findBest(lines, func(stuff float64) { profits = append(profits, stuff) })
	fmt.Println(profits)
}
func oppgave12() {
	answer := 0
	for i := 0; i < 100000000; i++ {
		if i%7 == 0 && i%5 > 0 {
			answer = answer + i
		}
	}
	fmt.Println(answer)
}

// Generate numbers until the limit max.
// after the 2, all the prime numbers are odd
// Send a channel signal when the limit is reached
func Generate(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 // signal that the limit is reached
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		if i%prime != 0 {
			out <- i
		}
	}
	out <- -1
}

func CalcPrimeFactors(number_to_factorize int) []int {
	rv := []int{}
	ch := make(chan int)
	go Generate(number_to_factorize, ch)
	for prime := <-ch; (prime != -1) && (number_to_factorize > 1); prime = <-ch {
		for number_to_factorize%prime == 0 {
			number_to_factorize = number_to_factorize / prime
			rv = append(rv, prime)
		}
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return rv
}

func oppgave13() {

	hits := 0
	theNumber := 0
	found := false
	for i := 1; !found; i++ {
		hit := false
		if i%2 == 0 || i%3 == 0 || i%5 == 0 {
			hit = true
			//						if (len(CalcPrimeFactors(i)) > 3) {
			//							hit = false
			//						}
			if i > 6 {
				for y := 6; y < i; y++ {
					if i%y == 0 && big.NewInt(int64(y)).ProbablyPrime(1) {
						hit = false
						break
					}
				}
			}
		}
		if hit {
			hits++
			theNumber = i
			fmt.Println("Hits: " + strconv.Itoa(hits))
			fmt.Println("The number: " + strconv.Itoa(theNumber))
		}
		if hits == 10000 {
			found = true
		}

		//				if (contains(factors, 2) && !contains(factors, 3) && !contains(factors, 5)) {
		//					hits++
		//					theNumber = i
		//				}
		//				if (!contains(factors, 2) && contains(factors, 3) && !contains(factors, 5)) {
		//					hits++
		//					theNumber = i
		//				}
		//				if (!contains(factors, 2) && !contains(factors, 3) && contains(factors, 5)) {
		//					hits++
		//					theNumber = i
		//				}

		//		if (i == 1) {
		//			hits++
		//			theNumber = i
		//		}

		//		if (i % 2 == 0 && i % 3 != 0 && i % 5 != 0) {
		//			hits++
		//			theNumber = i
		//		}
		//		if (i % 3 == 0 && i % 2 > 0 && i % 5 != 0) {
		//			hits++
		//			theNumber = i
		//		}
		//		if ( i % 5 == 0 && i % 2 > 0 && i % 5 > 0) {
		//			hits++
		//			theNumber = i
		//		}
	}

	fmt.Println("The final number: " + strconv.Itoa(theNumber))
}

func findBestProfit(lines []string) float64 {
	var bestProfit float64 = 0
	for index, element := range lines {
		var currentPrice, _ = strconv.ParseFloat(element, 64)
		for i := index; i < len(lines); i++ {
			secondPrice, _ := strconv.ParseFloat(lines[i], 64)
			currentProfit := secondPrice - currentPrice
			if currentProfit > bestProfit {
				bestProfit = currentProfit
			}
		}
	}
	return bestProfit
}

type floatSlice []float64

func (slice floatSlice) pos(value float64) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findParans(pos int, n int, open int, close int, f func()) {
	if close == n {
		f()
		return
	} else {
		if open > close {
			findParans(pos+1, n, open, close+1, f)
		}
		if open < n {
			findParans(pos+1, n, open+1, close, f)
		}
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
