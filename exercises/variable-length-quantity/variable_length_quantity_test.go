package variablelengthquantity

import (
	"bytes"
	"reflect"
	"testing"
)

const targetTestVersion = 3

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestDecodeVarint(t *testing.T) {
	for i, tc := range decodeTestCases {
		o, size, err := DecodeVarint(tc.input)
		if err != nil {
			var _ error = err
			if tc.output != nil {
				t.Fatalf("FAIL: case %d | %s\nexpected %#v got error: %q\n", i, tc.description, tc.output, err)
			}
		} else if tc.output == nil {
			t.Fatalf("FAIL: case %d | %s\nexpected error, got %#v\n", i, tc.description, o)
		} else if !reflect.DeepEqual(o, tc.output) {
			t.Fatalf("FAIL: case %d | %s\nexpected\t%#v\ngot\t\t%#v\n", i, tc.description, tc.output, o)
		} else if size != tc.size {
			t.Fatalf("FAIL: case %d | %s\n expected encoding size of %d bytes\ngot %d bytes\n", i, tc.description, tc.size, size)
		}
		t.Logf("PASS: case %d | %s\n", i, tc.description)
	}
}

func TestEncodeVarint(t *testing.T) {
	for i, tc := range encodeTestCases {
		if encoded := EncodeVarint(tc.input); bytes.Compare(encoded, tc.output) != 0 {
			t.Fatalf("FAIL: case %d | %s\nexpected\t%#v\ngot\t\t%#v\n", i, tc.description, tc.output, encoded)
		}
		t.Logf("PASS: case %d | %s\n", i, tc.description)
	}
}
