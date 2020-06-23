package mapobjectdb

import (
	"testing"
	"unicode/utf8"
)

func TestInvalidUtf8(t *testing.T) {
	if utf8.Valid([]byte{0xe1, 0x7f, 0xc7}) {
		t.Error("should be invalid")
	}
}

func TestValidUtf8(t *testing.T) {
	if !utf8.Valid([]byte("some valid string")) {
		t.Error("should be valid")
	}
}

func TestEmptyString(t *testing.T) {
	if !utf8.Valid([]byte("")) {
		t.Error("should be valid")
	}
}
