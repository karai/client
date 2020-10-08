package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"time"
)

// ascii Splash logo. We used to have a package for this
// but it was only for the logo so why not just static-print it?
func ascii() {
	fmt.Printf("\n\n")
	fmt.Printf(purple + "|  |/  / /  /\\  \\ |  |)  ) /  /\\  \\ |  |\n")
	fmt.Printf(brightpurple + "|__|\\__\\/__/¯¯\\__\\|__|\\__\\/__/¯¯\\__\\|__| \n")
	fmt.Printf(brightpurple + "v" + semverInfo() + white)
	fmt.Printf(brightgreen + " client\n")

}

// delay Delay for a number of seconds
func delay(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}

// printLicense Print the license for the user
func printLicense() {
	fmt.Printf(brightgreen + "\n" + appName + " v" + semverInfo() + white + " by " + appDev)
	fmt.Printf(brightgreen + "\n" + appRepository + "\n" + appURL + "\n")
	fmt.Printf(brightwhite + "\nMIT License\nCopyright (c) 2020-2021 RockSteadyTC")
	fmt.Printf(brightblack + "\nPermission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:\n\nThe above copyright notice and this permission notice shall be included in allcopies or substantial portions of the Software.\n\nTHE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.")
	fmt.Printf("\n")
}

// menuVersion Print the version string for the user
func menuVersion() {
	fmt.Printf("%s - v%s\n", appName, semverInfo())
}

// fileExists Does this file exist?
func fileExists(filename string) bool {
	referencedFile, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !referencedFile.IsDir()
}

// fileContainsString This is a utility to see if a string in a file exists.
func fileContainsString(str, filepath string) bool {
	accused, _ := ioutil.ReadFile(filepath)
	isExist, _ := regexp.Match(str, accused)
	return isExist
}

// timeStamp create a human readable timestamp string
func timeStamp() string {
	current := time.Now()
	return current.Format("2006-01-02 15:04:05")
}

// unixTimeStampNano UNIX epoch time in nano seconds
func unixTimeStampNano() string {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	return timestamp
}

// createMissingDirectory Check for a directory or create if missing
func createMissingDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		handle("Could not create directory: ", err)
	}
}

// checkDirs Check if directory exists
func checkDirs() {
	// This is a string array of folders to check.
	// Folder var names are in config.go
	folders := []string{
		configDir,
		configKeyDir,
		configHostsDir,
		configTxDir,
		configTxArchiveDir}

	// For each item in the folders list, check to see
	// if it exists and create the folder if it doesn't exist.
	for _, item := range folders {
		// fmt.Printf("\nChecking: %s", item)
		createMissingDirectory(item)
	}
}

// handle Ye Olde Error Handler takes a message and an error code
func handle(msg string, err error) {
	if err != nil {
		fmt.Printf(brightred+"\n%s: %s"+white, msg, err)
	}
}

// createFile Generic file handler
func createFile(filename string) {
	var _, err = os.Stat(filename)
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		handle("", err)
		defer file.Close()
	}
}

// writeFile Generic file handler
func writeFile(filename, textToWrite string) {
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	handle("", err)
	defer file.Close()
	_, err = file.WriteString(textToWrite)
	err = file.Sync()
	handle("", err)
}

// writeFileBytes Generic file handler
func writeFileBytes(filename string, bytesToWrite []byte) {
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	handle("", err)
	defer file.Close()
	createMissingDirectory(configDir)
	handle("", err)
}

// readFile read a file and output the contents as a string
func readFile(filename string) string {
	text, err := ioutil.ReadFile(filename)
	handle("Couldnt read the file: ", err)
	return string(text)
}

// readFileBytes read a file and output the contents as a byte
func readFileBytes(filename string) []byte {
	text, err := ioutil.ReadFile(filename)
	handle("Couldnt read the file: ", err)
	return text
}

// deleteFile Delete a file
func deleteFile(filename string) {
	err := os.Remove(filename)
	handle("Problem deleting file: ", err)
}

// validJSON Validate a JSON string
func validJSON(stringToValidate string) bool {
	return json.Valid([]byte(stringToValidate))
}

// validJSONBytes Validate a JSON byte
func validJSONBytes(bytesToValidate []byte) bool {
	return json.Valid(bytesToValidate)
}
