/* Algorithm :
1) Read the file.
2) Convert the byte format to String format.
3) In the string search for white spaces and the dot symbol and enter character to distinguish for words.
4) Count the number of words by finding the count.
5) Write the count into a file.
@Author: Chaitanya Sri Krishna Lolla.
*/

package main
import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
	

	dat,err := ioutil.ReadFile("input.txt")
	check(err)

	var dataInFile = string(dat)
	fmt.Print(dataInFile)

	var lengthData = len(dataInFile);
	fmt.Printf("\n%d\n",lengthData)

	var count int = 0;
	for i:=0 ; i<lengthData ; i++ {
		var character= dataInFile[i]
		if(character == ' '  || character=='.'){
			count = count+1

		}
	}

	fmt.Printf("%d",count)
	newcount := strconv.Itoa(count)
	d1 := []byte(newcount)
	err1:= ioutil.WriteFile("output.txt",d1,0644)
    check(err1)

    fmt.Printf("\nSuccessfully written the data into the file.")

}