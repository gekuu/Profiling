package main

import (
	"fmt"
	"os"
	"log"
	"runtime/pprof"
	"net/http"
	"runtime"
)

type Copyable interface {
	Copy() interface{}
}

func main() {

	val :=  runtime.NumCPU();
	fmt.Printf("Value : %v\n",val);
	runtime.GOMAXPROCS(runtime.NumCPU());
	fmt.Printf("Loading..!\n");
	f, err := os.Create("cpu.prof")
	if (err != nil) {
		log.Fatal(err.Error())
	}
	pprof.StartCPUProfile(f)
	for i := 0; i < 1000; i++ {
		http.Get("http://192.168.1.80:8002")
	}
	pprof.StopCPUProfile()
	fmt.Printf("Done!\n");
}