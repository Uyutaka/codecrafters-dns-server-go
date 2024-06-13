package util

import (
	"errors"
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
	// if len(value) > 2 {
	// return SixteenBit([]byte{0}), errors.New("value exceeds 16 bit limit")
	// }
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
	copy(SixteenBitArray[:], value[3:5])

	qdcount, err := NewSixteenBit(SixteenBitArray)
	if err != nil {
		return Header{}, err
	}
	copy(SixteenBitArray[:], value[5:7])

	ancount, err := NewSixteenBit(SixteenBitArray)
	if err != nil {
		return Header{}, err
	}
	copy(SixteenBitArray[:], value[7:9])

	nscount, err := NewSixteenBit(SixteenBitArray)
	if err != nil {
		return Header{}, err
	}
	copy(SixteenBitArray[:], value[9:11])
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

func ToBytes(header Header) [12]byte {
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
