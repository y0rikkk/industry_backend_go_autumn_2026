package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return ""
	}

	shift = ((shift % n) + n) % n

	out := make([]rune, 0, n)
	out = append(out, runes[shift:]...)
	out = append(out, runes[:shift]...)
	return string(out)
}
