package simplygolog

import (
	"log"
	"os"
	"time"
)

func SaveTime(functionName string, duration time.Duration)  {
  line := functionName + "," + duration.String()+"\n"
  err := os.WriteFile("sgl.csv", []byte(line), 0644)
  if err != nil{
    log.Println("Tried benching: ",  line)
    log.Println("Got: ",  err)
  }
}
