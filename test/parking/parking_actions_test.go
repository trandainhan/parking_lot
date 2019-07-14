package parking_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	. "parking_lot/internal/parking"
	"testing"
)

var (
	parkingLot *ParkingLot
	msg        string
)

func TestMain(m *testing.M) {
	parkingLot = NewParkingLot(3)
	code := m.Run()
	os.Exit(code)
}

func TestParking(t *testing.T) {
	msg = parkingLot.Park(&Car{RegNumber: "123", Colour: "White"})
	assert.Equal(t, msg, "Allocated slot number: 1", "Sucessfully allocate slot to the car")
	msg = parkingLot.Park(&Car{RegNumber: "456", Colour: "White"})
	assert.Equal(t, msg, "Allocated slot number: 2", "Sucessfully allocate slot to another car")
	parkingLot.Park(&Car{RegNumber: "789", Colour: "Black"})
	msg = parkingLot.Park(&Car{RegNumber: "000", Colour: "Red"})
	assert.Equal(t, msg, "Sorry, parking lot is full", "Should return can not allocate the car")
}

func TestLeave(t *testing.T) {
	msg := parkingLot.Leave(1)
	assert.Equal(t, msg, "Slot number 1 is free", "Sucessfully leaving a car")
}

func TestStatus(t *testing.T) {
	msg = parkingLot.Status()
	expected := `Slot No.    Registration No    Colour
2           456                White
3           789                Black`
	assert.Equal(t, msg, expected, "Should print correct status")
}

func TestSlotsNumberOfCarWithColour(t *testing.T) {
	msg := parkingLot.SlotsNumberOfCarWithColour("White")
	assert.Equal(t, msg, "2", "Should return correct slot")

	msg = parkingLot.SlotsNumberOfCarWithColour("Red")
	assert.Equal(t, msg, "Not found", "Should not found the car")

	parkingLot.Park(&Car{RegNumber: "123", Colour: "White"})
	msg = parkingLot.SlotsNumberOfCarWithColour("White")
	assert.Equal(t, msg, "1, 2", "Should not found the car")
}

func TestSlotsNumberOfReg(t *testing.T) {
	msg = parkingLot.SlotsNumberOfReg("456")
	assert.Equal(t, msg, "2", "Should return correct slot")

	msg = parkingLot.SlotsNumberOfCarWithColour("000")
	assert.Equal(t, msg, "Not found", "Should not found the car")
}

func TestRegNumberForCarWithColour(t *testing.T) {
	msg = parkingLot.RegNumberForCarWithColour("White")
	assert.Equal(t, msg, "123, 456", "Should return correct slot")

	msg = parkingLot.RegNumberForCarWithColour("000")
	assert.Equal(t, msg, "Not found", "Should not found the car")
}
