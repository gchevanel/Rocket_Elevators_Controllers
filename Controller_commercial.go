package main

import (
	"fmt"
	"sort"
	"time"
)

//this corporate elevator algorithme have been made in GoLang

//the Main function to enter your value are at the bottom

//RequestElevator and Assignelevator

//FloorNumber = the place where the customer is at
//RequestedFloor = the floor the customer want to go

// ElevatorController hold
type ElevatorController struct {
	no_of_batteries int
	batteries       []Battery
	no_of_columns   int
	user_direction  string
}

// Battery hold
type Battery struct {
	no_of_columns int
	column_list   []Column
}

// Column hold
type Column struct {
	column_no               int
	no_elevators_per_column int
	elevator_list           []Elevator
}

// Elevator hold
type Elevator struct {
	elevator_no        int
	elevator_position  int
	floor_queue        []int
	elevator_status    string
	elevator_direction string
	door_sensor        bool
	column             Column
}

// NewController is where Battery are create

func NewController(no_of_batteries int) ElevatorController {
	controller := new(ElevatorController)
	controller.no_of_batteries = 1
	for index := 0; index < no_of_batteries; index++ {
		battery := NewBattery(index)
		controller.batteries = append(controller.batteries, *battery)
	}
	return *controller
}

// NewBattery is where Column are create

func NewBattery(no_of_columns int) *Battery {
	battery := new(Battery)
	battery.no_of_columns = 4
	for index := 0; index < battery.no_of_columns; index++ {
		column := NewColumn(index)
		battery.column_list = append(battery.column_list, *column)
	}
	return battery
}

// NewColumn is where Elevator are create

func NewColumn(no_elevators_per_column int) *Column {
	column := new(Column)
	column.no_elevators_per_column = 5
	for index := 0; index < column.no_elevators_per_column; index++ {
		elevator := NewElevator()
		column.elevator_list = append(column.elevator_list, *elevator)
	}
	return column
}

// NewElevator is where Elevator are Define

func NewElevator() *Elevator {
	elevator := new(Elevator)
	elevator.elevator_position = 1
	elevator.floor_queue = []int{}
	elevator.elevator_status = "idle"
	elevator.elevator_direction = "up"
	elevator.door_sensor = true
	return elevator
}

// -------------------- List of all methods--------------------------

// here is the request made by people that want to go down
func (controller *ElevatorController) RequestElevator(FloorNumber, RequestedFloor int) Elevator {
	fmt.Println("Request elevator to floor : ", FloorNumber)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Button Light On")
	var column = controller.batteries[0].FindBestColumn(FloorNumber)
	controller.user_direction = "down"
	var elevator = column.FindBestElevator(RequestedFloor, controller.user_direction)
	elevator.Send_request(FloorNumber)
	elevator.Send_request(RequestedFloor)
	return elevator
}

// here is the request from people that want to go up to a floor X
func (controller *ElevatorController) AssignElevator(RequestedFloor int) Elevator {
	fmt.Println("Request elevator to floor : ", RequestedFloor)
	time.Sleep(3 * time.Millisecond)
	fmt.Println("Button Light On")
	column := controller.batteries[0].FindBestColumn(RequestedFloor)
	controller.user_direction = "up"
	var elevator = column.FindBestElevator(RequestedFloor, controller.user_direction)
	var FloorNumber = 1
	elevator.Send_request(FloorNumber)
	elevator.Send_request(RequestedFloor)
	return elevator
}

// here is where the best column are find
func (b *Battery) FindBestColumn(RequestedFloor int) Column { // not sure about *
	if RequestedFloor > 0 && RequestedFloor <= 22 {
		return b.column_list[3]
	} else if RequestedFloor > 22 && RequestedFloor <= 43 {
		return b.column_list[1]
	} else if RequestedFloor > 43 && RequestedFloor <= 64 {
		return b.column_list[2]
	} else if RequestedFloor > 64 && RequestedFloor <= 85 {
		return b.column_list[3]
	}
	return b.column_list[3]
}

// here is where the best elevator are found
func (c *Column) FindBestElevator(RequestedFloor int, user_direction string) Elevator {
	var selected_elevator = c.elevator_list[0]
	for _, e := range c.elevator_list {
		if RequestedFloor < e.elevator_position && e.elevator_direction == "down" && user_direction == "down" {
			selected_elevator = e
		} else if e.elevator_status == "idle" {
			selected_elevator = e
		} else if e.elevator_direction != user_direction && e.elevator_status == "moving" || e.elevator_status == "stopped" {
			selected_elevator = e
		} else if e.elevator_direction == user_direction && e.elevator_status == "moving" || e.elevator_status == "stopped" {
			selected_elevator = e
		}
	}
	return selected_elevator
}

// sendrequest receive information that people made
//in the requestelevator and assign elevator and sort my list
func (e *Elevator) Send_request(RequestedFloor int) {
	e.floor_queue = append(e.floor_queue, RequestedFloor)
	if RequestedFloor > e.elevator_position {

		sort.Ints(e.floor_queue)
	} else if RequestedFloor < e.elevator_position {

		sort.Sort(sort.Reverse(sort.IntSlice(e.floor_queue)))
	}
	e.Operate_elevator(RequestedFloor)
}

// here is where the task are separate depending on the direction
func (e *Elevator) Operate_elevator(RequestedFloor int) {
	if RequestedFloor == e.elevator_position {
		e.OpenDoor()
	} else if RequestedFloor > e.elevator_position {
		e.elevator_status = "moving"
		e.Move_up(RequestedFloor)
		e.elevator_status = "stopped"
		e.OpenDoor()
		e.elevator_status = "moving"
	} else if RequestedFloor < e.elevator_position {
		e.elevator_status = "moving"
		e.Move_down(RequestedFloor)
		e.elevator_status = "stopped"
		e.OpenDoor()
		e.elevator_status = "moving"
	}
}

// here is OpenDoor and CloseDoor
func (e *Elevator) OpenDoor() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("Door is Opening")
	time.Sleep(1 * time.Second)
	fmt.Println("Door is Open")
	time.Sleep(1 * time.Second)
	fmt.Println("Button Light Off")
	e.CloseDoor()
}
func (e *Elevator) CloseDoor() {
	if e.door_sensor == true {
		fmt.Println("Door is Closing")
		time.Sleep(1 * time.Second)
		fmt.Println("Door is Close")
		time.Sleep(1 * time.Second)
		fmt.Println("---------------------------------------------------")
		time.Sleep(1 * time.Second)
	} else if e.door_sensor {
		e.OpenDoor()
		fmt.Println("Door cant be close please make sur door is not obstruct")
	}
}

//here is move_up and move_down
func (e *Elevator) Move_up(RequestedFloor int) {
	fmt.Println("Column : ", e.column.column_no, " Elevator : #", e.elevator_no, " Current Floor :", e.elevator_position)
	for RequestedFloor > e.elevator_position {
		e.elevator_position += 1
		if RequestedFloor == e.elevator_position {
			time.Sleep(1 * time.Second)
			fmt.Println("---------------------------------------------------")
			fmt.Println("Column : ", e.column.column_no, " Elevator : #", e.elevator_no, " Arrived at destination floor : ", e.elevator_position)
		}
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Column : ", e.column.column_no, " Elevator : #", e.elevator_no, " Floor : ", e.elevator_position)
	}
}
func (e *Elevator) Move_down(RequestedFloor int) {
	fmt.Println("Column : ", e.column.column_no, " Elevator : #", e.elevator_no, " Current Floor :", e.elevator_position)
	for RequestedFloor < e.elevator_position {
		e.elevator_position -= 1
		if RequestedFloor == e.elevator_position {
			time.Sleep(1 * time.Second)
			fmt.Println("---------------------------------------------------")
			fmt.Println("Column : ", e.column.column_no, " Elevator : #", e.elevator_no, " Arrived at destination floor : ", e.elevator_position)
		}
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Column : ", e.column.column_no, " Elevator : #", e.elevator_no, " Floor : ", e.elevator_position)
	}
}

func main() {
	controller := NewController(1)
	fmt.Println(controller.batteries[0].column_list)
	controller.AssignElevator(36)
	controller.RequestElevator(33, 1)
}
