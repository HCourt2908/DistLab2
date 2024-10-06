package main
import (
	"flag"
//	"net/rpc"
//	"fmt"
//	"time"
//	"net"
)

var nextAddr string


func main(){
//	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
//	bottles := flag.Int("n",0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.
}
