package main

import (
	"fmt"
  "os"
  "go get github.com/joho/godotenv"
)

 //request data that will be sent to the tenderly api endpoint
  type RequestData struct{
    //simulation configuration
    save            bool  //if true, simulation is saved to dashboard
    save_if_fails   bool //if true, reverting simulation is saved to dashboard
    simulation_type string //full or quick(full is the default)

    //network configuration
    network_id  string //the network to run simulation on

    //Standard EVM tx object
    from      string 
    to        string
    input     string
    gas       uint64
    gas_price uint64
    value     uint64
    
  }

func main() {
  //load enviromental variables to the os and 
  //check for any error in loading them
	err := godotenv.Load()
  if err != nil {
    fmt.Println("Error Loading enviromental variables", err)
    return
  }

  //get the loaded enviromental variables and assign
  //them to the values below
  tenderlyUser := os.Getenv("TENDERLY_USER")
  tenderlyAccessKey := os.Getenv("TENDERLY_ACCESS_KEY")
  tenderlyProject := os.Getenv("TENDERLY_PROJECT")

}

//get user input
func getUserInput() {
  var saveSimulation bool
  var saveSimulationIfFails bool
  
  fmt.Println("Save Simulation(yes/no):\n")
  fmt.Scan(&saveSimulation)

  fmt.Println("Save Simulation if it fails(yes/no):\n")
  fmt.Scan(&saveSimulationIfFails)

  fmt.Println("Simulation type(full/quick):\n")
  fmt.Scan(&)
  
}
