package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sleepinggenius2/gosmi"
	"github.com/sleepinggenius2/gosmi/parser"
	"github.com/sleepinggenius2/gosmi/types"

	//        "github.com/alecthomas/repr"

	"io/ioutil"
	"path/filepath"
)

type arrayStrings []string

var modules arrayStrings
var paths arrayStrings
var paths_str string

var debug bool

func (a arrayStrings) String() string {
	return strings.Join(a, ",")
}

func (a *arrayStrings) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

/*
 *  https://oidref.com/1.3.6.1.4.1.9.1
 */

func main() {

	flag.BoolVar(&debug, "d", false, "Debug")
	flag.StringVar(&paths_str, "p", "../../mibs/vendor_mibs/", "Path to add")
	flag.Parse()

	//fmt.Printf("paths %s\n",paths_str)

	slice := strings.Split(paths_str, ":")
	for _, str := range slice {
		//fmt.Printf("[%s]\n", str)
		paths.Set(str)
	}

	//paths.Set("../../mibs/vendor_mibs/")

	for _, path := range paths {
		filelist := dirwalk(path)
		for _, filepath := range filelist {
			if debug {
				fmt.Printf(":%s\n", filepath)
			}
			module, err := parser.ParseFile(filepath)
			if err != nil {
				log.Fatalln(err)
				continue
			}
			_ = module
			if debug {
				fmt.Println(module.Name)
			}
			modules.Set(string(module.Name))
		}
	}

	err := Init()

	if err != nil {
		Exit()
	}

	oid := flag.Arg(0)
	//oid := "1.3.6.1.4.1.9.1.1753"

	if oid == "" {
		ModuleTrees()
	} else {
		Subtree(oid)
	}

	Exit()
}

func Init() error {
	gosmi.Init()

	for _, path := range paths {
		gosmi.AppendPath(path)
	}

	for i, module := range modules {
		moduleName, err := gosmi.LoadModule(module)
		if err != nil {
			fmt.Printf("Init Error: %s\n", err)
			return err
		}
		if debug {
			fmt.Printf("Loaded module %s\n", moduleName)
		}
		modules[i] = moduleName
	}

	if debug {
		path := gosmi.GetPath()
		fmt.Printf("Search path: %s\n", path)
		loadedModules := gosmi.GetLoadedModules()
		fmt.Println("Loaded modules:")
		for _, loadedModule := range loadedModules {
			fmt.Printf("  %s (%s)\n", loadedModule.Name, loadedModule.Path)
		}
	}

	return nil
}

func Exit() {
	gosmi.Exit()
	os.Exit(0)
}

func Subtree(oid string) {
	var node gosmi.SmiNode
	var err error
	if (oid[0] >= '0' && oid[0] <= '9') || oid[0] == '.' {
		node, err = gosmi.GetNodeByOID(types.OidMustFromString(oid))
	} else {
		node, err = gosmi.GetNode(oid)
	}
	if err != nil {
		fmt.Printf("Subtree Error: %s\n", err)
		return
	}

	subtree := node.GetSubtree()

	jsonBytes, _ := json.Marshal(subtree)
	fmt.Printf("\n")
	os.Stdout.Write(jsonBytes)
	fmt.Printf("\n")
}

func ModuleTrees() {
	for _, module := range modules {
		m, err := gosmi.GetModule(module)
		if err != nil {
			fmt.Printf("ModuleTrees Error: %s\n", err)
			continue
		}

		nodes := m.GetNodes()
		types := m.GetTypes()

		jsonBytes, _ := json.Marshal(struct {
			Module gosmi.SmiModule
			Nodes  []gosmi.SmiNode
			Types  []gosmi.SmiType
		}{
			Module: m,
			Nodes:  nodes,
			Types:  types,
		})
		fmt.Printf("\n")
		os.Stdout.Write(jsonBytes)
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
