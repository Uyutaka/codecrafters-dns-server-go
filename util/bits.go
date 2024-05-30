package util

import "errors"

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

type SixteenBit []byte

func NewSixteenBit(value []byte) (SixteenBit, error) {
	if len(value) > 2 {
		return SixteenBit([]byte{0}), errors.New("value exceeds 16 bit limit")
	}
	return SixteenBit(value), nil
}

type Header struct {
	id      SixteenBit
	qr      OneBit
	opcode  FourBit
	aa      OneBit
	tc      OneBit
	rd      OneBit
	ra      OneBit
	z       ThreeBit
	rcode   FourBit
	qdcount SixteenBit
	ancount SixteenBit
	nscount SixteenBit
	arcount SixteenBit
}
