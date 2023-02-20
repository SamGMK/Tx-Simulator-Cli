package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
  "bytes"

	"github.com/joho/godotenv"
)

//request data that will be sent to the tenderly api endpoint
  type RequestData struct{
    
    //simulation configuration
    save            bool    //if true, simulation is saved to dashboard
    save_if_fails   bool    //if true, reverting simulation is saved to dashboard
    simulation_type string  //full or quick(full is the default)

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
    fmt.Println("Error Loading enviromental variables:", err)
    return
  }

  //get the loaded enviromental variables and assign
  //them to the values below
  tenderlyUser := os.Getenv("TENDERLY_USER")
  tenderlyAccessKey := os.Getenv("TENDERLY_ACCESS_KEY")
  tenderlyProject := os.Getenv("TENDERLY_PROJECT") 

  //get user input and assign it to the request data struct
  var requestData = getUserInput()
  fmt.Println(requestData)

  //time the simulation to see how long it takes
  //this is useful to check performance and later optimize
  startTime := time.Now()
  defer func() { 
    timeElapsed := time.Since(startTime)
    fmt.Printf("Simulation took: %s\n",timeElapsed)
  }()

  //set up the HTTP request url for tenderly simulator
  url := fmt.Sprintf("https://api.tenderly.co/api/v1/account/%v/project/%v/simulate",tenderlyUser,tenderlyProject)

  requestBody, err:= json.Marshal(requestData)
  if err != nil {
    fmt.Println("Error encoding request data in json",err)
  }

  //make a new request
  request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
  if err != nil {
    fmt.Println("Error making a new request:", err)
  }

  //set the header
  request.Header.Set("X-Access-Key", tenderlyAccessKey)




} //main



//get user input and assign it to the request data struct
func getUserInput() RequestData {
  requestData := RequestData{}
  
  fmt.Println("Save Simulation(true/false):")
  fmt.Scan(&requestData.save)

  fmt.Println("Save Simulation if it fails(true/false):")
  fmt.Scan(&requestData.save_if_fails)

  fmt.Println("Simulation type(full/quick):")
  fmt.Scan(&requestData.simulation_type)

  fmt.Println("Network Id(eg: 1):")
  fmt.Scan(&requestData.network_id)

  fmt.Println("From(addr):")
  fmt.Scan(&requestData.from)
  
  fmt.Println("To(addr)")
  fmt.Scan(&requestData.to)
  
  fmt.Println("Input Data:")
  fmt.Scan(&requestData.input)
  
  fmt.Println("Gas:")
  fmt.Scan(&requestData.gas)
  
  fmt.Println("Gas Price:")
  fmt.Scan(&requestData.gas_price)
  
  fmt.Println("Value:")
  fmt.Scan(&requestData.value)
 
  return requestData
  
}



    
   
