package main
import(
  "fmt"
  "os"
  "bufio"
)
func main(){
usage()
filename:=string(os.Args[1])
fmt.Printf("filename: %s lines:# %d \n",filename,count_line(filename))
}

func usage(){
if len(os.Args)<2 {
    fmt.Printf("USAGE %v <filename> \n",os.Args[0])
    os.Exit(1)
  }
}

func count_line(filename string)(count int){
fp,err:=os.Open(filename)
if err!=nil{
  fmt.Printf("error opening file %s",filename)
  os.Exit(1)
}
scanner:=bufio.NewScanner(fp)
counter:=0
for scanner.Scan(){
  counter++
}
defer fp.Close()
return counter
}
