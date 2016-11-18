/* Algorithm:
1) For the given String initially consider the each character at once.
2) Take two pivot elements say 0 and Lenght of String -1
3) Loop over the values from l to r  and swap the corrosponding elements and form a string.
*/

package main

import "fmt"
func main() {
	
	var input string
	fmt.Printf("Please enter the string: ")
	fmt.Scanf("%s", &input)
	var array []byte
	array = make([]byte,len(input),len(input))
	for k:=0;k<len(input);k++ {
		array[k] = input[k]
	}
	permute(array,0, len(input)-1)
}

func permute(a []byte,l int,r int) {
	
   	if (l == r) {
     fmt.Printf("%s\n", a);
   	}   
   	if (l!=r) {
       for i:= l; i <= r; i++ {
       	  a[l], a[i] = a[i], a[l]
          permute(a, l+1, r)
          a[l], a[i] = a[i], a[l]
       }
   }
}