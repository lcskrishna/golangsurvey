package main 
import(
	"fmt"
	"time"
)

// A function which assigns messages to the channels.
func pinger(c chan string){
	for i:=0; ; i++ {
			c <- "ping"	
	}
}
// A function which sends the messages to channels as Pong.
func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}
// A function which prints the Channel messages.
func printer(c chan string){
	for {
		msg:= <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Println("Please enter the number")
	fmt.Scanln(&input)
	fmt.Println("You have entered the number: ",input)


}