package main

import (
	"fmt"
	"strings"
)

var report struct {
	rootNode *node
	nodeMap  map[string]*node
}

func init() {
	report.nodeMap = make(map[string]*node)
}

// node represent a unique lib appeared in go mod graph
type node struct {
	pkgNameWithVersion string // package name with version
	parent             []*node
	children           []*node
}

func traverse() {
}

// go mod graph | this lib
func main() {
	var lines = strings.Split(modGraph, "\n")
	for idx, line := range lines {
		pkgDep := strings.Split(line, " ")
		if len(pkgDep) != 2 {
			fmt.Println("wrong format", pkgDep)
		}

		// github.com/cch123@v1.0.1    github.com/google@v2.3.4
		pkg, dep := pkgDep[0], pkgDep[1]
		pkgNode := &node{pkgNameWithVersion: pkg}
		depNode := &node{pkgNameWithVersion: dep}
		if _, ok := report.nodeMap[pkg]; !ok {
			report.nodeMap[pkg] = pkgNode
			report.nodeMap[pkg].children = append(report.nodeMap[pkg].children, depNode)
		}

		if _, ok := report.nodeMap[dep]; !ok {
			report.nodeMap[dep] = depNode
			report.nodeMap[dep].parent = append(report.nodeMap[dep].parent, pkgNode)
		}

		if idx == 0 {
			// root pkg
			report.rootNode = pkgNode
		}
	}
	fmt.Printf("%#v\n", report)
}

var modGraph = `cch.com/c github.com/AlexStocks/goext@v0.3.2
cch.com/c github.com/AlexStocks/log4go@v1.0.7
cch.com/c github.com/apache/dubbo-go@v1.3.0
cch.com/c github.com/golang/lint@v0.0.0-20180702182130-06c8688daad7
cch.com/c github.com/samuel/go-zookeeper@v0.0.0-20190923202752-2cc03de413da
cch.com/c mosn.io/mosn@v0.11.0
github.com/AlexStocks/log4go@v1.0.7 github.com/mailru/easyjson@v0.7.1
github.com/AlexStocks/log4go@v1.0.7 github.com/mattn/go-isatty@v0.0.12
github.com/AlexStocks/log4go@v1.0.7 gopkg.in/yaml.v2@v2.2.8
gopkg.in/yaml.v2@v2.2.8 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
github.com/apache/dubbo-go@v1.3.0 github.com/Workiva/go-datastructures@v1.0.50
github.com/apache/dubbo-go@v1.3.0 github.com/afex/hystrix-go@v0.0.0-20180502004556-fa1af6a1f4f5`
