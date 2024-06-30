package util

import (
	"errors"
	"fmt"
	"strings"
)

const MaxOneBitValue OneBit = 0b1

type OneBit byte

func NewOneBit(value byte) (OneBit, error) {

	if value > byte(MaxOneBitValue) {
		return 0, errors.New("value exceeds 1 bit limit")
	}
	return OneBit(value), nil
}

const MaxThreeBitValue ThreeBit = 0b111

type ThreeBit byte

func NewThreeBit(value byte) (ThreeBit, error) {

	if value > byte(MaxThreeBitValue) {
		return 0, errors.New("value exceeds 3 bit limit")
	}
	return ThreeBit(value), nil
}

const MaxFourBitValue FourBit = 0b1111

type FourBit byte

func NewFourBit(value byte) (FourBit, error) {

	if value > byte(MaxFourBitValue) {
		return 0, errors.New("value exceeds 4 bit limit")
	}
	return FourBit(value), nil
}

type SixteenBit [2]byte

func NewSixteenBit(value [2]byte) (SixteenBit, error) {
	// TODO
	// - change parameter to uint
	//    - check the max int value
	return SixteenBit(value), nil
}

type Header struct {
	// 0,1
	id SixteenBit
	// 2
	qr     OneBit
	opcode FourBit
	aa     OneBit
	tc     OneBit
	rd     OneBit
	// 3
	ra    OneBit
	z     ThreeBit
	rcode FourBit
	// 4,5
	qdcount SixteenBit
	// 6,7
	ancount SixteenBit
	// 8,9
	nscount SixteenBit
	// 10, 11
	arcount SixteenBit
}

func DebugHeader(h Header) {
	fmt.Println("Debug Header")
	fmt.Printf("id:\t%x\n", h.id)
	fmt.Printf("qr:\t%x\n", h.qr)
	fmt.Printf("opcode:\t%x\n", h.opcode)
	fmt.Printf("aa:\t%x\n", h.aa)
	fmt.Printf("tc:\t%x\n", h.tc)
	fmt.Printf("rd:\t%x\n", h.rd)

	fmt.Printf("ra:\t%x\n", h.ra)
	fmt.Printf("z:\t%x\n", h.z)
	fmt.Printf("rcode:\t%x\n", h.rcode)
	fmt.Printf("qdcount:\t%x\n", h.qdcount)
	fmt.Printf("ancount:\t%x\n", h.ancount)

	fmt.Printf("nscount:\t%x\n", h.nscount)

	fmt.Printf("arcount:\t%x\n", h.arcount)
}

func NewHeader(value [12]byte) (Header, error) {
	if len(value) != 12 {
		return Header{}, errors.New("size is invalid")
	}

	// 0-1
	var SixteenBitArray [2]byte
	copy(SixteenBitArray[:], value[0:2])
	id, err := NewSixteenBit(SixteenBitArray)
	if err != nil {
		return Header{}, err
	}

	// 2
	qr, err := NewOneBit((value[2] >> 7) & 0x01)
	if err != nil {
		return Header{}, err
	}
	opcode, err := NewFourBit((value[2] >> 3) & 0x0F)
	if err != nil {
		return Header{}, err
	}
	aa, err := NewOneBit((value[2] >> 2) & 0x01)
	if err != nil {
		return Header{}, err
	}
	tc, err := NewOneBit((value[2] >> 1) & 0x01)
	if err != nil {
		return Header{}, err
	}
	rd, err := NewOneBit(value[2] & 0x01)
	if err != nil {
		return Header{}, err
	}

	// 3
	ra, err := NewOneBit((value[3] >> 7) & 0x01)
	if err != nil {
		return Header{}, err
	}
	z, err := NewThreeBit((value[3] >> 4) & 0x07)
	if err != nil {
		return Header{}, err
	}
	rcode, err := NewFourBit(value[3] & 0x0F)
	if err != nil {
		return Header{}, err
	}

	// 4-11
	copy(SixteenBitArray[:], value[4:6])
	qdcount, err := NewSixteenBit(SixteenBitArray)
	if err != nil {
		return Header{}, err
	}

	copy(SixteenBitArray[:], value[6:8])
	ancount, err := NewSixteenBit(SixteenBitArray)
	if err != nil {
		return Header{}, err
	}

	copy(SixteenBitArray[:], value[8:10])
	nscount, err := NewSixteenBit(SixteenBitArray)
	if err != nil {
		return Header{}, err
	}

	copy(SixteenBitArray[:], value[10:12])
	arcount, err := NewSixteenBit(SixteenBitArray)
	if err != nil {
		return Header{}, err
	}

	header := Header{
		id:      id,
		qr:      qr,
		opcode:  opcode,
		aa:      aa,
		tc:      tc,
		rd:      rd,
		ra:      ra,
		z:       z,
		rcode:   rcode,
		qdcount: qdcount,
		ancount: ancount,
		nscount: nscount,
		arcount: arcount,
	}
	return header, nil
}

func NewHeaderWithQdcountAndAncount(header Header, qdcount uint16, ancount uint16) (Header, error) {

	qdcountBytes, err := NewSixteenBit([2]byte{byte(qdcount >> 8), byte(qdcount & 0xff)})
	if err != nil {
		return Header{}, err
	}

	ancountBytes, err := NewSixteenBit([2]byte{byte(ancount >> 8), byte(ancount & 0xff)})
	if err != nil {
		return Header{}, err
	}

	header.qdcount = qdcountBytes
	header.ancount = ancountBytes
	return header, nil
}

func HeaderToBytes(header Header) [12]byte {
	var result [12]byte

	// Convert id to bytes
	result[0] = header.id[0]
	result[1] = header.id[1]

	// Pack qr, opcode, aa, tc, rd into the third byte
	result[2] = byte(header.qr)<<7 | byte(header.opcode)<<3 | byte(header.aa)<<2 | byte(header.tc)<<1 | byte(header.rd)

	// Pack ra, z, rcode into the fourth byte
	result[3] = byte(header.ra)<<7 | byte(header.z)<<4 | byte(header.rcode)

	// Convert qdcount to bytes
	result[4] = header.qdcount[0]
	result[5] = header.qdcount[1]

	// Convert ancount to bytes
	result[6] = header.ancount[0]
	result[7] = header.ancount[1]

	// Convert nscount to bytes
	result[8] = header.nscount[0]
	result[9] = header.nscount[1]

	// Convert arcount to bytes
	result[10] = header.arcount[0]
	result[11] = header.arcount[1]

	return result
}

func (h Header) GetQdcount() [2]byte {
	return [2]byte(h.qdcount)
}

func (h Header) GetAncount() [2]byte {
	return [2]byte(h.ancount)
}

func (h Header) GetNscount() [2]byte {
	return [2]byte(h.nscount)
}

func (h Header) GetArcount() [2]byte {
	return [2]byte(h.arcount)
}

func Reply(header Header) Header {
	newHeader := header
	newHeader.qr = 1
	newHeader.aa = 0
	newHeader.tc = 0
	newHeader.ra = 0
	if header.opcode == 0 {
		newHeader.rcode = 0
	} else {
		newHeader.rcode = 4
	}
	return newHeader
}

func DomainToByte(domain string) []byte {
	parts := strings.Split(domain, ".")
	hexBytes := make([]byte, len(domain)+2)

	index := 0
	for i, part := range parts {
		hexBytes[index] = byte(len(parts[i]))
		index++
		for _, char := range part {
			hexBytes[index] = byte(char)
			index++
		}
	}
	return hexBytes
}

const (
	_ uint8 = iota
	A
	NS
	MD
	MF
	CNAME
	SOA
	MB
	MG
	MR
	NULL
	WKS
	PTR
	HINFO
	MINFO
	MX
	TXT
	MAX_RECORD_TYPE uint8 = 16
)
const (
	_ uint8 = iota
	IN
	CS
	CH
	HS
	MAX_CLASS uint8 = 4
)

func NewRecordType(rType uint8) ([2]byte, error) {

	if MAX_RECORD_TYPE < rType {
		return [2]byte{}, errors.New("value exceeds request type limit")
	}
	var byteArray [2]byte
	byteArray[1] = rType

	return byteArray, nil
}

func NewClass(class uint8) ([2]byte, error) {
	if MAX_CLASS < class {
		return [2]byte{}, errors.New("value exceeds class limit")
	}
	var byteArray [2]byte
	byteArray[1] = class
	return byteArray, nil
}

type Question struct {
	domain     []byte
	recordType [2]byte
	class      [2]byte
}

func NewQuestion(domain string, recordType uint8, class uint8) (Question, error) {
	domainByte := DomainToByte(domain)
	recordTypeByte, err := NewRecordType(recordType)
	if err != nil {
		return Question{}, err
	}
	classByte, err := NewClass(class)
	if err != nil {
		return Question{}, err
	}
	return Question{domain: domainByte, recordType: recordTypeByte, class: classByte}, nil
}

func NewQuestionFromByte(question []byte) (Question, error) {

	domainEnd, err := NullIndex(question)
	if err != nil {
		return Question{}, err
	}
	domainByte := question[0:domainEnd]
	domainByte = append(domainByte, 0x00)

	recordTypeByte, err := NewRecordType(A)
	if err != nil {
		return Question{}, err
	}
	classByte, err := NewClass(IN)
	if err != nil {
		return Question{}, err
	}

	return Question{domain: domainByte, recordType: recordTypeByte, class: classByte}, nil
}

func NullIndex(buf []byte) (int, error) {
	start := 0
	end := -1

	for i := start; i < len(buf); i++ {
		if buf[i] == 0x0 {
			end = i
			break
		}
	}
	if end == -1 {
		return -1, errors.New("not found")
	}

	return end, nil
}
func DomainInQuestion(q Question) (string, error) {

	domain, err := byteToDomain(q.domain)
	if err != nil {
		return "", err
	}
	return domain, nil
}

func byteToDomain(input []byte) (string, error) {
	if input[len(input)-1] != 0x00 {
		return "", errors.New("invalid input (last element is not null)")
	}

	buf := append([]byte{}, input...)
	buf = buf[:len(buf)-1]

	var domain string
	i := 0
	for i < len(buf) {
		length := int(buf[i])
		if i+length >= len(buf) {
			return "", errors.New("invalid input")
		}
		if domain != "" {
			domain += "."
		}
		domain += string(buf[i+1 : i+1+length])
		i += length + 1
	}
	return domain, nil
}

func QuestionToBytes(question Question) []byte {

	var questionBytes []byte

	typeSlice := question.recordType[:]
	classSlice := question.class[:]

	questionBytes = append(questionBytes, question.domain...)
	questionBytes = append(questionBytes, typeSlice...)
	questionBytes = append(questionBytes, classSlice...)

	return questionBytes
}

type Answer struct {
	rr []ResourceRecord
}

type ResourceRecord struct {
	domain     []byte
	recordType [2]byte
	class      [2]byte
	ttl        [4]byte
	rdlength   [2]byte
	rdata      []byte
}

func NewTTL(time uint32) [4]byte {
	return [4]byte{
		byte(time >> 24),
		byte(time >> 16),
		byte(time >> 8),
		byte(time)}
}

func NewLength(length uint16) [2]byte {
	return [2]byte{
		byte(length >> 8),
		byte(length),
	}
}

func NewResourceRecord(domain string) (ResourceRecord, error) {
	domainByte := DomainToByte(domain)
	recordType, err := NewRecordType(A)
	if err != nil {
		return ResourceRecord{}, err
	}
	class, err := NewClass(IN)
	if err != nil {
		return ResourceRecord{}, err
	}
	ttl := NewTTL(60)
	length := NewLength(4)
	return ResourceRecord{domain: domainByte, recordType: recordType, class: class, ttl: ttl, rdlength: length, rdata: []byte{0x08, 0x08, 0x08, 0x08}}, nil
}

func NewAnswer(rr ResourceRecord) Answer {
	var ans Answer
	ans.rr = append(ans.rr, rr)
	return ans
}

func AnswerToBytes(answer Answer) []byte {
	var answerBytes []byte

	// TODO: validation of length
	rrData := answer.rr[0]

	typeSlice := rrData.recordType[:]
	classSlice := rrData.class[:]
	ttlSlice := rrData.ttl[:]
	rdlengthSlice := rrData.rdlength[:]

	answerBytes = append(answerBytes, rrData.domain...)
	answerBytes = append(answerBytes, typeSlice...)
	answerBytes = append(answerBytes, classSlice...)
	answerBytes = append(answerBytes, ttlSlice...)
	answerBytes = append(answerBytes, rdlengthSlice...)
	answerBytes = append(answerBytes, rrData.rdata...)
	return answerBytes
}

func QuestionBytes(buf []byte) ([]byte, error) {
	if len(buf) < 13 {
		return []byte{}, errors.New("buf less than 12")
	}

	start := 12
	end := 0

	for i := start; i < len(buf); i++ {
		if buf[i] == 0x0 {
			end = i
			break
		}
	}

	if end == 0 {
		return []byte{}, errors.New("invalid format")
	}

	end += 5

	return buf[start:end], nil
}
