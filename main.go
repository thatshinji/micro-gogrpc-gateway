package main

import (
	"fmt"
	tracepb "micro-gogrpc-gateway/server/proto/gen/go"
)

func main() {
	trace := tracepb.Trace{
		Time:     1231313,
		Location: "shanghai",
		StaySec:  100,
		VieePage: "home",
	}
	fmt.Println(trace)
}
