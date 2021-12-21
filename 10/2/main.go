package main

import (
	"errors"
	"log"
	"sort"
	"strings"
)

type Char struct {
	Open  string
	Close string
}

var points = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func main() {
	lines := strings.Split(input, "\n")
	clean := [][]string{}
	chars := map[string]Char{
		"()": {"(", ")"},
		"[]": {"[", "]"},
		"{}": {"{", "}"},
		"<>": {"<", ">"},
	}
	for _, line := range lines {

		s, err := stack(strings.Split(line, ""), chars)
		if err != nil {
			continue
		}

		//incomplete
		clean = append(clean, s)
	}

	for _, l := range clean {
		log.Println("clean", l)
	}
	counts := []int{}
	for _, line := range clean {
		count := 0
		for i := len(line) - 1; i >= 0; i-- {
			count *= 5
			count += points[line[i]]
		}
		counts = append(counts, count)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
	log.Println(counts[(len(counts)-1)/2])

}

func stack(line []string, chars map[string]Char) ([]string, error) {
	var depth = []string{}
	for _, char := range line {
		// log.Println("depth", depth, "progress", progress)
		c, err := FindChar(char, chars)
		if err != nil {
			return nil, err
		}

		if char == c.Open {
			depth = append(depth, c.Close)
		}
		if char == c.Close {
			if len(depth) == 0 {
				return nil, errors.New("invalid")
			}
			var x string
			x, depth = depth[len(depth)-1], depth[:len(depth)-1]
			if x != char {
				return nil, errors.New("invalid")
			}
		}
	}
	return depth, nil
}

func FindChar(c string, charMap map[string]Char) (Char, error) {
	for idx, val := range charMap {
		if strings.Contains(idx, c) {
			return val, nil
		}
	}
	return Char{}, errors.New("not found")
}

// var input = `[({(<(())[]>[[{[]{<()<>>
// [(()[<>])]({[<{<<[]>>(
// {([(<{}[<>[]}>{[]{[(<()>
// (((({<>}<{<{<>}{[]{[]{}
// [[<[([]))<([[{}[[()]]]
// [{[{({}]{}}([{[{{{}}([]
// {<[[]]>}<{[{[{[]{()[[[]
// [<(<(<(<{}))><([]([]()
// <{([([[(<>()){}]>(<<{{
// <{([{{}}[<[[[<>{}]]]>[]]`

var input = `[[<<<{<<<(<<(([]())(<><>))<[{}()][(){}]>>(([()[]]<{}()>)[<[]{}>[(){}]])>)>([{<<<()()>[<>()
[<[<{[[({[({{[(){}][()[]]}<[{}<>](<>[])>}{<[{}{}][<>()]]{{[]}{[][]}}}){(([<>{}]<(){}>))<({<>{}}{{}{}}){[[]<>
[<<<[{{[[{({{[<>{}][()]}}[[[[][]][{}<>]]])<{(({}<>)<()>)<[{}<>][()()]>}>}(<{[[()()]{{}()}]>[<<{}
([[{<{{({<[{[([][]){[]<>}>(<[][]>[<>()])}([<<>><{}<>>]{({}{})[{}[]]})]>})[[[<<(<()[]>{<>()
{([<(({[[<{{[[{}<>]({}[])><{(){}}{()()}>}<<[()[]][{}[]]><[<>{}]([]{})>>}([(<(){}>{{}()})<<{}{}>{[]{
[[(<{[<(<<([<{<>()}<<><>>><({}{}}{{}[]}>]{[{<>()}(<>{})]{{()[]}{(){}}}})><<{({<>[]}[<>[]]){[()[]]}}((<<>{}>
{([{({([((<<{(()())<<>[]>}<{()[]}({}<>)>>({<{}()>[<>{}]}[<(){}>{()<>}])>{[<{<>{}}>({()[]}<<>()>)]({({}<>)<{
{([((<{({[[{<{<>{}}>}{[{{}{}}({}{})]{[[]])}][<[<{}{}><()()>][(<>{})]><<{{}{}}({}{})>>]](<[([{}()]({}()))[(
(({<{[{{([<{[<()()>]<{{}()}{[][]}>}<{[<>()]{()[]}}<(()()>{{}()}>>>[<{{[]<>}}>]]{{[{{[][]}[[]{}]
[(((<{({<{<(({<><>}({}()))<[()()]{()<>}>)({(()<>)[(){}]}<{{}[]}[{}{}]>)]}{([{[[]()]{{}{}}}[((){}){<>
({{{((<(<<<[<({}<>]{<>{}}>{{()<>}{<>()}}][[[()()]{[]<>}]{[{}][[][]]}]>>>)>{{({({{{()<>}(<>{})}
({[[(<[[(((([({}<>)[()[]]]{<()()>[[][]]})([[<>{}][()()]]{{<>{}}[{}[]]})})[<{<(<>[])({}[])>}{<((){})>{{{}
{[<<{([[<<({{((){})(()[])}(<[]<>>{()[]}))<[[<>][[][]]][(<><>)<<>()>]>){{[<<>()>(<>[])]([()[]]{()
<[{({<<[{[[({[[]]([]{})}[(()())<<>{}>])([[<>()]<[]{}>][{<>()}<[]()>]}]](<<{{{}<>}<<>[]>}<(()())>>>{{<[[]]<<><
<[[<(({<({[[<[{}[]]{()()}>[[<>{}]]]<(<[][]>[{}{}]>>]})[({<[[{}[]](()())](<{}<>>[{}[]])>[(<<>{}>({}<>
[([([(<[({([[(()()){()()}]])}<[{<[<>{}]>{<[]()>{()()}}}][<<[[]}[[]<>]>([[]()](<>()))>[{{(){}}({
({[[<((<{[<{({[][]}<{}{}>)<[[]()][()()]>}>(<[[(){}][()[]]]{[<>][<><>]}>([[()[]]{<>{}}]{[<>]{<>[]}
[({<{<((({{<{{[]{}}(()[])}>}<{[[<>()][{}]]{<<>{}>[[]()]}}{<<[]{}>>[<<>()><<>()>]}>}{{{[(<>())(()[
<(<<{(<(([[[<[<>{}][()()]>({{}{}}[[]{}])]][<<<<>()><<>{}>>[{[][]}]>([{[][]}<()()>]<<{}<>>>)]]<<{({<>()}{[]
((<<(({{<{({<<<>()>[()()]>((<><>)<(){}>)}({<()<>><<>()>}{(<>())[[][]]}))<({(<><>)([]())}<<[]<>>(
({(<<[{<<<{<{([]{})<[]<>>}<<<>()>(<>{})>>[[{[][]}{()[]}](<<><>>)]}([[{{}{}>(()[])]{({}<>)<[][]>}](<<()<>
{{<[<[{([({[([(){}](()[]))([{}<>][[]{}])]}({[({}{})](({}{})[{}{}]))<{(()())}{[[]()](<><>)}>))]{(<{{[()
[{[{({<{<([<<<{}{}>[[]()]>{<[]<>>(()[])}>{([()<>]<()()>){[{}{}>(<>[])}}]((<{<>{}}<{}{}>>)))[[<<<()()
{({({<[([<[{[(<>[])<<>[]>]}{<(<>()){{}<>}>({<>{}}<<>[]>)}][(<[{}()][<><>]>([<>()]))]>{({(({}())<[]()>)}[
({((<{{([[{<{<<><>>([]<>)}{([]()]{<>()}}><{({}<>)(()())}{[()<>][{}[]]}>}(<{{()}[[]{}]}{[[][]][[]()]}>(([<
{{<[<[{[({{[{(()[])[[]<>]}<{[]<>}{[]<>}>]}(({[{}<>]{<>[]}}([{}()]))(<[{}[]][[]{}]>{(()<>)}))}{[[[([]())(()<
[{{[{[[([{(([((){})[{}()]>{{[][]}(<><>)})[<<{}{}><()<>>>({[]()}[{}{}])])}[[<[({}{}){[][]}]{(()())<{}
{[<(({[(({[{[<[][]>[<>()]][(<>()){<>[]}]}]([{[{}()}<<><>>}[<()[]><[]()>]]{<[()()][<>]>({<>{}}{()()})})
((<<[<(({{<({{()<>}(()())}{{[]{}}}){<[[]{}][[]<>]><<<>()>{{}<>}>}><({[()<>]})<<((){})[[]<>]>>>}{(<[[
(([<({[[{{((<(()<>)<<><>>><{()<>}[[][]]>))}{({<((){})(<><>)>}<{[<>[]][()()]}[<()><[]{}>]>)<[(((){})<{}[]>
<{[<{<{(<({(([{}<>]<()<>>)[{{}<>}(()<>)})(({()<>}<[]{}>)({()}))}[[({[]}<<>[]>)[[{}()][[]<>]]][{{
(<[([[((<<<(([{}{}])([()()](()<>))){[[{}{}]<()[]>]{<[]()>{{}{}}}}>{([[()()](<><>)]){{<(){}>}
([({<[([[<<<<(<>())({}[])><[{}<>](<><>)>><[<()[]><<>()>]{{<>[]}([]<>)}>>>]<[[<{[(){}]([]{})}([[]{}]
([{([[([[{([<{{}[]}(()<>)><<{}<>>[<>[]]>]{[(<><>)[<>()]][{{}{}}]})}{<[{([]())[[]]}<[[]{}][<>()]>]>([<[<><
[{<[([{([([({<<><>>(<><>)}{[{}{}]<()>})])[[[<{<>{}}[[]()]>{<<>>[[]<>]}]<<<<>{}>([][])>({()()}[[]<>])>]]]
(([{(({(<([<<(<>())[{}<>]>>({(()[])[<>()]})]<[([(){}]<{}<>>)<((){})<[][]>>]({(<>()){()<>}}[{()()}<(
{<{{{({{{{([<([]{})<<>[]>><{[][]}{[]<>>>]([<[]<>><()[]>]{[{}<>](<>[])}))<({(()<>)([][])})>}{[<<[{}
((<({[([{<<<{([][])[[]]}{<<>[]>[{}[]]}}>((<[()]{()()}>){[[<><>]([]{})]})><<[<{<><>}[()()]>(<[]{}>([]()))
[<[{[[[[{{{(([{}[]]<{}[]>))[{<[]<>>}[([][])([]())]]}(([([]{}]]<({}{})<<><>>>)[<{{}[]}<{}()>>[<
<([<{[{<[<{<<{(){}}[{}<>]>({[]{}}}>}<<({[][]}{[]<>})[(()<>)({}<>)]>>><[{({<>}([]<>))[<{}<>>
{[[{(<(<<([[<<[]()><{}()}>]<<{<>[]}([]{})>{[<>[]]<{}[]>}>][[[<{}()>({}<>)]]<{<<>>}<<[]{}>>>
([{<<<(<[([<([[]{}]<[][]>)<{<>[]}[[][]]>>{{{<>{}}[[][]]}{[{}()])}]{(<[<><>]<()<>>>){([{}()][<>[]])
[{[[<{[<<{{[([[]{}])]}<[<[{}()]{(){}}>[<()[]><()()>]]>}({{<<{}<>><{}[]>>}[<<<>{}>{(){}}>]})>><[
[{[{<[<(([[{{{()()}<()<>>}[[[]<>]<[]()>]}{{<{}{}><<>[]>}[([]<>)[{}{}]]}]{{<<()()>[[][]]><{()()}(()())>}}]
([<{{(<(<<(<([()[]]<<>[]>){{()()}(<>[])}>)<[<[<><>]{()[]}><<{}{}>{[][]}>][([()[]]<{}[]>)<<<><>>
[{<{[[{<([{[<[<><>]{()()}>[({}<>)]]}({[(<>{})<{}<>>]{[<>()][[]{}]}}[<([]<>)(<>[])>[[()()]{
(<{{(<({[(<{[([]())[[][]]][(<>())]}>)[<<[{{}<>}<<>[]>][({}())<[][]>]><[([][])]<{()()}({}{})>>>[([<(){}>{{}[]
[([<(<{{{{{[[(())<[]<>>]([()()])]<{{<><>}(()<>)}[<<>>[<>{}]]>}(({{<><>}{<>}}[([])<{}()}]){{<{}()><[]()
(<[[[({(<(<[{({}<>)<<><>>}]>(<[<[]{}>{{}()}][[<>()](<>)]>))<{{<{(){}}[<><>]>[<()<>>]}[(<<>{}><<
[([(<[[({<<([{(){}}([][])]{[[][]][{}<>]}){<({}()){<>{}}>[((){})[[][]]]}>((<<[]{}>>({<>[]}]){<(<>(
(({<(<(<({{({<[]()>})[<{<>{}}{<><>}><(<>[])(<><>)>]}{{{{<>{}}(())}}<[{{}<>}<{}{}>]{(()()){{}{}}}>}})>[(<{({[<
([<[([<{[[<({<()[]>}{{{}{}}[()[]]})(<[{}[]]{<>{}}><{{}[]}<[]<>>>)>{[[{{}[]}<[]{}>]{{<>{}}<{}()>]]<[{<>[]}[
{{[([<<{((({{{[]}(()[])}[({}{})(()<>)]}(((()){()[]})[<{}<>)<{}{}>])){[<<{}()>[()<>]>[(()<>)[()[]]]]})({{[
{<<[<({<[[<{((()[])({}())){{{}[]}<{}()>}}{<<[][]>[(){}]>((<>{})[()()])}>]{{{{[[]()](()())}(<
<<{[{<{([{{{[{()<>}}[<<>[]>]}}{(<[<><>]<<>{}>><[()<>]({}())>)<{<<>{}>[<><>]}[([]())<<>>]>}}][{[[{[<>[]
[(({{{{({({{<{[]()}>(({}{}){<>()})}})<(<(<[][]>)[<<>><[][]>]>(<(<>[])(<>())>))>})<<<[{<({}[]){()[]}>[
({{<(<{({{[[<[{}()]<[]<>>>([()<>]<[]<>>)][[[<>{}]{()()}]{{[]()}<()>}]][([[()<>][[][]]]<({}[]){<>}>)[<<{}
[<(<(<[([<<{<[[][]][<>{}]>}(<{{}()}{(){}}>)>>]){([<{(<[]{}>{()}){[{}()]<<><>>}}{{<<>()>}]><[<<[][]>[[]<>]>
<({[{(<(<[({((()())[<>()])[<{}[]>({}())]}{{([]())<<><>>}{[{}()]{{}{}}}})[(<{<><>}<[]{}>>[[[]
<[<[<{({{[<[(([]<>)<[]<>>)<{[]<>}{()[]}>]<<<{}<>><{}>><[{}{}]<[]()>>>><(({<>[]]<[]{}>))[{<[]{}><[]<>>}((
<<{<(<{{(<{<<<(){}>[[]<>]>>({[{}()]<{}>})}<[<{[][]]{<>[]}><{{}[]}<[]{}>>]([{{}()}]{<()><{}[]>})>>)}<[<[<
[<<{{<[[([<{({[]()}{[]<>})<[(){}]<()[]>>}<[{<>()}<{}<>>]>><<[<(){}><()[]>]{[[]()]}><{{()[]}{(){}}}{[[]
[<{((<{(({{<{{()<>}[<>()]}((()[]))>}([({<>()}[<>{}])<({}{})<()[]>>][<([]())<{}{}>>[({}{}){[]{}}]])}){
({[{{{<(<[[([{<>()}<{}<>>])({[[]<>]{{}<>}}[({}())<{}<>>])]([{<()()>}({()[]})])](<<{[{}{}][<>[]]}{[<>](()[
{<{{([<<[({({<(){}>[<><>]}[<<>{}>{{}[]}])}[[[<(){}>[[]()]][{[]()}]]<([()[]][[]<>])[<{}{}><<>[]>]>>)<<[
<<(<{{(({<<(<{{}{}}{{}<>}><{[]<>}{{}()}>)>{{<{[]()}[()()]>[<[]()>{{}{}}]}<(<[]()>{()()})[[[][]]]>
<{<[{{<[((({<<[]()>{()[]}>}{{[(){}]{{}()}>{[{}[]]((){})}})<([[{}()]{[]()}])>)){<[[[<<>[]>{[]}
[<(({[(<<<<<{[<>[]][[]<>]}(<<>()>{[]()})>><[(<<>[]>(()<>))<([]{})({}{})>]([[(){}]]{[{}[]]([]
<[[({[((<[{[({[]<>})({[]{}}[()()])]}<[[[<><>]<(){}>]{(<><>){[]<>}}][<{<>[]}{<>{}}>([<>()])]>]<((<{[][
<<[{([{<<<[[[{()[]}([]<>)]<({}()){[]()}>](<((){}){{}<>>><{{}()}{[]<>}>)]>>><<<<<[{[]}]>{({()<>}(
([([{<{{{[([({{}<>}{()()>)]<{[()()]}({(){}}[[]<>])>)[({<<>{}>}[[()][<>[]]]){[<()[]>]}]]}<[(<(({}())[(
([<[<{{({<{({[()[]]{[][]>}[[<>[]]<<>{}>])<<[[]]{<><>}>({[]()}[{}{}])>}[{<[{}[]]<{}()>><<()[
((<<({<<([{<<((){}){{}()}>>}([<{{}()}({}{})>[<[]>[{}{}]]]<[({}[])[<>{}]][(<>{})<(){}>]>)])(<[<{{[]()}[<>()]}[
[{([(<<([[<(({[]{}}[<>])[{<>{}}{<>{}}])((<[]>))>][{[[({}()){[]{}}]]<([()][{}{}])>}({(<[]()>(<>()
(<([((<<[{[<[[<>{}][{}{}]]>]{(<(<>[])<{}[]>>{(<>[])[()<>]})[{{<>()}{{}[]>}{{(){}}[<>[]]}]}
<{{[{{<[{{{[<[<><>]{()()}>((<><>)([]()))]<{<()()>[()]}<<[][]>{{}[]}>>>{[{{[]<>}}[{()<>}([][])]]}
<{<<[({((({<(<{}>[()[]])((()){{}{}})>({([]<>)[[]{}]}[([]()){[][]}])}{(<{()[]}(<>{})>{{<>[]}[[]{}]}){[(
(<{[<[{<[(<<<{[]()}[[]<>]>{(()<>)([]<>)}>(({{}()}))>[([{<>[]}]{[<><>]<[]{}>})(<<[]<>>{()()}><[<>{}]
{<{<<{<<{{{<(<{}><[]<>>)({(){}}[[][]])><{{[]()}{<>{}}}>}}{[[[{[][]}([]())][([]<>){<>}]]][<((<
{(<{[[[{([{[<<{}>[(){}]>[<[]()>([]())]]({({}[])[[][]]}({{}{}}(()[])))}{(([()[]]{{}[]})<{<>[]}<()[]>
[[{{<{(([{{{[(<><>)]({{}[]}[()])}<<<<>[]><{}())><([]{})>>}<<{[[]()]{{}{}}}{{[][]}[()<>]}>>}<<([{[
(<({[<{<({<[{<[]{}>[<>()]}]<[{()[]}<()[]>][([]{}){()()}]>>})>}[<([{{<({})<[]<>>>{{<><>)}}[[([])[
{[{({[[[[{(([{(){}}<[]<>>][{<>()}<[]()>])([(()[])]{{{}}([]<>)}))}{{{<[{}()][<>()]>{{()()}[(
{{{[[({{{{<[[{{}{}}((){})]<<<>{}><<>[]>>]>{<[[{}[]]{[]()}][<[]()>{{}}]>{[(<>{}){[]{}}]{(()[])((){})}}}}{
(<[{({<(([<((<{}{}>[[]{}])[<()()>[<>{}]])>(<[<[]()>{{}[]}]{<{}<>><()[]>}>{{[<>[]]<{}>}<{{}<>}
[{({<{[{[[{([<[]>[()()]]<{()[]}>)}[<<<[]>[()()]>>]]]<{<<{((){})<<><>>}[([]{}){(){}}]>>{([[<>
[{<[[(<(([([[[{}<>]({}[])]{<()<>>{{}{}}}))]{[([<{}<>>{()[]}][(()<>)([][])]){([[]()][{}{}])<[()<>]{
<<(<{[<{{[<([{()<>}<()()>]<<[]{}><[]<>>>)>]({[<([]{})[<><>]><{(){}}<[]()>>]})}[({[[<<>>(<>{})
{({<{[<<({{(([(){}](<><>)))({{<><>}[[][]]}{({}()){<>{}}})}{<[((){})]<([]())>>([[{}<>]{()[]}])}})([<[{[<><>
[[<[[<({<(<({({}<>)[[][]]}){<<()><[][]>>[({}())[()[]]]}>[(((()())([]())))])>({({{{()()}({}())}`
