package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("./hello.txt")
	fmt.Fprintf(w, string(file))
}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http serve faild  err:%v\n", err)
		return
	}
}
