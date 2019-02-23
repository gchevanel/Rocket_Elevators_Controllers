import time


class ElevatorController:
    def __init__(self, nb_of_floor, nb_of_elevator):
        self.nb_of_floor = nb_of_floor
        self.nb_of_elevator = nb_of_elevator
        self.column = Column(nb_of_floor, nb_of_elevator)
        print("Controller iniatiated")

#   ---------here is where we are going to find the best elevetir and send the resquest-----------------------------------------------------------------------------

    def RequestElevator(self, FloorNumber, Direction):
        time.sleep(1)
        print("---------------------------------------------------")
        print("Request elevator to floor : ", FloorNumber)
        time.sleep(1)
        print("---------------------------------------------------")
        print("Call Button Light On")
        time.sleep(1)
        elevator = self.find_best_elevator(FloorNumber, Direction)
        elevator.send_request(FloorNumber)
        return elevator
# ----------send the request floor to send request------------------------------------------------------

    def RequestFloor(self, elevator, RequestedFloor):
        time.sleep(1)
        print("---------------------------------------------------")
        print("Requested floor : ", RequestedFloor)
        time.sleep(1)
        print("---------------------------------------------------")
        print("Request Button Light On")
        time.sleep(1)
        elevator.send_request(RequestedFloor)

#  -------------determine the best elevar for the call floor-------------------------------------------------------------------------

    def find_best_elevator(self, FloorNumber, Direction):
        bestElevator = None
        shortest_distance = 1000
        for elevator in (self.column.elevator_list):
            if (FloorNumber == elevator.elevator_floor and (elevator.status == "stopped" or elevator.status == "idle" or elevator.status == "moving")):
                return elevator
            else:
                ref_distance = abs(FloorNumber - elevator.elevator_floor)
                if shortest_distance > ref_distance:
                    shortest_distance = ref_distance
                    bestElevator = elevator

                elif elevator.Direction == Direction:
                    bestElevator = elevator

        return bestElevator


class Elevator:
    def __init__(self, elevator_no, status, elevator_floor, elevator_direction):
        self.elevator_no = elevator_no
        self.status = status
        self.elevator_floor = elevator_floor
        self.elevator_direction = elevator_direction
        self.floor_list = []
#  -------------here is the stage before operate the elevator-------------------------------------------------------------------------

    def send_request(self, RequestedFloor):
        self.floor_list.append(RequestedFloor)
        self.compute_list()
        self.operate_elevator(RequestedFloor)

# -------------here is the stage before operate the elevator-------------------------------------------------------------------------

    def compute_list(self):
        if self.elevator_direction == "up":
            self.floor_list.sort()
        elif self.elevator_direction == "down":
            self.floor_list.sort()
            self.floor_list.reverse()
        return self.floor_list

 # -------------here is the stage before operate the elevator-------------------------------------------------------------------------

    def operate_elevator(self, RequestedFloor):
        while (len(self.floor_list) > 0):
            if ((RequestedFloor == self.elevator_floor)):
                self.Open_door()
                self.status = "moving"

                self.floor_list.pop()
            elif (RequestedFloor < self.elevator_floor):

                self.status = "moving"
                print("---------------------------------------------------")
                print("Elevator", self.elevator_no, self.status)
                print("---------------------------------------------------")
                self.Direction = "down"
                self.Move_down(RequestedFloor)
                self.status = "stopped"
                print("---------------------------------------------------")
                print("Elevator", self.elevator_no, self.status)
                print("---------------------------------------------------")
                self.Open_door()
                self.floor_list.pop()
            elif (RequestedFloor > self.elevator_floor):

                time.sleep(1)
                self.status = "moving"
                print("---------------------------------------------------")
                print("Elevator", self.elevator_no, self.status)
                print("---------------------------------------------------")
                self.Direction = "up"
                self.Move_up(RequestedFloor)
                self.status = "stopped"
                print("---------------------------------------------------")
                print("Elevator", self.elevator_no, self.status)
                print("---------------------------------------------------")

                self.Open_door()

                self.floor_list.pop()

        if self.floor_list == 0:
            self.status = "idle"


# -------------open and close door-----------------------------------------------------------------------------

    def Open_door(self):
        time.sleep(1)
        print("Open Door")
        print("---------------------------------------------------")
        print("Button Light Off")
        time.sleep(1)
        print("---------------------------------------------------")
        time.sleep(1)
        self.Close_door()

    def Close_door(self):
        print("close door")
        time.sleep(1)


# -------------MOVE THE ELEVATOR---------------------------------------------------------------------------------

    def Move_up(self, RequestedFloor):
        print("Floor : ", self.elevator_floor)
        time.sleep(1)
        while(self.elevator_floor != RequestedFloor):
            self.elevator_floor += 1
            print("Floor : ", self.elevator_floor)
            time.sleep(1)

    def Move_down(self, RequestedFloor):
        print("Floor : ", self.elevator_floor)
        time.sleep(1)
        while(self.elevator_floor != RequestedFloor):
            self.elevator_floor -= 1
            print("Floor : ", self.elevator_floor)

            time.sleep(1)


class Call_button:
    def __init__(self, FloorNumber, Direction):
        self.FloorNumber = FloorNumber
        self.Direction = Direction
        self.light = False


class Floor_button:
    def __init__(self, RequestedFloor):
        self.RequestedFloor = RequestedFloor


class Column:
    def __init__(self, nb_of_floor, nb_of_elevator):
        self.nb_of_floor = nb_of_floor
        self.nb_of_elevator = nb_of_elevator
        self.elevator_list = []
        for i in range(nb_of_elevator):
            elevator = Elevator(i, "idle", 1, "up")
            self.elevator_list.append(elevator)


# // // //------------- WORKINGGGG - -------------//// //

controller = ElevatorController(10, 2)
controller.column.elevator_list[0].elevator_floor = 2

controller.column.elevator_list[0].status = "moving"
controller.column.elevator_list[0].elevator_direction = "down"
controller.column.elevator_list[1].elevator_floor = 6

controller.column.elevator_list[1].status = "moving"
controller.column.elevator_list[1].elevator_direction = "down"

elevator = controller.RequestElevator(5, "up")
controller.RequestFloor(elevator, 7)
print("==============================")
print("scene 1 ended")
print("==============================")
# // // //---------------------------------------//// //

# //////------------- WORKING - -------------//// //
# controller = ElevatorController(10, 2)
# controller.column.elevator_list[0].elevator_floor = 10
# controller.column.elevator_list[0].status = "moving"
# controller.column.elevator_list[0].elevator_direction = "down"
# controller.column.elevator_list[1].elevator_floor = 3
# controller.column.elevator_list[1].status = "moving"
# controller.column.elevator_list[1].elevator_direction = "down"


# elevator = controller.RequestElevator(1, "up")
# controller.RequestFloor(elevator, 6)
# elevator = controller.RequestElevator(3, "up")
# controller.RequestFloor(elevator, 5)
# elevator = controller.RequestElevator(9, "down")
# controller.RequestFloor(elevator, 2)
# print("==============================")
# print("scene 2 ended")
# print("==============================")
#     // ////---------------------------------------//// //


#     //////------------- WORKINGGGG - -------------//// //

# controller = ElevatorController(10, 2)

# controller.column.elevator_list[0].elevator_floor = 10
# controller.column.elevator_list[0].status = "moving"
# controller.column.elevator_list[0].elevator_direction = "down"
# controller.column.elevator_list[1].elevator_floor = 3
# controller.column.elevator_list[1].status = "moving"
# controller.column.elevator_list[1].elevator_direction = "down"

# print(controller.column.elevator_list)
# elevator = controller.RequestElevator(10, "down")
# controller.RequestFloor(elevator, 3)

# elevator = controller.RequestElevator(3, "down")
# controller.RequestFloor(elevator, 2)
# print("==============================")
# print("scene 3 ended")
# print("==============================")
