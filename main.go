package main

import (
	"fmt"
	"log"
	"net/http"
)


func formhandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!=nil {
		fmt.Fprintf(w,"ParseForm() err :%v ",err);
		return;
	}
	fmt.Fprintf(w,"POST request successful")
	name:=r.FormValue("name");
	address:=r.FormValue("address");
	fmt.Fprintf(w,"name =%s\n",name);
	fmt.Fprintf(w,"address =%s\n",address)

}
/*
 * helloHanlder() handles request for path /hello
*/

func hellohandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello"{
		http.Error(w,"404 not Found",http.StatusNotFound)
		return ;
	}
	if r.Method != "GET"{
		http.Error(w,"method not supported",http.StatusNotFound)
		return;
	}
	fmt.Fprintf(w,"hello");

}
func main(){
	fileServer := http.FileServer(http.Dir("./static"));
	/*
	 * By default serves document index.html
     */
	http.Handle("/",fileServer);
	http.HandleFunc("/form",formhandler);
	http.HandleFunc("/hello",hellohandler);
	fmt.Printf("starting serever at port 8080\n");
	if err:= http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err);
	}
}