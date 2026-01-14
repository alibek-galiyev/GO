package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)

	printSomething(err)
	printSomething(true)
	printSomething(todo)
	printSomething(1)
	printSomething(3.1415)

	fmt.Println(add(1, 5))
	fmt.Println(add("1", "5"))

}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
	}
	strVal, ok := value.(string)
	if ok {
		fmt.Println("String:", strVal)
	}
	// 	switch value.(type) {
	// 	case int:
	// 		fmt.Println("Integer:", value)
	// 	case float64:
	// 		fmt.Println("Float:", value)
	// 	case string:
	// 		fmt.Println(value)
	// 	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Save succedded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
