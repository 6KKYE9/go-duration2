package duration2

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	cases := []struct {
		in   string
		want time.Duration
	}{
		{"1h30m", 90 * time.Minute},
		{"90m", 90 * time.Minute},
		{"100", 100 * time.Second},
		{"1天2小时30分钟", 24*time.Hour + 2*time.Hour + 30*time.Minute},
		{"30秒", 30 * time.Second},
	}
	for _, c := range cases {
		got, err := Parse(c.in)
		if err != nil {
			t.Fatalf("Parse(%q) err=%v", c.in, err)
		}
		if got != c.want {
			t.Errorf("Parse(%q)=%v want %v", c.in, got, c.want)
		}
	}
	if _, err := Parse("乱七八糟"); err == nil {
		t.Error("expected error for junk")
	}
}
