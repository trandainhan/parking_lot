package parking

import (
	"fmt"
	"strconv"
	"strings"
)

func (parkingLot *ParkingLot) Park(car *Car) string {
	if parkingLot.Allocated == parkingLot.Size {
		return "Sorry, parking lot is full"
	}
	spot := -1
	for i := 1; i <= parkingLot.Size; i++ {
		if parkingLot.Slots[i] == nil {
			spot = i
			parkingLot.Slots[i] = car
			parkingLot.Allocated++
			break
		}
	}
	return fmt.Sprintf("Allocated slot number: %d", spot)
}

func (parkingLot *ParkingLot) Leave(number int) string {
	parkingLot.Slots[number] = nil
	parkingLot.Allocated--
	return fmt.Sprintf("Slot number %d is free", number)
}

func (parkingLot *ParkingLot) Status() string {
	result := "Slot No.    Registration No    Colour"
	for i := 1; i <= parkingLot.Size; i++ {
		car := parkingLot.GetCar(i)
		if car != nil {
			result = result + fmt.Sprintf("\n%-12d%-19s%-s", i, car.RegNumber, car.Colour)
		}
	}
	return result
}

func (parkingLot *ParkingLot) SlotsNumberOfCarWithColour(colour string) string {
	var slots []string
	for i := 1; i <= parkingLot.Size; i++ {
		car := parkingLot.GetCar(i)
		if car != nil && car.Colour == colour {
			slots = append(slots, strconv.Itoa(i))
		}
	}
	if len(slots) == 0 {
		return "Not found"
	}
	return strings.Join(slots, ", ")
}

func (parkingLot *ParkingLot) SlotsNumberOfReg(reg string) string {
	for i := 1; i <= parkingLot.Size; i++ {
		car := parkingLot.GetCar(i)
		if car != nil && car.RegNumber == reg {
			return strconv.Itoa(i)
		}
	}
	return "Not found"
}

func (parkingLot *ParkingLot) RegNumberForCarWithColour(colour string) string {
	var regNumbers []string
	for i := 1; i <= parkingLot.Size; i++ {
		car := parkingLot.GetCar(i)
		if car != nil && car.Colour == colour {
			regNumbers = append(regNumbers, car.RegNumber)
		}
	}
	if len(regNumbers) == 0 {
		return "Not found"
	}
	return strings.Join(regNumbers, ", ")
}
