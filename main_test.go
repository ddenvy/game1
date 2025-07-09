package main

import (
	"testing"
)

func TestLook(t *testing.T) {
	player = &Player{
		CurrentRoom: &Room{
			Name:        "Test Room",
			Description: "This is a test room.",
			Items:       []Item{{Name: "Test Item"}},
			Exits:       make(map[string]*Room),
		},
		Inventory: []string{},
	}

	expected := "This is a test room.\nItems in the room: Test Item "
	result := look()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestWalk(t *testing.T) {
	player = &Player{
		CurrentRoom: &Room{
			Name:        "Start Room",
			Description: "This is the start room.",
			Items:       []Item{},
			Exits: map[string]*Room{
				"north": {
					Name:        "North Room",
					Description: "This is the north room.",
					Items:       []Item{},
					Exits:       make(map[string]*Room),
				},
			},
		},
		Inventory: []string{},
	}

	result := walk("north")
	expected := "You have entered North Room."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = walk("south")
	expected = "You can't go that way."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestTake(t *testing.T) {
	player = &Player{
		CurrentRoom: &Room{
			Name:        "Test Room",
			Description: "This is a test room.",
			Items:       []Item{{Name: "Test Item"}},
			Exits:       make(map[string]*Room),
		},
		Inventory: []string{},
	}

	result := take("Test Item")
	expected := "You have taken the Test Item."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = take("Nonexistent Item")
	expected = "Item not found in the room."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestWear(t *testing.T) {
	player = &Player{
		CurrentRoom: &Room{
			Name:        "Test Room",
			Description: "This is a test room.",
			Items:       []Item{{Name: "Test Item"}},
			Exits:       make(map[string]*Room),
		},
		Inventory: []string{"Test Item"},
	}

	result := wear("Test Item")
	expected := "You are now wearing the Test Item."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = wear("Nonexistent Item")
	expected = "You don't have that item in your inventory."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestUse(t *testing.T) {
	player = &Player{
		CurrentRoom: &Room{
			Name:        "Test Room",
			Description: "This is a test room.",
			Items:       []Item{{Name: "Test Item"}},
			Exits:       make(map[string]*Room),
		},
		Inventory: []string{"Test Item"},
	}

	result := use("Test Item", "Target")
	expected := "You used the Test Item on Target."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = use("Nonexistent Item", "Target")
	expected = "You don't have that item in your inventory."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestInvalidCommands(t *testing.T) {
	player = &Player{
		CurrentRoom: nil,
		Inventory:   []string{},
	}

	result := look()
	expected := "You are not in a room."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = walk("north")
	expected = "You are not in a room."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = take("Item")
	expected = "You are not in a room."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = wear("Item")
	expected = "You are not in a room."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	result = use("Item", "Target")
	expected = "You are not in a room."
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
