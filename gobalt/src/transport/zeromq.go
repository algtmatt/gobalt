package transport

import (
	"fmt"
	"github.com/pebbe/zmq4"
	"io/ioutil"
)
import "github.com/vmihailenco/msgpack"

// FIXME transport and constructor declared in http transport

// FIXME create golang interface for transport types and rename back to `Fetch`

func (t *Transport) Publish(publication map[string]interface{}) {
	context, _ := zmq4.NewContext()
	socket, _ := context.NewSocket(zmq4.REQ)

	defer context.Term()
	defer socket.Close()

	socket.Connect("tcp://localhost:4506")
	b, err := msgpack.Marshal(publication)
	fmt.Printf("%v", b)
	if err != nil {

	}
	socket.SendBytes(b, 0)
}

func BuildLoad(pubStructure map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"enc":  "clear",
		"load": pubStructure,
	}
}

func BuildPub(tgt string, fun string) map[string]interface{} {

	var a [0]string
	b := make(map[string]string)
	b["delimiter"] = "false"

	pubStructure := map[string]interface{}{

		"tgt_type": "glob",
		// Remember that the key rotates on master restart!
		"key":    string(getKey()),
		"tgt":    tgt,
		"cmd":    "publish",
		"user":   "sudo_mikeplace",
		"jid":    "20181221194122747334", // FIXME
		"fun":    fun,
		"ret":    "",
		"arg":    a,
		"kwargs": b,
	}

	return pubStructure
}

func getKey() (key []byte) {
	key, err := ioutil.ReadFile("/var/cache/salt/master/.root_key")
	if err != nil {
		panic(err)
	}
	return key
}
