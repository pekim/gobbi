package generate

import (
	"fmt"
	"sync"
)

type repositorySpec struct {
	Version string
	Name    string
}

type RepositorySpecs []repositorySpec

func (ss RepositorySpecs) Generate() {
	ss.ansi("25l") // hide cursor

	var wg sync.WaitGroup

	namespaces := make(namespaces)
	rr := []*repository{}

	progressLineLen := 2 * len(ss)
	ss.prepareProgressLine(progressLineLen)

	var mu sync.Mutex
	for _, spec := range ss {
		wg.Add(1)

		go func(spec repositorySpec) {
			defer wg.Done()

			r := repositoryFromFile(spec)

			mu.Lock()
			defer mu.Unlock()

			ss.incrementProgress()
			namespaces[r.Namespace.Name] = r.Namespace
			rr = append(rr, r)
		}(spec)
	}
	wg.Wait()

	// Deep initialise all namespaces.
	// In particular provide descendants with their namespace.
	for _, r := range rr {
		r.Namespace.init(namespaces)
	}

	// Generate files for all namespaces
	unsupportedCount := 0
	for _, r := range rr {
		wg.Add(1)

		go func(r *repository) {
			r.Namespace.generate()

			mu.Lock()
			defer mu.Unlock()

			unsupportedCount += r.Namespace.unsupportedCount
			ss.incrementProgress()

			wg.Done()
		}(r)
	}
	wg.Wait()

	ss.clearProgressLine(progressLineLen)
	ss.ansi("25h") // show cursor
	fmt.Printf("%d unsupported\n", unsupportedCount)
}

func (ss RepositorySpecs) ansi(sequence string) {
	if !isTerminal() {
		return
	}

	fmt.Printf("\u001B[%s", sequence)
}

func (ss RepositorySpecs) moveToStartOfLine() {
	ss.ansi("G") // start of line
}

func (ss RepositorySpecs) lineOfChars(char string, len int) {
	if !isTerminal() {
		return
	}

	ss.moveToStartOfLine()
	for i := 0; i < len; i++ {
		fmt.Print(char)
	}
	ss.moveToStartOfLine()
}

func (ss RepositorySpecs) prepareProgressLine(len int) {
	ss.lineOfChars("\u2591", len) // light shade block
}

func (ss RepositorySpecs) clearProgressLine(len int) {
	ss.lineOfChars(" ", len)
}

func (ss RepositorySpecs) incrementProgress() {
	if !isTerminal() {
		return
	}

	fmt.Print("\u2592") // medium shade block
}
