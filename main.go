package main

import (
	"fmt"
	tracepb "micro-gogrpc-gateway/server/proto/gen/go"
	"time"
)

func main() {
	trace := tracepb.Trace{
		Time:     int64(time.Now().Unix()),
		Location: "shanghai",
		StaySec:  100,
		VieePage: "home",
	}
	fmt.Println(&trace)
}
