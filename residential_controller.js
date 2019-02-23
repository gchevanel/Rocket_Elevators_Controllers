//-------------------------------------------------------------------------------------------------------------------
function test_elevator() {
  var controller = init_elevator_system(10, 2);

  //////------------- WORKINGGGG --------------//////

  //    controller.column.elevator_list[0].elevator_floor = 2; // set elevator 1 floor
  //    controller.column.elevator_list[0].status = "moving";
  //    controller.column.elevator_list[0].elevator_direction = "down";
  //    controller.column.elevator_list[1].elevator_floor = 6; // set elevator 2 floor
  //    controller.column.elevator_list[1].status = "moving";
  //    controller.column.elevator_list[1].elevator_direction = "down";

  //    var elevator = controller.RequestElevator(5, "up");
  //    controller.RequestFloor(elevator, 7);
  //   console.log("==============================");
  //   console.log("scene 1 ended");
  //   console.log("==============================")

  //////---------------------------------------//////

  //////-------------  WORKING  --------------//////

  controller.column.elevator_list[0].elevator_floor = 10;
  controller.column.elevator_list[0].status = "moving";
  controller.column.elevator_list[0].elevator_direction = "down";
  controller.column.elevator_list[1].elevator_floor = 3;
  controller.column.elevator_list[1].status = "moving";
  controller.column.elevator_list[1].elevator_direction = "down";

  var elevator = controller.RequestElevator(1, "up");
  controller.RequestFloor(elevator, 6);
  elevator = controller.RequestElevator(3, "up");
  controller.RequestFloor(elevator, 5);
  elevator = controller.RequestElevator(9, "down");
  controller.RequestFloor(elevator, 2);
  console.log("==============================");
  console.log("scene 2 ended");
  console.log("==============================");

  //////---------------------------------------//////

  //////------------- WORKINGGGG --------------//////

  // controller.column.elevator_list[0].elevator_floor = 10;
  // controller.column.elevator_list[0].status = "moving";
  // controller.column.elevator_list[0].elevator_direction = "down";
  // controller.column.elevator_list[1].elevator_floor = 3;
  // controller.column.elevator_list[1].status = "moving";
  // controller.column.elevator_list[1].elevator_direction = "down";

  // console.log(controller.column.elevator_list)
  // var elevator = controller.RequestElevator(10, "down"); /* IF change 4 pour 1, NO PROBLEM */
  // controller.RequestFloor(elevator, 3);

  // elevator = controller.RequestElevator(3, "down");
  // controller.RequestFloor(elevator, 2);
  //   console.log("==============================");
  //   console.log("scene 3 ended")
  //   console.log("==============================");

  //////---------------------------------------//////

  // ----------------INIT SYSTEM--------------------------------------------------------------
}
function init_elevator_system(nb_of_floor, nb_of_elevator) {
  var controller = new ElevatorController(nb_of_floor, nb_of_elevator);

  return controller;
}
class Column {
  constructor(nb_of_floor, nb_of_elevator) {
    this.nb_of_floor = nb_of_floor;
    this.nb_of_elevator = nb_of_elevator;
    this.elevator_list = [];
    for (let i = 0; i < this.nb_of_elevator; i++) {
      let elevator = new Elevator(i, "idle", 1, "up");
      this.elevator_list.push(elevator);
    }
  }
}

class Elevator {
  constructor(elevator_no, status, elevator_floor, elevator_direction) {
    this.elevator_no = elevator_no;
    this.status = status;
    this.elevator_floor = elevator_floor;
    this.elevator_direction = elevator_direction;
    this.floor_list = [];
  }
  //    ---------------Send the request to the compute list then to operate-----------------------------------------------------------------------------

  send_request(RequestedFloor) {
    this.floor_list.push(RequestedFloor);
    this.compute_list();
    this.operate_elevator(RequestedFloor);
  }
  //    ---------------compute list-----------------------------------------------------------------------------

  compute_list() {
    if (this.elevator_direction === "up") {
      this.floor_list.sort();
    } else if (this.elevator_direction === "down") {
      this.floor_list.sort();
      this.floor_list.reverse();
    }
    return this.floor_list;
  }
  //    --------------here is operate system where all the action and animation will be done-----------------------------------------------------------------------------

  operate_elevator(RequestedFloor) {
    while (this.floor_list > 0) {
      // READ nextfloor FROM floor_list COMMENT???
      if (RequestedFloor === this.elevator_floor) {
        this.Open_door();
        this.status = "moving";

        this.floor_list.shift();
      } else if (RequestedFloor < this.elevator_floor) {
        this.status = "moving";
        console.log("---------------------------------------------------");
        console.log("Elevator" + this.elevator_no, this.status);
        console.log("---------------------------------------------------");
        this.Direction = "down";
        this.Move_down(RequestedFloor);
        this.status = "stopped";
        console.log("---------------------------------------------------");
        console.log("Elevator" + this.elevator_no, this.status);
        console.log("---------------------------------------------------");
        this.Open_door();
        this.floor_list.shift();
      } else if (RequestedFloor > this.elevator_floor) {
        sleep(1000);
        this.status = "moving";
        console.log("---------------------------------------------------");
        console.log("Elevator" + this.elevator_no, this.status);
        console.log("---------------------------------------------------");
        this.Direction = "up";
        this.Move_up(RequestedFloor);
        this.status = "stopped";
        console.log("---------------------------------------------------");
        console.log("Elevator" + this.elevator_no, this.status);
        console.log("---------------------------------------------------");

        this.Open_door();

        this.floor_list.shift();
      }
    }
    if (this.floor_list === 0) {
      this.status = "idle";
    }
  }
  Request_floor_button(RequestedFloor) {
    this.RequestedFloor = RequestedFloor;
    this.floor_light = floor_light;
  }
  Call_floor_button(FloorNumber, Direction) {
    this.FloorNumber = FloorNumber;
    this.Direction = Direction;
  }
  //    ---------------open and close door----------------------------------------------------------------------------

  Open_door() {
    sleep(1000);
    console.log("Open Door");
    console.log("---------------------------------------------------");
    console.log("Button Light Off");
    sleep(1000);

    console.log("---------------------------------------------------");
    sleep(1000);
    this.Close_door();
  }
  Close_door() {
    console.log("close door");
    sleep(1000);
  }

  // -------------MOVE THE ELEVATOR UP---------------------------------------------------------------------------------

  Move_up(RequestedFloor) {
    console.log("Floor : " + this.elevator_floor);
    sleep(1000);
    while (this.elevator_floor !== RequestedFloor) {
      this.elevator_floor += 1;
      console.log("Floor : " + this.elevator_floor);

      sleep(1000);
    }
  }

  Move_down(RequestedFloor) {
    console.log("Floor : " + this.elevator_floor);
    sleep(1000);
    while (this.elevator_floor !== RequestedFloor) {
      this.elevator_floor -= 1;
      console.log("Floor : " + this.elevator_floor);

      sleep(1000);
    }
  }
}

class ElevatorController {
  constructor(nb_of_floor, nb_of_elevator) {
    this.nb_of_floor = nb_of_floor;
    this.nb_of_elevator = nb_of_elevator;
    this.column = new Column(nb_of_floor, nb_of_elevator);
    // console.log(this.column);

    console.log("Controller iniatiated");
  }

  //    -------here is the request elevator where find best elevator is done----------------------------------------------------------------------------------

  RequestElevator(FloorNumber, Direction) {
    sleep(1000);
    console.log("---------------------------------------------------");
    console.log("Request elevator to floor : ", FloorNumber);
    sleep(1000);
    console.log("---------------------------------------------------");
    console.log("Call Button Light On");
    sleep(1000);

    let elevator = this.find_best_elevator(FloorNumber, Direction);
    elevator.send_request(FloorNumber);
    return elevator;
  }
  //    -----------request floor-----------------------------------------------------------------------------

  RequestFloor(elevator, RequestedFloor) {
    sleep(1000);
    console.log("---------------------------------------------------");
    console.log("Requested floor : ", RequestedFloor);
    sleep(1000);
    console.log("---------------------------------------------------");
    console.log("Request Button Light On");
    sleep(1000);
    elevator.send_request(RequestedFloor);
    // Elevator.operate_elevator(RequestedFloor);
  }
  //    -----------here where the find best elevator is done-----------------------------------------------------------------------

  find_best_elevator(FloorNumber, Direction) {
    console.log("find_best_elevator", FloorNumber, Direction);

    let bestElevator = null;
    let shortest_distance = 1000;
    for (let i = 0; i < this.column.elevator_list.length; i++) {
      let elevator = this.column.elevator_list[i];

      if (
        FloorNumber === elevator.elevator_floor &&
        (elevator.status === "stopped" ||
          elevator.status === "idle" ||
          elevator.status === "moving")
      ) {
        return elevator;
      } else {
        let ref_distance = Math.abs(FloorNumber - elevator.elevator_floor);
        if (shortest_distance > ref_distance) {
          shortest_distance = ref_distance;
          bestElevator = elevator;

          if (elevator.Direction === Direction) {
            bestElevator = elevator;
          }
        }
      }
    }
    return bestElevator;
  }
}

//    -------------------------------------------------------------------------------------------------------------------

function sleep(milliseconds) {
  var start = new Date().getTime();
  for (var i = 0; i < 1e7; i++) {
    if (new Date().getTime() - start > milliseconds) {
      break;
    }
  }
}
