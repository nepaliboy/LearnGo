package main
import (
  "fmt"
  "os"
  "bufio"
)

func main(){
usage()
fp:=os.Args[1]
cat(string(fp))
}


func cat (filename string){

//opening a file by filepointer,, returns io.reader
  fp,err:=os.Open(filename)
  if err!=nil{
    fmt.Printf("error reading file",err)
    os.Exit(1)
  }
defer fp.Close()
//bufio reads from file io.reader
scanner:=bufio.NewScanner(fp)
//scans default by line
for scanner.Scan(){
  line:=scanner.Text()
  fmt.Println(line)
}
if scanner.Err()!=nil{
  fmt.Printf("error scanning")
}

}
//will be implemented with flag in future
func usage(){
  if len(os.Args)<2  {
  fmt.Printf("exiting invalid input \n")
  fmt.Printf("\n USAGE:./catv1 filename  \n")
  os.Exit(1)
  }
}
