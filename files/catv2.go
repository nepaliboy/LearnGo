//simple cat program that uses flag and bufio
//introduce flag <usage> ./catv1 -type=file <filename> ....<filenamen>
package main
import (
  "fmt"
  "flag"
  "os"
  "bufio"
)
func main(){
usage()
cat()
}

//cat function
func cat(){
  for i,filename:=range flag.Args(){
    fmt.Printf("Opening File %v, index %d \n",string(filename),i)
    fp,err:=os.Open(filename)
    if err!=nil{
    fmt.Printf("error opening file",err)
    os.Exit(1)
}

scanner:=bufio.NewScanner(fp)
  for scanner.Scan(){
    line:=scanner.Text()
    fmt.Printf("%s \n",line)
  }
  }
}

func usage(){
option:=flag.String("type","file","default is file input")
flag.Parse()
if *option=="file" && len(flag.Args())<1{
  fmt.Printf("no files provided \n")
  flag.Usage()
  }
}
