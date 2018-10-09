package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Слушаю")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		fmt.Fprintf(writer,"Ответ из GO %v",request.Form)
	})
	http.ListenAndServe(":3000",nil)
}
