package main

import "fmt"

func main() {
	var coffeeSizes [3]string
	coffeeSizes = [3]string{"Large", "Medium", "Small"}

	fmt.Println(coffeeSizes)

	coffeeTypes := [3]string{"Espresso", "Latte", "Cappuccino"}

	fmt.Println(coffeeTypes)
	fmt.Printf("Address: %p\n", &coffeeTypes)
	fmt.Println("Length of the array:", len(coffeeTypes))

	dessertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}
	fmt.Println(dessertMenu)

	slice := dessertMenu[1:3]
	fmt.Println(slice)

	slice = dessertMenu[:]
	fmt.Println(slice)
	slice = append(slice, "Cheesecake")
	fmt.Println(slice)

	slice[0] = "Pie"
	fmt.Println(dessertMenu)
	fmt.Println(slice)

	menu := [3]string{"Tea", "Coffee", "Juice"}
	sliceMenu := menu[:2]
	fmt.Println("Before slice change:")
	fmt.Println("menu:", menu)
	fmt.Println("slice:", sliceMenu)
	sliceMenu[0] = "Cocoa"
	fmt.Println("After slice change:")
	fmt.Println("menu:", menu)
	fmt.Println("slice:", sliceMenu)

	ratings := []int{5, 4, 5, 5, 3}
	fmt.Println(ratings)

	ratings[3] = -99
	fmt.Println(ratings)

	coffeeTypesSlice := make([]string, 3)
	fmt.Println("coffeeTypes:", coffeeTypesSlice)
	coffeeTypesSlice[0] = "Cappuccino"
	coffeeTypesSlice[1] = "Latte"
	coffeeTypesSlice[2] = "Espresso"
	fmt.Println("coffeeTypes reassigned:", coffeeTypesSlice)

	fmt.Println("--------------------------------------")

	movies := []string{"I.T.", "E.T."}
	fmt.Println(movies)
	fmt.Println("Length:", len(movies), "Capacity:", cap(movies))
	fmt.Printf("Memeory location: %p\n", &movies)
	fmt.Printf("Memeory location of first element: %p\n", &movies[0])

	fmt.Println("--------------------------------------")

	movies = append(movies, "Titanic")
	fmt.Println(movies)
	fmt.Println("Length:", len(movies), "Capacity:", cap(movies))
	fmt.Printf("Memeory location: %p\n", &movies)
	fmt.Printf("Memeory location of first element: %p\n", &movies[0])

	fmt.Println("--------------------------------------")

	movies = append(movies, "Avatar")
	fmt.Println(movies)
	fmt.Println("Length:", len(movies), "Capacity:", cap(movies))
	fmt.Printf("Memeory location: %p\n", &movies)
	fmt.Printf("Memeory location of first element: %p\n", &movies[0])

	fmt.Println("--------------------------------------")

	movies = append(movies, "Terminator")
	fmt.Println(movies)
	fmt.Println("Length:", len(movies), "Capacity:", cap(movies))
	fmt.Printf("Memeory location: %p\n", &movies)
	fmt.Printf("Memeory location of first element: %p\n", &movies[0])

	fmt.Println("--------------------------------------")

	names := make([]string, 0, 5)
	names = append(names, "Aliba")
	names = append(names, "Galym")
	names = append(names, "Madi")
	fmt.Println("len of names:", len(names))
	fmt.Println("cap of names:", cap(names))

	fmt.Println("--------------------------------------")

	desserts := [3]string{"Cupcake", "Eclair", "Ice cream"}
	sliceDesserts := desserts[:1]
	fmt.Println(sliceDesserts)
	fmt.Println("Len of sliceDesserts:", len(sliceDesserts))
	fmt.Println("Cap of sliceDesserts:", cap(sliceDesserts))

	sliceDesserts = append(sliceDesserts, "Macaron")
	fmt.Println(desserts)

	sliceDesserts = append(sliceDesserts, "Cheesecake")
	fmt.Println(desserts)

	sliceDesserts = append(sliceDesserts, "Apple Pie")
	fmt.Println(sliceDesserts)

	sliceDesserts[0] = "Chocolate"
	fmt.Println(desserts)
	fmt.Println(sliceDesserts)

	fmt.Println("--------------------------------------")

	nbaPlayers := []string{"Kobe", "AI", "T-Mac", "Shaq", "Ray Allen"}
	fmt.Println(nbaPlayers)
	fmt.Println("Len of slice:", len(nbaPlayers))
	fmt.Println("Cap of slice:", cap(nbaPlayers))

	idxToRemove := 1
	nbaPlayers = append(nbaPlayers[:idxToRemove], nbaPlayers[idxToRemove+1:]...)
	fmt.Println(nbaPlayers)

	nbaPlayers = deleteByIndex(1, nbaPlayers)
	fmt.Println(nbaPlayers)
}

func deleteByIndex(idx int, slice []string) []string {
	return append(slice[:idx], slice[idx+1:]...)
}
