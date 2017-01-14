package main
import (
  "fmt"
  "os"
 "strings"
)

//declaring x1 global variable
var x1 string

//main function
func main (){
usage()

//standardinput of string from command line and converting string array to string
input:=os.Args[1:]
sinput:=strings.Join(input," ")
fmt.Printf("%v \n",sinput)
fmt.Printf("%v \n",reverse(sinput))
}

//function to reverse the string
func reverse(s string)(rs string) {
//string is a sequence of runes - runes are fundamental data type, integer type
  a:=[]rune(s)
  for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1{
    a[i],a[j]=a[j],a[i]
  }
return string(a)
}

//will be implemented with flag in future
func usage(){
  if len(os.Args)<2  {
  fmt.Printf("exiting \n")
  fmt.Printf("\n USAGE:./reversestring abc  \n")
  os.Exit(1)
  }
}
