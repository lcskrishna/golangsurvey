/* Algorithm:
1) Take pivot element as first element and last element.
2) Scan the first element and last element.
3) Check for particular '[', '(','{' the end characters would be ']',')','}'
3) If any of the other not matches, print NO and break., else continue and finally print YES.
@Author: Chaitanya Sri Krishna Lolla.
*/

package main
import "fmt"
func main() {

    var givenString string ;
    fmt.Printf("Please enter the string for checking balanced paranthesis: \n")
    fmt.Scanf("%s\n",&givenString)

    var pivotStart int = 0
    var pivotEnd int = len(givenString)-1

    var isBalanced bool = true;
    var length int = len(givenString)

if(length%2==0){
    for pivotStart < pivotEnd{
        
        var startCharacter byte = givenString[pivotStart]
        var endCharacter byte = givenString[pivotEnd]
        fmt.Printf("%c\n",startCharacter)
        fmt.Printf("%c\n",endCharacter)

        if(startCharacter == '['){
            if(endCharacter !=']'){
                isBalanced = false;
                break;
            }
           
            
        }else if (startCharacter== '('){
            if(endCharacter!=')'){
                isBalanced = false;
                break;
            }

        }else if(startCharacter == '{'){
            if(endCharacter!='}'){
                isBalanced= false;
                break;
            }
        }else{
            isBalanced = false;
            break;
        }

        pivotStart++;
        pivotEnd--;
    }

        if(isBalanced){
            fmt.Printf("YES")
        }else{
            fmt.Printf("NO")
        }
}else{
    fmt.Printf("NO")
}

}