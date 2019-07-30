package main

import (
	"fmt"
)

func main() {
	fmt.Println("testprogram")
	DoStuff()
}

type GenerateReportHelper struct {
	ProjectRoot    string
	RunLocal       bool
	BazelTargets   []string
	IncludeFiles   string
	CiJobName      string
	NoCoverageFile string
}

const(
	HELLO = "hello"
	WORLD = "world"
)

func unexportedFunction() {

}

// Whatever does other stuff
func Whatever() {

}

//AnExportedFunction TODO: document exported function
func AnExportedFunction() {

}

//DoStuff TODO: document exported function
func DoStuff() {

}

// DoOtherStuff does other stuff
func DoOtherStuff() {

}
