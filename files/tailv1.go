package main
import(
  "fmt"
  "flag"
  "bufio"
  "os"
)


func main(){
opt:=usage()

//parsing n files
for _,filename:= range flag.Args(){
tail(filename,opt)
}
}

//tail function - would be able to tail n files as arguments
func tail(filename string,lines int){
fp,err:=os.Open(filename)
if err!=nil{
  fmt.Printf("error reading file %v",err)
  os.Exit(1)
}
defer fp.Close()
//counting total lines
scanner:=bufio.NewScanner(fp)
  counter:=0
  for scanner.Scan(){
    counter++
  }
//calculating last n lines from opt
line:=counter-lines
fp1,err:=os.Open(filename)
if err!=nil{
  fmt.Printf("error reading file")
  os.Exit(1)
}
scanner1:=bufio.NewScanner(fp1)
counter1:=0
//printing last n lines
for scanner1.Scan(){
  counter1++
  text:=scanner1.Text()
  if counter1>line {
    fmt.Printf("%d \t",counter1)
    fmt.Printf("%s \n",text)
  }
}
defer fp1.Close()

}
//usage function
func usage()(x int){
  option:=flag.Int("l",10,"default printing last ten lines\n")
 flag.Parse()
  if *option==10 {
      flag.Usage()

  }
  return *option
}
