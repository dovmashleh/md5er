package md5

import (
	"crypto/md5"
	"io"
	"reflect"
	"testing"
)

func TestAsByteSlice(t *testing.T) {
	tests := []struct {
		name string
		mes  string
	}{
		{
			name: "simple md5",
			mes:  "md5",
		},
		{
			name: "longer than 64b",
			mes:  "das ist ein very long sting, at least should be longer than 64b, but who's counting.... Probably ich musste, aber ish bin to lazy, so i keep writing that crap until i'm sure it's longer than 64b",
		},
		{
			name: "utf8 string",
			mes:  "We need some multibyte string, так что вот немного русских букв, יְהוָֹה шоб помучаться с переходом на левостороннее, да пара 4-байтных смайлов 😈 ⛧ 🤘",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteMes := []byte(tt.mes)
			md5er := New()
			got := md5er.AsByteSlice(byteMes)
			h := md5.New()
			io.WriteString(h, tt.mes)
			want := h.Sum(nil)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("AsByteArray() = %v, want %v", got, want)
			}
		})
	}
}
