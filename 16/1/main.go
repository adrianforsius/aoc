package main

import (
	"log"
	"strconv"
	"strings"
)

var conv = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func main() {
	var binary []string
	for _, c := range input {
		binary = append(binary, strings.Split(conv[c], "")...)
	}
	// num, err := strconv.ParseInt(input, 16, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("num", num)

	// val := fmt.Sprintf("%08b", num)
	// binary := strings.Split(val, "")
	log.Println(binary)
	packet, _ := packets(binary)
	log.Printf("packet %+v", packet)
	log.Println("sum", sum(packet, 0))
}

func sum(p packet, count int) int {
	count += int(p.version)
	for _, pack := range p.packets {
		count = sum(pack, count)
	}
	return count
}

type packet struct {
	version    int64
	typeID     int64
	val        int64
	lenghtType string
	packets    []packet
}

func packets(binary []string) (packet, int) {
	log.Println("binary", binary, "len", len(binary))
	versionBits := binary[0:3]
	versionString := strings.Join(versionBits, "")
	version, err := strconv.ParseInt(versionString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("version", version)

	typeIDBits := binary[3:6]
	typeBitsString := strings.Join(typeIDBits, "")
	typeID, err := strconv.ParseInt(typeBitsString, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("typeID", typeID)
	packet := packet{
		version: version,
		typeID:  typeID,
	}

	// var count int
	if typeID == 4 {
		valStr, count := parseLit(binary[6:])
		val, err := strconv.ParseInt(strings.Join(valStr, ""), 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		packet.val = val
		log.Println("literal *************", val)
		return packet, count + 6
	} else {
		lengthType := binary[6:7][0]
		log.Println("length type", lengthType)
		modeLengthBits := modeLength[lengthType]
		log.Println("mode lenght bits", modeLengthBits)
		subPacketsBits := binary[7 : 7+modeLengthBits]
		subPacketLength, err := strconv.ParseInt(strings.Join(subPacketsBits, ""), 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println("sub packet length", subPacketLength, modeLengthBits, int(subPacketLength)+7+modeLengthBits, len(binary))
		// log.Println(binary)
		if lengthType == "0" {
			log.Println("subPacketSize", subPacketLength)
			// literals := parseOpMode0(binary[7+modeLengthBits : 7+modeLengthBits+int(subPacketLength)])
			// var literalDecimal []int
			// for _, l := range literals {
			// 	literal, err := strconv.ParseInt(strings.Join(l, ""), 2, 64)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	literalDecimal = append(literalDecimal, int(literal))
			// }

			// log.Println(literalDecimal)
			start := 7 + modeLengthBits
			count := start
			for count < start+int(subPacketLength) {
				log.Println("+++++++++++++", "count", count, "lenght", start+int(subPacketLength), "start", start)
				p, c := packets(binary[count : start+int(subPacketLength)])
				packet.packets = append(packet.packets, p)
				count += c
				log.Println("++++++ incr +++++++", count)
			}
			log.Println("count =========", count)
			return packet, count
		}
		if lengthType == "1" {
			// log.Println("subPacketLength", subPacketLength)
			// var literalDecimal []int
			// rest := binary[7+modeLengthBits : 7+modeLengthBits+int(subPacketLength)]
			count := 0
			start := 7 + modeLengthBits
			for i := 0; i < int(subPacketLength); i++ {
				log.Println("next chunck", binary[start+count:], len(binary[start+count:]))
				p, c := packets(binary[start+count:])
				packet.packets = append(packet.packets, p)
				// literal, _ := parseLit(part)
				// l, err := strconv.ParseInt(strings.Join(literal, ""), 2, 64)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// literalDecimal = append(literalDecimal, int(l))
				count += c
				// log.Println("loop count", i)
			}
			// log.Println("sub packets total length", count, start)
			// log.Println("count =========", count)
			return packet, count + 7 + modeLengthBits
		}
	}
	// if len(binary) > count {
	// 	return packets(binary[count:])
	// }
	return packet, 0
}

var modeLength = map[string]int{
	"0": 15,
	"1": 11,
}

func parseLit(val []string) ([]string, int) {
	count := 0
	var concat []string
	for {
		continueBit, part := val[count : count+1][0], val[count+1:count+5]
		if len(val) < 5 {
			val = val[count+5:]
		}
		concat = append(concat, part...)
		count += 5
		if continueBit == "0" {
			return concat, count
		}
	}
}

func parseOpMode0(val []string) [][]string {
	var vals [][]string
	c := 0
	for {
		// versionBits := binary[0:3]
		// versionString := strings.Join(versionBits, "")
		// version, err := strconv.ParseInt(versionString, 2, 64)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println("version", version)

		// typeIDBits := binary[3:6]
		// typeBitsString := strings.Join(typeIDBits, "")
		// typeID, err := strconv.ParseInt(typeBitsString, 2, 64)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		c += 6
		log.Println("============", c, val[c:])
		log.Println("val", val)
		v, count := parseLit(val[c:])
		vals = append(vals, v)
		c += count
		if c == len(val) {
			return vals
		}
	}
}

// var input = `D2FE28`

// var input = `38006F45291200`

// var input = `EE00D40C823060`
// var input = `8A004A801A8002F478`

// var input = `620080001611562C8802118E34`

// var input = `C0015000016115A2E0802F182340`

// var input = `A0016C880162017C3686B18A3D4780`

// puzzle input
var input = `00569F4A0488043262D30B333FCE6938EC5E5228F2C78A017CD78C269921249F2C69256C559CC01083BA00A4C5730FF12A56B1C49A480283C0055A532CF2996197653005FC01093BC4CE6F5AE49E27A7532200AB25A653800A8CAE5DE572EC40080CD26CA01CAD578803CBB004E67C573F000958CAF5FC6D59BC8803D1967E0953C68401034A24CB3ACD934E311004C5A00A4AB9CAE99E52648401F5CC4E91B6C76801F59DA63C1F3B4C78298014F91BCA1BAA9CBA99006093BFF916802923D8CC7A7A09CA010CD62DF8C2439332A58BA1E495A5B8FA846C00814A511A0B9004C52F9EF41EC0128BF306E4021FD005CD23E8D7F393F48FA35FCE4F53191920096674F66D1215C98C49850803A600D4468790748010F8430A60E1002150B20C4273005F8012D95EC09E2A4E4AF7041004A7F2FB3FCDFA93E4578C0099C52201166C01600042E1444F8FA00087C178AF15E179802F377EC695C6B7213F005267E3D33F189ABD2B46B30042655F0035300042A0F47B87A200EC1E84306C801819B45917F9B29700AA66BDC7656A0C49DB7CAEF726C9CEC71EC5F8BB2F2F37C9C743A600A442B004A7D2279125B73127009218C97A73C4D1E6EF64A9EFDE5AF4241F3FA94278E0D9005A32D9C0DD002AB2B7C69B23CCF5B6C280094CE12CDD4D0803CF9F96D1F4012929DA895290FF6F5E2A9009F33D796063803551006E3941A8340008743B8D90ACC015C00DDC0010B873052320002130563A4359CF968000B10258024C8DF2783F9AD6356FB6280312EBB394AC6FE9014AF2F8C381008CB600880021B0AA28463100762FC1983122D2A005CBD11A4F7B9DADFD110805B2E012B1F4249129DA184768912D90B2013A4001098391661E8803D05612C731007216C768566007280126005101656E0062013D64049F10111E6006100E90E004100C1620048009900020E0006DA0015C000418000AF80015B3D938`
