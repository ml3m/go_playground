package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
    fmt.Println("Hello, World!")
    start := 100
    sum := 0
    i := 0
    for i <= start{
        sum += i
        i++
    }
    fmt.Println("sum = ",sum)
    fmt.Println("pi = ",math.Pi)

    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Printf("%s.\n", os)

    }

    fmt.Println("\n")

    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }
    
    fmt.Println("\n")
    
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }

    my_time := time.Now()
    fmt.Println("current time in Romania/Bucharest :",my_time.In(time.UTC))
    
   // loc, err := time.LoadLocation("Europe/Bucharest")
   // if err != nil {
   //     fmt.Println("Error:", err)
   //     return
   // }

   // currentTime := time.Now().In(loc)
   // fmt.Println("Current time in Romania Bucharest", currentTime.Format("2006-01-02 15:04:05"))

   // fmt.Println("\n")
   // 
   // tt := time.Now()
   // fmt.Println("Location:", tt.Location(), ":Time:", tt)

   // utc, err := time.LoadLocation("Europe/Bucharest")
   // if err != nil {
   //     fmt.Println("err:", err.Error())
   //     return
   // }

   // berlinTime := tt.In(utc)
   // fmt.Println("Location:", berlinTime.Location(), ":Time:", berlinTime)
   //

 	//defer fmt.Println("world")
	//fmt.Println("hello")   

    //fmt.Println("\n")

	//ss := []string{"foo", "bar", "baz"}
	//fmt.Println(Index(ss, "bar"))

}
