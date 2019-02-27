using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Timers;

namespace Commercial_controller
{

    public class ElevatorController
    {
        public int no_of_floor;
        public int no_of_elevator_per_column;
        public int no_of_column;
        public string user_direction;
        public Battery battery ;
        //public List<int> shortest_distance;
        public List<int> shortest_list;

        public ElevatorController(int no_of_floor, int no_of_column, int no_of_elevator_per_column, string user_direction)
        {
            this.no_of_floor = no_of_floor;
            this.no_of_column = no_of_column;
            this.no_of_elevator_per_column = no_of_elevator_per_column;
            this.user_direction = user_direction;
            this.battery = new Battery(this.no_of_column);


    }
    public Elevator RequestElevator(int FloorNumber, int RequestedFloor)
        {
            System.Threading.Thread.Sleep(200);
            Console.WriteLine("Request elevator to floor : " + FloorNumber);
            Console.WriteLine("---------------------------------------------------");
            System.Threading.Thread.Sleep(200);
            Console.WriteLine("Call Button Light On");
            Console.WriteLine("---------------------------------------------------"); 
            var column = battery.Find_best_column(FloorNumber);
            user_direction = "down";
            var elevator = column.Find_elevator(FloorNumber, user_direction);
            if(elevator.elevator_floor > FloorNumber)
            {
                elevator.Send_request(FloorNumber, column.column_no);
                elevator.Send_request(RequestedFloor, column.column_no);
            }

            else if (elevator.elevator_floor < FloorNumber)
            {
                elevator.Move_down(RequestedFloor, column.column_no);
                elevator.Send_request(FloorNumber,column.column_no);
                elevator.Send_request(RequestedFloor,column.column_no);
            }

            Console.WriteLine("Button Light Off");

            return elevator;
            }


        public Elevator AssignElevator(int RequestedFloor)
        {
            System.Threading.Thread.Sleep(200);
            Console.WriteLine("Requested floor : " + RequestedFloor);
            Console.WriteLine("---------------------------------------------------");
            System.Threading.Thread.Sleep(200);
            Console.WriteLine("Call Button Light On");
            Console.WriteLine("---------------------------------------------------");


            Column column = battery.Find_best_column(RequestedFloor);
            user_direction = "up";
            Elevator elevator = column.Find_elevator(RequestedFloor, user_direction);
            elevator.Open_door();

            elevator.Send_request(RequestedFloor,column.column_no);



            return elevator;
        }

    }
    public class Elevator
    {
        public int elevator_no;
        public string status;
        public int elevator_floor;
        public string elevator_direction;
        public List<int> floor_list;

        public Elevator(int elevator_no, string status, int elevator_floor, string elevator_direction)
        {
            this.elevator_no = elevator_no;
            this.status = status;
            this.elevator_floor = elevator_floor;
            this.elevator_direction = elevator_direction;
            this.floor_list = new List<int>();
        }
        public void Send_request(int RequestedFloor,char column_no)
        {
            floor_list.Add(RequestedFloor);
            if (RequestedFloor > elevator_floor)
            {
                floor_list.Sort((a, b) => a.CompareTo(b)); // EXEMPLE STACK OVERFLOW
            }
            else if (RequestedFloor < elevator_floor)
            {
                floor_list.Sort((a, b) => -1 * a.CompareTo(b));

            }

            Operate_elevator(RequestedFloor, column_no);
        }
        public void Operate_elevator(int RequestedFloor,char column_no)
        {
            if (RequestedFloor == elevator_floor)
            {
                Open_door();
                this.status = "moving";

                this.floor_list.Remove(0);
            }
            else if (RequestedFloor < this.elevator_floor)
            {
                status = "moving";
                Console.WriteLine("Button Light Off");
                Console.WriteLine("---------------------------------------------------");
                Console.WriteLine("Column : " +column_no +" Elevator : " + this.elevator_no + " " + status);
                this.elevator_direction = "down";
                Move_down(RequestedFloor, column_no);
                this.status = "stopped";
                Console.WriteLine("---------------------------------------------------");
                Console.WriteLine("Column : " + column_no + " Elevator : " + this.elevator_no + " " + status);
                this.Open_door();
                this.floor_list.Remove(0);
            }
            else if (RequestedFloor > this.elevator_floor)
            {
                System.Threading.Thread.Sleep(300);
                this.status = "moving";
                Console.WriteLine("Button Light Off");
                Console.WriteLine("---------------------------------------------------");
                Console.WriteLine("Column : " + column_no + " Elevator : " + this.elevator_no + " " + status);
                this.elevator_direction = "up";
                this.Move_up(RequestedFloor, column_no);
                this.status = "stopped";
                Console.WriteLine("---------------------------------------------------");
                Console.WriteLine("Column : " + column_no + " Elevator : " + this.elevator_no + " " + status);

                this.Open_door();

                this.floor_list.Remove(0);
            }
       
        }

        public  void Open_door()
        {
            System.Threading.Thread.Sleep(300);

            Console.WriteLine("---------------------------------------------------");

            Console.WriteLine("Open Door");
            Console.WriteLine("---------------------------------------------------");
            System.Threading.Thread.Sleep(300);

            this.Close_door();
        }
        public void Close_door()
        {
            Console.WriteLine("close door");
            System.Threading.Thread.Sleep(300);
            Console.WriteLine("---------------------------------------------------");

        }


        public void Move_up(int RequestedFloor, char column_no)
        {
            Console.WriteLine("Column : " + column_no + " Elevator : #" + elevator_no + "  Floor : " + this.elevator_floor);
            System.Threading.Thread.Sleep(300);
            while (this.elevator_floor != RequestedFloor)
            {
                this.elevator_floor += 1;
                Console.WriteLine("Column : " + column_no + " Elevator : #" + elevator_no + "  Floor : " + this.elevator_floor);

                System.Threading.Thread.Sleep(100);
            }
        }

        public void Move_down(int RequestedFloor, char column_no)
        {
            Console.WriteLine("Column : " + column_no + " Elevator : #" + elevator_no + "  Floor : " + this.elevator_floor);
            System.Threading.Thread.Sleep(100);
            while (this.elevator_floor != RequestedFloor)
            {
                this.elevator_floor -= 1;
                Console.WriteLine("Column : " + column_no + " Elevator : #" + elevator_no + "  Floor : " + this.elevator_floor);

                System.Threading.Thread.Sleep(100);
            }
        }






    }
    public class Column
    {
        public char column_no;
        public int no_of_floor;
        public int no_of_elevator_per_column;
        public List<Elevator> elevator_list;
        public List<int> call_button_list;


        public Column(char column_no, int no_of_floor, int no_of_elevator_per_column)
        {
            this.column_no = column_no;
            this.no_of_floor = no_of_floor;
            this.no_of_elevator_per_column = no_of_elevator_per_column;
            elevator_list = new List<Elevator>();
            call_button_list = new List<int>();
            for (int i = 0; i < this.no_of_elevator_per_column; i++)
            {
                Elevator elevator = new Elevator(i, "idle", 1, "up");
                elevator_list.Add(elevator);
            }
        }
        public Elevator Find_elevator(int RequestedFloor, string user_direction)
        {
            Elevator bestElevator = null;
            //int shortest_distance = 300;
            int shortest_list = 85;
            foreach (var elevator in elevator_list)
            {
                if (elevator.elevator_floor > RequestedFloor && elevator.elevator_direction == "down" && user_direction == "down")
                {
                    if (shortest_list > elevator.floor_list.Count)
                    { 
                        shortest_list = elevator.floor_list.Count;
                        bestElevator = elevator;
                    }
                }

                else if (elevator.status == "idle")
                {
                    bestElevator = elevator;
                }

                else if (elevator.status == "moving" || elevator.status == "stopped" && elevator.elevator_direction != user_direction)
                {
                    if (shortest_list > elevator.floor_list.Count)
                    {
                        shortest_list = elevator.floor_list.Count;
                        bestElevator = elevator;
                    }
                }
                else if(shortest_list > elevator.floor_list.Count)
                {
                    shortest_list = elevator.floor_list.Count;
                    bestElevator = elevator;
                }
            }
            return bestElevator;
        }

} 

    public class Battery
    {
        public int no_of_column;
        public List<Column> column_list;
        //public int FloorNumber;
        //public List<int> column_list;

        public Battery(int no_of_column)
        {
            this.no_of_column = no_of_column;
            column_list = new List<Column>();



            char cols = 'A';
            for (int i = 0; i < this.no_of_column; i++, cols++)
            {
                
                Column column = new Column(cols, 85, 5);

                column.column_no = cols;
                column_list.Add(column);

            }
        }
        public Column Find_best_column(int RequestedFloor)
        {
            Column best_column = null;
            string col = "e";
            foreach (Column column in column_list)
            {
                if (RequestedFloor > 0 && RequestedFloor <= 20)
                {
                    best_column = column_list[0];
                    col = "Column A";
                }
                else if (RequestedFloor > 20 && RequestedFloor <= 42)
                {

                    best_column = column_list[1];
                    col = "Column B";


                }
                else if (RequestedFloor > 42 && RequestedFloor <= 64)
                {
                    best_column = column_list[2];
                    col = "Column C";


                }
                else if (RequestedFloor > 64 && RequestedFloor <= 85)
                {
                    best_column = column_list[3];
                    col = "Column D";


                }

            }
            Console.WriteLine(col);
            return best_column;
        }
}





    public class CommercialCS
    {
        public static void Main(string[] args)
        {
            ElevatorController controller = new ElevatorController(85, 4, 5, "down");
            controller.AssignElevator(66);

            Elevator elevator = controller.RequestElevator(80,1);
            controller.AssignElevator(66);
            elevator = controller.RequestElevator(29,1);
            controller.AssignElevator(32);
            elevator = controller.RequestElevator(4,1);
            controller.AssignElevator(2);






        }



    }

}

//Méthode 1: RequestElevator(FloorNumber, RequestedFloor)
//Méthode 2: AssignElevator(RequestedFloor)
