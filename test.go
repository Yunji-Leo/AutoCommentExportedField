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

const ColumnBranch = "branch"
const ColumnFilePath = "file_path"
const ColumnReportUri = "report_file_u_r_i"
const ColumnRepo = "repository"
const ColumnLineCovered = "lines_covered"
const ColumnTimestamp = "timestamp"
const ColumnSha = "sha"
const ColumnLineMissed = "lines_missed"
const ColumnModule = "module"
const ColumnPackage = "package"
const ColumnCoverageType = "coverage_type"
const FileMetricInsertionSuccess = "file_metrics.insertion.success"
const FileMetricInsertionFailure = "file_metrics.insertion.failure"
