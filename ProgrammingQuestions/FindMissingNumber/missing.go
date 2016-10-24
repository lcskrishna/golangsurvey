package main

import "fmt"

func main() {

    var n int=0
    fmt.Printf("Please enter the size of an array:\n");
    fmt.Scanf("%d\n",&n)

    var arr[10] int
    var newarray[10] int
    var xorsum int =0;
    fmt.Println("Enter the array values:")
    for i := 0; i < n; i++ {     
        n, err := fmt.Scanf("%d\n", &arr[i])
        if err != nil {
            fmt.Println(n, err)
        }

        xorsum = xorsum ^ arr[i]
    }

    fmt.Println("Enter the array to find missing element:")
    var newxorsum int = 0
    for j:=0;j<n-1 ;j++{
        n,err:= fmt.Scanf("%d\n",&newarray[j])
        if err!= nil {
            fmt.Println(n,err)
        }

        newxorsum = newxorsum ^ newarray[j]
    }

    var missingElement int= xorsum ^ newxorsum;

    fmt.Printf("\nThe missing element is : %d",missingElement)
    
}