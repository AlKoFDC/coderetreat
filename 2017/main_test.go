package main

import "testing"

func TestShouldDraw8(t *testing.T) {
	expect := ` _ ` + `
|_|
|_|`
	result := ToLCD(8)
	if result != expect {
		t.Fatalf("expected\n%s\nbut got\n%s", expect, result)
	}
	t.Logf("\n%s", result)
}

func TestShouldDraw88(t *testing.T) {
	result := ToLCDMultiple(8, 8, 0, 8)
	expect := ` _  _  _  _ ` + `
|_||_|| ||_|
|_||_||_||_|`
	if result != expect {
		t.Fatalf("expected\n%s\nbut got\n%s", expect, result)
	}
	t.Logf("\n%s", result)
}
