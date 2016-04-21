package main

type Str struct {
	AccField string
}

func (h *Str) Index(n int) byte {
	return h.AccField[n]
}

func NewStr(field string) Str {
	return Str{field}
}

func main() {
	s := NewStr("123")
	println(string(s.Index(0)))
}
