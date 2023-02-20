package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

  type Data struct{
    
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
  //this is useful to check performance 
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

  //create a new client and send HTTP request
  client := http.DefaultClient
  response, err := client.Do(request)
  if err != nil{
    fmt.Println("Error receiving response:", err)
  }

  //defer the closing of the response body until the end of the main execution
  defer response.Body.Close()

  responseBody, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println("Error reading response body:", err)
  }

  var responseData map[string]interface{}
  err = json.Unmarshal(responseBody, &responseData)
  if err != nil{
    fmt.Println("Error unmarshalling response body:", err)
  }

  transaction := responseData["transaction"].(map[string]interface{})
  fmt.Printf("Simulated transaction info:\nhash: %s\nblock number: %v\ngas used: %v\n\n", transaction["hash"], transaction["block_number"], transaction["gas_used"])

fmt.Println("Events:")
events := transaction["transaction_info"].(map[string]interface{})["logs"]
eventsJSON, err := json.MarshalIndent(events, "", "  ")
if err != nil {
	fmt.Println("Error marshalling events data:", err)
	return
}
fmt.Println(string(eventsJSON))

} //main
//0x137019e3c8aB7574c44B0ce0324cB03CbDede737


//get user input and assign it to the request data struct
func getUserInput() Data {
  requestData := Data{}
  
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



    
   
