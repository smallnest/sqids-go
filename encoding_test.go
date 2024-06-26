package sqids

import (
	"math"
	"reflect"
	"testing"
)

const (
	minUint64Value = uint64(0)
	maxUint64Value = uint64(math.MaxUint64)
)

func BenchmarkEncodeDecode(b *testing.B) {
	numbers := []uint64{1, 2, 3, 4, 5}

	s, err := New()
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		id, err := s.Encode(numbers...)
		if err != nil {
			b.Fatal(err)
		}

		decodedNumbers := s.Decode(id)
		if !reflect.DeepEqual(numbers, decodedNumbers) {
			b.Errorf("Could not encode/decode `%v`", numbers)
		}
	}
}

func TestEncodingSimple(t *testing.T) {
	id := "86Rf07"
	numbers := []uint64{1, 2, 3}

	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	generatedID, err := s.Encode(numbers...)
	if err != nil {
		t.Fatal(err)
	}

	if id != generatedID {
		t.Errorf("Encoding `%v` should produce `%v`, but instead produced `%v`", numbers, id, generatedID)
	}

	decodedNumbers := s.Decode(generatedID)
	if !reflect.DeepEqual(numbers, decodedNumbers) {
		t.Errorf("Could not encode/decode `%v`", numbers)
	}
}

func TestEncodingDifferentInputs(t *testing.T) {
	numbers := []uint64{minUint64Value, 0, 0, 1, 2, 3, 100, 1_000, 100_000, 1_000_000, maxUint64Value}

	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	id, err := s.Encode(numbers...)
	if err != nil {
		t.Fatal(err)
	}

	decodedNumbers := s.Decode(id)
	if !reflect.DeepEqual(numbers, decodedNumbers) {
		t.Errorf("Could not encode/decode `%v`", numbers)
	}
}

func TestEncodingIncrementalNumbers(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	ids := map[string][]uint64{
		"bM": {0},
		"Uk": {1},
		"gb": {2},
		"Ef": {3},
		"Vq": {4},
		"uw": {5},
		"OI": {6},
		"AX": {7},
		"p6": {8},
		"nJ": {9},
	}

	for id, numbers := range ids {
		generatedID, err := s.Encode(numbers...)
		if err != nil {
			t.Fatal(err)
		}

		if id != generatedID {
			t.Errorf("Encoding `%v` should produce `%v`, but instead produced `%v`", numbers, id, generatedID)
		}

		decodedNumbers := s.Decode(id)
		if !reflect.DeepEqual(numbers, decodedNumbers) {
			t.Errorf("Decoding `%v` should produce `%v`, but instead produced `%v`", id, numbers, decodedNumbers)
		}
	}
}

func TestEncodingIncrementalNumbersSameIndex0(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	ids := map[string][]uint64{
		"SvIz": {0, 0},
		"n3qa": {0, 1},
		"tryF": {0, 2},
		"eg6q": {0, 3},
		"rSCF": {0, 4},
		"sR8x": {0, 5},
		"uY2M": {0, 6},
		"74dI": {0, 7},
		"30WX": {0, 8},
		"moxr": {0, 9},
	}

	for id, numbers := range ids {
		generatedID, err := s.Encode(numbers...)
		if err != nil {
			t.Fatal(err)
		}

		if id != generatedID {
			t.Errorf("Encoding `%v` should produce `%v`, but instead produced `%v`", numbers, id, generatedID)
		}

		decodedNumbers := s.Decode(id)
		if !reflect.DeepEqual(numbers, decodedNumbers) {
			t.Errorf("Decoding `%v` should produce `%v`, but instead produced `%v`", id, numbers, decodedNumbers)
		}
	}
}

func TestEncodingIncrementalNumbersSameIndex1(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	ids := map[string][]uint64{
		"SvIz": {0, 0},
		"nWqP": {1, 0},
		"tSyw": {2, 0},
		"eX68": {3, 0},
		"rxCY": {4, 0},
		"sV8a": {5, 0},
		"uf2K": {6, 0},
		"7Cdk": {7, 0},
		"3aWP": {8, 0},
		"m2xn": {9, 0},
	}

	for id, numbers := range ids {
		generatedID, err := s.Encode(numbers...)
		if err != nil {
			t.Fatal(err)
		}

		if id != generatedID {
			t.Errorf("Encoding `%v` should produce `%v`, but instead produced `%v`", numbers, id, generatedID)
		}

		decodedNumbers := s.Decode(id)
		if !reflect.DeepEqual(numbers, decodedNumbers) {
			t.Errorf("Decoding `%v` should produce `%v`, but instead produced `%v`", id, numbers, decodedNumbers)
		}
	}
}

func TestEncodingMultiInput(t *testing.T) {
	numbers := []uint64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
		26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73,
		74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97,
		98, 99}

	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	id, err := s.Encode(numbers...)
	if err != nil {
		t.Fatal(err)
	}

	decodedNumbers := s.Decode(id)
	if !reflect.DeepEqual(numbers, decodedNumbers) {
		t.Errorf("Could not encode/decode `%v`", numbers)
	}
}

func TestEncodingEmptySlice(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	id, err := s.Encode()
	if err != nil {
		t.Fatal(err)
	}

	if id != "" {
		t.Errorf("Could not encode empty slice")
	}
}

func TestEncodingEmptyString(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(s.Decode(""), []uint64{}) {
		t.Errorf("Could not decode empty string")
	}
}

func TestEncodingInvalidCharacter(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(s.Decode("*"), []uint64{}) {
		t.Errorf("Could not decode with invalid character")
	}
}

// TestEncodingOutOfRange - no need since `[]uint64` handles ranges
// func TestEncodingOutOfRange(t *testing.T) {}
