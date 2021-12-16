package problems

// ----------------------------------------------------------------------------
// AOC 2021
// ----------
//
// Day 16 - Packet Decoder
// ----------------------------------------------------------------------------

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc2021/lib"
)

type Day16Msg []uint8

// ---------------------------- common interface for package types -----------------------------
type Day16Packet interface {
	toString() string
	toFormula() string
	versionSum() uint64
	evaluate() uint64
}

// ---------------------------- literal package -----------------------------
type Day16LiteralPacket struct {
	version uint64
	value   uint64
}

func (p *Day16LiteralPacket) toString() string {
	return fmt.Sprintf("version: %v, type: literal, value: %v\n", p.version, p.value)
}

func (p *Day16LiteralPacket) toFormula() string {
	return fmt.Sprint(p.value)
}

func (p *Day16LiteralPacket) versionSum() uint64 {
	return p.version
}

func (p *Day16LiteralPacket) evaluate() uint64 {
	return p.value
}

// ---------------------------- operator package -----------------------------
type Day16OperatorPacket struct {
	version    uint64
	ptype      uint64
	subpackets []Day16Packet
}

func (p *Day16OperatorPacket) toString() string {
	out := fmt.Sprintf("version: %v, type: op (%v), # of packets: %v\n", p.version, p.ptype, len(p.subpackets))
	for _, sub := range p.subpackets {
		out += sub.toString()
	}
	return out
}
func (p *Day16OperatorPacket) versionSum() uint64 {
	sum := p.version
	for _, c := range p.subpackets {
		sum += c.versionSum()
	}
	return sum
}

func (p *Day16OperatorPacket) evaluate() uint64 {
	switch p.ptype {
	case 0: // sum package
		var sum uint64 = 0
		for _, sub := range p.subpackets {
			sum += sub.evaluate()
		}
		return sum

	case 1: // product package
		var prod uint64 = 1
		for _, sub := range p.subpackets {
			prod *= sub.evaluate()
		}
		return prod

	case 2: // minimum package
		var min uint64 = p.subpackets[0].evaluate()
		for _, sub := range p.subpackets[1:] {
			v := sub.evaluate()
			if v < min {
				min = v
			}
		}
		return min

	case 3: // maximum package
		var max uint64 = p.subpackets[0].evaluate()
		for _, sub := range p.subpackets[1:] {
			v := sub.evaluate()
			if v > max {
				max = v
			}
		}
		return max

	case 5: // gt (>) package
		sub1 := p.subpackets[0].evaluate()
		sub2 := p.subpackets[1].evaluate()
		if sub1 > sub2 {
			return 1
		} else {
			return 0
		}

	case 6: // lt (<) package
		sub1 := p.subpackets[0].evaluate()
		sub2 := p.subpackets[1].evaluate()
		if sub1 < sub2 {
			return 1
		} else {
			return 0
		}
	case 7: // ==  package
		sub1 := p.subpackets[0].evaluate()
		sub2 := p.subpackets[1].evaluate()
		if sub1 == sub2 {
			return 1
		} else {
			return 0
		}
	default:
		panic("Unknown package!")
	}
}

func (p *Day16OperatorPacket) toFormula() string {
	switch p.ptype {
	case 0: // sum package
		str := p.subpackets[0].toFormula()
		if len(p.subpackets) > 1 {
			for _, sub := range p.subpackets[1:] {
				str += " + " + sub.toFormula()
			}
		}
		return fmt.Sprintf("(%v)", str)

	case 1: // product package
		str := p.subpackets[0].toFormula()
		if len(p.subpackets) > 1 {
			for _, sub := range p.subpackets[1:] {
				str += " * " + sub.toFormula()
			}
		}
		return fmt.Sprintf("(%v)", str)

	case 2: // minimum package
		str := p.subpackets[0].toFormula()
		if len(p.subpackets) > 1 {
			for _, sub := range p.subpackets[1:] {
				str += "," + sub.toFormula()
			}
		}
		return fmt.Sprintf("min(%s)", str)

	case 3: // maximum package
		str := p.subpackets[0].toFormula()
		if len(p.subpackets) > 1 {
			for _, sub := range p.subpackets[1:] {
				str += "," + sub.toFormula()
			}
		}
		return fmt.Sprintf("max(%s)", str)

	case 5: // gt (>) package
		return fmt.Sprintf("(%s > %s)?", p.subpackets[0].toFormula(), p.subpackets[1].toFormula())

	case 6: // lt (<) package
		return fmt.Sprintf("(%s < %s)?", p.subpackets[0].toFormula(), p.subpackets[1].toFormula())
	case 7: // ==  package
		return fmt.Sprintf("(%s == %s)?", p.subpackets[0].toFormula(), p.subpackets[1].toFormula())
	default:
		panic("Unknown package!")
	}
}

// ---------------------------------------------------------
type Day16 struct {
	solution1  uint64
	solution2  uint64
	msg        Day16Msg
	packetTree []Day16Packet
}

func (p *Day16) GetName() string {
	return "AoC 2021 - Day 16 - Packet Decoder"
}

func (p *Day16) Init() {
	// Read input
	// lines := lib.ReadInputLines("input/day16-test.txt")
	lines := lib.ReadInputLines("input/day16-input.txt")

	// test data:
	// line 0: literal value: 011111100101, which is 2021
	// line 1: operator packet (hexadecimal string 38006F45291200) with length type ID 0 that contains two sub-packets
	// line 2: operator packet (hexadecimal string EE00D40C823060) with length type ID 1 that contains three sub-packets
	// line 3: 8A004A801A8002F478 represents an operator packet (version 4) which contains an operator packet (version 1) which contains an operator packet (version 5) which contains a literal value (version 6)
	// line 4: 620080001611562C8802118E34 represents an operator packet (version 3) which contains two sub-packets; each sub-packet is an operator packet that contains two literal values
	// line 5: C0015000016115A2E0802F182340 has the same structure as the previous example, but the outermost packet uses a different length type ID
	// line 6: A0016C880162017C3686B18A3D4780 is an operator packet that contains an operator packet that contains an operator packet that contains five literal values
	// line := lines[3]
	line := lines[0]

	// p.msg is an array of 0 / 1 values, as uint8
	p.msg = make(Day16Msg, 0)
	for _, i := range line {
		bits, err := strconv.ParseUint(string(i), 16, 4)
		if err != nil {
			panic(err)
		}
		toBits := lib.AsBits(bits, 4)
		p.msg = append(p.msg, toBits...)
	}
}

// reads the next bits as literal 5-bit-values, WITHOUT
// version and type (so the first bit pt is pointing to must
// be the 1st value's start indicator (1))
// returning:
// - the resulting Literal Package, without version nr set
// - the moved pointer (points to the NEXT bit AFTER the package)
func (m Day16Msg) readLiteralPackage(pt uint64) (*Day16LiteralPacket, uint64) {
	last := false
	var value uint64 = 0
	valueBits := make([]uint8, 0)
	for {
		if m[pt] == 0 {
			last = true
		}
		pt++
		valueBits = append(valueBits, m[pt:pt+4]...)
		pt += 4
		if last == true {
			break
		}
	}
	value = lib.BitsToUint64(valueBits)
	pkg := Day16LiteralPacket{
		value: value,
	}
	return &pkg, pt
}

// reads the next bits as operator package, WITHOUT
// version and type (so the first bit pt is pointing to must
// be the 1st bit of the operator package without version and type.
//
// returning:
// - the resulting operator package, without version nr set
// - the moved pointer (points to the NEXT bit AFTER the package)
func (m Day16Msg) readOperatorPackage(pt uint64) (*Day16OperatorPacket, uint64) {
	typeId := m[pt]
	pt++
	opPackage := Day16OperatorPacket{}

	if typeId == 0 {
		// next 15 bits are total length of subpackets
		if pt+15 >= uint64(len(m)) {
			return nil, pt
		}
		lenBits := m[pt : pt+15]
		pt += 15
		totLen := lib.BitsToUint64(lenBits)
		bits := m[pt : pt+totLen]
		pt += totLen
		newMessage := Day16Msg(bits)
		// bits now contain a bunch of packages to read.
		packets, _ := newMessage.readPackets()
		if packets == nil {
			panic("Empty operator package")
		}
		opPackage.subpackets = packets
		return &opPackage, pt
	} else {
		// next 11 bits are the nr of sub-packets
		if pt+11 >= uint64(len(m)) {
			return nil, pt
		}
		nrOfPackets := lib.BitsToUint64(m[pt : pt+11])
		pt += 11
		opPackage.subpackets = make([]Day16Packet, 0)
		// read n packages:
		for i := 0; i < int(nrOfPackets); i++ {
			// read n packages:
			pkg, newPt := m.parse(pt)
			pt = newPt
			if pkg != nil {
				opPackage.subpackets = append(opPackage.subpackets, pkg)
			}
		}
		return &opPackage, pt
	}
}

// parses the msg from a certain point, reading ONE packet
// and return it. The pt should point on the very first
// bit of the packet (so the begin of the version number).
// If the given packet is an operator packet,
// it reads its content recusively.
// returning:
// - the resulting packet struct
// - the moved pointer (points to the NEXT bit AFTER the package)
func (m Day16Msg) parse(pt uint64) (Day16Packet, uint64) {
	if pt >= uint64(len(m)) {
		return nil, pt
	}

	// 1: 3 bits ==> version number
	if pt+3 >= uint64(len(m)) {
		return nil, pt
	}

	part := m[pt : pt+3]
	pt += 3

	versionNr := lib.BitsToUint64(part)

	// 2: 3 bits ==> packet type
	if pt+3 >= uint64(len(m)) {
		return nil, pt
	}
	part = m[pt : pt+3]
	pt += 3
	pType := lib.BitsToUint64(part)

	switch pType {
	case 4: // literal package
		pkg, newPt := m.readLiteralPackage(pt)
		pkg.version = versionNr
		return pkg, newPt
	default: // every other package is an operator package
		pkg, newPt := m.readOperatorPackage(pt)
		if pkg == nil {
			return nil, newPt
		}
		pkg.ptype = pType
		pkg.version = versionNr
		return pkg, newPt
	}
}

func (m Day16Msg) readPackets() ([]Day16Packet, uint64) {
	packets := make([]Day16Packet, 0)
	var pt uint64 = 0
	for {
		if pt >= uint64(len(m)) {
			break
		}
		packet, retPt := m.parse(pt)
		pt = retPt
		if packet != nil {
			packets = append(packets, packet)
		} else {
			break
		}
	}
	return packets, pt
}

// Writing a little parser.
// We need to go over the msg and parse bit by bit,
// keeping track of the actual bit array pointer,
// also considering recursive parsing sub-packets.
//
// This leads to a nice litte AST with mathematical
// operations.
func (p *Day16) Run1() {
	p.packetTree, _ = p.msg.readPackets()

	var versionSum uint64 = 0
	for _, p := range p.packetTree {
		versionSum += p.versionSum()
	}

	p.solution1 = versionSum
}

// Walk the AST created in Part 1
func (p *Day16) Run2() {
	if len(p.packetTree) != 1 {
		panic("Oops - more than 1 root package....")
	}
	root := p.packetTree[0]
	p.solution2 = root.evaluate()
}

func (p *Day16) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day16) GetSolution2() string {
	fmt.Printf("Formula:\n\n%s\n\n", p.packetTree[0].toFormula())
	return fmt.Sprintf("%v\n", p.solution2)
}
