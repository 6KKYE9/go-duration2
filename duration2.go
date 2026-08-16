package duration2

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 支持 time.ParseDuration 的全部单位，外加中文口语单位：
//
//	天 d，小时 h，分钟 m，秒 s，毫秒 ms，微秒 us/µs，纳秒 ns
//	中文：天/小时/分钟/秒/毫秒，例如 "1天2小时30分钟"
//
// 也接受纯数字（按秒处理）。解析失败返回错误。
func Parse(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty input")
	}
	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}
	// 纯数字当秒
	if n, err := strconv.ParseFloat(s, 64); err == nil {
		return time.Duration(n * float64(time.Second)), nil
	}
	// 中文单位
	if d, err := parseChinese(s); err == nil {
		return d, nil
	}
	return 0, fmt.Errorf("cannot parse duration: %q", s)
}

func parseChinese(s string) (time.Duration, error) {
	var total time.Duration
	// 单位 -> 纳秒倍数
	units := []struct {
		name string
		mul  time.Duration
	}{
		{"天", 24 * time.Hour},
		{"小时", time.Hour},
		{"分钟", time.Minute},
		{"毫秒", time.Millisecond},
		{"秒", time.Second},
	}
	i := 0
	for i < len(s) {
		// 读数字（支持小数）
		start := i
		for i < len(s) && (isDigit(rune(s[i])) || s[i] == '.') {
			i++
		}
		if i == start {
			return 0, fmt.Errorf("expected number at %q", s[i:])
		}
		num, err := strconv.ParseFloat(s[start:i], 64)
		if err != nil {
			return 0, err
		}
		// 读单位
		matched := false
		for _, u := range units {
			if strings.HasPrefix(s[i:], u.name) {
				total += time.Duration(num * float64(u.mul))
				i += len(u.name)
				matched = true
				break
			}
		}
		if !matched {
			return 0, fmt.Errorf("unknown unit at %q", s[i:])
		}
	}
	return total, nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
