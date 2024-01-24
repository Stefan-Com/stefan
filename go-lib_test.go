package lib

import "testing"

func TestGetHelloMsg(t *testing.T) {
	name := "Jeramy"
	expected := "Hello Jeramy!"

	if GetHelloMsg(name) != expected {
		t.Errorf("GetHelloMsg(%v), didn't return %v", name, expected)
	}
}
