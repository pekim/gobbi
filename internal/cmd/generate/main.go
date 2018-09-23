package main

import (
	"fmt"
	"github.com/pekim/gobbi/internal/generate"
	"time"
)

func main() {
	start := time.Now()

	girs := generate.GirNewRoot("Gio", "2.0")

	parseDuration := time.Now().Sub(start)
	start = time.Now()

	for _, gir := range girs {
		gir.Generate()
	}

	generateDuration := time.Now().Sub(start)

	durationMessage("parse", parseDuration)
	durationMessage("generate", generateDuration)
}

func durationMessage(action string, duration time.Duration) {
	fmt.Printf("  %-8s %.0fms\n", action, duration.Seconds()*1000)
}
