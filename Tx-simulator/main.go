package main

import (
	"fmt"
  //"os"
  //"github.com/joho/godotenv"
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
	/**err := godotenv.Load()
  if err != nil {
    fmt.Println("Error Loading enviromental variables", err)
    return
  }

  //get the loaded enviromental variables and assign
  //them to the values below
  tenderlyUser := os.Getenv("TENDERLY_USER")
  tenderlyAccessKey := os.Getenv("TENDERLY_ACCESS_KEY")
  tenderlyProject := os.Getenv("TENDERLY_PROJECT") **/

  //get user input and assign it to the request data struct
  var requestData = getUserInput()
  fmt.Println(requestData)

}

//get user input
func getUserInput() RequestData {
  var saveSimulation bool
  var saveSimulationIfFails bool
  var simulationType string
  var networkId string
  var fromAddr string
  var toAddr string
  var inputData string
  var gasToSend uint64
  var gasPriceToUse uint64
  var valueToSend uint64

  fmt.Println("Save Simulation(true/false):")
  fmt.Scan(&saveSimulation)

  fmt.Println("Save Simulation if it fails(true/false):")
  fmt.Scan(&saveSimulationIfFails)

  fmt.Println("Simulation type(full/quick):")
  fmt.Scan(&simulationType)

  fmt.Println("Network Id(eg: 1):")
  fmt.Scan(&networkId)

  fmt.Println("From(addr):")
  fmt.Scan(&fromAddr)
  
  fmt.Println("To(addr)")
  fmt.Scan(&toAddr)
  
  fmt.Println("Input Data:")
  fmt.Scan(&inputData)
  
  fmt.Println("Gas:")
  fmt.Scan(&gasToSend)
  
  fmt.Println("Gas Price:")
  fmt.Scan(&gasPriceToUse)
  
  fmt.Println("Value:")
  fmt.Scan(&valueToSend)

  requestData := RequestData{
    save:saveSimulation,          
    save_if_fails:saveSimulationIfFails,  
    simulation_type:simulationType,
    network_id:networkId,
    from:fromAddr,       
    to:toAddr,        
    input:inputData,  
    gas:gasToSend,     
    gas_price:gasPriceToUse,
    value:valueToSend,    
  }
  
  return requestData
  
}



    
   
