package main

import (
	"log"
	"strings"
)

type rule struct {
	pair    string
	element string
	index   int
}

func main() {
	parts := strings.Split(input, "\n\n")
	template := strings.Split(parts[0], "")
	ruleLines := strings.Split(parts[1], "\n")
	rules := []rule{}
	for _, line := range ruleLines {
		parts := strings.Split(line, " -> ")
		// rules = append(rules, rule{strings.Split(parts[0], ""), parts[1]})
		rules = append(rules, rule{parts[0], parts[1], 0})
	}

	for x := 0; x < 10; x++ {
		newTemplate := make([]string, len(template))
		progress := make([]string, len(template))
		copy(newTemplate, template)
		copy(progress, template)
		i := 0
		for len(progress) > 1 {
			one, two, newProgress := progress[0], progress[1], progress[1:]
			progress = newProgress

			i += 1
			element := getRule(rules, one+two)
			if element != "" {
				newTemplate = append(newTemplate[:i], append([]string{element}, newTemplate[i:]...)...)
				i += 1
			}
		}
		template = newTemplate
		if x < 3 {
			log.Println(template)
		}

	}
	occ := make(map[string]int)
	for _, t := range template {
		if _, ok := occ[t]; !ok {
			occ[t] = 1
		} else {
			occ[t] += 1
		}
	}

	max := 0
	min := 0
	count := 0
	for _, o := range occ {
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
	log.Println("max min", max-min)

}
func getRule(rules []rule, element string) string {
	for _, rule := range rules {
		if element == rule.pair {
			return rule.element
		}
	}
	return ""
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
