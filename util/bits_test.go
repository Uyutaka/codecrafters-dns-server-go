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
			{
				[12]byte{
					0x1b, 0xc3,
					0x01, 0x00,
					0x12, 0x21,
					0x12, 0x34,
					0x56, 0x78,
					0x9a, 0xbc},
				Header{
					id: SixteenBit{0x1b, 0xc3},
					qr: OneBit(0), opcode: FourBit(0), aa: OneBit(0), tc: OneBit(0), rd: OneBit(1), ra: OneBit(0x0), z: ThreeBit(0x0), rcode: FourBit(0x0),
					qdcount: SixteenBit{0x12, 0x21},
					ancount: SixteenBit{0x12, 0x34},
					nscount: SixteenBit{0x56, 0x78},
					arcount: SixteenBit{0x9a, 0xbc},
				},
			},
		}
		for _, tt := range tests {

			got, err := NewHeader(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeader(%x) = %x; want %x", tt.input, got, tt.want)
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

			got := HeaderToBytes(tt.input)
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

func TestNewRecordType(t *testing.T) {
	t.Run("NewRecordType", func(t *testing.T) {
		tests := []struct {
			input uint8
			want  [2]byte
		}{
			{
				A,
				[2]byte{0x00, 0x01}, // 1
			},
			{
				TXT,
				[2]byte{0x00, 0x10}, // 16
			},
		}
		for _, tt := range tests {
			got, _ := NewRecordType(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecordType(%d) = %08b; \nwant %08b", tt.input, got, tt.want)
			}
		}
	})
	t.Run("NewRecordType_Failure", func(t *testing.T) {
		tests := []struct {
			input uint8
			want  [2]byte
		}{
			{
				MAX_RECORD_TYPE + 1,
				[2]byte{0x00, 0x00},
			},
			{
				100,
				[2]byte{0x00, 0x00},
			},
		}
		for _, tt := range tests {

			got, err := NewRecordType(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecordType(%d) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("NewRecordType(%d) = %08b; do not want nil", tt.input, got)
			}
		}
	})
}
func TestNewClass(t *testing.T) {
	t.Run("NewClass", func(t *testing.T) {
		tests := []struct {
			input uint8
			want  [2]byte
		}{
			{
				IN,
				[2]byte{0x00, 0x01}, // 1
			},
			{
				HS,
				[2]byte{0x00, 0x04}, // 4
			},
		}
		for _, tt := range tests {
			got, _ := NewClass(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClass(%d) = %08b; \nwant %08b", tt.input, got, tt.want)
			}
		}
	})
	t.Run("NewClass_Failure", func(t *testing.T) {
		tests := []struct {
			input uint8
			want  [2]byte
		}{
			{
				MAX_CLASS + 1,
				[2]byte{0x00, 0x00},
			},
			{
				100,
				[2]byte{0x00, 0x00},
			},
		}
		for _, tt := range tests {

			got, err := NewClass(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClass(%d) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("NewClass(%d) = %08b; do not want nil", tt.input, got)
			}
		}
	})
}

func TestQuestionBytes(t *testing.T) {
	t.Run("AnswerBytes", func(t *testing.T) {
		tests := []struct {
			input []byte
			want  []byte
		}{
			{
				[]byte{'<', '-', '-', 'h', 'e', 'a', 'd', 'e', 'r', '-', '-', '>', 'n', 'a', 'm', 'e', 0x0, 't', 'y', 'c', 'l', 'a', 'n', 's', 'w', 'e', 'r'},
				[]byte{'n', 'a', 'm', 'e', 0x0, 't', 'y', 'c', 'l'},
			},
			{
				[]byte{'<', '-', '-', 'h', 'e', 'a', 'd', 'e', 'r', '-', '-', '>', 'n', 'a', 'm', 'e', 'h', 'o', 'g', 'e', 0x0, 't', 'y', 'c', 'l', 'a', 'n', 's', 'w', 'e', 'r', 'h', 'o', 'g', 'e'},
				[]byte{'n', 'a', 'm', 'e', 'h', 'o', 'g', 'e', 0x0, 't', 'y', 'c', 'l'},
			},
		}
		for _, tt := range tests {
			got, _ := QuestionBytes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuestionBytes(%c) = %c; \nwant %c", tt.input, got, tt.want)
			}
		}
	})
	t.Run("AnswerByte_Failure", func(t *testing.T) {

		tests := []struct {
			input []byte
			want  []byte
		}{
			{
				// less than 13 bytes (= only header section)
				[]byte{'h', 'e', 'a', 'd', 'e', 'r', 's', 'e', 'c', 't', 'i', 'o'},
				[]byte{},
			},
			{
				// without null byte
				[]byte{'h', 'e', 'a', 'd', 'e', 'r', 's', 'e', 'c', 't', 'i', 'o', 'n', 'n', 'a'},
				[]byte{},
			},
		}
		for _, tt := range tests {
			got, err := QuestionBytes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuestionBytes(%c) = %c; \nwant %c", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("QuestionBytes(%c) = %c; do not want nil", tt.input, got)
			}
		}
	})
}

func TestByteToDomain(t *testing.T) {
	t.Run("ByteToDomain", func(t *testing.T) {
		tests := []struct {
			input []byte
			want  string
		}{
			{
				[]byte{0x0c, 'c', 'o', 'd', 'e', 'c', 'r', 'a', 'f', 't', 'e', 'r', 's', 0x2, 'i', 'o', 0x00},
				"codecrafters.io",
			},
			{
				[]byte{0x06, 'g', 'o', 'o', 'g', 'l', 'e', 0x03, 'c', 'o', 'm', 0x00},
				"google.com",
			},
		}
		for _, tt := range tests {
			got, _ := byteToDomain(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteToDomain(%c) = %s; \nwant %s", tt.input, got, tt.want)
			}
		}
	})
	t.Run("ByteToDomain_Failure", func(t *testing.T) {

		tests := []struct {
			input []byte
			want  string
		}{
			{
				[]byte{0x0e, 'c', 'o', 'd', 'e', 'c', 'r', 'a', 'f', 't', 'e', 'r', 's', 0x2, 'i', 'o'},
				"",
			},
			{
				[]byte{0x06, 'g', 'o', 'o', 'g', 'l', 'e', 0x04, 'c', 'o', 'm'},
				"",
			},
		}
		for _, tt := range tests {
			got, err := byteToDomain(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteToDomain(%c) = %s; \nwant %s", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("ByteToDomain(%c) = %s; do not want nil", tt.input, got)
			}
		}
	})
}

func TestNullIndex(t *testing.T) {
	t.Run("NullIndex", func(t *testing.T) {
		tests := []struct {
			input []byte
			want  int
		}{
			{
				[]byte{'c', 'o', 'd', 'e', 0x00},
				4,
			},
			{
				[]byte{'c', 'o', 0x00, 'e'},
				2,
			},
		}
		for _, tt := range tests {
			got, _ := NullIndex(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullIndex(%c) = %d; \nwant %d", tt.input, got, tt.want)
			}
		}
	})
	t.Run("NullIndex_Failure", func(t *testing.T) {

		tests := []struct {
			input []byte
			want  int
		}{
			{
				[]byte{'c', 'o', 'd', 'e'},
				-1,
			},
			{
				[]byte{},
				-1,
			},
		}
		for _, tt := range tests {
			got, err := NullIndex(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullIndex(%c) = %d; \nwant %d", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("NullIndex(%c) = %d; do not want nil", tt.input, got)
			}
		}
	})
}
