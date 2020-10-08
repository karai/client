package main

func main() {

	// Check for operating system
	osCheck()

	// Check that directories are present
	checkDirs()

	// Load keyset
	keys := initKeys()

	// Announce
	ascii()

	// Accept input from user
	inputHandler(keys)
}
