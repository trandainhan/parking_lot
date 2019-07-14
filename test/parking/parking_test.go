package parking_test

import (
	"github.com/stretchr/testify/assert"
	. "parking_lot/internal/parking"
	"testing"
)

func TestNewParkingLot(t *testing.T) {
	parkingLot := NewParkingLot(1)
	if assert.NotNil(t, parkingLot, "parkingLot should be created") {
		assert.Equal(t, 1, parkingLot.Size, "parkingLot created with the right size")
	}
	assert.Nil(t, parkingLot.Slots[1], "parkinglot have an empty slot for parking")
}
