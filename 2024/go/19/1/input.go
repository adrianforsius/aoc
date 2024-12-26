package main

const exampleInput = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

const input = `uugbwr, rwgbb, wgb, wwb, uggrbb, rburr, bggbb, rrggr, rggr, www, rbub, wug, rbbr, rbb, wwr, wwubwg, rggurw, bb, wbbbuwb, wgw, urur, rbwgwu, wwbb, ubbrbr, gwb, uwwwg, gwr, rruw, ugbruu, wgg, uuub, gug, uuuub, rwbgw, rwu, urbggu, bbwugur, bgug, bbuu, uwuw, grbr, wwwgbbg, uuu, wguwugr, uburu, ggrwg, wuubrw, wrwbr, ggb, gbr, burwu, wubgw, rb, buwww, ubur, wuwr, ubbw, wu, grwurr, ggggbr, ggg, rurwuggu, brgggug, rww, wwwwg, uwb, grr, wwwrrb, buu, uwub, gww, rwwgww, guurrg, wwurr, wgggug, gwbu, bgrgbwr, wwrrrg, bwruru, rgubu, rrwr, ggwbu, uuw, bgur, grg, brg, uwugu, urrbrbg, bwbgb, burwb, wrrbbgg, rrw, urgu, guu, ruub, grwuuw, br, bgrgr, wrbbg, rbw, ugwuw, wbrw, grgbr, bwu, ubgu, uburg, bgb, uggu, rgwu, grrugwrb, uwbburgr, rbgbww, rubrwu, gugrbw, wgub, bbgugwg, ugw, uuru, wuuwrw, uurww, bbwrr, rwrrg, gur, wwuu, rrr, gugugr, wwg, rbgu, urg, bwgbw, brgw, ggrgbub, guwb, gbg, wwuub, b, guwwgrbb, rbgrgub, wwrb, wgr, rrrrg, guuu, uwrg, rrg, bbrwuw, wruug, grguwgrg, rurubr, wrurwgru, ww, rgggu, rwrbugr, burrbbr, rw, urw, uwu, uwg, rwg, rwwgug, gburw, urgwrwr, uwwgu, uwwbug, ugur, wgu, brw, bwbw, wguguwu, uwuu, rgr, uguww, brwr, wrgrbu, uuguwb, gb, gwwr, brr, gw, bggbg, wbwg, wrg, grugwg, bbb, grrgrbr, bgu, bw, rgbwbuu, wrb, wruwbg, wrurggb, wubggg, bwrr, gwrurgr, uuguubu, buw, ggww, gguwrwb, bwgburb, ubb, ub, rrug, rugbwgb, wrrbw, bgg, rbwrgr, guub, bubgbgwr, bwwu, uur, wrwu, uwrgburg, rwgb, wrw, wbrb, gugwwggu, rgg, gg, rrwb, ugrwwb, rbgwr, rbr, bbr, gwg, wruggu, ggr, wgbwrbu, gbwb, wbgb, bub, wuw, ugbub, wub, wrrb, rwww, brrbuw, rrwrb, gbu, buubrg, rubr, wgbbb, bgbg, rbu, uwbr, ubrbr, ru, wr, gurgu, bug, bgr, grgw, rrb, wrugg, gbwggwu, rurrrg, grb, rrggb, bru, rbbrg, gggu, guww, ugg, brwrbubb, ug, g, wuuw, uww, urbrr, gbb, wwbg, bbbr, rwurur, wbugbu, bbu, rgwgu, ggu, bbgr, grgbb, ruuurr, wgwg, rbrb, wrubbww, wbwugr, wbw, gugwuw, guw, brgub, wbbru, uubw, ubbwubr, ruuwrrr, bubrg, ubr, uub, w, wru, bgrwwbwg, wg, bggrgwu, gu, gwrw, u, rgw, uruu, uugur, ggbr, rbg, rgug, gbw, wb, uubbg, gwbgrb, rbgru, ubrgb, uw, grw, ubgug, gwbbub, wugrg, uwgw, wubwbwr, wwrbur, uu, rggbwr, brwugg, ubruugg, rwubbw, rug, ggw, rwb, rrru, bbgw, uwgww, ugb, gbwbwgw, urr, wbr, ugrb, bg, gwrbgwwu, uguggw, gwgbur, ubw, wur, bgw, ggru, uwurw, gr, bgbrub, gubb, wgrb, gggw, bbgubbgr, bggrw, wrbwgg, ggbb, rur, rgu, gguru, brgu, ubbb, rwrubrbu, ubwb, ggub, bbgrr, gwgu, uug, bbg, wrr, ugbbr, uwbbu, guwgw, rwbru, burbgb, gugurru, uwwg, grbg, ruw, bwg, ur, bwrwbr, buggu, gruurubw, uwrw, wwu, wuubr, bubuu, uguubbu, uwwwu, uubbwbgu, uwgwg, wuwb, rbur, wrbwr, uru, brb, ubwr, gru, gbbubwug, bwwgr, rub, bur, rgwgwb, gub, wbru, uwr, wbb, wbu, ugr, rwbww, bgurw, bww, grrg, wgbg, grrrbb, bwr, gugr, wgbrr, rrwu, bgbbrg, wuwrwurr, rbgg, rwrgwwg, guuurguu, bbw, rr, wuu, ubrw, rubbubw, urb, rgrbw, ubu, rwgg, gbru, bbrbbru, uurrbr, wubrwur, wbg, gwu, wgugrug, guug, bgwgugb, urrbr, bwbggrgb, bbwr, rru, bgrrug, ruu

urwrbwuruugwwgwwrbwbbrrbruggwbburwrrgbrwbrubggbgbg
urwrbggbwwbrbggwrbbuwwuguwbwbgrwwrrgwrbggruubwbbwbbgruwu
uwugbbggwbgggrruwgbrrrurrgubbrburgggwguugrrubrrbguubwwr
rrwwrgbugbbugruwgrbrrurubburbbbrurgbwrbugbubww
ubwwrgwuurgwbrguugbwruwgugburuugbwugguwwrwb
bwurwbbbbgbwgwbuwbwruuuuubuwurrwbrbrbgugrgrrguwgw
bbwuwbgrubrgruwubbuuwwrggwrguwwugwruubrbbbgwwwwrrrwrwwwubr
bwwwrwuwwgwrrbubbgubwurbrrggrbgbwuubuwwugrubugwrb
uwbbwbbrurugggruwrwwgbwgwrwbwgwbgugbgwuwrggrbubgur
ubwrrgubuggwrgwwwbwwgwurwgwurwgbrurbrgurwgbubu
wwrgwrbwbgwwwubrwggwwrwruurrwggbrgbubuwguwrrwbwrb
uwgrgugbwruggubuugbburuugrrrrbwgwgbgwurwugburgrrubbbwbbguw
rgbwgrrrubwwwggugwbgbbbgruwuugrubwbugbbbwrbgbgrrbggw
uwgwwggbwbgwguugbwbbbwbrggrggwbugbruwbrbbrgwbgbrbbbubwb
guwrgrwwbuuuwguurbubrgugbwbwburbrgrwrgrbuuwurgwugwgggubb
rubwgbwgwgbrgurruwbwurrgrrbrwbwwrwwbrubwbubwurgbubwrugrg
gwrgbbrubwrwurrrwbwrwwwugwbrgbrwgwruguubrggrrurb
guwwruurbwwuguuwuurbugurubwgubwgwwurrbbbgwbwbrb
ggbwbubgwrugurggwwrgrwbwuwugruwbrwbwgubwbbuwgbgwgbbburrw
urwwrbgwwwwgbgwrbgurrrwubrgbururrrubbwwubgubbrugbbwbggrgb
grgwuggurwuuwggbuuwrgwbbgbbguwuwbubwuwbgburruuu
gwgbrbrugwrwbggruwuwubgrwrrbgbuggugrwrbugurburwuwrbggu
gbuugwwrruguuuwgbugrurggbgwgwburrgrgbburgr
rwrgbwbrrbrburgbrrbuwbburgruruwuwgrbruwrbwrgwubb
rgrbgwgwbbrbrbgwubrwruuggrwgrbuwbrurbwgbgwrrrrubwugbwrur
ubbwgwugurrwurwwbbuuggugggggbuuuuwurwgurwruubwburwgrgr
wrbwgrrgbrugugwrwwbwbbgggbgrgggwuggwbgrrwwgggwu
gbwrbburwbbuwburrrgguuwggbugwgbwgrbgrgwubbbrbwubb
ubrwgwurbuuwbwgruuuwwbwwwruugbwwubggwwgrrrwwbrbb
wbrrwgguggrwrbbrbgrgbbwbbrugbbgrgbgubgbruwrbbrgrurrrbgb
rgbubruuwwwbwbbwrggurrwwgwgbgrbuggbugwgrbrwrguggwww
wgrgbbrrgbuuggrbbwbuuwwbuurugugwwbbwuruwbuwbrg
grubbwrbruwrwuwbbwrbrrggrrgrurrgruuubbrrbbuu
rubbguguuurgugburuugrgbbrbubwbbbwrbwwgbbgruguwrrrwuguguu
rgbbruuburrgubuugrgrbbgwbrrrggrggubbwgbgrgubrbwug
ggrgrubgbrrrrwgugrubugwgrrbuguuubgwbbbuguuubbbww
rbugrwgrbwwgggrwwruuwrwbbruwbbwbuuugbgubrwwb
rgburbuuuggrwgbuuwbbwbuwubgwrgrwgwbrgbbbguguurbw
wrubuwubbbuurggrwgbubbbguuugguugbrwrrbrruguurbgbuwuwwrugu
wrgwgwbrwwurgurruuuuwbubbrwurrgrwwwuwbggubbwwbgu
gubbbuurrgwgrbbuwggrgwugrrrggrgwgwgwwuwuwbuwgruruubrrubg
burrruuwgbgbwgburwwrwrgbbrrbgrgrbwguuwuuwrugurubrbuwuuw
rgbbgwgbuwuuwwrwugbbubwugruuwwbgrrbrwbwwbbrgwwgu
rwuurgrwubgrurbwgwbgwrbbbguwbggbruuuuwgrgrubwrgwbbbubugrur
rgburrubgggrruwuurguubrwgwggugbrrguurrggbwruwbgurbrubuwr
uuruuwuuwwbwuruwugwrwbwwrrwrgwrrwwgbbbugwuwrwwbr
ubbrurrrgbbrbgburuuubwgwuuwwuuwrwrgbuwwuwugwggugggrgwbw
wgbruwgubggurbuwwwurruubbgbubrbgwbruwugrwgbgwrgbgbbbgg
grrwrrwbguwugugwuruwuurrubgbbgrbuuwwbrgwrrrbbwguggwbubb
buwrbwrbwububrurgugwrgrgwbrbbrrugbrgbggwuwgbrrwugg
rgbrwubgrwbuggwwgurgwuguuwrbgbrwwgrwrurrrrwr
bwbubwrubwgrgugugbugruuugugbwbwrgurwubbggwuub
rgwuwwwgubruguwwuruuubwrbgrggwgurgwubguggguww
ugggwwuuugbruggwbbwwrugbwbrgbrbwuwwggrruwbubgwgrubwruuww
brrbwrgubbuggbgwurgrwbrgrrwwugrrurubrbbwgbwubuguub
rubbrwwbrguugurwrruggbbwggrubruggrbbbbubrrwuwuwbgugwrrugb
bruggrbuurgruuuuuggugbgrrwgrwwwggrrrwubuggbb
rgbwrbwugwrbruubbgugwbwbwbwgrwbrbrubuwgugbg
rwbuubbwrburuwbgbubrugrwrugwurrwbgubgubrrrbgww
wrgrrbuuurbruggrbrggggwubwbugwbbbugwuwrrbrugwrrgw
uwbwrrbbbubbggurbuwurwgwbwwbuwuuwrurbbwgubbwwrb
bgwwurgbugubrwbubgbgwrgubrwrrbgrgwuuggbuuwgubuwrwb
bbwubgwwbuuurwrwwwubrrgbugugruurubwgugbbwbgwub
bwbrrrugggrbbwbbuwwubwrrgbrgwrugwrwuwbrgbggbu
uugbbbuwgubrbgbrbrwbbwwrrwbbbrugwuurgggrubwggwuwwwrr
grgggguuuwuuuwgwuwrgbbwrwrbrrwrbggbbbuwwbgrgbwbbgr
urggbbguruwbggbbwuwuuururruwgbrurgrwrwwuubbrrgruwwrw
rgbugrguwgrgggrruwbbuburrrbgbwuggbwgrubgwgugrbwrbbubr
wrwgbrugwrbwgrbbrggwuwbwrwggwgguubwrwburwrgbwurbbuwbrggw
rgbwbbgubbgrrubrurrurrgggbbgurrwrwrbgbrbwururubgwbwwwr
uugwgguwbrrgubbbuurwugrbbbbrrubgbuwgubuwrrwrbub
rgbwubbrbgwrgguuggbwwgrggwgrbrwrrgrbwwbguwbrbuguugrurw
uwbuwrgburggwurggbwugbbuwgrrggwwgbrgwbbguu
wubgrggrrgrwgggubggubggwbbugrruwuwwwbrbbwwruugugbwbwuwbwr
bgbbwrbuwrrubuwbrbbwggbbuurbuurugwuwrburrrgrwg
gruwrwbbrbururwrrruubbbgrruguugubbwbwbguwugrubwg
uubguurwugwbbwgwuuwugrwbwrugrgwgggbubrwwrwbwuu
gwgrwgbwuwuruwwrbrgbwurgbwrubugwwbrbgbggbubguwrbbugbbw
rrbgwwurgwbruwrwbburruurrwrbwbubbrwburwgwggbgb
ubwbuguugbwgurgubggbugrrrwgwwgggrggbwgbwbbruwgbwrur
uwbrwbwburbbbwwgbgugrrwggwgrguuggurubugwuwggbu
ubuwuuggwrbgwwubugburbgrgurrrwbbgbwwbrburgrgrrgwbbwwgburg
wrwrgububgbgrrugbuugbbuwgrggburrbrubuwwrbgurbgwrwwuurrwgbr
rgbgbgurgbwggugburrbbwbguwrgburgruuwwubbrggwubrwrruu
bgubbrgrwbwgbrbuububbrgwuurgbburwrrwuuggwgwuurwurrwrggw
uubbugwbwrggrbrugwurwruwrgbbgbubwguugubwwurrbbgggrrurbr
wbwrrwgwbgbggbuwguggwgwwrurwbwrwugruwgguwgbwrgbwggbwuu
ggwuuuuubrwrugbgrrrbbbbrrrguwwbuwwuuwggwurruwwgugruwwubbu
rgbwurbbwurwbwgwwruuwbwgbwwrruuuuwwwwggbbrrgguwwgrbbwrbr
rburuubwggguwgugrbbuwgbbgbbruwrgwuwbgbwrwbbrgwrurwruruuubu
uugrbubuwrgrrbbgwbwwrbwbubwgrwwbrwgggruguwrgggbrbgbbubwr
gwwbrrbgbubbrgbrrbwrwuwruwgggrbrgwuggrrrurgrgrwgggubggwbub
wubbrrbwuuurgrguruwwrgggrugbubggugbbbguugrbwurruwu
ggrrwubrgurgwbugbgrwbbuggugwbbggbuwwbgwgwu
rgwbgrwwuwbbuwgggruuburwwbrrbwruuubggrubrbwrwwrwwwrr
uwurbwuwrugwrwgrruwbwwwwbbuwubwrwguwwgrbbgw
gwubgurrrggbuguwbbugrggruruwurbwrguuwwwbbwwruruwbrru
bwgwwgbwbwwrbgruwgbuburwgbbrwguwbggruwwwurb
gurrrbwwbrguwbubuwwbggwuuguurwbrbwgguugrbwwgrbrb
wubgggrwwgwwruguuwbuwwwbbggbbrrwwrgubwrggbuguwwgwwgburrbw
rgbruwrbbubrubbwgrwbrrruwwuwruwrbgrwwbwgguguwbgbguuuurrbruwu
rgggguwgrrwggwrurubbwbrwwwguuwwgrwguwuugurrgrg
wrrgrrrrwbuuguwbwurrwwrbgwbbrrbrbruwuwgbuwwbr
wbwwbgbrwgubbgwgbgrbubwwbruugwgrubrwruwuugw
ubrrgbwgwggguwugbbgrbrrurgbwbbgubbgrbburuu
uuruwwubuwubwruwgrggbgwgbgurugrbrwuruuuwrgwwgbg
rbgbbubwwuuggggwrwrbggwwubggggrwwubbbrbguwguugrr
ruwbwubbrbgwuggwrbrwbubuwurggbubuwrwgbwbbrgrgrurbbgg
rbbgruubggurgbrgrbuuuuurrruguwwugubuwugrrwgrubwugubbuwrgu
gwrwwrgrbwwbgbwrwrurubwgwwrwubgrggrrwuggwwuurrrwurrg
rgwubwwwuwgrbuwgbwbrgwbruuggbbrburguwwuuwgbuwugubgbrb
wbrbbguruwuwguuuruwgwubgbbgbwugbrwrubwwubruwug
gburubbwwwwurrwwbbbgrgwuugwrrruugwrbrgwgbgwuwgbgwbwg
rgrugggugrbrgwwwgubwgurguwrugbbburwwbugggwrbgwrurbr
rbubrwwrbwbgrrrwrbbuwwuwwrgbwwgbgrbgwwwrurubrgrgbbbugwwg
wwbrwwubrgwugubgbbruuwrbwuwrbwburrwuuubgrbubwuwb
wrbgrwbgugurwggbgwwwrgrgwbwbbugubrgggbbwwuwuwb
rgbbgwbrrwubugruurubwbuugurwbrggubbggwbrgwuuburbwbrbwgb
bbugubrbruuwbrggrrugwrbwrurbgbubwbuwbwrubr
gurgubbwggwwuurggwrrrbwrbwwbbgwgbwrguwbuubwguwu
bwwwbgbwwrwbbbwgruuuwwwuwrgwuggurgbwuuwwwbu
grbrrwbwgwrwgrbwrgwgggrbgbbruguubwubgurbgrrw
grbrbwbuubrrgbwrgbugrbgrubuurubwgguwuuuuwgwrggwugbrgu
gbgrgbugggwuwwbruwguruguggguurbbbggwuggwwgurgwbu
uggurrgbwubrrrwrrwbgburubwgbuwgwrrbwggbuguggrwbuwgugwrrw
rwuguburbwwuugugwurbwugwgrwbugwwwbwggwuggb
rgbwbbbrrrguwbrwubburuwguggguuurguugwubwwrbgw
uggrwggwwuubugrbugwwwbbugggrwggwbuurbrgrgbwwwgbuurgbb
grgwgrgwrrbrwwwwrgugruugrrbwubrwgguwrurrrwwgggrgggg
wrrwrrrbguwwrruubgubugrgrbuuggbuuuruubguwrrwbwwwgrwggww
gbrwwwwrwuggruwbwbwrrgbbrubgubwgwwbbrubuuwgurg
rgbwgrrbwwurbgbwurbgrugugugwgwuubruuruubru
bbbrrrwruwruwrurbgugbbgrgbuwwuruwrbbgrguwgrgwgbgggggw
wuwuuuurgbrwgbwbruuwgbbbwrbwubuwuuwbruwrrbwuwbbwgbrwgubbb
rbwgwgrbwgrubgbubuwrbbwwwgwrurgwwbrwgruuur
gwbrwgwruggrwrubbbuwrububgrwrbwbbgbuwuugbwgwgbguug
wwgugrruwwgrbwugbrubuuuburwgggwgwrbwrrbgugr
rrwurbbbrbgbbwgburbwbrwbububrwrbuburwrruruugbwuuububbbgg
ugwwbrbwwuuburbubggbwrbgwwgbrrgbuwgwubugrgugwu
rgburwuwruwrrbgwrwrurwgrubuuburbbgwwgwuurgubrwbuggu
ubwruuwuwbguggubbugrwrurwwbuwggguggburwrrwrb
wrubburggrbubwwbwgrbwrwggwbbburuuwrrbbrubruuwrrgbuggr
urgwuwrrrggwbuuwbrwwrwubrwwgwbrgrbbrbubgwuwruuubrrbwbguuwg
bwwwbubwbwwbrguurbgruwwbgburbwwgwwgwgwrwwgggb
ruurbbbbbbwwrrrbuwgwrbrbbbwrbbwrugrwgbgrbrwgbuu
uuwwrrurrbbugbwgububguwurrrguugwwbubwwubburuwrwwgbu
gwwgwwuwwwwrbguwrubgbwbuwbwguwuurbwugwrwgwu
wuuggwwgubwbrbgbuugurubwruuurgrwrrruwgurrubuuwugguwbg
wbubruburrgrbubwrrurburwrubgrugwrwgwuburbgbuurrbu
rgbgbgwwwgwrbgwwuwrrugubgubwurwgugguurguwwbgruuuugwubu
uggugubbuubrguggurrruuwbrrbuuwbrwrwugbbuurwggrruubburr
gugububrgbrugubwbuwuwwruuwbrrguuwgbwuwbgbgrwbrggrrw
guuruugggugwuubrgrubgbwwgwurubwbbwuwuuurggwbbgrrwwgrwwu
wbbbrgrbbbbgwggburubbbwuwrruwrgururwbrwuubbggurwurrbgggu
bubwwrwgubrrbgbgurrggrrgurugbrbwggbgbbwrgwbbgbgb
ubuuwrgbwgbbugrrbbwbruwrubgurrbgrbbrbrrbwwurgwbgbrgb
grbubbwggbggwggwwggbrrugwrgubwbuwbbuwbbgbwgwrbrrgbbbggw
bggwrbubwrbbwbrwgbuwruruwuwububguwbgbbbggw
gwrgwbbubwurwwburrwuwbburbrurwwurrbuugbgbbggwugwgurru
wrwwwurwbbbugbbugggurrgwuburwrbgurwbgwbugbrgggwuubbbb
burgggwwruuuubwgrbggrwgbrbgrrgbbggwuubbgrbug
gbbrgrgbrwrgbbrugwrurguubgbbrbgrbrwubuuuruwuwbrwugrgg
grggbgruwugbuburuwbgrrrrgubwbrrrwbubruwwrbrrbbuubuwb
rgbubbuwugrbbuurruugwuwrwurrrrwbbwwgbrurrrwgurrwg
urgwbgrrgrbubwrwuwgbubbrgwbwwbwrrrrwuruwruwgwrurrwgwurb
rrwwwggurgbwwubbuuuruguuuwbbwbrrggwwuuburubbwgub
bwggbgubbubwrggwwbbwwrbuwbwubbrbbubgubggrgggbbgbrwb
bwbwgugbugrgbrwwgrbrbugwubbwurgruguwbrbburwrugwurgruuug
wgwwrrgwuuburwbubuuuggurgubbbuwwbrrbgwwuggwgbbrgwurg
grwgrrwrgwurrbbwbrwwwwrwuwuwgwubbrwugbwgbuuwurbrbguug
uuuuuuggwuugwurgbrbrbbwbrbbubwburwbwugruugwuwbbrbuuwwu
wbuuuuwwwbguwbggbgrrwwrwurgwgbbgbbggwbgbgbrbgugwbgrrgurg
buurrrbgbwrbbrguggrwbgrbubgbugwbwwruwgubgwwb
uuubwgggbgrguwbbbuwbwbuubuuwgbrrubgrrrwguubub
bbrrwwwggwbwggubuuwrgrbbgwururgurbwrgggrrgwwgwgburgugrr
rgbgggrrbwrbwgguuwwwbgbruwwwgrgbwrrwuwuwgww
grrguruuggubbguuwwrbbuwubgrbubrbbuubwbbburbwguwrrguw
ggbrrwrwgwrbwuwburubwbgbrguuuwrugrrwrwugwwubb
wwugwurruwubwwguugrrbbuwuurugwbugbrrwgwguugwuwbbrwwbuguuu
wrbwgguwgguwwgbruugrwuuuwgrggbuwbguuwuuwbguw
uurwbwwrwuuwburguubugrbggrgbbwrrubwguggbrbrrwgrwwbguu
rgbwwuwguuubbgurbbrggbgrbgugwwgguuruwubgbbgrbuuwwbwbbw
bwbwbrbrrbgwwbwgbugrgugrbwburrrgugbuuurrrwruuguwg
ugbgwrrrbrwggrrgrwuugbrgrrbubgwrwuuwugrbbrrrwgwrrrggrguwgb
ggwwbrwrurwrburrrrwwwgwrwrruuwbbrgwuwbbbwwrrbb
bwwwrurbbwrrrguwuwwuruwgubuwrugugwwbwrbggurwbuwb
urbbuubwurwwrrwurgwwggrubwwbrururgrrbwuubuub
uuubgwwbbgbgurbgwbugugurugbwruruwggbwrugurrwu
rgbwbrbwrrbugrrwurgwwuguurwwbbwubrbwwbgwugbgwgbwuwbwgguub
rgbwrgwwuwgrruuwugwurwbwbuwrgrurbggbwuruggwubwgguwwggbruwg
bwrgubgugwbuguwuwwwwbwbgwuwgbgbbbgwwuwwrbubbbbuwbwu
wruruwuuwwurbwbuuguruwrwbuguwrwwrbrwurguubugurbbbgwwrwb
rbrrgguwwwuuurwguubbgurwbrugrggbruwrrbubrrbuuwbgrbwug
rbubgbwwwgwrugbwubgwrgggwugwguwgrgggbbbrgbbgbuur
uuuuwuwgwuwbgrrwuwrrwgubbrbbbwbgrbwrbwwwrrbuurggrrg
gbbwwrrrwbgguwggruugurrwbwgrwwgbgwbgguuwbbuubwbrr
wuuuggbubwgbbgurruggbubwurwwrruggrguuwburwrwguguubgwbw
ubrugwgwbgwrggubbbgrwurwbubrrurruugrgrbgbbuww
uubgguwrrgwbbbgbguwubbwbruugggrubrrwggrruuwgrrwwurrgr
rwrbwwrgwbrwgggugbugwrwgwbubgwbguurbgubwbwubwwuugbgg
wurwbwurugbgbrrbwurwrruburgwbrwbrbbburuwggrwbuurubrg
rrgrwugwwuguwuugwuurbbrrgguwbgruwbbgwburwrrwbwbuurwugru
ruuwurrbwugbugwwbwwugbbgwgwrruubuuwbuwwbugwwbrbrwgrb
uwwburgurrubwrugwgwwwurwgbwguuubgbgwbgwwrwuuu
wruuubuuwuwggbwwrbwwrurbggwuuwbwuwurwwuururrrwbbggu
bubggbrwgrugruwuuwubbwwugugurruuwgbbrgbwbuggbubwurbgwg
gbubugggrurguwrgbgrwgwwwbuwrgruwwrwrgubgrguwbrwruugbubr
bggwuwgbggbrwgurgrubrubgugburwgbwwwggrbuggwburgbgrbbuwwgru
bgwgubrwuwwrguuruwugbrwwrwruggbwwuwrubgrgrbbbruuggwb
wwgubugbrguwgbrwgwbwbbwrgbrgbrbguubbbgbrgbb
gwgrgugburwbwgbwgwrrruwwrrwrbbbuggwbrguubrrbuwrgurgbwugubg
bwrugrbwbrurbgrbwurrguwbbbuwrwwwbwrrrwbwbrubwruwb
bbwwbrguwgburrggbbrwuwwwwwgbrrwbrgwuuugrwuwuguwugugbw
ggrurwubgrwbrggwuurrruuuuwuuuwrburbrrrbgwwgg
bwwggguubbgbwbwbrwbbgwwugwuwrrgbgbgrrwuuguurbg
urrgurrbwbwrurgbgwbrrbbwwburrgbugbwgrbgbruuwwbbwgbgwu
gwwuuwuggurbuuurrububuwwbgrwbburbbbuubruugbu
uwbruuubwwwgugbrwwugruugwggrrrbuwbgugbrwubu
ruwruuwrurwguuwwgggurbwgbggrgurggubwwbbguug
urrbrgrggurbwwrrwgubrrguwrurbbruwrwrguuuugwbggbrwurgu
ggrwggrgbgwuubrurggbggwwgwuuwuguwwrrubrubrrgrbbru
bggbbwuubwbburgbgwgrrbuuwgwurbggguwgrbwwwuwrguwrwbrrgrrgr
rubrguggwwuwggbwrubbrbrugwgbrwbgwwuwrwbbrrububbgwr
urrruugbrrgbwgugbbwwwruuubrbwgrggrwwbwgbgwrwbwr
bgbwrgwbrwrggrgwgggwurwugrrrwgbwubrwbrgwug
bwbgwrururruguuurbbbwwwgwugbwwgrwuuurguruwrwuguw
wbrwgrwrrbwuwbgguwrgrubbrugrrrbrwuuuwbgubbgrugrgwuwuwr
urgrwbrbggrugrbbuubgrbrwggwbrurubrwubrbbwbwbuwuwruwb
uguubbbrgugwwrrurrwwbgwgrrwgwuuugurrwwbburwwbwuguruwwb
bwbgggwwrgbwugguwrgbrggwrbbbrwwuwgugbrbwguugwgr
uuuwurrurwwgbrbwgrggwgruburgbuurwrwugwggugrgrrurrr
bubwgurbrgrbwbbgugggwuugrgrbrwrbgwuwuurggwwwgubrg
wgbwrbwuwbgrruruuwgggrurrwwwrbrwrbggbrrwgwwburr
bgwbwbrrgubgruwuwwbgwubgurwbrbrwrbguugbwrubugguguuruggbgr
buubrrwrubrbubbgbburuuwgrgbggugrbuwgbubgrwbwubrr
brgwbbgwrrgbggubrbwggubbbuuruuwrruuruurbuguubrgrrbgrwbwuu
bbbgrrggruwrggggbbgbbbbrrrbwggwbbwwrwgwurwruwguwr
gwbugwbubbuuggrgbwwgrwwbgbrbbwwrbggrubgggbuggggbwbruwgu
rrubuwwwubbrwubrgurubwrgubgwgbgwwgubggrbbuw
gwgrrugwwrrgrwrguwuurwugwgugwbbwurruwgbgrbb
bgrrwrubwwbwgrgbgwguwruurgwgggguubwurbbuubwuwwuggrrwrb
grbrrbwggwubbrbbubuuwbrgurugrgbrrrubrbrrbbguwgw
brbwwwrgbgwubugugwwbgburbubbgbwubrrwgwgbgugbubg
uwbbrwbbwrrggwwwuwggrgugubrwgubwwuuubbwugw
rgbuuggugggwggrgrruwgrrrbbwrugbugbwubgwwrrrrwbrbuguug
gbbruurrggbugugbrwbrrrrrwwwbrgbgubuwwbrgwuwrrbgwwrbww
gwuruurruwwbuuwbgguggrbgrwwbwggbwrwwwwrgubrgbwbwbggwgww
rwwwubgbbuurwwuuuwbwrgbrbwruwurwwwbrwuwrwurrr
gwwrrbwrgwbuguububwwrgurbwurugwrgwguuuwbuuwrbwgbb
rbgwwurugggwugbwbwurbwrugrwbbwggrrbwwbbbgurrrbrrrrubwguu
rbrbbbgbbbuuwbrgrbgbuurwbwgrurwggrrurwwrbrgwwwrbgrgu
rgbrgrgwwgwuuuwbbbbrrwgruuwgbggguuggugrgugrgbrgbrbbgugur
wbbgugbwgrrurubwgwuwuwbbbgwrgbgubrgurgrbwgwrrg
rwrwwugwuuwwbwwugugruguuurguuwrwrrwwrrrbbbbrggu
rggrugrrgbbggwrurgbgwgbbwrwbburwrrwwruruuubg
rgbruburwrbwwbwbrbwwgrrbbwrrrrurbwbuugwwru
urwuuwggbbrrwbgwuurwurbbbugruuwrrwrgwwbubwrbburrrrwrrggwg
rgbguggububwbrbwwgbuurbwugbrbbubuwruurguguwbwrbrurwbb
rrbrbrurwbrrwwuuuwbwbgrubbwgrwbwwurgubuwrwwbwwuuu
wgbrbguwrgbgbbuwrrurbwgwrruwwbuguuugrugbubwgrur
uwrrgrbuubggggruruwrwuugwburuubbgrrbwgwgwbgwubr
bbrwubrrgrrgurrubrrgrrwuuwrwbwrwbugwwwururrrwuuwurubrbg
rgbgggubbubuugggrrugwrbruggwgrbwubuubggubwwuwwuu
gwgrugbruwrwgbburbuwgbuburwguugrbwwrgwubgruruug
rgbwbggrgbgrbwruggguuwgwgrrwwguwwruguuwgwrwggbb
wbgwgwuwgwuwwugwrgbuwrubbggbrrgurugwwbgwrugbrwb
rrubugrguuwgruuwrurububwrrwwuggbbgwbrrwwurbwbbggugwbgww
rgbgbubwbwbrbrbgubrgrrgubugubgubbguwbbgugrw
gwwgrbuurbubwrurbwggbrurgwuguwrbrwgrwurbbguurgbrwrwbb
wbrubuwuguurgrbwrwgrrbwgwrwbgwwgurrwuwbggwrurugb
rbwbbwgwwwrwrgwrwbrbwwgubguwwbbrrgwgbbwbggguwrwwrgbgwu
uwuwbbbbrgwgbwrrwrwububuwwbbrwgbrwgbbubwugbwr
gwbuubbgbrbbwbrwwwbrwrrbwwgggwbugbwrubrrgrurwrrrwugruru
rrwwbrgwbgggbwgwbrguggwggwguggugbruuwggwuwuwwruurbrwbwuuug
rgrbbbrwuwwgwrwbugburbrbuwgburrrgggbrwbwwrbgbubgggwwgwbwr
bgbwbrrrurgbbbbuwrgwuwuuuuuwggbrrbrwggrwuwwugrbbr
rgbubgbwubruuwuurgrrbwwbgwubgurbrbgugrgwggwgwggrrbbrruwugrr
grbrurrrrubbgbgbbrbbruubwgruwubgbugwwgrrwub
gburbwubwwbuwbwbgwbbgruruwgruuuwuuurburwuguwubbrwur
wwgwbbruwguurwbggrrugbrrbwwgruwruwuugwrwub
rrgrwbwurwurgrbggwwbgrbwwwgbwruugguwgubugrwrrurgbubrggubbb
wgwrrwbgrruurbbugrguubrgrrrrbwuugbrruuugwwwuubbbwwgwbbr
rgwrbwbbbggwbgrwrwggrgbgrugrwrurrburgrwgwrbbrwbr
ubburwwuuurrgrgwwbbwbbwuuugwburrbrgwrbgrbgwrgwggwubwburub
ubwurwuggwwrgwrrrrurgwrggwwwugbruurwrwgbrugwggbwuggrrwurrg
urrrgwguwbgguwbburgruwwbbgwguubruwgbugbrwbggwbwu
brrwwwbuuwuubgggwburrrggrrgwruubgrbbggrwgg
rggwwuuubgrgbrguggrrbgururbuuubrbrwuwrbwgrugbruuugw
rbgwbgbuburrbwrgrwguwuuggrbbbuurgwuguwbugrwwuwgrbw
wwbbwrgwbrgwrbrrggbgggwuubwbugggbbwwubggubwwguu
wurbuuuuurugggugurbggrrubrwwgwugrwwugrrguuwrgrgbr
uwwrbwuburwubbuwuurwbugrwgbgbgrwwrgrbwrwrw
uwbgbgbwurbrbggugggubbwrurwgrugrwgwwwrrrgggrgrgbrr
rgwwrbrgggrbrbuuwbbgbrbwwwrgrbwuuwuwbbgbgbgugrruwbwrbwur
rgbbbbbwrbubrbbwbbubbrgrubrwugrguurwuuuuug
ubburrrwbrggrwbwwrguuuurguuguuwwbwrbggbwrruugwbbwrbwbww
rwbrwubgrwwwbrwgrgugbgrgbugurbbgrrbgrugrubugwwrbuburrurbu
ubwwruwrgbwguuwwubgubuwrbrurbwrrwubrbbbwrrurbrb
wwrgbrubbuwwbruburuurbrubrbbguwgwgrbruggbw
uruugbbgwugugwwggugrrrggwrrggbuwurwgrwruuuwwbubbuwbgrbgguu
ubgrgwwrwgrggggrbbgbrwggwbrwgrwggbrrgruwwg
ubwbugwwuwwuwrgwwuwgrruubrubrgwuuubwbbrrubrbb
uurggbruurrrwgwuwwbgubbbburbbubwuguuwwruuug
rgbwbrwrbubbwbwwgbuugrrbugbbwwwgbuugbbrubrgbgwbwwrwbbgrbgrrr
bbrubugbbbbrggubburgbrgbbggwguwwguugubrbbgrwbbgg
gurbrwwwrrubbbburubgurubbbgbgwurrruuurubgubwrurrug
gwwrrbwgwrwbuubrbgwrggwuubgubbwwubrbrrwuburugrbggrrubb
urbubrrbgwugbrwbuubgwwwwbrgugrrbgrbwwgurrg
wurbwgwugbbwwugwgggburrguruwwubggbuwwbguggbugbwbrru
gurwguwgbgwuubwbgrwuwgrbwbwbbrruruwrgbwgwggggbgwrgburgrrrw
rgbuwbwrwrgrurubuwbubgbgwrwugugubuuwurrugubggwwggugrgurgrwrr
brrrbbbbrgugwgwubwgwuugwbbuwbuurgubgwrbubugwgrbub
wgbwruwguurwgugbugwbwuugwuwwggwbbwwubrwgruwbgurbub
rgbrbwwwwrrgwugbuubbuugugrgbbgwwbuguwubuuubgwbwwbuubggbbrgr
rgbbguurrguwurbrggguwbguugwurwuugugwguruubbwbguwuwu
brwububwuubuuububbguwuguurwguwuurrrwwbrbggr
burwgrwurgrwwrrubrbwgwgbwgugugrbbgrrrrubrbrwwwwrbbubgbw
burwbrubbuwggrwugugbrbguuguubggbbrbwbruugggwwrbbw
rbgwrbuuurrruuwwbbruuguubgwubububwbwgbgbuwwrruug
wbrwugguwwbrwruwbrurwrbggrwgrwrguubgugrbrwggwuur
grwgwbgwbwuwwwbugwuguugwrruugurwwruurrrrrbguubbrubru
bbubgrrwuuugbwgrbrwwubgbwurrurbwwrrbgwwbrbgbggb
rggbuuwubugrwwgbugubwggugruwwbguurgwwrrugrbwrgruw
brruuwuwurwuubbwurwbbbwuwwgwbwbuwbrruwwrrgr
urugugbuuuuugbrgrgbbwuwgbguururrgwrurgrbgbbrgwwbrbburw
uwbbubwurwbgwgwruurgrwbwgrbwugrbbbrrrbgwruwwbgbrwrbubb
wwbrwgrbbwrbugbbwbbwbrwuuubrwguwbugbububgbruwuuu
rbrgwwrrbuwwbwwwrgwbubruubburwwwggbuwwbbbbbbugbwuurbrggg
rbbrbgrgrguwwbbrubrwbwwgbrwuruubgwrbrbuubwbuurrwg
rrgbwggbwwuuwrrgrgwruwwbwuggrbggbrrrbrgwruwwgwrbgwggw
bwggugbbwrrugrrgugguuwrguubbbgwbgrugwrubwrbwgg
rurgwugwwbbruugbugwbbgbbwgwwwgrwurggbuuggwrwgubrb
ubwuurrgurgruruwururbwbbwuwrrbwwuwwwggbbrbggbubg
gwuurrgrbbururgrwburuuwwbbrrbwbgggburrbbbwwrguuwb
gwrgwbuubbugruwrwrrbrwbguuwububrgbwrwugwwrw
bwubwbwgbgwurgwubwurgrbgbrubbbguurugrwuuggw
rgbrwrubrbugurwbwubrbgggrbwwrubgruuruwbwrgwru
bwwbgwguwbbbuuubuwwgwbgwbgwrbuuugrwbguuruubuwrurb
brgbwgrubbwrrwwubrgugbbrbbuuuuurwbbwrgbggwuwww
rwwwwguurgrbgbbwbrugubbruuwwwwbgruugwgrbrugwgbwwurbuwu
gbgrgugruwwwuwbbrbwbgurgbrgbgbrgbrbuwbwwgbgwwggwrgwgbrgub
rgbuwbbrrggrurrwbuwbwrrwwgrwwgrgbwbururwuggurgrrrb
wugugwwuurburubuwurguguubrwrbuwuwrrurbwgggwwuubrbuwrwbwu
rgbrbrurwrbbwgrrwbgwugbuwbbwbubwrwgbwuwggruubgurrgubbbru
bgubrrrbrubgrwuuugrrwrwbgbbgugbwwrrwwrurrgguwrggruwwuwww
uwurbwwggwbwuuubwbggwgbuwbubrbgwubbbwwgugbbbubgrwwubr
wbgwbwggrburbrbggbburwwwubuwrububggwrgbwwggwwgurrgurw
rwbgrwurwbwrwurwgguguwbbwrgggggwgbwbwwuuggggbrrubwrwrrrb
ruuuwrbwrwwrbgguurwwrrguurwurbgwuuggwwgwgugwgwubwbrbrubug
gwrbwrurbgrggwgbggrgrbgrrubbggwrugbbgrrrgrbuw
bggubbuuuwgbwggwwbbbuguwbugggrgurgrbwguubwrbrwuw
uwgbbrbbbrwwrggbwwurugwgurgwwwbuuwgbugbrggggwgbwuugggwg
rgugugrbrbbrrrgbguuwwgurrbbguurrwbgwugwwwbubwuuwuur
ggrgrwurrbrgrggbrrggubgrrubuwugggwrbgbbuguwb
bguuubgwwgbbruwuwgbguwrrwrrwbrubbuuubgubuu
uuwubuurgburuuubgbgbgbwgrubrbubgwuwbbwrggwrr
ubbbuwuruuubuwbrbgurbgrruugrrgubgurbruwgbubrbuwgrbrubgg
bwgbbbwuuggbrrubbbubggrrburubgbruwrurbrggbrwbrw
wurbbrugubugurugruugubgubbrwrrrgbbwubrwgwgggbbwbgurww
wrgbuburgrbbbwgrrbgrrrguuugwrbugwgrbbgbbbgww
rruwwrrgwrgrrggurbgwurrrgwgurgruwgrwbrbgrbgwguububrbbw
wwbbrugrwrbwrbbuuguurgwrbrrugwwburgugwbuwr
ugwggggbruuwggbubuuubggbggrbbrbuwbbbwuwuwwrbguwwwgru
gbugbgbwgubgbgggwbgwwgbgwuwuuwbwgrbuuwugwuwggurrubgrb
uwbbbggugburgrbrruguwbwuurggrbgugbbwwwuggugwggbguuwwwrr
ruwbgbrrbuuuwwuwrgbwgwbrburrbburrbwbgwwguwuwrbgbrruuuwrg
wugbbwrrwggbuugrwurrubbgrurgrrgugrwgrgwbbgwbbw
ggguwuuwburrurrwbgugbbbwbuburgurrwbwbwrugrbgbrwbgw
rgbwrwuruwgrwwgwggbwwrgwgwrbrugwrbgwurbbg
gwbwwgrbuggrrgbuwwrgwwgrgwwrgruuuburwbrbwgwrgb
brwurggwgbguwggwgrwuwbubbrgwbbrbrurbruwbwrbbuggrgwbb
ububwwrubwrugbgwwggrrrrbbbrwurwbbbwrwwbwrbggbgbug
gbwbgwbwgubwrguuwbbrrwugwuggwgrgruuuwguwuugguwuugwwgrwwg
brbrbwguuggwrbwuwwuwgurwwwubwbrwgrwgbbrubrggwwbgrbbbrugu
guwwbrbggrbgbrrwrugwrrgugrwrwgbrbwubwwwgguwrgburgwurbw
wubbwwgbrwrwrugrruubbugbuuurubgrgguurbuwwbrrwwbwugw
rbbugruurubuuuugrwrgggrrbubwwrbbwururwruruurb
gwrrwbbgubbrbbugwbgwbrbrrgrrbwuuwuwwwbwrbgrgrugwbbbwuwu
wuurrwbgrwwbwugbrbuubgbrurgwugrururwuwuwguwgubburwggb
bguruwuubwuwgwgrguwurgrbggubruwrbgwbburwrbgbrubbwgubbgrr
buwuwrgrrrrbgwgbruwrgbuggurwgbgubguubwugurg
gwwuwruwgrbburrrwugbbbbbbwgubwuugbgbuwwwbgurrbwgrruuwuw
wruwrgbwgugrrgrgggwrrgwwbuwgrurgbgggrugwbrwrwubrbwb
gwguuuuguurwwrgrggubwggurubbbguwwgbrrwrgrguw
guwurrburubggwbguwbbrrwbbgbggbgbuguugugrbbwgggubuu
bbrurrgwuwuwbgbgbrugrrrurwrrgwwguurrbbrbrbgwwbbg
wbrgbrrbbbgwuwruwrgurguubbwbguwbrrbrugwwgbgbgb
ubgugguwwuwguwwwgbbwggbbrgrwgurrrrrrgbuuwgr
wuwbruwbwwbrwwgguurgrwurbgwrugbburbwruruwuubrurggbggugguw
rrggwgrububwgubuwwurruwwugwuuurgugggubwgugbgrbbr
buuruggrwggbbubgguwggbbgubgrrbwgwruwguwgwrbbwgwbrurwru
gwgwwbwbrgburubwrbugubgrrwubruggwwuggurgrbub
uwgugrbuwrugwrwwburgrwuuwggbuwggguwuburubgurrgbruuub
gbgggugbwburrrurbrbgrwuubuwgwguurbgwbgwuwububwg
rubbbwbwwrbwwbgrbbwuubrguwbwbwuuwgubguubugbgrugubgbrwg
bggbbwgrgwuuwgugrbbwgrrubruwrubbbubbugrbwwwbrrruwgb
gwugwgrwbuurwuurwgbbrgguubwuuwwuuurwrwgwgggbwgrb
rugbrurwuggurwrrgugbgwwrwuwrwwuugwbbrbubgbgwwurggbgurgbur
gbbgguruurrubrwubugbrrugubgrgwwrugbrguurwrb
grwubwgrrwwbuugrgwwwrwrbrubgwwubrwgwbgbggrruuwbrb`