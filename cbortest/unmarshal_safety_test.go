package cbortest

import (
	"bytes"
	"encoding/hex"
	"errors"
	"reflect"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

type SafetyEmbedded struct {
	Value int
}

func TestUnmarshalSafety(t *testing.T) {
	tests := []struct {
		name      string
		hex       string
		newTarget func() any
		want      any
		wantError bool
	}{
		{"scalar slice field", "a16656616c75657301", func() any { return &struct{ Values []int }{} }, nil, true},
		{"map slice element", "81a10102", func() any { return &[][]int{} }, nil, true},
		{"scalar nested array", "8101", func() any { return &[][]int{} }, nil, true},
		{"private assignable field", "a16673656372657401", func() any { return &struct{ secret int8 }{} }, nil, true},
		{"private convertible field", "a16673656372657401", func() any { return &struct{ secret int }{} }, nil, true},
		{"nil embedded pointer", "a16556616c756501", func() any { return &struct{ *SafetyEmbedded }{} }, nil, true},
		{"fixed array", "8101", func() any { return &[1]int{} }, [1]int{1}, false},
		{"short fixed array", "820102", func() any { return &[1]int{} }, nil, true},
		{"populated slice", "8101", func() any { v := []int{0}; return &v }, []int{1}, false},
		{"larger slice keeps tail", "8101", func() any { v := []int{0, 9}; return &v }, []int{1, 9}, false},
		{"growing slice", "820102", func() any { v := []int{0}; return &v }, []int{1, 2}, false},
		{"nested fixed array", "818101", func() any { return &[][1]int{} }, [][1]int{{1}}, false},
		{"exported field", "a16556616c756501", func() any { return &SafetyEmbedded{} }, SafetyEmbedded{Value: 1}, false},
		{"initialized embedded pointer", "a16556616c756501", func() any { return &struct{ *SafetyEmbedded }{&SafetyEmbedded{}} }, struct{ *SafetyEmbedded }{&SafetyEmbedded{Value: 1}}, false},
	}
	for _, tt := range tests {
		for _, streaming := range []bool{false, true} {
			t.Run(tt.name+map[bool]string{false: "/bytes", true: "/stream"}[streaming], func(t *testing.T) {
				defer func() {
					if p := recover(); p != nil {
						t.Fatalf("unexpected panic: %v", p)
					}
				}()
				data, err := hex.DecodeString(tt.hex)
				if err != nil {
					t.Fatal(err)
				}
				target := tt.newTarget()
				if streaming {
					err = cbor.NewDecoder(bytes.NewReader(data)).Unmarshal(target)
				} else {
					err = cbor.UnmarshalTo(data, target)
				}
				if tt.wantError {
					if !errors.Is(err, cbor.ErrUnmarshal) {
						t.Fatalf("got %v, want ErrUnmarshal", err)
					}
					return
				}
				if err != nil {
					t.Fatal(err)
				}
				if got := reflect.ValueOf(target).Elem().Interface(); !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("got %#v, want %#v", got, tt.want)
				}
			})
		}
	}
}

func TestUnmarshalStructValueSlice(t *testing.T) {
	target := struct{ Values []int }{[]int{0}}
	data := []byte{0xa1, 0x66, 'V', 'a', 'l', 'u', 'e', 's', 0x81, 0x01}
	if err := cbor.UnmarshalTo(data, target); err != nil {
		t.Fatal(err)
	}
	if target.Values[0] != 1 {
		t.Fatalf("got %v, want [1]", target.Values)
	}
}
