package main

import "testing"

func TestStart(t *testing.T) {
	// call start to ensure it exists
	if got := start("Hello"); got != "Hello" {
		t.Errorf("Got %s; Expected Hello", got)
	}
}

func TestStartBasicResponseFrech(t *testing.T) {
	expected := "Salut, ça va ?"
	if got := start("Salut"); got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

// START FR_ES
func TestStart2Frech(t *testing.T) {
	expected := "Salut, ça va ?"
	if got := start2("Salut"); got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

func TestStart2Spanish(t *testing.T) {
	expected := "Hola, ¿Cómo estás?"
	if got := start2("Hola"); got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

// END FR_ES

// START TABLE
type testData struct {
	tag      string // identify the test case
	input    string
	expected string
}

var testTable = []testData{
	testData{
		tag:      "French",
		input:    "Salut",
		expected: "Salut, ça va ?",
	},
	testData{
		tag:      "Spanish",
		input:    "Hola",
		expected: "Hola, ¿Cómo estás?",
	},
}

func TestStart2(t *testing.T) {
	for _, td := range testTable {
		if got := start2(td.input); got != td.expected {
			t.Errorf("Tag: %s: Expected: %s; Got: %s", td.tag, td.expected, got)
		}
	}
}

// END TABLE
