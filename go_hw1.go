package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
)
type Result struct{
			CurrentObservation struct{
				TempC     float64 `json:"temp_c"`
				IconURL   string  `json:"icon_url"`
			}`json:"current_observation"`		
}

func main(){
		
		uss,err :=ioutil.ReadFile("NSYSU.json")
		
		if err!=nil{
				log.Fatal(err)
		}
		var result Result
		err =json.Unmarshal([]byte(uss),&result)
		if err!=nil{
			panic(err)
		}
		fmt.Printf("%f\n%s\n", result.CurrentObservation.TempC, result.CurrentObservation.IconURL);
}