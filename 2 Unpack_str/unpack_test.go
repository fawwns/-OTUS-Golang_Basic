package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde", wantErr: false},
		{input: "abccd", expected: "abccd", wantErr: false},
		{input: "3abc", expected: "", wantErr: true},
		{input: "45", expected: "", wantErr: true},
		{input: "aaa10b", expected: "", wantErr: true},
		{input: "aaa0b", expected: "aab", wantErr: false},
		{input: `qwe\4\5`, expected: "qwe45", wantErr: false},
		{input: `qwe\45`, expected: "qwe44444", wantErr: false},
		{input: `qw\ne`, expected: "", wantErr: true},
	}

	for _, s := range tests {
		res, err := Unpack(s.input)
		if s.wantErr {
			if err == nil {
				t.Errorf("Unpack(%q) expected error, got nil", s.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unpack(%q) unexpected error: %v", s.input, err)
			}
			if res != s.expected {
				t.Errorf("Unpack(%q) = %q; want %q", s.input, res, s.expected)
			}
		}
	}
}
