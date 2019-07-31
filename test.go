package main

import (
	"fmt"
)

func main() {
	fmt.Println("testprogram")
	DoStuff()
}

//GenerateReportHelper TODO: document exported type
type GenerateReportHelper struct {
	ProjectRoot    string
	RunLocal       bool
	BazelTargets   []string
	IncludeFiles   string
	CiJobName      string
	NoCoverageFile string
}

const (
	//HELLO TODO: document exported value
	HELLO = "hello"
	//WORLD TODO: document exported value
	WORLD = "world"
)

//BAR TODO: document exported value
const BAR = "bar"

//FOO TODO: document exported value
var FOO string

func unexportedFunction() {

}

// Whatever does other stuff
func Whatever() {

}

//AnExportedFunction TODO: document exported function
func AnExportedFunction() {
	XXX, YYY := "x", "y"
	fmt.Println(XXX, YYY)
}

//DoStuff TODO: document exported function
func DoStuff() {

}

// DoOtherStuff does other stuff
func DoOtherStuff() {

}

//FooMethod TODO: document exported function
func (f *FOO) FooMethod() {

}

type itest interface {
}
