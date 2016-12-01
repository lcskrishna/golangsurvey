package main
import "time"
import "fmt"
func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go func() {
    for {
      c1 <- "from 1"
      time.Sleep(time.Second * 2)
    }
  }()

  go func() {
    for {
      c2 <- "from 2"
      time.Sleep(time.Second * 3)
    }
  }()

  go func() {
    for {
      select {
      case msg1 := <- c1:
        fmt.Println(msg1)
      case msg2 := <- c2:
        fmt.Println(msg2)
      }
    }
  }()

  fmt.Println("Please enter your number: ")
  var input string
  fmt.Scanln(&input)

  fmt.Println("You have entered the number: ",input)
  fmt.Println("Now you will exit from the program")
}