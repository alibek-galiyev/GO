package main

import (
	"fmt"
)

type AnyValue interface{}

type Customer struct {
	Name string
}

type Staff struct {
	Role string
}

type Greeter interface {
	Greet() string
}

func (c Customer) Greet() string {
	return fmt.Sprintf("Customer %s says: Hello! How are you?", c.Name)
}

func (s Staff) Greet() string {
	return fmt.Sprintf("Staff (%s) says: Welcome to the Brew&Beans!", s.Role)
}

type PaymentMethod interface {
	Pay(amount float64) string
}

type CardInfoProvider interface {
	CardInfo() string
}

type GiftCard struct {
	Code    string
	Balance float64
}

func (g GiftCard) Pay(amount float64) string {
	if amount > g.Balance {
		return "Not enough balance"
	}
	return fmt.Sprintf("Paid $%.2f using gift card.", amount)
}

func (c GiftCard) CardInfo() string {
	return fmt.Sprintf("Current card: %s, with balance: $%.2f.\n", c.Code, c.Balance)
}

type Tea struct {
	Type string
	Size string
}

type Describable interface {
	Description() string
}

func (t Tea) Description() string {
	return fmt.Sprintf("A %s cup of %s tea.\n", t.Size, t.Type)
}

type Menu []string

type Order struct {
	Customer string
	Item     string
	Quantity int
}

type Barista interface {
	PrepareCoffee() string
}

type SeniorBarista struct {
	Name string
}

type JuniorBarista struct {
	Name string
}

func (s SeniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s prepared a caramel latte\n", s.Name)
}

func (s JuniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s prepared a hot chocolate\n", s.Name)
}

type CoffeeMachine interface {
	Brew() string
	Clean() bool
}

type CapsuleMachine struct {
	Brand string
	Model string
	Price int
}

func (machine CapsuleMachine) Brew() string {
	return fmt.Sprintf(
		"%s - %s has brewed one cup of coffee",
		machine.Brand,
		machine.Model,
	)
}

func (machine CapsuleMachine) Clean() bool {
	return true
}

type DripMachine struct {
	Brand string
	Model string
	Price int
	Time  float64
}

func (machine DripMachine) Brew() string {
	return fmt.Sprintf(
		"%s - %s has brewed a one cup of coffee",
		machine.Brand,
		machine.Model,
	)
}

func (machine DripMachine) Clean() bool {
	return true
}

func main() {
	var myMachine1 CoffeeMachine
	myMachine1 = CapsuleMachine{
		Brand: "Nespresso",
		Model: "F1000",
		Price: 1000,
	}
	fmt.Println(myMachine1.Brew())
	fmt.Printf("Cleaning the machine: %t\n", myMachine1.Clean())

	fmt.Println("----------------------------------------------")

	var myMachine2 CoffeeMachine
	myMachine2 = DripMachine{
		Brand: "Delonghi",
		Model: "T3000",
		Price: 9999,
		Time:  .15,
	}
	fmt.Println(myMachine2.Brew())
	fmt.Printf("Cleaning the machine: %t\n", myMachine2.Clean())

	fmt.Println("----------------------------------------------")
	aliba := SeniorBarista{Name: "Alibek"}
	var madi Barista = JuniorBarista{"Madishok"}
	fmt.Println(aliba.PrepareCoffee())
	fmt.Println(madi.PrepareCoffee())

	fmt.Println("----------------------------------------------")

	ServeDrink(aliba)
	ServeDrink(madi)

	fmt.Println("----------------------------------------------")

	madi = SeniorBarista{"Madishok"}

	ServeDrink(aliba)
	fmt.Println("----------------------------------------------")
	ServeDrink(madi)

	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")

	order := Order{"Alibek", "Latte", 2}
	fmt.Println(order)

	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")

	var menu Menu = []string{"Coffee", "Tea", "Cocoa"}
	fmt.Println("Menu:", menu)

	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")

	var d Describable = Tea{Type: "Green", Size: "Large"}
	fmt.Println(d.Description())

	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")

	card := GiftCard{Code: "01039834", Balance: 234234.23423}
	var pay PaymentMethod = card
	var info CardInfoProvider = card
	fmt.Println(info.CardInfo())
	fmt.Println(pay.Pay(344))

	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")

	greeters := []Greeter{
		Customer{Name: "Alibek"},
		Staff{Role: "CEO"},
		Customer{Name: "Madi"},
	}

	for _, v := range greeters {
		fmt.Println(v.Greet())
	}

	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")
	fmt.Println("----------------------------------------------")

	var any AnyValue = "Coffee"
	fmt.Println(any)

	any = 10
	fmt.Println(any)

	any = []string{"Latte", "Espresso", "Cappuccino"}
	fmt.Println(any)

	var anotherAny interface{} = "Latte"
	fmt.Println(anotherAny)

	anotherAny = 10
	fmt.Println(anotherAny)

	valuesOfDifferentTypes := []interface{}{
		"Latte",
		50.5,
		true,
		[3]int{1, 2, 3},
	}

	for _, v := range valuesOfDifferentTypes {
		fmt.Println(v)

		LogAnyValuer("Alibek")
		LogAnyValuer(true)
		LogAnyValuer(100)

		LogAnyValue(10.2)
		LogAnyValue([]string{"Aliba", "Madi", "Galym"})
	}

}

func LogAnyValuer(v interface{}) {
	fmt.Println(v)
}

func LogAnyValue(v any) {
	fmt.Println(v)
}

func ServeDrink(b Barista) {
	fmt.Println(b.PrepareCoffee())
	fmt.Println("Barista served a coffee to a client!")
}

func (o Order) String() string {
	return fmt.Sprintf("Order: %s has ordered %s (%d)", o.Customer, o.Item, o.Quantity)
}

func (m Menu) String() string {
	var res string = "["
	for i, v := range m {
		if i == len(m)-1 {
			res += v
		} else {
			res += v + ","
		}

	}
	res += "]"
	return res
}
