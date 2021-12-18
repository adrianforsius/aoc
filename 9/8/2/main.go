package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Entry struct {
	Signal []string
	Output []string
}

func main() {
	lines := strings.Split(input, "\n")
	var entrys []Entry
	for _, line := range lines {
		parts := strings.Split(line, "|")
		signal := strings.Split(strings.TrimSpace(parts[0]), " ")
		output := strings.Split(strings.TrimSpace(parts[1]), " ")
		entrys = append(entrys, Entry{signal, output})
	}

	leds := map[string]int{
		"a": -1,
		"b": -1,
		"c": -1,
		"d": -1,
		"e": -1,
		"f": -1,
		"g": -1,
	}

	outcomes := map[int][]int{
		0: zero, 1: one, 2: two, 3: three, 4: four, 5: five, 6: six, 7: seven, 8: eight, 9: nine}
	start := map[int]string{
		0: "",
		1: "",
		2: "",
		3: "",
		4: "",
		5: "",
		6: "",
		7: "",
		8: "",
		9: "",
	}
	count := 0
	for _, e := range entrys {
		for _, word := range e.Signal {
			if len(word) == 2 {
				start[1] = word
			}
			if len(word) == 7 {
				start[8] = word
			}
			if len(word) == 4 {
				start[4] = word
			}
			if len(word) == 3 {
				start[7] = word
			}
		}
		d, s := diff(start[1], start[7])

		leds[d[0]] = 0
		for _, six := range ofSize(e.Signal, 6) {
			_, s1 := diff(s[0], six)
			_, s2 := diff(s[1], six)
			if len(s1) == 0 || len(s2) == 0 {
				start[6] = six
			}
		}

		for _, five := range ofSize(e.Signal, 5) {
			_, s1 := diff(s[0], five)
			_, s2 := diff(s[1], five)
			if len(s1) == 1 && len(s2) == 1 {
				start[3] = five
			}
		}
		for _, five := range filter(ofSize(e.Signal, 6), start[6]) {
			_, s := diff(start[3], five)
			if len(s) == len(start[3]) {
				start[9] = five
			}
		}
		start[0] = filter(filter(ofSize(e.Signal, 6), start[6]), start[9])[0]

		diff8and9, _ := diff(start[8], start[9])
		for _, five := range filter(ofSize(e.Signal, 5), start[3]) {
			_, s := diff(diff8and9[0], five)
			if len(s) == 1 {
				start[2] = five
			} else {
				start[5] = five
			}
		}
		diff8and0, _ := diff(start[8], start[0])
		leds[diff8and0[0]] = 3
		diff8and6, _ := diff(start[8], start[6])
		leds[diff8and6[0]] = 2
		diff2andPos2, _ := diff(start[1], diff8and6[0])
		leds[diff2andPos2[0]] = 5
		diff3andDisc, _ := diff(start[3], start[7]+diff8and0[0])
		leds[diff3andDisc[0]] = 6
		diff9andDisc, _ := diff(start[8], start[4]+d[0]+diff3andDisc[0])
		leds[diff9andDisc[0]] = 4
		diff8andDisc, _ := diff(start[8], start[3]+diff9andDisc[0])
		leds[diff8andDisc[0]] = 1

		digits := []string{}
		for _, word := range e.Output {
			translated := []int{}
			for _, letter := range word {
				num, err := strconv.Atoi(fmt.Sprint(leds[string(letter)]))
				if err != nil {
					log.Fatal(err)
				}
				translated = append(translated, num)
			}

			digit := ""
			for num, o := range outcomes {
				if equal(translated, o) {
					digit = fmt.Sprint(num)
				}
			}
			digits = append(digits, digit)
		}

		num, err := strconv.Atoi(strings.Join(digits, ""))
		if err != nil {
			log.Fatal(err)
		}
		count += num
	}
	log.Println("count", count)
}

func equal(a, b []int) bool {
	log.Println("enter equal with a and b", a, b)
	if len(a) != len(b) {
		return false
	}
	count := 0
	for _, A := range a {
		for _, B := range b {
			if A == B {
				count += 1
			}
		}
	}
	log.Println("actual count", count, len(a))
	return len(a) == count
}

func filterStr(entries string, f string) string {
	out := ""
	for _, e := range entries {
		if string(e) != f {
			out += string(e)
		}
	}
	return out
}

func filter(entries []string, f string) []string {
	out := []string{}
	for _, e := range entries {
		if e != f {
			out = append(out, e)
		}
	}
	return out
}

func ofSize(entries []string, size int) []string {
	out := []string{}
	for _, word := range entries {
		if len(word) == size {
			out = append(out, word)
		}
	}
	return out
}

func diff(a string, b string) ([]string, []string) {
	diffMap := make(map[string]int)
	sameMap := make(map[string]int)
	for _, A := range a {
		if in(string(A), b) {
			sameMap[string(A)] = 1
		} else {
			diffMap[string(A)] = 1
		}
	}
	for _, B := range b {
		if in(string(B), a) {
			sameMap[string(B)] = 1
		} else {
			diffMap[string(B)] = 1
		}
	}

	diff := []string{}
	for x := range diffMap {
		diff = append(diff, x)
	}
	same := []string{}
	for x := range sameMap {
		same = append(same, x)
	}
	return diff, same
}

func in(a string, b string) bool {
	for _, B := range b {
		if string(B) == a {
			return true
		}
	}
	return false
}

var (
	zero  = []int{0, 1, 2, 4, 5, 6}
	one   = []int{2, 5}
	two   = []int{0, 2, 3, 4, 6}
	three = []int{0, 2, 3, 5, 6}
	four  = []int{1, 2, 3, 5}
	five  = []int{0, 1, 3, 5, 6}
	six   = []int{0, 1, 3, 4, 5, 6}
	seven = []int{0, 2, 5}
	eight = []int{0, 1, 2, 3, 4, 5, 6}
	nine  = []int{0, 1, 2, 3, 5, 6}
)

// var input = `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`

var input = `bgafcde gfcd agc ebdgac adfceb bafeg efgca cgdfae cg ecadf | fabgced gc agc cdfg
gbdacfe gcabfd cdb dcfba bfacg cgad fadeb feabcg cd gcbfed | bdagcef dcb cdag gbfca
dgcbafe dbfca fbaed be cedb gefad dcfeab facdgb eba gbface | eb gadfcbe cfbad gfbeca
ebc cb aedbf agcef badecg gaebfc bcgf adbcfge ceabf daecgf | cb bce efdab ecbaf
fedbc cebad gfcbd fec gcdfab ecbfga dacgbfe gfed fe gefbdc | bfecag cef egdf fgde
bafedc baefg dbfga daegcb gae egbcfa eg cefab fgce decbafg | abdgf cgfe cedgba befga
dcba fagbed cgbfed dgfbeac da dag acbgde fcaeg becgd acgde | agd bacd dga gbecad
ec aebgcfd fecd bagfec efagd edfgab cgdfea dgcab ecdga eac | ce ec ce gbacefd
aegdbfc fe dacbf aefbgd aecgb cdfe ebadcf ecbfa cbfdga aef | dfce fe fcde afebdgc
bgaf gcbad cagdfb gb dfbeca degfcb bfadc cfgeadb bdg agcde | fdegacb acbefgd bdg gafb
gdecba dcae ec dbfeag bgead dgbce gec cgefab dcbgefa cdbfg | afdbgce agbfde abcgefd ec
gacbefd fgbedc egafdc adf fcbae afedb ad bfeadg adbg dfgeb | ad degcfa ad dfbge
fadebc baef eadfcg ebgdc eda cfgabd afbdc ea edcab cdgbefa | aefb ead ea ebaf
badgecf caefd aebfdc fecbdg fadcg fed de ebad abgecf bceaf | ed gcfebd aebd gfdca
bcdae aefcbg fgdea bf cdabegf adcebf fcbd fbead afb begadc | gbcafe faegbc fba eabdc
bfcdeg ecd eabgc adeg ed dbaec bcfad ecabfg dbgcfea dbacge | de bdfcaeg ecgfbd bcdea
deacf bf eagfdb gfecadb facbeg bfa cbdage abfce fbcg gecab | aebcf ecfagb abegfd bgeac
eag dgbecf gdfae caed dfebgca ea ebcgfa dfbag gecfda dgcfe | ae ea deac fcdega
edbga egdabc cgfdbe gabfe efcdgab gbdefa fb abdf fbg gecfa | cgaebd cbgdea acgfe begdca
cgfedb egbdc agcfd cdagb eabgdf ab abce adb agcebfd dabegc | ba bda abd deagfb
fc agfbce fdebg dgafeb ebfdc cbf cfedgba fdcg gcefdb eacdb | gbfaec fgcbed gbceafd bedagf
bdcgefa acdfbe ed gbcea cfgbad ecbda edfcag afcbd ead bdef | dfbe de dea dea
gfbeca gcefdb fegacd cgbda fbde ed bgedc begcf agdcefb ced | gbacd ced dec bcegfd
dgbacf gfabec agd egdaf badcegf ad ecgaf defgb cade dfegca | cdae dga bgcaef ad
dcf cfadb bcagd afebd fbgdec egfcbda fcea gfeabd fc ebdfac | gabdc ebfad aefc eacbdf
cfgdba fcdab cgfbade adeb afdebc egcfdb de aedfc cfage edf | ed eadb dgebcf cefdabg
bcae befga dgebfc fbagce cbefg agb fcedgab ba bdgacf dgfae | beca dcaegbf gab geafb
dgafbec fedba cbega aebfdc eabfdg bdgfce edc fdac cd adceb | abedf dfeabg dc gdbefc
adegfbc gbedcf cfabd gacedb abcfe fdbacg dgabc df dfga bfd | fdag dbf dafg fd
dfbcag ebfcad ed agecdb bfacd egbfa badfe afbcegd aed cdef | afebd ecdf cfed eadfbgc
fca dbcaf ebdcf fadbg cdbefa adec ac dacbgfe fbcdge cgefba | ac edca edac ca
agdfce cedbfag agc dcbaeg gc bcdga bgdfa cfaebd ebdac cbge | ceadb becg ebadc gc
aefbc aecdg efgcad cdebag bcg agfdbc bceag bdge bg fecdgab | bgecda bcdaeg gb caebf
cfgebd acedf ag gbad gcfbad acg bgecfa bdfcg gbadecf cdfga | dgba dgcbf gac bgfdc
cdgbafe cdgbe gd agde bcaged dfebc geacb cdg bgeacf gbdcfa | efgcba edag gecfdba cgd
bcg bcaf cegdaf dfacg dgacb fgcebda bgdcaf efbcgd bc gadbe | gcb bcg caebdfg cfab
eg cdgfae aecfd fgbaedc gbdfc facdeb cfegd ceag efg efdgab | fge acefbgd egf efg
fcdg gefda cf edabc cedgfa cfa acefd cfebadg gdbafe bgecfa | fc fca cfabged fca
ed bcadegf egdfb fgcdb edbafg facbeg abegf deba fadgec dge | eadb ged eadb adfegb
bgedc cf cefadg ecbgdf gbfae cbdf fcg ecbgad abfgced becfg | bdcf dbceg fgcbde cdbeg
cg gcaf defcag cdefg gdafe dgfabe gec dabegfc bcegad ebdcf | feadg degcf gcaf dcgefa
fdeac gdcebf daecbg bgfec fdeagbc gca gecfa aefcbg fbag ag | ga decfagb ag agfb
bc fbdc bagecdf gfdcab acdgb dcaeg dbgfa cgbaef ebgfda gbc | bdfc cb cfbd fbcd
efagc bafegd abecg eab eb fgbeca gbedcfa bfec fcgade cgbad | abfgdce degacf fcaebg cdabfeg
egbfcd afde ecgab fac fa edabcf cbdagfe afbec cfabgd fdebc | deaf feda cgbdfa fa
gbafd defca ebcf eb cefbad dcebag dcgfae edfab deb bgedcfa | be deb gdbafce be
afdgceb cf gbdeaf acdbfe abcef acf deafb fedgac cdbf cgeba | fac cfdb bdfcae fdcb
cfdba ca afbgd eabc cda afbcdeg cfedb ecfgbd efcgad bcfeda | bfgad abec abce caeb
cfebad debfacg acgfeb cab fedcb afcdb cdae febcdg ac dgbaf | bac ca faecbgd cgbfae
eg dfcegb dge abedcgf facdg bafdgc degacf dafeg agec adebf | egd fegda fedbcag eagdf
age gcbad feacb cgbdfae ge gfbe agebc fcadeg cbafde cgebaf | fagecdb bfeg baecg fceabd
bfcdge fgcbead ebdgac ba dagbfe gab dabc bcegd gface cebga | cbda gbedc egcbafd gab
dbcegfa fadceg fcage bcage gb bdagfc abdec gbfe bgafce bcg | ebfdagc ebfg ebgf egfb
gcdefa bagf dgebf bdefc bgcafde dbefga edabg gfe gf dcgeba | fg efg agfb gebfd
fabdc gdeabf dag gbace adcgb gd eagfcb gaedbc egdfabc egdc | adgbc bcadf gd cedfabg
deagcf ce dfcaeb gfadceb bgade cde daegc cgdafb gadcf efcg | dbaecf fdecbga ce ecd
dfcbag abceg db dgcefa dabec bcegdaf cdb bdef defac cfbdea | gdafec dcb febd bdef
ead fgceabd edgcaf da fgda fgcea ecdgb cgaed fdbace cbagef | ad eafcdb bfecagd ad
ab caefg agbcef efacb gecdfba ecfagd fbdec bdecga cba fgab | eafgc abfg cdagbfe dcageb
afbcge ebafcgd dacegb cfbea afbg agc egcfa ag cgdfe bdceaf | adfbceg fagb gfba aecgbf
acg cdegbf ecfadg cgdab fcbadg cdbgfae ca agedb cbgdf facb | cdeabfg ca afbc dagfec
dca dbfae bagdcf aceg ac afdegc cegfd fcedgb cbdaegf cedfa | dac acge cedgf cfdbgae
gdcafe ega cgbfde cgba afdbe bacegd ga begdc dcebagf edgab | abgc ga bacg bdacge
begfda gdafce fg gdbf gbaced agf decagbf aebdg gabfe cebfa | fag acdbeg agcdfe dgbefca
gfcdbea fecab cbadge afgce gc fgdace fdcg cga eagdf eagbfd | efgacbd gc dgfc dacbeg
gdbefc feadg dbagef afd abde gebfd cefga bdgafc da efdacbg | cbegfda adf bade cdgfbe
aefbdgc eabcfg cgefa degbf cedfg cd cadg cde fgecad fecadb | dbfegca egcdabf gcda dc
baed bcgda abdegc decag ega dgbfaec gacbfe ae cgfde dbacgf | aebd age ea bgdecfa
abecg faegcd egfcd efcga cdaf baedgf efa fcbaedg fcgedb fa | afe af aef fa
abgfce cbedgaf fbedc eacdfb dcf fd faegcd fecba dfba gdceb | faecdg gdceb dfc gcbde
fgce fgadb egdabc edg defbg eg bfced fdbace bdfgeac gbcfde | gdcbfea deg afbedc dgcaeb
eg dgcfe efdgab bfdec aecg acbegfd bdcafg adcfg efg daefgc | fge aecg gfe ecagfdb
acdfe bedcfg abgec cfeabd aedbfcg fecdag fg fcg afecg gdaf | cdaefg fdag cgf gadcfbe
aedc gbdac cbagfe bfdcg gdbfea acb cbgaed ac eadgb gecadfb | gbfdc gbdaec fbgeadc decabg
gbc afgdb cfbdga badgef gadbce egabcdf geacf bcdf cb gabfc | fgaec fcagdeb gbcfa bc
cfgabd cfdbe eafgdc badcfeg dgfab cbag eafbdg gc dfgbc cdg | fagecd cgd afdbge cg
eadcbf fcag gc gbc gadcefb gebad febdgc egcbaf ecbag bacef | dgbfec egadcbf gcfa bcegfda
acbeg acgedb eg dgbfce cdagb cgdbaf gaed bafce cfebadg ecg | dbgac eabcf eabfdgc eagd
bgefda eacg abdcfe eabcd abegcd ebg cbedg cbfdg dgfecab eg | ecag eg ge afebcd
dgecbf bg dfcba abge gadebc edagc bcg ecdfga fbgecad bgcad | cgb gbea ebgcdf gbae
dcabg gdfbca debc aed ed fabeg afdecg gbeda caegbd agbfcde | fcbdgea cedb dea efcdagb
edbcafg cgaed gc gbdc acfebd cge gceabd abecd afegd bgafce | gc facegbd edbcaf cg
gadbce gbdaf acebgf bd abcgfd edfga cdfb gbd cfgab dgafbce | gfdab fdegbac eabcfdg bgcdae
eacgf fb cegbdf dgbfca cfb dfba agfcb cgabd ebcdgfa gedacb | acbdg dbaf ceafg bf
gebac egdacf fdceg daef af gfa dagfbc fcgbead afgce cebfgd | afg bcage gcdeabf cgfea
cdabfg fgbec gcdafe ceafgb egc ec beac febdg bdegcfa fabcg | cafegdb beca ecgbfa bace
gda feabdg ebgaf bgdfc bgfcae bfeagdc edab ad adgbf fceagd | gad ad afgbde dfaebg
bfc bgfced edbfag fcea cf afgbdce cgbda abgef fcgba aebcfg | bcf efac eacf fc
bdeag bgdcf ec ebgdfac gcbde gaec dec cafdeb afdebg cebdag | gcbfd agedb acfdbe ecd
befad gd edgbf bcgafd cebgfd abgcef gfd ecgd cgdfaeb ecfgb | cegd cdeg efadbgc egcdbaf
gfdbae bc cegbfd gadcf cdaefgb gfbed fbcgd beagcd gbc becf | gedcfba edcagb fbcgdae cfbe
ecf cbdaf bdfcaeg fagdec efba bcdfe bfecda gcedb ef cbadgf | ef fbedc dbceg edgfcab
cabde ae ecfdba cbdfe faed dabcg cfbgaed gbfcea begdcf abe | adfe aeb cabgd ae
gefcabd ecabd gebafc ag ecfgbd gacdbf dfcbg fdga cagbd acg | gebcadf gefbac dgaf dagf
bgde efcgd cebgfa dbacgf cgfeb gdc fdaec dfbceg gd fegcadb | efbcadg cgd gbdcfa facebg
gefcbd bgfdcea fg dcbag fbgcd fgeb defcag fcg becfad dfcbe | cefbda cfg fegbdc fg
bcfed bcfdge aedbfg fad decbfa facge cdba efdca ad dbfagce | eabfgd fadce begfdca gecdbaf
edfgbca dbega fdb gfecab fd ebcfda cfda feabd becfa edgbfc | df ebgad cabegfd adcf
daegbc fbacg edfg befdac efbcd edafbgc cdg bcgfd fdegcb gd | dg becdf dbeafcg gadefcb
egdacf gde ed dfebcag acgdfb cfadg agebdf cgeab gecda dfce | fedagc gabfecd eadgfbc cfgebda
dfcea afbde fdgac gefbac eabcdf ace bced ec ebagfd adgbcfe | ce fcdga baecfg cea
adebg gdfcae gb acegd abg begc bdaef bfgcda dgaefbc abgdec | bg cbafgd ecbg gba
bfacde egfadcb cb fcdea eacbf cedb dcgeaf bgadfc abc ebgfa | adfebc dceb bdcafe efdac
faegbd dgfce dcgafe fgadecb bcgd bcfge bfaec gb cdbfeg fgb | aecfb fbg fcgade dafgbec
fgacbe fc ebcadg agfdb fdbca cagbdfe cdfe abcde fcb bfacde | cabed cfde bfdac dgbaf
defabc gfcbe fea bdacfg adgecf fabce ae aebd bfecdga cdafb | fea ea bgfaced bead
dfgae dcafe dce dgebfca abfcd ecab bfcdea gcfebd gafcbd ec | bcdagf cde gbedcfa bfdace
cea gaedfb acdf ac ebacgd gefdabc befda cfbae egcfb efcabd | gcaebfd cbaef dfac ca
gebac befadgc bedcgf cgb gc eagbdc cebaf cadg bgfaed aegdb | cbg cbg cbgadef acdg
gfcdeb bfdegca beac befda bcf dcfga fbcda cb adecfb gbefda | fbc dbafc cbf cfb
cbdfeag becagf acgeb be cgbda gfcdbe efba acfeg ecb eadgcf | aebf eb bagcfe eb
febg aecbdg fedacb ecg ge fbgaedc dfgce afcgd fcbed efcdgb | aedcbg baegcfd ecg bedcf
bcfedg gdbcae febd fcbag fgbcd bd ecdgfa dbc fdcagbe gdcfe | efgbadc fcbedg abfecgd cgbedf
feagc cb dcgeabf fadbgc cgfdae fbdea feagcb bcafe cbeg bac | bgce bc cgadbf cba
efagbd dageb dcgbeaf gc acg gdecab gecd fdbca gbcfea acbgd | eagbd egfabcd gfbace gca
ae cgae dea eacgfd fadgc gfebd afdcbe dfeag acbfgd dacbgef | dgbafc cbdfgae ae ead
egb gcfe abcfdeg begfcd abcdfg dabec eg bdfcg edgcb bedagf | bdecg beg gbdcef dbeac
fcgae dfcebag aefdb geacbf cda dagbcf cd adcfe fedacg gdce | bcdafg adefb fdeba gecbfa
afegd gadfb fcgadeb gbcade gdabef cefgad dfbe bdg fagbc bd | dabegc bfgca bd dgb
dfceba faecd bd afbedgc ebad fcadb bafgc dcb efgdcb gfaedc | faedc bead egcdbaf bd
cgb bdcfeg cg gdbeca cgef gdbcfea febdg fdcgb dafgeb bacfd | cgfe cgb gc gfec
cadgfb cefdbag cg begacf dfgbea bgc adgbf bgcad fdgc eabcd | cg fgceab ebdacgf ecbdgaf
bgdefc decfg cfedga cefgb bfdcea edgfacb bgcd fegba ceb cb | cfgbe dcaefg cbgfe dbcg
gefbda acde gcbaf cd cfd bdcgfae daefg gdcbfe cgefad agdcf | gcebafd efgcda cfd aecd
gedfc cead egbfda fgcdabe de fcgea gfdbc deg cagdef gebacf | deg bfgead gdfbc gde
fbgae bfecg dbcaef gdcaeb cbfdge gbc decgbfa edbfc fdgc gc | gc gbcdef fdcg gfcbe
fc caf cadbe dgabfc gebdfa fecg fagbe cebaf acbegdf cgabfe | egcf ceagbf gcfe ebgafc
abfcd cgbfa bcgafd gadefc abfde cdf dc agfbced cbegfa gcbd | cd cfd fgeacd cdf
bdgae gfbce cbagfd gedbcaf ebcgd gbcefa dc fcde cegdfb dbc | dcegb fgdcbea deabcfg bdc
ebcagf df bgcdefa fdceba cgdf dbacgf eadbg afd abgcf gabfd | ceagdbf dagcbf fd dcfbaeg
dga cadbgef efbcgd bgdca eacd cedbga gdbec gdeabf fbacg da | dag adg edac edgcb
febcg gbadce aedbcf acbed cbgae cag gadb degbcfa ga fcgdae | bgad abdgfec acbed acg
gdcaf dcbfag gdefa abcdg bcfg cf fedgacb dbaegc fcedba afc | cf afc gcfb eafdgbc
gdcaef gcabfd bcefgda cde gcdfb gbecfd ec dfaeb cefbd egbc | egbfdc ce gbdcfa afbcgd
fdea fga fa adcgbef abfcge gbeda fagdb abedgc cgbdf bdefag | fa gaecfb beadgc cbfgd
ba gcbafd decgba bfga fbadc bda dfbeacg ecfgbd cedfa fdbgc | gfab gbdfc adb ba
bd ecdgab aedfg bde fbcae decfba cgafbe fdeab dbfc fdgceba | db dbe cdfb bdgfaec
bd gcefbd ebcd dafbgc ebfgd decgaf edcfg fbd fbaeg gafcdbe | cbeagdf cebd cedb cedb
edagf cdaefbg ead efgdb bgafed cfdag bfea dabgce dfbceg ae | gcafd feab dgfca feba
gecfda bgcdf ba abd edgabf bcdeaf gedfa fdbga baeg abegcdf | ab bad bda fgebcda
egadb bec dbefga cbafg cageb ce ebdgca dceg gdebfca cfbead | ec ebc egdc ecb
cbgfe fadcbg agde cbdge egdacb degcabf bed faebdc ed dcabg | eadg facdbe bcdeg debfac
fabe fb bgf edcgaf bgacd bcfegd fcdegab adgef afbdg gbfade | bfg dbcgfae edacbgf bf
bcdfe becgd cbafegd decag geb dgfb bfcade dgcefb bg gfbaec | bdgce gb dcbge gdfb
egcadf bfg fgbdac faedg ebaf gcedb begdf fb adfebg cafdgeb | faedbcg bafcgde bf becgd
edafgb ecbgd gcdfaeb ef gbfadc cgfade gef fdegc fgcda efca | agfecdb dcabgef ef dgcaf
cbdgea fegcad bagdc abdfgc age gebac afbcdge egbfc adeb ea | ae ega ecagbd gbacd
ecdfb edafbc fe afdcb efc degcb aedf acbgfd ecgfab fecabdg | fe dbegcaf fe cedgfba
edbfgc caeb ca cbdafg bcdge bgcdae edfga eacgd gca cgabfed | facebdg ca aedbgc ac
gfeba aegc cbfega ea bea dfgbe edfacgb fadgbc gbafc aefcdb | gbecfda bcdfag gcea afcegdb
fgcea dc egfdba cdafe gfdecb fbaedc ebdfa dacb bdafgec dec | ced bcfdaeg cefda fagcbed
dbefa defacg agd dcbfg bdfagce ga efgbdc abgc fdgab bacfdg | cbga dga gedcfab fdabcg
eacbf eac dfagbce cfga fbgea ca bfced agebdc fbedga fcbega | dabfgce gfca ebfacg gbaecf
cfagb cgfdea gdc agfcbe fdeacbg bcfdg dg badg fdgacb bcdef | cgd gfacb gd gdba
cfga gbadef decabfg cbged baefcd ca cda dbagf fdbcag dbgca | afgebdc bdgeaf gafc cfag
cgfbd cdafgb gbcead dbg afdb bgecf adfcg ecbgdfa bd gdacfe | fgacd cgdfa bdaf dbg
abf fb febd dbafg febcga gbead egdacb fedgba fbagdce gdcaf | bf fb dbceafg fba
fd dgfabe ecafdg deabcg cbefa fed dbgea bdefa bgdf dbfgeca | fde bdfg fdgb gfeabcd
fgeadb abfed dgfbcea bafgd agcfd gfb bdfaec bg dgbe cegbaf | egdb fegadbc dcgaf bg
fdabc acfbgde dg fcage cdfgab dfebag dga gdafc cdgb eafbcd | gfcad aegdcbf dag afbdc
dbfe efdag fgdcae cgbfaed bgfead acebg fb bgf egabf abfcdg | ebdf gfb ebfd fgabed
bafegd egcfadb fbc bc fcebgd becaf afbed agecf beacfd bdca | cb dbca cfb afgbcde
gbefdc bfead gfcdab gbedca fc cabegdf bedfc fdc cefg cgbde | fdbceg cgfe fc bdcaefg
fg dbfea fedagbc cageb aecfdg dfgb dbfaec gafbed begaf gaf | gfecad dgcfea dafeb gfbd
dabfg eb egacdfb bfcdga faedc gdeb fbagec aeb bgefad dfeba | bafdcg gfbda gebd be
adb dgeabc bfdg gaecfd afdgc cbdgfa bd dcgeafb dafbc bafce | bd fcagbde defagcb dbfg
fge bagcf dfaebg fgebc eg cdeg fabcedg bedcf dfcegb acfbde | egcd ecgd cebafd cgdaefb
bafdgc gecabf caegf ae beaf gfacb decgf cfedbag aec cdabeg | dcfge ae egcdafb eca
fcgba ebgaf agefd agcdfb beac gdbcfe beg bcdfgae gacfbe eb | cfebgd abecdgf ceab abcfdge
abcgfd bac feabd adcgeb fcgb bc debgcfa fbdac gceafd fagdc | fbcg gcfad afbed cfbda
afdbegc dea abegd agecbd afbgdc dgbfe cabgd afcbed ceag ae | daecbfg cdbag acge ea
ebdcg cbfg cdgebf bge adbec eabgdf baecfdg fgadce gedfc bg | dcbae fbcg beg gfbc
cfbga cefabg ebfca abeg ae bcedfga gbacfd eac bdefc dcafge | abge fedbc bacgf eac
eb aedcg dbfcge cegdbaf gfbe bed dagcbf dacebf dcfgb dbceg | gbef dbe fegb eb
cgfe ced dgeba ce cegad abedcf cdaefg cfagbd cadfg gbcedaf | ecgad gcabefd dfbcgae ce
dbagc afdc dgfaeb bcfage gbedfac cafbdg cdgeb ac cab gdafb | cab dafbcg dacf acdgfeb
dgbae bfa eafgb bafegd fgaec bcfadg fb dbfe egbadc gbecadf | cdegbfa dagfcbe fdeb fab
be aecgdf gbdacef bagfc bdef beagf beg dbefag gedfa cbgade | be eb afegb beg
cfdagb ac gcfdae cbad bfgda bcfedga begcf afc ebafdg cfbga | deagcf abdc ca cagfb
eg efdbga adbcfg ebfcga aefgd bdfga adcbgef defac gbed gae | edbg gedb bfdeag dbcgaf
bacegf gbfca cfa ceab gafeb ca gcfdb gadfec dbafge gbdfaec | caf fbgedac fca afc
fbagde cbeg afecg gdaecbf fedcga gab efacbg gb acbfg cfdba | gba ecbg agb abgcf
abfg dafcgb gcafd bdcgaef ab cba ebdcf bgdaec cfaegd fcbda | acb ecadbg cgdabe fgab
agcbd cegdba feagb fd daf gbfad bacdfg cgdfabe dcefab fgcd | df dagbc fgcd fd
gdfeabc bcd debaf cb gedfab bcefd edfgc bdceaf aecb cdgbfa | ebac beac cbea abce
cbgdef eg beg dbgfc bfgcead gbefd efcg edgbca eabfd acbfdg | gbe gbe cefg gbe
fedg begcad gdb bafdge gd egbcaf gdbaf bfacd eacgdbf eafgb | cbefag baegf gefcdab bedgac
gbce dcafgb agfdceb efdac acg agcfe aefgb befagd cegfab gc | fbedgac abdefg cga cag
gabc bceadf agfcd agd ag fcdgbea gabdef gedfc dbcfa agbfdc | gacb gad bgceadf cbgafde
afedc fabced gfcbed feb cegdbaf dbfa ebfac dacfeg fb gbace | feb afdb ebf acfde
cbed edgcf ecabdfg ceg cfbgd abgfdc gebfdc ec efbagc gedfa | ec decb ecdb gec
egd gfbeac gedcb de dabgfe dcfbg cbefdag eadc dcegba acegb | ed adgfceb gde abcged
efadcgb fbdce fagedc agebfd gbafe abedf da bgad ead ceagbf | efagb ecdbf ade aebcgfd
cgeb bfaed cbedga cbaed bc bca afgcbd dgface bgfecda decga | fgadcb bca dabfe cab
dgaec deagcf egf cdeabg abgdf gdfceab dfcbeg gfead efac ef | agebcd efdbcag aefc bgefcad
fbeac gebcaf efa fgce bgfdea dgbcfae gcbdea adcbf ef egacb | gebfacd bfadcge ef eagbc
acdbgfe adfg edfcbg efacb egacf gca dgaceb egdcf ga gefcda | fdga cfgae cga ga
dfebcg gdaebf dagcf gbaecd dcgebfa bf cfbe bgf bgdce fdbgc | cgbefd bcfe gbf ecbf
dcafeg ca acf fcdeb abcfdge facdgb bfagd dfaegb bagc dfbac | dfabgce dfaegc gbafde bacg
cgefb cbe cfab afdbeg cb faebg gcfde cefbag ebadgc afdbgec | egbdcfa cbe cafb fecagbd`
