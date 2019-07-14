package parking_test

import (
	"github.com/stretchr/testify/assert"
	. "parking_lot/internal/parking"
	"testing"
)

func TestProcessComand(t *testing.T) {
	msg := ProcessCommand("park KA-01-HH-1234 White")
	assert.Equal(t, msg, "Allocated slot number: 1", "Should call park")
}
