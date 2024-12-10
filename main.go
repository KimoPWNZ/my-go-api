package main  

import (  
    "fmt"  
    "net/http"  
)  

func addHandler(w http.ResponseWriter, r *http.Request) {  
    a := r.URL.Query().Get("a")  
    b := r.URL.Query().Get("b")  
    fmt.Fprintf(w, "Addition: %s + %s = %s", a, b, a+b)  
}  

func main() {  
    http.HandleFunc("/add", addHandler)  
    http.ListenAndServe(":8080", nil)  
}
