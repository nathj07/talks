package conversation

import "testing"

func TestGreeting(t *testing.T) {
	// call Greeting to ensure it exists
	if got := Greeting("Hello"); got != "Hello" {
		t.Errorf("Got %s; Expected Hello", got)
	}
}

// FR_ES OMIT
func TestGreetingFrech(t *testing.T) {
	expected := "Salut, ça va ?"
	if got := Greeting("Salut"); got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

func TestGreetingSpanish(t *testing.T) {
	expected := "Hola, ¿Cómo estás?"
	if got := Greeting("Hola"); got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

// FR_ES OMIT

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

// TABLE OMIT
var testTable = []struct {
	tag      string // identify the test case
	input    string
	expected string
}{
	{
		tag:      "French",
		input:    "Salut",
		expected: "Salut, ça va ?",
	},
	{
		tag:      "Spanish",
		input:    "Hola",
		expected: "Hola, ¿Cómo estás?",
	},
}

// TABLE OMIT

func TestGreetingV2All(t *testing.T) {
	for _, td := range testTable {
		if got := GreetingV2(td.input); got != td.expected {
			t.Errorf("Tag: %s: Expected: %s; Got: %s", td.tag, td.expected, got)
		}
	}
}

// Extras for demo
/*
// "Bug"
{
	tag: "English Informal",
	input: "Hello, old friend",
	expected: "Good to see you again",
},
// Increase converage
{
	tag:      "Default",
	input:    "Good morning",
	expected: "Good morning",
},
// Add extra options
{
	tag:      "Shop Greeting",
	input:    "Hello",
	expected: "How may I help you?",
},
{
	tag:      "Simple French",
	input:    "Bonjour",
	expected: "Salut",
},*/
