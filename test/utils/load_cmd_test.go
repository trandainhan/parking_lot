package utils_test

import (
	"github.com/stretchr/testify/assert"
	. "parking_lot/internal/utils"
	"testing"
)

func TestLoadCmd(t *testing.T) {
	cmds := LoadCmd("./fixtures/file_input.txt")
	assert.Equal(t, len(cmds), 15, "should load correct number of commands")
	assert.Equal(t, cmds[0], "create_parking_lot 6", "should load correct command")
}
