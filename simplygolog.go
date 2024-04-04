package simplygolog

import (
	"log"
	"os"
	"time"
)

func SaveTime(functionName string, duration time.Duration)  {
  line := functionName + "," + duration.String()+"\n"
  f, err := os.OpenFile("sgl.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    log.Println(err)
  }
  defer f.Close()
  if _, err := f.WriteString(line) ; err != nil {
    log.Println("Tried benching: ",  line)
    log.Println("Got: ",  err)
  }
}

//PrintSaveTime will print out what is being saved before saving it
func PrintSaveTime(functionName string, duration time.Duration){
  log.Println("functionName", ",", duration.String())
  SaveTime(functionName, duration)
}
