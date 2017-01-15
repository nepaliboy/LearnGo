package main
import (
"fmt"
"os"
"bufio"
)

func main(){
usage()
filename:=string(os.Args[1])
fmt.Printf("filename:%s %d \n",filename,count_word(filename))
}

//function to count words
func count_word(filename string)(count int){
fp,err:=os.Open(filename)
if err!=nil{
  fmt.Printf("error opening file %v", err)
  os.Exit(1)
}
scanner:=bufio.NewScanner(fp)
//setting the split function for scanning operations by words
//default is line
scanner.Split(bufio.ScanWords)
counter:=0
for scanner.Scan() {
  counter++
  }
return counter
}
//function usage, will be upgraded to flags
func usage(){
  if len(os.Args)<2 {
    fmt.Printf("USAGE: %v <filename> \n",os.Args[0])
    os.Exit(1)
  }
}
