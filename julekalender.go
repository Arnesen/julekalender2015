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
	// oppgave10()
	// oppgave18()
	// oppgave19()
	// oppgave20()

	oppgave22()

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
	for i := 0; i < len(lines) / 2; i++ {
		leftSide = append(leftSide, lines[i])
	}
	if len(leftSide) < 4 {
		return
	}
	f(findBestProfit(leftSide))
	findBest(leftSide, f)

	var rightSide []string
	for i := (len(lines) / 2); i < len(lines); i++ {
		rightSide = append(rightSide, lines[i])
	}
	if len(rightSide) < 4 {
		return
	}
	f(findBestProfit(rightSide))
	findBest(rightSide, f)
	return
}

func oppgave3() {
	fmt.Println("Oppgave3")
	currentDay := 0
	numberOfOccurrences := 0
	numberOfDays := 365
	numberOfYears := 2015
	for year := 0; year < numberOfYears; year++ {
		if year % 4 == 0 && year % 100 == 1 || year % 400 == 0 {
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
		if i % 7 == 0 && reverse % 7 == 0 {
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
	var columntitle []rune
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
	var profits []float64
	findBest(lines, func(stuff float64) { profits = append(profits, stuff) })
	fmt.Println(profits)
}
func oppgave12() {
	answer := 0
	for i := 0; i < 100000000; i++ {
		if i % 7 == 0 && i % 5 > 0 {
			answer = answer + i
		}
	}
	fmt.Println(answer)
}

func oppgave18() {

	fmt.Println("Oppgave18")
	letters := []int{2907, 6165, 6129, 3468, 2040, 4331, 7935, 5683, 6004, 9694, 8092, 188, 5796, 1184, 8873, 3200, 1981, 9556, 9981, 1387, 7802, 8387, 9970, 7326, 5372, 28, 628, 3408, 6, 3425, 3071, 6021, 9989, 5077, 824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247, 5164, 2026, 4, 4468, 9524, 8, 9227, 8969, 1746, 5593}
	sort.Sort(ByFactor(letters))
	fmt.Println(letters)
	number := ""
	for _, element := range letters {
		number = number + strconv.Itoa(element)
	}
	fmt.Println(number)
}

func oppgave19() {
	fmt.Println("Oppgave19")
	fmt.Println(countWays(30, 3))
}
func oppgave20() {
	fmt.Println("Oppgave19")
	str := "FJKAUNOJDCUTCRHBYDLXKEODVBWTYPTSHASQQFCPRMLDXIJMYPVOHBDUGSMBLMVUMMZYHULSUIZIMZTICQORLNTOVKVAMQTKHVRIFMNTSLYGHEHFAHWWATLYAPEXTHEPKJUGDVWUDDPRQLUZMSZOJPSIKAIHLTONYXAULECXXKWFQOIKELWOHRVRUCXIAASKHMWTMAJEWGEESLWRTQKVHRRCDYXNTLDSUPXMQTQDFAQAPYBGXPOLOCLFQNGNKPKOBHZWHRXAWAWJKMTJSLDLNHMUGVVOPSAMRUJEYUOBPFNEHPZZCLPNZKWMTCXERPZRFKSXVEZTYCXFRHRGEITWHRRYPWSVAYBUHCERJXDCYAVICPTNBGIODLYLMEYLISEYNXNMCDPJJRCTLYNFMJZQNCLAGHUDVLYIGASGXSZYPZKLAWQUDVNTWGFFYFFSMQWUNUPZRJMTHACFELGHDZEJWFDWVPYOZEVEJKQWHQAHOCIYWGVLPSHFESCGEUCJGYLGDWPIWIDWZZXRUFXERABQJOXZALQOCSAYBRHXQQGUDADYSORTYZQPWGMBLNAQOFODSNXSZFURUNPMZGHTAJUJROIGMRKIZHSFUSKIZJJTLGOEEPBMIXISDHOAIFNFEKKSLEXSJLSGLCYYFEQBKIZZTQQXBQZAPXAAIFQEIXELQEZGFEPCKFPGXULLAHXTSRXDEMKFKABUTAABSLNQBNMXNEPODPGAORYJXCHCGKECLJVRBPRLHORREEIZOBSHDSCETTTNFTSMQPQIJBLKNZDMXOTRBNMTKHHCZQQMSLOAXJQKRHDGZVGITHYGVDXRTVBJEAHYBYRYKJAVXPOKHFFMEPHAGFOOPFNKQAUGYLVPWUJUPCUGGIXGRAMELUTEPYILBIUOCKKUUBJROQFTXMZRLXBAMHSDTEKRRIKZUFNLGTQAEUINMBPYTWXULQNIIRXHHGQDPENXAJNWXULFBNKBRINUMTRBFWBYVNKNKDFR"
	strS := "ABCDA"
	smallestHit := str
	for i := 0; i < len(str); i++ {
		for c := 1; c <= len(str) - i; c++ {
			subStr := str[c : c + i]
			if len(subStr) > 4 && len(subStr) < len(smallestHit) {
				hit := true
				for _, element := range []rune(strS) {
					if !containsRune([]rune(subStr), element) {
						hit = false
					}
					if hit && !containsTwoRunes([]rune(subStr), []rune(strS)[0]) {
						hit = false
					}
				}
				if hit {
					smallestHit = subStr
				}
			}
		}
	}
	fmt.Println(smallestHit)
}

func palindrome(word string) (ops int) {
	for i, j := 0, len(word) - 1; i < j; i, j = i + 1, j - 1 {
		n := int(word[i]) - int(word[j])
		if n < 0 {
			n = -n
		}
		ops += n
	}
	return ops
}

func oppgave22() {
	fmt.Println("oppgave22")
	fmt.Println(palindrome("evdhtiqgfyvcytohqppcmdbultbnzevdbakvkcdpbatbtjlmzaolfqfqjifkoanqcznmbqbeswglgrzfroswgxoritbw"))
}


func countWays(s int, m int) int {
	return countWaysUtil(s + 1, m)
}
func countWaysUtil(n int, m int) int {
	if n <= 1 {
		return n
	}
	res := 0
	for i := 1; i <= m && i <= n; i++ {
		res += countWaysUtil(n - i, m)
	}
	return res
}

//ByFactor Sorts number by how large they are combined
type ByFactor []int

func (s ByFactor) Len() int {
	return len(s)
}

func (s ByFactor) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByFactor) Less(i, j int) bool {
	a := strconv.Itoa(s[i]) + strconv.Itoa(s[j])
	b := strconv.Itoa(s[j]) + strconv.Itoa(s[i])

	first, _ := strconv.Atoi(a)
	second, _ := strconv.Atoi(b)
	fmt.Println(first)
	fmt.Println(second)
	if first < second {
		return false
	}
	return true
}

// Generate numbers until the limit max.
// after the 2, all the prime numbers are odd
// Send a channel signal when the limit is reached
func Generate(max int, ch chan <- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 // signal that the limit is reached
}

// Filter and copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan <- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		if i % prime != 0 {
			out <- i
		}
	}
	out <- -1
}

// CalcPrimeFactors calculates prime factors to a number
func CalcPrimeFactors(numberToFactorize int) []int {
	rv := []int{}
	ch := make(chan int)
	go Generate(numberToFactorize, ch)
	for prime := <-ch; (prime != -1) && (numberToFactorize > 1); prime = <-ch {
		for numberToFactorize % prime == 0 {
			numberToFactorize = numberToFactorize / prime
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
		if i % 2 == 0 || i % 3 == 0 || i % 5 == 0 {
			hit = true
			//						if (len(CalcPrimeFactors(i)) > 3) {
			//							hit = false
			//						}
			if i > 6 {
				for y := 6; y < i; y++ {
					if i % y == 0 && big.NewInt(int64(y)).ProbablyPrime(1) {
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
	}

	fmt.Println("The final number: " + strconv.Itoa(theNumber))
}

func findBestProfit(lines []string) float64 {
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

// SliceIndex gets index of element in slice
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

func containsRune(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func containsTwoRunes(s []rune, e rune) bool {
	hits := 0
	for _, a := range s {
		if a == e {
			hits++
		}
	}
	if hits > 1 {
		return true
	}
	return false
}

func findParans(pos int, n int, open int, close int, f func()) {
	if close == n {
		f()
		return
	}
	if open > close {
		findParans(pos + 1, n, open, close + 1, f)
	}
	if open < n {
		findParans(pos + 1, n, open + 1, close, f)
	}
}

// Reverse a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

// RuneSlice for comparing runes
type RuneSlice []rune

func (p RuneSlice) Len() int { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
