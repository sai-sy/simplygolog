package simplygolog

import (
	"testing"
	"time"
)

func TestSaveTime(t testing.T){
  startTime := time.Now()
  f := "TestSaveTime"
  time.Sleep(2*time.Second)
  SaveTime(f, time.Since(startTime))
}
