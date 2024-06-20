package util

import (
	"reflect"
	"testing"
)

func TestNewOneBit(t *testing.T) {
	t.Run("NewOneBit", func(t *testing.T) {
		tests := []struct {
			input byte
			want  OneBit
		}{
			{0b00000000, 0b00000000},
			{0b00000001, 0b00000001},
		}
		for _, tt := range tests {

			got, err := NewOneBit(tt.input)
			if got != tt.want {
				t.Errorf("NewOneBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err != nil {
				t.Errorf("NewOneBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
	t.Run("NewOneBit_Failure", func(t *testing.T) {
		tests := []struct {
			input byte
			want  OneBit
		}{
			{0b00000010, 0b00000000},
			{0b00000011, 0b00000000},
		}
		for _, tt := range tests {

			got, err := NewOneBit(tt.input)
			if got != tt.want {
				t.Errorf("NewOneBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("NewOneBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
}

func TestNewThreeBit(t *testing.T) {
	t.Run("NewThreeBit", func(t *testing.T) {
		tests := []struct {
			input byte
			want  ThreeBit
		}{
			{0b00000000, 0b00000000},
			{0b00000001, 0b00000001},
			{0b00000111, 0b00000111},
		}
		for _, tt := range tests {

			got, err := NewThreeBit(tt.input)
			if got != tt.want {
				t.Errorf("NewThreeBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err != nil {
				t.Errorf("NewThreeBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
	t.Run("NewThreeBit_Failure", func(t *testing.T) {
		tests := []struct {
			input byte
			want  ThreeBit
		}{
			{0b00001000, 0b00000000},
			{0b00001001, 0b00000000},
		}
		for _, tt := range tests {

			got, err := NewThreeBit(tt.input)
			if got != tt.want {
				t.Errorf("NewThreeBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("NewThreeBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
}

func TestNewFourBit(t *testing.T) {
	t.Run("NewFourBit", func(t *testing.T) {
		tests := []struct {
			input byte
			want  FourBit
		}{
			{0b00000000, 0b00000000},
			{0b00000001, 0b00000001},
			{0b00000111, 0b00000111},
			{0b00001111, 0b00001111},
		}
		for _, tt := range tests {

			got, err := NewFourBit(tt.input)
			if got != tt.want {
				t.Errorf("NewFourBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err != nil {
				t.Errorf("NewFourBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
	t.Run("NewFourBit_Failure", func(t *testing.T) {
		tests := []struct {
			input byte
			want  FourBit
		}{
			{0b0010000, 0b00000000},
			{0b0010001, 0b00000000},
		}
		for _, tt := range tests {

			got, err := NewFourBit(tt.input)
			if got != tt.want {
				t.Errorf("NewFourBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("NewFourBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
}

func TestNewSixteenBit(t *testing.T) {
	t.Run("NewSixteenBit", func(t *testing.T) {
		tests := []struct {
			input [2]byte
			want  SixteenBit
		}{
			{[2]byte{255, 255}, [2]byte{255, 255}},
			{[2]byte{3, 4}, [2]byte{3, 4}},
			{[2]byte{0, 0}, [2]byte{0, 0}},
			{[2]byte{0}, [2]byte{0, 0}},
		}
		for _, tt := range tests {

			got, err := NewSixteenBit(tt.input)
			if got != tt.want {
				t.Errorf("NewSixteenBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err != nil {
				t.Errorf("NewSixteenBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
}

func TestNewHeader(t *testing.T) {
	t.Run("NewHeader", func(t *testing.T) {
		tests := []struct {
			input [12]byte
			want  Header
		}{
			{[12]byte{255, 25, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Header{id: [2]byte{255, 25}}},
			{[12]byte{0, 0, 0b11111101, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				Header{qr: 1, opcode: 0b1111, aa: 1, tc: 0, rd: 1}},
		}
		for _, tt := range tests {

			got, err := NewHeader(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeader(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err != nil {
				t.Errorf("NewHeader(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
}

func TestToBytes(t *testing.T) {
	t.Run("ToBytes", func(t *testing.T) {
		tests := []struct {
			input Header
			want  [12]byte
		}{
			{Header{id: [2]byte{255, 25}}, [12]byte{255, 25, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			{
				Header{qr: 1, opcode: 0b1111, aa: 1, tc: 0, rd: 1},
				[12]byte{0, 0, 0b11111101, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		}
		for _, tt := range tests {

			got := ToBytes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeader(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
		}
	})
}

func TestDomainToByte(t *testing.T) {
	t.Run("DomainToByte", func(t *testing.T) {
		tests := []struct {
			input string
			want  []byte
		}{
			{"google.com", []byte{0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x03, 0x63, 0x6f, 0x6d, 0x00}},                                    // \x06google\x03com\x00
			{"codecrafters.io", []byte{0x0c, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x73, 0x02, 0x69, 0x6f, 0x00}}, // \x0ccodecrafters\x02io\x00
		}
		for _, tt := range tests {
			got := DomainToByte(tt.input)
			if !reflect.DeepEqual(got, tt.want) {

				t.Errorf("DomainToByte(%s) = %08b; \nwant %08b\n got: %d\n want: %d", tt.input, got, tt.want, len(got), len(tt.want))
			}
		}
	})
}
