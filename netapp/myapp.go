//sample program a) simple web server b)system programming - process status
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", myreply)
	http.HandleFunc("/mystock", mystock)
	log.Printf("\n starting:Brahmanda \n")
	log.Printf("Ready To Serve \n ")
	http.ListenAndServe("0.0.0.0:8081", nil)

}

//function myreply
func myreply(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App Server Name: brahmanda")
	//implementing system programming to capture process status
	out, err := exec.Command("ps").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, "Starting: Brahmanda")
	fmt.Fprintf(w, "\n \n")
	fmt.Fprintf(w, string(out))
	fmt.Fprintf(w, "\n \n")
}

//function mystock
func mystock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "your stock worth 50")
}
