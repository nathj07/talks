package conversation

import "testing"

func TestGreeting(t *testing.T) {
	// call start to ensure it exists
	if got := Greeting("Hello"); got != "Hello" {
		t.Errorf("Got %s; Expected Hello", got)
	}
}

// START FR_ES
func TestGreetingV2Frech(t *testing.T) {
	expected := "Salut, ça va ?"
	if got := GreetingV2("Salut"); got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

func TestGreetingV2Spanish(t *testing.T) {
	expected := "Hola, ¿Cómo estás?"
	if got := GreetingV2("Hola"); got != expected {
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

// END TABLE

func TestGreetingV2All(t *testing.T) {
	for _, td := range testTable {
		if got := GreetingV2(td.input); got != td.expected {
			t.Errorf("Tag: %s: Expected: %s; Got: %s", td.tag, td.expected, got)
		}
	}
}

// Extras for demo
/*
testData{
	tag:      "English",
	input:    "Good morning",
	expected: "Good morning",
},
testData{
	tag:      "Shop Greeting",
	input:    "Hello",
	expected: "How may I help you?",
},
testData{
	tag:      "SimpleFrench",
	input:    "Bonjour",
	expected: "Salut",
},*/
