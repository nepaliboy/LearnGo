package main
import(
  "fmt"
  "os"
  "flag"
  "bufio"
)
func main(){
usage()
}
func usage(){
option:=flag.String("a","w","default word")
flag.Parse()
//output from wc
w,l:=wc(*option)
//printing word, line or both
if *option=="w" {
  fmt.Printf("# of words %d \n",w)
} else if  *option=="l"{
  fmt.Printf("# of lines %d \n ",l)
 } else if *option=="wl" { //need help
  fmt.Printf("# of word :lines: \n %d :%d ", w,l)
 } else {
   flag.Usage()
   os.Exit(1)
 }
}
func wc(a string)(countw int, countl int){
for _,filename:= range flag.Args() {
fmt.Printf("Reading %v ......\n",filename)
fp,err:=os.Open(filename)
if err!=nil{
  fmt.Printf("error reading file",err)
  os.Exit(1)
}
defer fp.Close()
scanner:=bufio.NewScanner(fp)
switch a {
case "w":
  scanner.Split(bufio.ScanWords)
 word:=0
 for scanner.Scan(){
  word++
  }
return word,0
case "l":
  line:=0
  for scanner.Scan(){
  line++
  }
  return 0,line
case "wl":
  //need help
  line,word:=0,0
  return line,word
}
}
return 0,0
}
