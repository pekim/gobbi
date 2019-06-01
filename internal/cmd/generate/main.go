package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	generateLibraries()
	outputScc()
}

func generateLibraries() {
	noFormat := ""
	if len(os.Args) >= 2 && os.Args[1] == "noformat" {
		noFormat = "noformat"
	}

	start := time.Now()

	// build main
	buildCmd := exec.Command("go", "build", "-o", "internal/cmd/generate-from-gir/main", "internal/cmd/generate-from-gir/main.go")
	output, err := buildCmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(output))
		os.Exit(1)
	}

	girs := []struct {
		name    string
		version string
	}{
		{"Atk", "1.0"},
		{"cairo", "1.0"},
		{"Gdk", "3.0"},
		{"GdkPixbuf", "2.0"},
		{"Gio", "2.0"},
		{"GLib", "2.0"},
		{"GObject", "2.0"},
		{"Gtk", "3.0"},
		{"GtkSource", "3.0"},
		{"Pango", "1.0"},
		{"PangoCairo", "1.0"},
		{"PangoFT2", "1.0"},
	}

	for _, gir := range girs {
		generateCmd := exec.Command("internal/cmd/generate-from-gir/main", gir.name, gir.version, noFormat)
		output, err := generateCmd.Output()

		fmt.Print(string(output))
		if err != nil {
			panic(err)
		}
	}

	end := time.Now()
	fmt.Printf("\ngeneration %.0fms\n\n", end.Sub(start).Seconds()*1000)
}
