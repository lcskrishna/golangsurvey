package main
import(
	"fmt"
	"math/rand"
	"time"
)

// A Function which is being executed by go routine.
func helperFunction(n int){
	for i:=0; i<10;i++ {

		// Used for printing the number.
		fmt.Println(n,":",i)

		// Adding delay to the routine to demonstrate the concurrency using go routines better.
		amt:= time.Duration(rand.Intn(350))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	 for i := 0; i < 10; i++ {
    	go helperFunction(i)
  	}
	var input string
	fmt.Scanln(&input)
	fmt.Println("You have entered the number: ",input)
}