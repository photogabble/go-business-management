package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Factory struct {
	resources         []int // R(I)
	resourceCost      []int // C(I)
	finishedProducts  []int // F(I)
	productValue      []int // P(I)
	cash              int   // CH
	manufacturingCost int   // M
	month             int   // T
}

var rng *rand.Rand

func rnd(min, max int) int {
	return rng.Intn(max-min) + min
}

func input(prompt string, validator func(string) bool) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if validator(text) {
			return text
		}
	}
}

func selectInput(prompt string) int {
	command := input(fmt.Sprintf("%s (Q to return) ? ", prompt), func(s string) bool {
		return len(s) == 1 && slices.ContainsFunc(
			[]rune{'1', '2', '3', 'Q'},
			func(r rune) bool { return rune(strings.ToUpper(s)[0]) == r },
		)
	})
	if rune(strings.ToUpper(command)[0]) == 'Q' {
		return -1
	}

	i, _ := strconv.Atoi(command)
	return i
}

func numericInput(prompt string) int {
	num, _ := strconv.Atoi(input(prompt, func(s string) bool {
		return len(s) > 0 && strings.IndexFunc(s, func(r rune) bool {
			return r < '0' || r > '9'
		}) == -1
	}))

	return num
}

func initFactory() *Factory {
	f := Factory{
		month:             1,
		cash:              500,
		manufacturingCost: 2,
		resources:         make([]int, 3),
		finishedProducts:  make([]int, 3),
		resourceCost:      make([]int, 3),
		productValue:      make([]int, 3),
	}

	for i := 0; i < 3; i++ {
		f.resourceCost[i] = rnd(10, 20)
		f.productValue[i] = rnd(50, 90)
	}

	return &f
}

// display: outputs current game state for the player
func (f *Factory) display() {
	fmt.Println("Item:  Materials: Product:")
	for i := 0; i < 3; i++ {
		fmt.Printf(
			"%d%7d $%d%7d $%d\n",
			i+1,
			f.resources[i],
			f.resourceCost[i],
			f.finishedProducts[i],
			f.productValue[i],
		)
	}

	fmt.Printf("Month %d, you have $%d\n", f.month, f.cash)
	fmt.Printf("Manufacturing costs are $%d/unit\n", f.manufacturingCost)
}

func (f *Factory) purchase() {
	for {
		id := selectInput("Which material to purchase")
		if id < 0 {
			return
		}

		amount := numericInput(fmt.Sprintf("That costs $%d/unit, you have $%d. How many to purchase? ", f.resourceCost[id-1], f.cash))

		cost := amount * f.resourceCost[id-1]

		if cost > f.cash {
			fmt.Printf("Purchasing %d units would cost %d, you have insufficient funds!\n", amount, cost)
			continue
		}

		f.cash -= cost
		f.resources[id-1] += amount
	}
}

func (f *Factory) manufacture() {
	for {

		id := selectInput("Which material to manufacture")
		if id < 0 {
			return
		}

		amount := numericInput(fmt.Sprintf("Manufacturing costs $%d/unit, you have $%d. How many to manufacture? ", f.manufacturingCost, f.cash))

		cost := amount * f.manufacturingCost

		if cost > f.cash {
			fmt.Printf("Manufacturing %d units would cost %d, you have insufficient funds!\n", amount, cost)
			continue
		}

		hasMaterials := true

		for i := 0; i < 3; i++ {
			if i == id-1 {
				continue
			}
			if f.resources[i] < amount {
				hasMaterials = false
				break
			}
		}

		if hasMaterials == false {
			fmt.Println("You have insufficient materials to manufacture that much!")
			continue
		}

		for i := 0; i < 3; i++ {
			if i == id-1 {
				continue
			}
			f.resources[i] -= amount
		}

		f.cash -= cost
		f.finishedProducts[id-1] += amount
	}
}

func (f *Factory) sell() {
	for {
		id := selectInput("Which product to sell")
		if id < 0 {
			return
		}

		amount := numericInput(fmt.Sprintf("You have %d units, of that product, they sell for $%d/unit. How many to sell? ", f.finishedProducts[id-1], f.productValue[id-1]))

		if amount > f.finishedProducts[id-1] {
			fmt.Println("You have insufficient products to sell that much!")
			continue
		}

		f.cash += f.productValue[id-1] * amount
		f.finishedProducts[id-1] -= amount
	}
}

func (f *Factory) update() {
	j := 0
	for i := 0; i < 3; i++ {
		for j < 10 || j > 20 {
			j = f.resourceCost[i] + rnd(-2, 2)
		}
		f.resourceCost[i] = j

		j = 0
		for j < 50 || j > 90 {
			j = f.productValue[i] + rnd(-5, 5)
		}
		f.productValue[i] = j
	}

	j = 0
	for j < 1 || j > 9 {
		j = f.manufacturingCost + rnd(-2, 2)
	}
	f.manufacturingCost = j
}

func (f *Factory) netWorth() int {
	netWorth := f.cash
	for i := 0; i < 3; i++ {
		netWorth += f.productValue[i] * f.finishedProducts[i]
		netWorth += f.resourceCost[i] * f.resources[i]
	}
	return netWorth
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(source)

	state := initFactory()

	for state.month < 12 {
		state.display()

		command := input("Transaction (O,B,M,S) ? ", func(s string) bool {
			return len(s) == 1 && slices.ContainsFunc(
				[]rune{'O', 'B', 'M', 'S'},
				func(r rune) bool { return rune(strings.ToUpper(s)[0]) == r },
			)
		})

		switch rune(strings.ToUpper(command)[0]) {
		case 'B':
			state.purchase()
			break
		case 'M':
			state.manufacture()
			break
		case 'S':
			state.sell()
		}

		state.update()
		state.month++
	}

	fmt.Printf("Your net worth is $%d", state.netWorth())
}
