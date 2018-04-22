package hha

import (
	"testing"
)

func TestCreation(t *testing.T) {
	h := Create(1 << 5)
	if len(h.array) != 1<<5 {
		t.Error("wtf.")
	}
	t.Log(h)
}

func TestWrite(t *testing.T) {
	h := Create(1 << 6)
	h.OverWrite(10, []byte("abcasd;fkaje;afisdjfaklsdkcvieowvzlnd;vlksdjafiea;sldkf"))
	t.Log(h.Length(), h)
}

func TestBase64(t *testing.T) {
	h := Create(1 << 6)
	h.OverWrite(10, []byte("I'm a hackable array!"))
	x := h.Base64()
	t.Log(x)

}

func TestStringer(t *testing.T) {
	h := Create(1 << 6)
	t.Logf("%s", h)
}
