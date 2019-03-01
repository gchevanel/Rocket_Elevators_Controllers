package main

import "fmt"

// ElevatorController HOLD
type ElevatorController struct {
	no_of_floor    int
	no_of_elevator int
	no_of_column   int
	user_direction string
}

// Battery HOLD
type Battery struct {
	no_of_column int
	column_list  []Column
}

// Column HOLD
type Column struct {
	column_no                 int
	no_of_floor               int
	no_of_elevator_per_column int
	elevator_list             []Elevator
}

// Elevator HOLD
type Elevator struct {
	elevator_no        int
	elevator_floor     int
	floor_list         []int
	status             string
	elevator_direction string
	sensor             bool
}

func (controller *ElevatorController) NewElevatorController(no_of_floor,
	no_of_column, no_of_elevator int,
	user_direction string) *ElevatorController {
	// fmt.Println("NewElevatorController")
	controller.no_of_floor = no_of_floor
	controller.no_of_column = no_of_column
	controller.no_of_elevator = no_of_elevator
	controller.user_direction = user_direction
	nb := Battery{}
	nb.NewBattery(4)
	return controller
}

func (b *Battery) NewBattery(no_of_column int) *Battery {

	for index := 0; index < no_of_column; index++ {
		column := &Column{column_no: index, no_of_floor: 85, no_of_elevator_per_column: 5}
		b.column_list = append(b.column_list, *column)
	}
	fmt.Println(b.column_list)

	nc := Column{}
	nc.NewColumn(0, 85, 5)
	return b
}
func (c *Column) NewColumn(column_no, no_of_floor, no_of_elevator_per_column int) *Column {

	for index := 0; index < no_of_elevator_per_column; index++ {
		elevator := &Elevator{elevator_no: index, elevator_floor: 1, status: "idle", elevator_direction: "down", sensor: true}
		c.elevator_list = append(c.elevator_list, *elevator)
	}
	fmt.Println(c.elevator_list)

	ne := Elevator{}
	ne.NewElevator(0, 1, "idle", "up", true)
	return c
}

func (e *Elevator) NewElevator(elevator_no, elevator_floor int, status, elevator_direction string, sensor bool) *Elevator {
	e.elevator_no = elevator_no
	e.elevator_floor = elevator_floor
	e.floor_list = []int{}
	e.status = status
	e.elevator_direction = elevator_direction
	return e
}

func (b *Battery) FindBestColumn(RequestedFloor int) Column {

	if RequestedFloor > 0 && RequestedFloor <= 22 {
		return b.column_list[0]
	} else if RequestedFloor > 22 && RequestedFloor <= 43 {
		return b.column_list[1]
	} else if RequestedFloor > 43 && RequestedFloor <= 64 {
		return b.column_list[2]
	} else if RequestedFloor > 64 && RequestedFloor <= 85 {
		return b.column_list[3]

	}
	return b.column_list[0]
}

func (c *Column) FindBestElevator(RequestedFloor int, user_direction string) *Elevator {
	var best_elevator = nil
	for _, elevator := range c.elevator_list {

		if  {

		} else if elevator.status == "idle" {
		}
	}

}

func main() {
	econt := ElevatorController{no_of_floor: 85, no_of_elevator: 20, no_of_column: 4, user_direction: "down"}
	econt.NewElevatorController(85, 4, 5, "down")

}
