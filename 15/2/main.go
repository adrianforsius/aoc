package main

import (
	"container/heap"
	"log"
	"strconv"
	"strings"
)

type node struct {
	pos     []int
	visited bool
	weight  int
	lowest  int
	parent  *node
}

func main() {
	lines := strings.Split(input, "\n")
	grid := [][]node{}
	for y, line := range lines {
		numLine := []node{}
		nums := strings.Split(line, "")
		for x, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			numLine = append(numLine, node{[]int{y, x}, false, num, num, nil})
		}
		grid = append(grid, numLine)
	}

	newGrid := make([][]node, len(grid)*5)
	for i := 0; i < len(newGrid); i++ {
		newGrid[i] = make([]node, len(grid)*5)
	}
	// n := len(grid)
	log.Println("size", len(newGrid[0]))
	log.Println("size", len(newGrid[len(newGrid)-1]))
	n := len(grid)
	for i := 0; i < 5*n; i++ {
		for j := 0; j < 5*n; j++ {
			newGrid[i][j] = node{
				weight: grid[i%n][j%n].weight,
				pos:    []int{i, j},
			}
		}
	}
	// for _, row := range newGrid {
	// 	// log.Println(row)
	// 	vals := []int{}
	// 	for _, val := range row {
	// 		vals = append(vals, val.weight)
	// 	}
	// 	log.Println(vals)
	// }

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for y := 0; y < n; y++ {
				for x := 0; x < n; x++ {
					newGrid[i*n+y][j*n+x].weight = (newGrid[i*n+y][j*n+x].weight+i+j-1)%9 + 1
				}
			}
		}
	}

	// log.Println("===================")
	// for _, row := range newGrid {
	// 	// log.Println(row)
	// 	vals := []int{}
	// 	for _, val := range row {
	// 		vals = append(vals, val.weight)
	// 	}
	// 	log.Println(vals)
	// }

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		for row := 0; row < len(grid); row++ {
	// 			for col := 0; col < len(grid[0]); col++ {
	// 				highest := i
	// 				if j > i {
	// 					highest = j
	// 				}
	// 				x := col + j*len(grid)
	// 				y := row + i*len(grid)
	// 				if y < 3 && x < 3 {
	// 					log.Println(row, col, grid[col][row].weight)
	// 				}

	// 				newGrid[y][x] = grid[col][row]
	// 				newGrid[y][x].weight = newGrid[y][x].weight + highest
	// 				newGrid[y][x].weight = newGrid[y][x].weight % 9
	// 				if newGrid[y][x].weight == 0 {
	// 					newGrid[y][x].weight = 1
	// 				}
	// 				newGrid[y][x].pos[0] = y
	// 				newGrid[y][x].pos[1] = x
	// 				// log.Printf("row %d col %d, n %+v, %d, %d", i, j, newGrid[y][x], row, col)
	// 			}
	// 		}
	// 	}
	// }

	pq := make(PriorityQueue, 0)
	q := &pq
	heap.Init(q)
	heap.Push(q, &newGrid[0][0])
	newGrid = FindPath(newGrid, q)
	// for _, row := range newGrid {
	// 	// log.Println(row)
	// 	vals := []int{}
	// 	for _, val := range row {
	// 		vals = append(vals, val.weight)
	// 	}
	// 	log.Println(vals)
	// }

	log.Println(newGrid[len(newGrid)-1][len(newGrid)-1].lowest)
}

func printParents(n node, count *int) {
	// log.Println(n, "count", count)
	if n.parent != nil {
		*count++
		printParents(*n.parent, count)
	}
}

func FindPath(grid [][]node, q *PriorityQueue) [][]node {
	if q.Len() == 0 {
		return grid
	}

	n := heap.Pop(q).(*node)
	y, x := n.pos[0], n.pos[1]
	if grid[y][x].visited {
		return FindPath(grid, q)
	}

	grid[y][x].visited = true
	queue := make([]node, 0)
	if x != 0 {
		queue = append(queue, grid[y][x-1])
	}

	if x+1 < len(grid[0]) {
		queue = append(queue, grid[y][x+1])
	}

	if y != 0 {
		queue = append(queue, grid[y-1][x])
	}

	if y+1 < len(grid) {
		queue = append(queue, grid[y+1][x])
	}

	// log.Printf("queue %+v", queue)
	for _, neigboor := range queue {
		// log.Printf("next %+v, %+v", neigboor, n)
		next := grid[neigboor.pos[0]][neigboor.pos[1]]
		if next.visited {
			continue
		}
		if next.parent == nil || grid[y][x].lowest+next.weight < next.lowest {
			grid[neigboor.pos[0]][neigboor.pos[1]].parent = &grid[y][x]
			grid[neigboor.pos[0]][neigboor.pos[1]].lowest = next.weight + grid[y][x].lowest
		}
		heap.Push(q, &grid[neigboor.pos[0]][neigboor.pos[1]])
	}
	return FindPath(grid, q)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].lowest < pq[j].lowest
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[:n-1]
	return item
}

// var input = `1163751742
// 1381373672
// 2136511328
// 3694931569
// 7463417111
// 1319128137
// 1359912421
// 3125421639
// 1293138521
// 2311944581`

// puzzle input
var input = `9211721995398984823216199411311153152823357222811429112119221712176869272771135219495921176931988996
6632171351773132146884416975191918371137224144193771793912898863217419114196131892432139894169918218
9479521581911111241635179421411251317118763198279239927245149311971162612514425931183828668218451615
2799883538791161712421321215469831181832839573641299118119399389191992134461723567119512463934372862
7282911193661181886222521419365211346192563458259817925831313772244927912198535916963121911143145294
7614421564771345699136125411336941242471217736129198182121166939992929134242129115972131543886852442
5941872221148212759291291116631761153795991497256619169111283211161754939313151921885442411736928569
9123357419994724953111348498171371612472182316614117126994541982319998172331519248231238156111672964
1793111433116111717139179889917339411911111781137459123542929732221323213243778444994114521418912211
9599488259943918112119511324131313758145149481139556171281917832677986918336221981132216419838826923
2956771922813182998122434294185999481631267125112943151717238532359271953669798373246231149549677717
1412599822261914692793521648113911618413619921546893663539972268573129829791355399119114442113137199
8578722132131828966612211194992693171143881911161114661199811619919233226399886114799711223242215442
5165649652429116569861462312372874823923822319291648113938659182588595925492121359142594262829395934
9541911493899653531495338864629491794698697898463521358598137418158192457261911822927812111471292349
7429364138518115763191312177152451298513693393118832839735519284624151128122957386811151669127948794
8615934214181734643172613319995512889281941991712169211863647721437113142128248586189624229381111374
2918418921229288328473992115143932731885921143239762144164311299189511181119111816152915855462568719
9119293772519511337882218617499212262121549564512722927988491234712227611914618111722981319829398499
4523541192833191177596112643924111598741511593332131543471973928281591121833919113121113198575625411
2998335296491912914371431631267299259415579179432348119789614296272999149411624311372549422611551899
9781146981221224189368891714257919186915814674711926718622116212431743929137867427861146671217961671
4124157586349536791324361421921314652151419819192366661388828719114692811895131921247315241211165111
8312487711455414177171141322364274232517711292241491412117382831949485128118146253221922376492141411
5541292237521217379624533691549722937891724771281199981763812212729941439637585818256199172782121118
5892271515541571911791711285991668421294361671614413691194611219391547761238621922221611286751295712
7382862881885721433838891311918128792952948179179822729213156879721543855226626428885917574219561272
9839266152518315917619293212421333429254911438632476892451116219228214741161295361918211692241811482
2915171145993822142447377583731791967444691921247791113138538692255243121982881219311719149747627679
1116281733911412767756115128135192564845238215387575193132878198132325415992918423124212142512193195
4814779112928892321922261812164451519357735431712369239545111184425323125233741153613799832939226339
1117597897548248261215551422339772941193111312335591735412544412683393937219156322242649743814239512
1294311137922178121411292381743429423451324655517338821297957266781164724369143122232318226348238118
9211262511644444692263486213297465387978199213375652933173212711251264471911947442665428571922117196
8154419379395691912143138575212134632218165647199439491999193762894616911141948429362112433442818849
2693981111123246187211284584528492482381731919914642689987915119272994213579734791474599111668193727
8159681921121984956281438821919281979115987197988136918256661949612119943185322419132233852118921283
9728195912573151196721531117911992937943931292171313112411934614211189951712582391212157812794143114
7378834631134314416159385516195582151652733965543193162113919839984599175821851542141863572211111273
3151252721819924983938992611192712434847937474196654719921249841113214115925481558819918212361213672
2197231399528959498167971798141171427942329522292962597897439719799147494139685511272852288391619111
9239667496116398826282351159399496182684931155125231298461513623824278571715145996988766119111234592
8166169286339958612835714698635367616115726469152242374393117832112521117398977979611156158781211694
3277615115712627939151384361339466619681719141299119191262117381511591442175835999293318191354112855
1319383697112152117312211393165991922319793917392922519263185117332282587449137112171146379115823211
3224192125166826331311792711698225545724988966166298862931412221853197226178112799923155261414446462
9936716682982571861999825941741584429199369512534569689213213149153189195142772456239622191671912867
3153336496418879367193131154761212141283694427935991849778818281811722967199499172941221818111193611
1196468712323116135516296192189314642529119318142152521384124849688249812115184672951228312711221761
4925614148261211315916119899961216123721113991498437411349225811983592125346193217948711492474311911
9546593259361575284647351161813916122957412819937128343983211334341715474979629251311891811372281521
8712315781148329488379124111171869293612911668114216128122891432633667852327324137844197192111581737
2791296111783727899149416171612591431717951469517191622511943612233411185685854161415627528999621398
5592192631435182821111896841831991117451211597129221184472911984798115698822717683497316525145738648
1589332661372311511822246978921271411428674275516193113211829934295111231218559721211191142145244114
8678433291412349113791319411172216793485999141149413241612189114126529295726861669246641531222663392
9131881713726999839182449211771899568199152869311929135892545734183183733466371291419225151256232892
1417993118992128127151821611691327642589411512825913769229369496299491711613185333679683281838832118
1799111753291769233338642727891568212541892824711241151961311652161421175396935991411171933334217441
6113151991918271919273998811187396531375776181552289275813223372249117378431112429593632321853919411
8183276792153288631275188332915889418667212384194265918162325219229515375941394222512149148924613296
1416118913612978261561481631519185368359811592342789978956488163931243582535932496126511518184513511
1142897483426589738325146644691151577211287851444741442829261239154228233254339757462151935175431165
7671239318932121152252923935948533915134624338621119225913629391449531196712211881582629414969261199
2767953181951127327217681258592284293483391141118859279287791545924711322997681997143397121111482951
7121539121913947129691178729388513517893262952213199268231113828321266333238819115216162971219168818
8291923195939181843196238173691311211297177786591891164198218238296232464272431339297236811338113218
7149925296699481183346264814126163921317597184167216935912325139284193791194311496229182334581647511
1463121249711691914518682732461711316522678991282959595519524411629196171169217916517774284581961826
1912981254399292912114814228236312932141917815313812342661811819494533372112232181896643773835197251
2821196211341211176641216813891699772713812119918364838911391156469125511653148712626438931882423395
1133287984916171253816577183143349322793831996344418415232531945812829841487911417163237188711272537
7182492671863824212961115819127964197836819231178517872121831117991753344798926814137455714618229318
1311278161515249225291813119973962368169152197141156188615617338132177254373476795516816511118491642
1123193391558563483912311144713996192161482179923628873142817365231579492341196319155111156168193841
1411557115971111732257814212895353833613322537241328128171642119157925321161931117193135847535135597
7673549322277127386191411329155691119966925819881985154131998963918291711941825522349492791571286232
5191217791229527141374656855569411475547219947649119711229374137995414187829119359338917351211111391
4971631622187692488669761366711739321571717213122391951494694114111549812595111269215634117216712914
1311927358495122389571311912223238178181111688122629142846181374224667157199191137922411127243335952
3912591252312323974688146894541186511814832111935171228312991951161296436797111729722469146166144916
4163211151732959355115825912591228141213188722626981511412891119588297199473617415632819941651313832
2154117611375111597994783729212581114383729819171173455183344169513649171262112811817211144861213814
9915439694676431596885727544191155621598225166241725118552371413974734757639112995141917119141211711
3711453282797862299328182251929728336131381289267666123219858997317675991639421358364851863911454293
6614991938832137119182917221419119746181149219823693688941311344216335114145119199586533322874129522
3532112184179229497816318987467147511242172444389195861228313161741232149119964299131161236286933729
8765996448871489229112222919155591895145293543119522149917711313197239146198271266196947716712221617
4196622453197923474135732511346986817911188941119364282922941391338441291532786526299255451956951383
4623151772181827336112418794171595518137211529999296841259255111914817641418371335929675961457596912
3538242183751962911117966962315961446223121321368184734136224485119661738352274469281315256211175191
6639257518919164884271788811918189113368317882126621444656192186653625742218863881289611599985712871
2635156288199864445822545919165222128111493192326149782751738697642117516713351357132199118674292239
1646911331119821598241623452291251954723411141691328852829612997978613518131989239672231311226174731
1633111149692637121852427281623518481796158919325138939764915131718134261629971227883951556717531827
3922823175378757992171657356259729139333211179596169224374212246865456139147114722129842381149641211
6414123578648949375269599975911354889411927442231426151121151494531512239927351649956896214388958251
9627871171972922874152123416193689719289946935988174889999521225157114783814418341136678516112142142
1341114328117517394181175811132754116211655987121153834243689118151649774381424693462372627451331821
1632919311883994217232352862953438319333462223123468212213816241414949812219219618811481527563624821`
