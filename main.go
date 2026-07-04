package main

import (
	"fmt"
	"os"
	"strings"
)

type Pet struct {
	Name      string
	Animal    string
	Hunger    int
	Happiness int
	Health    int
	Energy    int
}

func (p *Pet) ShowStatus() {
	emoji := "🐱"
	if p.Animal == "dog" {
		emoji = "🐶"
	} else if p.Animal == "bunny" {
		emoji = "🐰"
	}

	fmt.Printf("\n❤️  %s %s  ❤️\n", emoji, p.Name)
	fmt.Printf("Hunger    : %d/10\n", p.Hunger)
	fmt.Printf("Happiness : %d/10\n", p.Happiness)
	fmt.Printf("Health    : %d/10\n", p.Health)
	fmt.Printf("Energy    : %d/10\n", p.Energy)
	fmt.Println("-----------------------------")
}

func main() {
	fmt.Println("🌟 Welcome to Cute Pet Simulator! 🌟\n")

	pet := Pet{
		Name:      "Muffin",
		Animal:    "cat",
		Hunger:    5,
		Happiness: 6,
		Health:    8,
		Energy:    7,
	}

	for {
		pet.ShowStatus()

		fmt.Println("What do you want to do?")
		fmt.Println("1. Feed 🍎")
		fmt.Println("2. Play  🎾")
		fmt.Println("3. Sleep  😴")
		fmt.Println("4. Quit")

		fmt.Print("\nEnter choice (1-4): ")

		var choice string
		fmt.Scanln(&choice)

		switch strings.TrimSpace(choice) {
		case "1":
			pet.Hunger = max(0, pet.Hunger-3)
			pet.Happiness++
			fmt.Println("🍎 Yum! Muffin enjoyed the food!")
		case "2":
			pet.Happiness += 3
			pet.Energy = max(0, pet.Energy-2)
			fmt.Println("🎾 Woohoo! You played with Muffin!")
		case "3":
			pet.Energy += 4
			pet.Hunger++
			fmt.Println("😴 Muffin is sleeping... zzz")
		case "4":
			fmt.Println("Goodbye! Thanks for taking care of Muffin ❤️")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice! Please try 1, 2, 3 or 4.")
		}

		// Keep stats in reasonable range
		if pet.Happiness > 10 {
			pet.Happiness = 10
		}
		if pet.Energy > 10 {
			pet.Energy = 10
		}
	}
}

// Helper function (Go doesn't have built-in max for int until recent versions)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
