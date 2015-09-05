package conversation

// Greeting simply replies with the greeting given
func Greeting(greeting string) string {
	return greeting
}

// GreetingV2 takes a greeting string and at the very least returns that as the response.
// it the greeting is supported then we get a relevant, in language response.
func GreetingV2(greeting string) string {

	switch greeting {
	case "Salut":
		return "Salut, ça va ?"
	case "Hola":
		return "Hola, ¿Cómo estás?"
	}
	// return at least the given greeting to seem polite
	return greeting
}

// Extras for demo
/*case "Hello":
	return "How may I help you?"
case "Bonjour":
	return "Salut"
*/
