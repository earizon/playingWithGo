package main

import (
  "fmt"
)

type MyCustomError struct {
     data1, data2 string
     data3        int
}

func (e MyCustomError) Error() string {
     return fmt.Sprintf("%s %s %d:" ,
             e.data1, e.data2, e.data3)
} 

func returnMyCustomError() (int, error) {
   return -1, MyCustomError{ "data1", "data2", 3 }
}

func throwPanic() {
   panic ("Let's panic now");
}

func main() {
   defer func() /* catch( */ {
      if r := recover(); r!=nil {
         fmt.Printf("recovered\n")
      }
   }() 

   defer func()/* finally */ {
     fmt.Printf("cleaning all resources\n")
   }()

   _, err := returnMyCustomError() 
   if  err != nil {
      fmt.Printf("%s\n",err.Error())
   }
   throwPanic()
}
