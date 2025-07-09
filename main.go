package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name string
}

type Room struct {
	Name        string
	Description string
	Items       []Item
	Exits       map[string]*Room
	OnEnter     func() string
}

type Player struct {
	CurrentRoom *Room
	Inventory   []string
}

var player *Player
var rooms map[string]*Room

func inGame() bool {
	return player != nil && player.CurrentRoom != nil
}

func handleCommand(command string) string {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "No command provided."
	}

	cmd := parts[0]
	switch cmd {
	case "look":
		return look()

	case "walk":
		if len(parts) < 2 {
			return "Where do you want to walk?"
		}
		return walk(parts[1])

	case "take":
		if len(parts) < 2 {
			return "What do you want to take?"
		}
		return take(parts[1])

	case "wear":
		if len(parts) < 2 {
			return "What do you want to wear?"
		}
		return wear(parts[1])

	case "use":
		if len(parts) < 3 {
			return "What do you want to use and on what?"
		}
		return use(parts[1], parts[2])

	default:
		return "Unknown command."
	}
}

func look() string {
	if player == nil || player.CurrentRoom == nil {
		return "You are not in a room."
	}
	room := player.CurrentRoom
	desc := room.Description + "\nItems in the room: "
	if len(room.Items) == 0 {
		desc += "none."
	} else {
		for _, item := range room.Items {
			desc += item.Name + " "
		}
	}
	return desc
}

func walk(direction string) string {
	if player == nil || player.CurrentRoom == nil {
		return "You are not in a room."
	}
	room := player.CurrentRoom
	nextRoom, exists := room.Exits[direction]
	if !exists {
		return "You can't go that way."
	}
	player.CurrentRoom = nextRoom
	if nextRoom.OnEnter != nil {
		return nextRoom.OnEnter()
	}
	return "You have entered " + nextRoom.Name + "."
}

func take(itemName string) string {
	if player == nil || player.CurrentRoom == nil {
		return "You are not in a room."
	}
	room := player.CurrentRoom
	for i, item := range room.Items {
		if item.Name == itemName {
			player.Inventory = append(player.Inventory, item.Name)
			room.Items = append(room.Items[:i], room.Items[i+1:]...)
			return "You have taken the " + itemName + "."
		}
	}
	return "Item not found in the room."
}

func wear(itemName string) string {
	if player == nil || player.CurrentRoom == nil {
		return "You are not in a room."
	}
	for _, item := range player.Inventory {
		if item == itemName {
			return "You are now wearing the " + itemName + "."
		}
	}
	return "You don't have that item in your inventory."
}

func use(itemName, target string) string {
	if player == nil || player.CurrentRoom == nil {
		return "You are not in a room."
	}
	for _, item := range player.Inventory {
		if item == itemName {
			return "You used the " + itemName + " on " + target + "."
		}
	}
	return "You don't have that item in your inventory."
}

func main() {
	// Initialize player and rooms
	player = &Player{
		CurrentRoom: nil,
		Inventory:   []string{},
	}

	rooms = make(map[string]*Room)

	// Define rooms
	rooms["start init rooms"] = &Room{
		Name:        "Start Room",
		Description: "You are in the starting room.",
		Items:       []Item{{Name: "Key"}},
		Exits:       make(map[string]*Room),
		OnEnter: func() string {
			return "You have entered the start room."
		},
	}

	player.CurrentRoom = rooms["start init rooms"]

	// Add more rooms and exits as needed
	rooms["living room"] = &Room{
		Name:        "Living Room",
		Description: "A cozy living room with a fireplace.",
		Items:       []Item{{Name: "Book"}},
		Exits:       make(map[string]*Room),
		OnEnter: func() string {
			return "You have entered the living room."
		},
	}

	rooms["start init rooms"].Exits["north"] = rooms["living room"]
	rooms["living room"].Exits["south"] = rooms["start init rooms"]

	fmt.Println("Добро пожаловать в игру!")
	fmt.Println("Доступные команды:")
	fmt.Println("look                - осмотреться")
	fmt.Println("walk <направление>  - идти в направлении (например, walk north)")
	fmt.Println("take <предмет>      - взять предмет")
	fmt.Println("wear <предмет>      - надеть предмет")
	fmt.Println("use <предмет> <на что> - использовать предмет")
	fmt.Println("exit                - выйти из игры")
	fmt.Println("\nСписок комнат:")
	for name, room := range rooms {
		fmt.Printf("- %s (ключ: %s)\n", room.Name, name)
	}
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()
		if command == "exit" {
			fmt.Println("Exiting the game.")
			break
		}
		response := handleCommand(command)
		fmt.Println(response)
	}
}
