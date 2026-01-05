package main

import (
	"fmt"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	if index < 0 || index >= len(j.entries) {
		panic("Index out of range")
	} else {
		j.entries = append(j.entries[:index], j.entries[index+1:]...)
	}
}

// func (j *Journal) Save(filename string) {
// 	_ = io.Writer
// }

func main() {
	j := Journal{}
	j.AddEntry("Aliba")
	j.AddEntry("Madi")
	j.AddEntry("Galym")
	fmt.Println(j)
	j.RemoveEntry(1)
	fmt.Println(j)
}
