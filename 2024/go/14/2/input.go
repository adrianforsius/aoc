package main

const exampleInput = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

const input = `p=40,73 v=-96,64
p=74,78 v=91,-65
p=100,86 v=98,62
p=61,29 v=95,-68
p=50,53 v=-50,-63
p=43,77 v=-42,-37
p=47,44 v=34,66
p=11,21 v=72,66
p=9,61 v=-26,77
p=46,100 v=-10,-82
p=64,15 v=-5,-44
p=50,97 v=-56,37
p=37,35 v=-95,80
p=33,8 v=60,39
p=56,81 v=38,7
p=60,22 v=-11,69
p=47,36 v=-55,33
p=29,51 v=-39,-69
p=93,89 v=-96,-14
p=67,51 v=-80,-78
p=68,102 v=-46,46
p=66,51 v=-9,40
p=95,28 v=73,-1
p=41,86 v=-97,81
p=15,54 v=16,-33
p=68,47 v=-7,-20
p=75,34 v=40,23
p=4,38 v=27,73
p=14,78 v=-82,-49
p=23,68 v=63,87
p=82,69 v=83,24
p=4,38 v=62,14
p=60,65 v=46,71
p=12,14 v=13,-31
p=14,17 v=20,-67
p=65,64 v=-4,-42
p=2,19 v=-68,89
p=4,12 v=-22,-81
p=17,95 v=20,-20
p=68,49 v=37,91
p=37,1 v=-92,-68
p=25,81 v=16,81
p=96,60 v=-75,57
p=57,37 v=-1,3
p=21,65 v=49,71
p=38,58 v=-93,-16
p=98,52 v=77,-70
p=66,85 v=91,54
p=16,57 v=-79,80
p=6,87 v=-27,-2
p=27,60 v=56,34
p=66,55 v=-52,-63
p=57,59 v=31,39
p=75,74 v=-10,34
p=31,83 v=-32,-49
p=95,34 v=-19,-7
p=81,68 v=-63,-99
p=43,68 v=63,14
p=43,78 v=84,-90
p=78,60 v=38,-86
p=77,58 v=37,-23
p=44,83 v=12,-39
p=50,61 v=-99,-3
p=91,51 v=24,73
p=0,23 v=-67,19
p=52,6 v=-50,42
p=79,29 v=-66,96
p=66,82 v=46,41
p=21,27 v=-85,-34
p=33,60 v=-95,7
p=60,53 v=-58,30
p=58,91 v=-34,-28
p=49,16 v=1,-84
p=7,2 v=-68,46
p=99,97 v=75,-55
p=95,39 v=37,46
p=17,37 v=-55,59
p=90,7 v=-19,85
p=4,95 v=-88,23
p=66,26 v=-6,-4
p=33,19 v=-39,-4
p=59,60 v=46,-13
p=31,67 v=13,80
p=70,102 v=-9,-18
p=95,32 v=-72,-87
p=73,40 v=-11,93
p=8,17 v=-80,-34
p=47,79 v=-48,51
p=74,59 v=-73,79
p=50,5 v=-51,72
p=44,25 v=-89,86
p=0,89 v=-67,-25
p=12,81 v=43,59
p=49,7 v=3,2
p=25,99 v=-39,25
p=42,91 v=-88,-48
p=6,28 v=83,44
p=53,57 v=-52,-23
p=45,30 v=53,53
p=98,4 v=26,65
p=69,1 v=82,-48
p=50,8 v=-99,-98
p=2,11 v=86,-33
p=58,29 v=-10,23
p=25,97 v=72,42
p=67,96 v=95,85
p=61,97 v=97,38
p=33,25 v=83,-50
p=45,1 v=-45,95
p=67,98 v=42,-78
p=30,28 v=7,-45
p=74,48 v=98,-30
p=12,6 v=74,75
p=53,5 v=48,-93
p=13,97 v=78,-45
p=96,10 v=-72,-41
p=23,13 v=14,-21
p=1,46 v=-24,10
p=40,39 v=96,40
p=14,98 v=21,55
p=22,102 v=-99,-39
p=78,98 v=-63,-48
p=65,88 v=62,-74
p=89,29 v=-64,89
p=41,74 v=89,67
p=81,51 v=90,69
p=7,96 v=-27,35
p=83,3 v=-64,-21
p=72,52 v=94,-83
p=43,54 v=-15,-31
p=10,92 v=-27,48
p=30,0 v=-87,28
p=15,18 v=-9,63
p=11,65 v=-30,84
p=13,27 v=62,-47
p=63,63 v=53,-9
p=67,80 v=-14,15
p=15,9 v=-98,64
p=58,24 v=96,-34
p=37,21 v=6,9
p=57,56 v=50,84
p=75,9 v=35,84
p=48,1 v=-60,58
p=70,51 v=47,34
p=5,69 v=76,-9
p=63,4 v=50,79
p=25,17 v=-36,3
p=87,42 v=16,2
p=69,27 v=42,66
p=8,42 v=70,-30
p=0,52 v=87,2
p=3,3 v=-23,75
p=45,27 v=47,-77
p=23,100 v=-13,-99
p=69,19 v=-7,-4
p=38,70 v=6,-89
p=4,22 v=-28,49
p=86,91 v=89,-2
p=90,100 v=-32,56
p=2,38 v=-81,99
p=44,32 v=-96,-40
p=96,77 v=51,-22
p=5,71 v=93,-89
p=90,89 v=24,30
p=27,65 v=-89,-29
p=33,56 v=10,67
p=34,71 v=18,-47
p=66,68 v=44,11
p=55,37 v=-91,15
p=73,6 v=-14,-91
p=45,97 v=-54,-21
p=34,88 v=34,-27
p=65,14 v=67,-2
p=40,5 v=-85,-14
p=8,95 v=15,91
p=5,35 v=-29,-57
p=71,83 v=-11,34
p=79,40 v=-6,-26
p=62,78 v=-3,51
p=12,15 v=-16,-25
p=64,12 v=77,-99
p=33,91 v=68,-45
p=86,77 v=-24,-6
p=55,27 v=83,-91
p=85,63 v=38,40
p=4,4 v=-74,-22
p=50,19 v=-89,20
p=45,20 v=-98,29
p=11,50 v=19,-23
p=43,82 v=5,-82
p=72,14 v=40,65
p=68,35 v=47,-7
p=7,91 v=-78,45
p=28,32 v=-87,41
p=47,48 v=-51,4
p=12,45 v=71,30
p=12,76 v=83,-3
p=74,21 v=-11,-27
p=43,32 v=59,16
p=21,12 v=-39,32
p=76,88 v=92,68
p=42,56 v=-44,-30
p=27,76 v=65,-9
p=61,2 v=-57,95
p=42,77 v=-90,88
p=27,74 v=-39,51
p=55,36 v=-59,50
p=3,9 v=82,-37
p=64,78 v=-57,-62
p=71,83 v=-51,94
p=91,63 v=8,-29
p=43,68 v=-46,25
p=88,38 v=87,-96
p=31,65 v=-35,-74
p=13,86 v=-30,-95
p=7,43 v=-31,-20
p=65,89 v=-60,81
p=10,55 v=69,48
p=66,85 v=-34,14
p=64,95 v=-99,62
p=53,46 v=39,2
p=79,3 v=-61,55
p=56,50 v=-82,-21
p=31,76 v=73,-92
p=13,102 v=-34,-88
p=31,100 v=-43,-98
p=17,4 v=-82,-24
p=45,53 v=55,-33
p=61,24 v=8,-95
p=75,96 v=-12,-78
p=64,7 v=-55,-1
p=82,27 v=-66,36
p=50,63 v=-99,47
p=38,63 v=44,-64
p=62,27 v=3,-4
p=78,40 v=67,20
p=7,24 v=84,-67
p=5,25 v=-28,-97
p=33,101 v=-40,-45
p=38,9 v=-94,-94
p=76,13 v=-55,55
p=48,47 v=21,-61
p=95,6 v=-65,99
p=12,51 v=69,-53
p=88,4 v=82,-71
p=41,76 v=-37,-12
p=51,45 v=-11,4
p=66,77 v=96,25
p=92,100 v=33,5
p=96,64 v=-29,78
p=57,54 v=-60,-20
p=22,33 v=71,49
p=12,35 v=-27,-57
p=88,17 v=97,-1
p=38,61 v=61,20
p=14,39 v=17,45
p=73,70 v=68,4
p=57,41 v=-45,-77
p=73,35 v=-68,66
p=95,70 v=-72,54
p=51,44 v=-57,-73
p=47,2 v=46,-68
p=55,65 v=-59,-82
p=71,15 v=-78,67
p=5,96 v=72,-88
p=10,30 v=73,-17
p=89,84 v=-66,-82
p=21,4 v=-36,-18
p=15,44 v=67,43
p=67,48 v=44,50
p=70,16 v=7,-21
p=62,5 v=-74,49
p=60,52 v=96,90
p=81,80 v=-57,-42
p=2,86 v=-25,-98
p=48,31 v=49,-97
p=25,69 v=17,24
p=68,7 v=-52,-5
p=20,1 v=-30,-54
p=95,25 v=55,68
p=76,101 v=-8,98
p=35,5 v=12,-28
p=80,90 v=37,-91
p=8,51 v=25,-53
p=71,58 v=-51,-33
p=36,19 v=-46,-80
p=80,8 v=37,-38
p=26,96 v=-80,14
p=19,56 v=-34,-33
p=58,93 v=12,66
p=23,65 v=-28,70
p=5,89 v=-77,-98
p=99,67 v=-66,-29
p=33,94 v=-94,52
p=71,10 v=-7,-21
p=67,38 v=45,30
p=30,27 v=-92,-86
p=40,55 v=-99,40
p=25,78 v=-46,9
p=98,0 v=-14,29
p=21,51 v=70,60
p=64,56 v=46,50
p=83,80 v=20,-87
p=32,60 v=-89,-86
p=52,61 v=-52,-46
p=49,20 v=-98,99
p=19,61 v=19,17
p=9,99 v=71,12
p=9,63 v=-80,24
p=16,0 v=18,-38
p=7,18 v=24,-5
p=89,46 v=-89,-11
p=31,64 v=-86,-51
p=60,13 v=54,-37
p=83,46 v=-70,-16
p=4,3 v=-25,-18
p=45,67 v=79,56
p=68,61 v=-15,-13
p=13,64 v=-23,-56
p=81,78 v=-92,-18
p=1,71 v=-17,-49
p=9,95 v=-64,28
p=69,21 v=44,69
p=25,74 v=-86,94
p=71,1 v=34,2
p=58,37 v=-10,-4
p=32,81 v=-46,25
p=21,87 v=-2,-92
p=78,70 v=35,-96
p=86,70 v=32,84
p=42,68 v=72,-68
p=76,69 v=36,-99
p=72,79 v=69,-6
p=92,22 v=-72,99
p=74,9 v=-12,-94
p=52,33 v=-43,45
p=34,79 v=-41,-59
p=26,56 v=10,50
p=59,46 v=-22,23
p=9,49 v=68,-59
p=78,11 v=62,40
p=88,25 v=43,34
p=61,64 v=-53,27
p=17,13 v=-32,-81
p=100,45 v=26,70
p=30,52 v=13,20
p=9,93 v=-81,18
p=68,3 v=93,15
p=77,56 v=93,60
p=48,5 v=-42,-77
p=30,55 v=-87,-33
p=42,87 v=6,78
p=7,3 v=51,-8
p=69,79 v=-61,30
p=51,45 v=-49,35
p=20,85 v=11,65
p=25,58 v=17,-66
p=45,2 v=-39,-5
p=88,38 v=69,-98
p=10,2 v=-81,32
p=26,31 v=44,-25
p=29,15 v=-67,52
p=59,73 v=-54,21
p=99,43 v=32,-48
p=99,36 v=-71,13
p=54,35 v=52,63
p=44,77 v=-3,54
p=32,82 v=-65,-50
p=40,1 v=-46,85
p=12,71 v=20,-69
p=90,78 v=78,87
p=18,52 v=62,-70
p=40,23 v=30,63
p=76,27 v=-75,76
p=88,62 v=91,97
p=31,88 v=64,-8
p=95,33 v=96,-26
p=21,63 v=16,-46
p=35,15 v=6,-54
p=73,22 v=-10,46
p=43,11 v=63,95
p=99,18 v=-66,-67
p=0,71 v=74,44
p=72,66 v=-8,14
p=100,30 v=30,-7
p=57,17 v=-41,73
p=5,3 v=80,58
p=77,100 v=84,78
p=61,78 v=99,8
p=59,51 v=45,-53
p=85,6 v=78,-17
p=49,56 v=-99,34
p=25,7 v=64,2
p=81,1 v=44,-1
p=19,50 v=-80,13
p=1,67 v=75,91
p=22,63 v=65,64
p=25,73 v=75,60
p=73,64 v=-11,84
p=27,13 v=14,49
p=4,27 v=-77,3
p=52,41 v=-7,30
p=34,83 v=-84,35
p=15,62 v=-81,87
p=49,101 v=13,77
p=60,52 v=98,-56
p=95,61 v=-22,-96
p=37,79 v=11,-98
p=82,48 v=20,70
p=28,57 v=63,70
p=28,48 v=-43,4
p=41,60 v=58,-89
p=94,12 v=-27,-58
p=22,66 v=-33,14
p=14,47 v=19,86
p=100,58 v=-70,-13
p=1,17 v=1,76
p=11,33 v=-54,93
p=10,46 v=67,76
p=69,11 v=50,-87
p=37,12 v=-96,32
p=1,23 v=69,33
p=31,22 v=54,-44
p=60,5 v=95,89
p=6,84 v=25,-59
p=2,22 v=76,89
p=63,14 v=71,25
p=54,33 v=-42,55
p=89,34 v=-76,-50
p=35,88 v=41,-34
p=84,63 v=-12,67
p=25,27 v=13,-70
p=47,60 v=-95,17
p=70,27 v=-5,24
p=88,97 v=84,28
p=47,52 v=50,-12
p=24,99 v=66,-71
p=35,67 v=7,64
p=41,31 v=7,-79
p=60,73 v=47,-82
p=83,37 v=54,12
p=14,45 v=19,-10
p=81,80 v=12,-53
p=92,51 v=38,-55
p=14,26 v=-75,82
p=89,38 v=-33,34
p=34,74 v=56,47
p=9,31 v=-35,-63
p=62,98 v=-74,-62
p=14,54 v=-68,29
p=70,41 v=66,-55
p=51,52 v=-57,34
p=0,35 v=61,-49
p=99,54 v=76,-26
p=75,3 v=-82,-26
p=43,95 v=7,28
p=1,51 v=27,-66
p=18,15 v=25,92
p=5,53 v=78,-53
p=58,95 v=49,-18
p=34,49 v=-39,90
p=19,75 v=60,91
p=61,43 v=39,3
p=74,51 v=-13,23
p=32,26 v=3,66
p=30,6 v=-33,99
p=66,86 v=61,1
p=25,87 v=45,-81
p=22,76 v=-86,71
p=90,12 v=32,99
p=74,98 v=93,25
p=16,94 v=68,58
p=12,64 v=-4,-96
p=87,8 v=-75,-4
p=28,33 v=-81,40
p=6,34 v=74,-44
p=65,85 v=-56,41
p=71,67 v=46,80
p=73,85 v=-9,-12
p=11,64 v=64,50
p=73,46 v=48,-23
p=62,22 v=4,31
p=10,22 v=21,92
p=73,67 v=-1,95
p=97,2 v=30,-58
p=96,42 v=-29,80
p=26,53 v=63,20
p=30,7 v=61,-11
p=4,90 v=-77,-48
p=23,55 v=58,14
p=42,68 v=58,77
p=65,68 v=60,-28
p=77,87 v=-15,71
p=40,14 v=-41,-21
p=78,42 v=-20,-20
p=29,95 v=71,76
p=38,62 v=-40,67
p=98,86 v=-25,38
p=3,37 v=27,63
p=41,70 v=39,21`
