package parking

type ParkingLot struct {
	Size      int
	Allocated int
	Slots     map[int]*Car
}

type Car struct {
	RegNumber string
	Colour    string
}

func NewParkingLot(size int) *ParkingLot {
	slots := make(map[int]*Car)
	for i := 1; i <= size; i++ {
		slots[i] = nil
	}
	return &ParkingLot{
		Size:      size,
		Allocated: 0,
		Slots:     slots,
	}
}

func (parkingLot *ParkingLot) GetCar(slot int) *Car {
	return parkingLot.Slots[slot]
}
