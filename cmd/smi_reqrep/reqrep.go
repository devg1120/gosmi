//
// To use:
//
//   $ go build .
//   $ url=tcp://127.0.0.1:40899
//   $ ./reqrep node0 $url & node0=$! && sleep 1
//   $ ./reqrep node1 $url
//   $ kill $node0
//
package main

import (
	"fmt"
	"os"
	"time"
        "bytes"

	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/rep"
	"go.nanomsg.org/mangos/v3/protocol/req"

	// register transports
	_ "go.nanomsg.org/mangos/v3/transport/all"

	"encoding/json"
	//"flag"
	 "log"
	"strings"

	"github.com/sleepinggenius2/gosmi"
	"github.com/sleepinggenius2/gosmi/parser"
	"github.com/sleepinggenius2/gosmi/types"

	//        "github.com/alecthomas/repr"

	"io/ioutil"
	"path/filepath"

        //"reflect"
)

type arrayStrings []string

var modules arrayStrings
var paths arrayStrings
var paths_str string

var debug bool

type SmiType struct {
      BaseType string
      Decl string
      Description string
      Enum int
      Format string
      Name string
      Ranges []int
      Reference string
      Status string
      Units string
}

type Smi struct {
    Kind string
    Name string
    OidLen int
    Oid   []int
    Type  SmiType
    Description string
}



func prettyprint(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
  }

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

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

func node0(url string) {
        debug = true 
        paths_str := "../../mibs/vendor_mibs/"
        slice := strings.Split(paths_str, ":")
	for _, str := range slice {
		//fmt.Printf("[%s]\n", str)
		paths.Set(str)
	}

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

	err2 := Init()

	if err2 != nil {
		Exit()
	}
	fmt.Printf("node0 Init end\n")

	var sock mangos.Socket
	var err error
	var msg []byte
	if sock, err = rep.NewSocket(); err != nil {
		die("can't get new rep socket: %s", err)
	}
	if err = sock.Listen(url); err != nil {
		die("can't listen on rep socket: %s", err.Error())
	}
	for {
		// Could also use sock.RecvMsg to get header
		msg, err = sock.Recv()
		if err != nil {
			die("cannot receive on rep socket: %s", err.Error())
		}
		if string(msg) == "DATE" { // no need to terminate
			fmt.Println("NODE0: RECEIVED DATE REQUEST")
			d := date()
			fmt.Printf("NODE0: SENDING DATE %s\n", d)
			err = sock.Send([]byte(d))
			if err != nil {
				die("can't send reply: %s", err.Error())
			}
                } else {
			fmt.Println("NODE0: RECEIVED SMI REQUEST")
                        subtree := Subtree(string(msg))
                        jsonBytes, _ := json.Marshal(subtree)
			fmt.Printf("NODE0: SENDING SMI \n" )
			err = sock.Send(jsonBytes)
			if err != nil {
				die("can't send reply: %s", err.Error())
			}

		}
	}
}

func node1(url string) {
	var sock mangos.Socket
	var err error
	var msg []byte

	if sock, err = req.NewSocket(); err != nil {
		die("can't get new req socket: %s", err.Error())
	}

	//if err = sock.Dial(url); err != nil {
	//	die("can't dial on req socket: %s", err.Error())
	//}
        
        for {
	  if err = sock.Dial(url); err == nil {
	  	break
	  }
	  fmt.Printf(".")
          time.Sleep(time.Second * 2)
        } 
	fmt.Printf(".\n")

/*

https://densan-hoshigumi.com/nw/general-snmp-trap
https://densan-hoshigumi.com/nw/general-snmp-trap
https://densan-hoshigumi.com/nw/general-snmp-trap

シンボル名	OID	Trap通知契機
coldStart	1.3.6.1.6.3.1.1.5.1	電源の投入
warmStart	1.3.6.1.6.3.1.1.5.2	再起動コマンドによる再起動
linkDown	1.3.6.1.6.3.1.1.5.3	インターフェースがDown状態に変化
linkUp	1.3.6.1.6.3.1.1.5.4	インターフェースがUp状態に変化
authenticationFailure	1.3.6.1.6.3.1.1.5.5	不正なコミュニティ名のSNMPポーリングを受信


*/
	fmt.Printf("NODE1: SENDING DATE REQUEST %s\n", "SMI")
	//if err = sock.Send([]byte("DATE")); err != nil {
	//if err = sock.Send([]byte("1.3.6.1.4.1.9.1.1753")); err != nil {
	//if err = sock.Send([]byte("1.3.6.1.6.3.1.1.4")); err != nil {
	//if err = sock.Send([]byte("1.3.6.1.6.3.1.1.5.1")); err != nil {
	if err = sock.Send([]byte(".1.3.6.1.4.1.9.9.13.3.0.2")); err != nil {
		die("can't send message on push socket: %s", err.Error())
	}
	if msg, err = sock.Recv(); err != nil {
		die("can't receive date: %s", err.Error())
	}
	fmt.Printf("NODE1: RECEIVED DATE \n")
	//fmt.Printf("%s\n", msg)

        //var nodes []gosmi.SmiNode
        //json.Unmarshal([]byte(msg), &nodes)
        // fmt.Printf("%v\n", nodes)

	//fmt.Printf("%s\n", string(msg))
        b, _ := prettyprint(msg)
        fmt.Printf("%s\n", b)

    // JSONデコード
    var smis []Smi
    if err := json.Unmarshal(msg, &smis); err != nil {
        log.Fatal(err)
    }
    // デコードしたデータを表示
    for _, p := range smis {
      fmt.Printf("%s : %s : %d : %v : %s\n", p.Kind,p.Name,p.OidLen ,p.Oid, p.Type.Name)
      fmt.Printf("%s\n",p.Description)
    }
	sock.Close()
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
//		modules[i] = moduleName
                _ = i
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

func Subtree(oid string) (nodes []gosmi.SmiNode) {
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

        /*
	jsonBytes, _ := json.Marshal(subtree)
	fmt.Printf("\n")
	os.Stdout.Write(jsonBytes)
	fmt.Printf("\n")
	fmt.Printf("\n")

        fmt.Printf("subtree TypeOf: %s\n",reflect.TypeOf(subtree))
        fmt.Printf("subtree Kind  : %s\n",reflect.TypeOf(subtree).Kind())
*/
        return subtree

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
func main() {
	if len(os.Args) > 2 && os.Args[1] == "node0" {
		node0(os.Args[2])
		os.Exit(0)
	}
	if len(os.Args) > 2 && os.Args[1] == "node1" {
		node1(os.Args[2])
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, "Usage: reqrep node0|node1 <URL>\n")
	os.Exit(1)
}
