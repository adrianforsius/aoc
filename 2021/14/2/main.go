package main

import (
	"log"
	"strings"
)

func main() {
	parts := strings.Split(input, "\n\n")
	template := strings.Split(parts[0], "")
	ruleLines := strings.Split(parts[1], "\n")
	rules := make(map[string]string)
	for _, line := range ruleLines {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	pairs := make(map[string]int)
	for i := 1; i < len(template); i++ {
		var one, two string
		one, two = template[i-1], template[i]
		pairs[one+two]++
	}

	result := make(map[string]int)
	for i := 0; i < len(template); i++ {
		letter := template[i]
		result[letter]++
	}

	log.Println(result)
	for i := 0; i < 40; i++ {
		newPairs := make(map[string]int)
		for pair, val := range pairs {
			r, ok := rules[pair]
			if ok {
				newPairs[string(pair[0])+r] += val
				newPairs[r+string(pair[1])] += val
				result[r] += val
			}
			// if i < 3 {
			// 	log.Println("result", result, newPairs)
			// }
		}
		// if i < 3 {
		// 	log.Println("iteration", i, result, newPairs)
		// }
		pairs = newPairs
	}

	log.Println(result)
	max := 0
	min := 0
	count := 0
	for _, o := range result {
		if count == 0 {
			max = o
			min = o
		}
		if o > max {
			max = o
		}
		if o < min {
			min = o
		}
		count += 1
	}
	log.Println("max min", max, min, max-min)

}

// var input = `NNCB

// CH -> B
// HH -> N
// CB -> H
// NH -> C
// HB -> C
// HC -> B
// HN -> C
// NN -> C
// BH -> H
// NC -> B
// NB -> B
// BN -> B
// BB -> N
// BC -> B
// CC -> N
// CN -> C`

// puzzle input
var input = `CNBPHFBOPCSPKOFNHVKV

CS -> S
FB -> F
VK -> V
HO -> F
SO -> K
FK -> B
VS -> C
PS -> H
HH -> P
KH -> V
PV -> V
CB -> N
BB -> N
HB -> B
HV -> O
NC -> H
NF -> B
HP -> B
HK -> S
SF -> O
ON -> K
VN -> V
SB -> H
SK -> H
VH -> N
KN -> C
CC -> N
BF -> H
SN -> N
KP -> B
FO -> N
KO -> V
BP -> O
OK -> F
HC -> B
NH -> O
SP -> O
OO -> S
VC -> O
PC -> F
VB -> O
FF -> S
BS -> F
KS -> F
OV -> P
NB -> O
CF -> F
SS -> V
KV -> K
FP -> F
KC -> C
PF -> C
OS -> C
PN -> B
OP -> C
FN -> F
OF -> C
NP -> C
CK -> N
BN -> K
BO -> K
OH -> S
BH -> O
SH -> N
CH -> K
PO -> V
CN -> N
BV -> F
FV -> B
VP -> V
FS -> O
NV -> P
PH -> C
HN -> P
VV -> C
NK -> K
CO -> N
NS -> P
VO -> P
CP -> V
OC -> S
PK -> V
NN -> F
SC -> P
BK -> F
BC -> P
FH -> B
OB -> O
FC -> N
PB -> N
VF -> N
PP -> S
HS -> O
HF -> N
KK -> C
KB -> N
SV -> N
KF -> K
CV -> N
NO -> P`
