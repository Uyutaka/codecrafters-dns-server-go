package util

import (
	"bytes"
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
			input []byte
			want  SixteenBit
		}{
			{[]byte{255, 255}, []byte{255, 255}},
			{[]byte{3, 4}, []byte{3, 4}},
		}
		for _, tt := range tests {

			got, err := NewSixteenBit(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("NewSixteenBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err != nil {
				t.Errorf("NewSixteenBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})

	t.Run("NewSixteenBit_Failure", func(t *testing.T) {
		tests := []struct {
			input []byte
			want  SixteenBit
		}{
			{[]byte{255, 255, 1}, []byte{0}},
		}
		for _, tt := range tests {

			got, err := NewSixteenBit(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("NewSixteenBit(%08b) = %08b; want %08b", tt.input, got, tt.want)
			}
			if err == nil {
				t.Errorf("NewSixteenBit(%08b) = %08b; want nil", tt.input, got)
			}
		}
	})
}
