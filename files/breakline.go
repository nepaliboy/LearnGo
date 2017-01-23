//breaks the line with command line delimiter
//writes to a new file
package main
import(
  "fmt"
  "strings"
  "os"
  "bufio"
  "flag"
)
func main (){

usage()
parse_string(os.Args[1],os.Args[2],os.Args[3])
}

func parse_string(filename string,deli string,filenameo string){
fp,err:=os.Open(filename)
if err!=nil{
  fmt.Printf("error reading file %v",err)
  os.Exit(1)
}
defer fp.Close()
//creating a new file to write
fp2,err:=os.Create(filenameo)
if err!=nil{
  fmt.Printf("error creating afile %v",err)
}

defer fp2.Close()
//using scanner to read file pointer
scanner:=bufio.NewScanner(fp)
for scanner.Scan(){
  line:=scanner.Text()
//split the string by delimeter and return string[]
//fmt.Printf("%v",scanner)
 x:=strings.Split(line,deli)
fmt.Printf("%s",x)
str1:=" "
 for _,str:=range x{
   str1=str1+" "+str
 }
 fmt.Printf("%s",str1)
 fmt.Fprintf(fp2,"%v\n",str1)
  }
}
//usage
func usage(){
options:=flag.String("d",";","default value")
flag.Parse()
if len(os.Args)<4{
  flag.Usage()
  os.Exit(1)
}
  if *options==";"{
      fmt.Printf("delimiter is ;")

}
}
