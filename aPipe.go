package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// aPad holds a workorder, num and
// work payload b.
type aPad struct {
	num int
	b   []byte
}

func (p *aPad) String() string {
	return fmt.Sprintf("%d: %s", p.num, p.b)
}

// aPipe represents a pipeline of work.
// A pipeline accepts a pointer to aPad
// and returns a pointer to aPad.
func aPipe() {
	go func() {
		var (
			words = strings.Fields("The quick Brown fox jumps over the Lazy dog")
			n     int
		)
		for {
			p := new(aPad)
			p = pInit(p, words, &n)
			p = pTitle(p)
			p = pLower(p)
			fmt.Printf("output %s\n", p)
			time.Sleep(3 * time.Second)
		}
	}()
}

func pInit(p *aPad, words []string, n *int) *aPad {
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	(*n)++

	p.num = *n
	p.b = []byte(strings.Join(words, " "))
	return p
}
func pLower(p *aPad) *aPad {
	p.b = bytes.ToLower(p.b)
	return p
}
func pTitle(p *aPad) *aPad {
	p.b = bytes.ToTitle(p.b)
	return p
}
