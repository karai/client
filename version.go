package main

// semverInfo Version string constructor
func semverInfo() string {
	majorSemver := "0"
	minorSemver := "2"
	patchSemver := "1"
	wholeString := majorSemver + "." + minorSemver + "." + patchSemver
	return wholeString
}
