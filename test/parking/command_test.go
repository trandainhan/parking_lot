package parking_test

import (
	"github.com/stretchr/testify/assert"
	. "parking_lot/internal/parking"
	"testing"
)

func TestMultipleCommand(t *testing.T) {
	cmd := []string{"create_parking_lot 1", "park KA-01-HH-1234 White"}
	msg := ProcessMultipleCommand(cmd)
	expected := "Created a parking lot with 1 slots\nAllocated slot number: 1"
	assert.Equal(t, msg, expected, "Test multiple command")
	ResetParkingLot()
}

func TestProcessComand(t *testing.T) {
	msg := ProcessCommand("leave 1")
	assert.Equal(t, msg, "ParkingLot haven't been setup yet")
}

func TestProcessComand_Create(t *testing.T) {
	msg := ProcessCommand("create_parking_lot 1")
	assert.Equal(t, msg, "Created a parking lot with 1 slots")
}

func TestProcessComand_Park(t *testing.T) {
	msg := ProcessCommand("park KA-01-HH-1234 White")
	assert.Equal(t, msg, "Allocated slot number: 1")
}

func TestProcessComand_Leave(t *testing.T) {
	msg := ProcessCommand("leave 1")
	assert.Equal(t, msg, "Slot number 1 is free")
}

func TestProcessComand_Status(t *testing.T) {
	ProcessCommand("park KA-01-HH-1234 White")
	msg := ProcessCommand("status")
	expected := `Slot No.    Registration No    Colour
1           KA-01-HH-1234      White`
	assert.Equal(t, msg, expected)
}

func TestProcessComand_SlotNumOfColour(t *testing.T) {
	msg := ProcessCommand("slot_numbers_for_cars_with_colour White")
	assert.Equal(t, msg, "1")
}

func TestProcessComand_SlotNumOfReg(t *testing.T) {
	msg := ProcessCommand("slot_number_for_registration_number KA-01-HH-1234")
	assert.Equal(t, msg, "1")
}

func TestProcessComand_RefOfCar(t *testing.T) {
	msg := ProcessCommand("registration_numbers_for_cars_with_colour White")
	assert.Equal(t, msg, "KA-01-HH-1234")
}
