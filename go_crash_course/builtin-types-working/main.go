package main

import (
	"fmt"
	"log"
	"myapp/channels"
	"sort"
)

var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myUint uint

var myFloat float32
var myFloat64 float64

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Anim interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs.\n", a.Name, a.NumberOfLegs)
}

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {

	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := ""

	log.Println(myString)
	myString = "John"

	var myBool = true
	log.Println(myBool)

	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in the array is", myStrings[0])

	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My car is a %d %s %s.\n", myCar.Year, myCar.Make, myCar.Model)

	x := 10
	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)
	fmt.Println("After function called x is now", x)

	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element zero is", animals[0])

	fmt.Println("First two animals are:", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)

	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5
	intMap["six"] = 6

	// for key, value := range intMap {
	// 	fmt.Println(key, value)
	// }

	// delete(intMap, "four")

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "isn't in map")
	}

	intMap["two"] = 22

	fmt.Println(intMap)

	z := addTwoNumbers(2, 4)
	fmt.Println(z)

	m := multiplyTwoNumbers(9, 8)
	fmt.Println(m)

	myTotal := sumMany(2, 3, 4, 5)
	fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()

	channels.GoRoutinesExample()

	doggy := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(&doggy)

	var kitten Cat
	kitten.Name = "cat"
	kitten.Sound = "meow"
	kitten.NumberOfLegs = 4
	kitten.HasTail = true

	Riddle(&kitten)

	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t.\n", name, age, rightHanded)

	ageInTenYears := age + 10

	isATeenager := age >= 13

	fmt.Printf("In 10 years %s will be %d years old.\n", name, ageInTenYears)
	fmt.Printf("%s is a teenager: %t\n", name, isATeenager)

	apples := 18
	oranges := 9

	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// if apples == "10" {
	// 	fmt.Print("yes")
	// }

	fmt.Printf("%d > %d: %t\n", apples, oranges, apples > oranges)
	fmt.Printf("%d >= %d: %t\n", apples, oranges, apples >= oranges)
	fmt.Printf("%d <= %d: %t\n", apples, oranges, apples <= oranges)
	fmt.Printf("%d < %d: %t\n", apples, oranges, apples < oranges)

	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 years older")
		} else {
			fmt.Println(x.Name, "is 30 years younger")
		}

		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "is over 30 and makes over 50K")
		} else {
			fmt.Println(x.Name, "is less 30 and makes under 50K")
		}
		if x.Age > 30 || x.Salary > 30000 {
			fmt.Println(x.Name, "is over 30 or makes over 30K")
		} else {
			fmt.Println(x.Name, "is less 30 or makes under 30K")
		}

		if x.Age > 30 || x.Salary < 50000 && x.FullTime {
			fmt.Println(x.Name, "matches our unclear criteria")
		}
	}

}

func changeValueOfPointer(num *int) {
	*num = 25
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

func addTwoNumbers(x, y int) (sum int) {
	sum = x + y
	return
}

func multiplyTwoNumbers(x, y int) int {
	return x * y
}

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}

func Riddle(a Anim) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
