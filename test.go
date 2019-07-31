package main

import (
	"fmt"
)

func main() {
	fmt.Println("testprogram")
	DoStuff()
}

type GenerateReportHelper struct {
	ProjectRoot	string
	RunLocal	bool
	BazelTargets	[]string
	IncludeFiles	string
	CiJobName	string
	NoCoverageFile	string
}

const (
	HELLO	= "hello"
	WORLD	= "world"
)

const BAR = "bar"

var FOO string

func unexportedFunction() {

}

// Whatever does other stuff
func Whatever() {

}

func AnExportedFunction() {
	XXX, YYY := "x", "y"
	fmt.Println(XXX, YYY)
}

func DoStuff() {

}

// DoOtherStuff does other stuff
func DoOtherStuff() {

}
