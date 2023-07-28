package main

import (
	"fmt"
)

func printStatus(water int, milk int, beans int, cups int, money int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", beans)
	fmt.Printf("%d disposable cups\n", cups)
	fmt.Printf("$%d of money\n", money)
	fmt.Println("")
}

func buyCoffee(water *int, milk *int, beans *int, cups *int, money *int) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - black, back - to main menu:")
	var action string
	fmt.Scan(&action)

	if action == "back" {
		fmt.Println("Going to back to main menu...")
	} else {
		if *cups > 0 {
			switch action {
			case "1":
				makeEspresso(water, beans, cups, money)
			case "2":
				makeLatte(water, milk, beans, cups, money)
			case "3":
				makeCappuccino(water, milk, beans, cups, money)
			case "4":
				makeBlackCoffee(water, beans, cups, money)
			default:
				fmt.Println("Unknown coffee type!")
			}
		} else {
			fmt.Println("Sorry, not enough cups!")
		}
	}
	fmt.Println("")
}

func makeEspresso(water *int, beans *int, cups *int, money *int) {
	if checkSupply(water, 250, "water") && checkSupply(beans, 16, "coffee beans") {
		fmt.Println("I have enough resources, making you a coffee!")
		*money += 4
		*cups--
	}
}

func makeLatte(water *int, milk *int, beans *int, cups *int, money *int) {
	if checkSupply(water, 350, "water") && checkSupply(beans, 20, "coffee beans") && checkSupply(milk, 75, "milk") {
		fmt.Println("I have enough resources, making you a coffee!")
		*money += 7
		*cups--
	}
}

func makeCappuccino(water *int, milk *int, beans *int, cups *int, money *int) {
	if checkSupply(water, 200, "water") && checkSupply(beans, 12, "coffee beans") && checkSupply(milk, 100, "milk") {
		fmt.Println("I have enough resources, making you a coffee!")
		*money += 6
		*cups--
	}
}

func makeBlackCoffee(water *int, beans *int, cups *int, money *int) {
	if checkSupply(water, 400, "water") && checkSupply(beans, 30, "coffee beans") {
		fmt.Println("I have enough resources, making you a coffee!")
		*money += 8
		*cups--
	}
}

func checkSupply(supply *int, refSupply int, name string) bool {
	if *supply >= refSupply {
		*supply -= refSupply
		return true
	} else {
		fmt.Printf("Sorry, not enough %s!\n", name)
		return false
	}
}

func fillSupplies(water *int, milk *int, beans *int, cups *int) {
	fmt.Println("Write how many ml of water you want to add:")
	addSupply(water)
	fmt.Println("Write how many ml of milk you want to add:")
	addSupply(milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	addSupply(beans)
	fmt.Println("Write how many disposable cups you want to add:")
	addSupply(cups)
}

func addSupply(supply *int) {
	var additional int
	fmt.Scan(&additional)
	*supply += additional
}

func takeOutMoney(money int) int {
	fmt.Printf("I gave you $%d\n", money)
	fmt.Println("")
	return 0
}

func main() {
	// define status variables of coffee machine
	var water, milk, beans, cups, money = 400, 540, 120, 9, 550

	for {
		// read action from stdin
		fmt.Println("Write action (buy, fill, take, remaining, exit): ")
		var action string
		fmt.Scan(&action)
		fmt.Println("")

		if action == "exit" {
			break
		} else {
			switch action {
			case "buy":
				buyCoffee(&water, &milk, &beans, &cups, &money)
			case "fill":
				fillSupplies(&water, &milk, &beans, &cups)
			case "take":
				money = takeOutMoney(money)
			case "remaining":
				printStatus(water, milk, beans, cups, money)
			default:
				fmt.Println("Invalid option!")
			}
		}

	}
}
