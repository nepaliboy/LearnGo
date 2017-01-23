package main
import (
  "fmt"
  "bufio"
  "strings"
  "os"
)

func main(){
  usage()
  replaceword(os.Args[1],os.Args[2],os.Args[3],os.Args[4])
}

func usage(){
if len(os.Args)<4 {
  fmt.Printf("usage:%s <inputfile> <outputfile> <searchword>  <replaceword>",os.Args[0])
  os.Exit(1)
}
}

//function to replaceword
func replaceword(filename string ,fileoutput string, replacew string, newreplace string){
  fileinput,err:=os.Open(filename)
  if err!=nil{
    fmt.Printf("error reading input file %v",err)
    os.Exit(1)
  }
  defer fileinput.Close()
scanner:=bufio.NewScanner(fileinput)

//creating outputfile
fileo,err:=os.Create(fileoutput)
  if err!=nil{
    fmt.Printf("error creating output file %v",err)
    os.Exit(1)
  }

  for scanner.Scan(){
//converts to string
         line:=scanner.Text()
//replacing word -1 - all replacements
         replace:=strings.Replace(line,replacew,newreplace,-1)
//writing to a new file
      fmt.Fprintf(fileo,"%v \n",replace)
  }
  defer fileo.Close()
}
