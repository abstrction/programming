package main

import "fmt"

func findSubstring(s, t string) bool {
	if len(s) < len(t) {
		return false
	}

	i := 0
	j := 0
	tj := 0

	for j < len(s) {
		if i >= len(t) || tj >= len(s) {
			break
		}
		if t[i] == s[tj] {
			i++
			tj++

			if i >= len(t) {
				return true
			}

			continue
		}

		j++
		tj = j
		i = 0
	}

	return false
}

func permuts(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	l := len(s)
	// to prevent duplicate
	ps := map[string]struct{}{}
	for bit := 0; bit < l; bit++ {
		ts := []byte(s)
		for i := 0; i < l; i++ {
			t := ts[bit]
			ts[bit] = ts[i]
			ts[i] = t

			ps[string(ts)] = struct{}{}
		}
	}

	r := []string{}
	for p := range ps {
		r = append(r, p)
	}

	return r
}

func main() {
	s := "aaab"
	t := "aab"
	ps := permuts(t)

	for _, p := range ps {
		r := findSubstring(s, p)
		fmt.Println(s, p, r)
	}
}