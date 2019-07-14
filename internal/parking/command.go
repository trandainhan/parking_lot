package parking

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	CREATE_PARKING_LOT             = "create_parking_lot"
	PARK                           = "park"
	LEAVE                          = "leave"
	STATUS                         = "status"
	SLOTS_NUMBER_OF_CAR_WITH_COLOR = "slot_numbers_for_cars_with_colour"
	SLOTS_NUMBER_OF_REG            = "slot_number_for_registration_number"
	REG_NUMBER_FOR_CARS_WITH_COLOR = "registration_numbers_for_cars_with_colour"
	EXIT                           = "exit"
)

func ProcessMultipleCommand(cmds []string) string {
	var result []string
	for _, cmd := range cmds {
		response := ProcessCommand(cmd)
		result = append(result, response)
	}
	return strings.Join(result, "\n")
}

var parkingLot *ParkingLot

func ProcessCommand(cmd string) string {
	cmdSplited := strings.Split(cmd, " ")
	command := cmdSplited[0]
	switch command {
	case CREATE_PARKING_LOT:
		value := cmdSplited[1]
		size, _ := strconv.Atoi(value)
		if parkingLot != nil {
			return fmt.Sprintf("Already created parking lot with %d slots", size)
		}
		parkingLot = NewParkingLot(size)
		return fmt.Sprintf("Created a parking lot with %d slots", size)
	case PARK:
		err := isParkingLotSetup(parkingLot)
		if err != nil {
			return err.Error()
		}
		car := &Car{
			RegNumber: cmdSplited[1],
			Colour:    cmdSplited[2],
		}
		return parkingLot.Park(car)
	case LEAVE:
		err := isParkingLotSetup(parkingLot)
		if err != nil {
			return err.Error()
		}
		value, _ := strconv.Atoi(cmdSplited[1])
		return parkingLot.Leave(value)
	case STATUS:
		err := isParkingLotSetup(parkingLot)
		if err != nil {
			return err.Error()
		}
		return parkingLot.Status()
	case SLOTS_NUMBER_OF_CAR_WITH_COLOR:
		err := isParkingLotSetup(parkingLot)
		if err != nil {
			return err.Error()
		}
		colour := cmdSplited[1]
		return parkingLot.SlotsNumberOfCarWithColour(colour)
	case SLOTS_NUMBER_OF_REG:
		err := isParkingLotSetup(parkingLot)
		if err != nil {
			return err.Error()
		}
		colour := cmdSplited[1]
		return parkingLot.SlotsNumberOfReg(colour)
	case REG_NUMBER_FOR_CARS_WITH_COLOR:
		err := isParkingLotSetup(parkingLot)
		if err != nil {
			return err.Error()
		}
		colour := cmdSplited[1]
		return parkingLot.RegNumberForCarWithColour(colour)
	default:
		return "Command not found"
	}
}

func GetParkingLot() *ParkingLot {
	return parkingLot
}

func isParkingLotSetup(parkingLot *ParkingLot) error {
	if parkingLot == nil {
		return errors.New("ParkingLot haven't been setup yet")
	}
	return nil
}
