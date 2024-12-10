package main  

import (  
    "net/http"  
    "net/http/httptest"  
    "testing"  
)  

func TestAddHandler(t *testing.T) {  
    req, err := http.NewRequest("GET", "/add?a=2&b=3", nil)  
    if err != nil {  
        t.Fatal(err)  
    }  

    rr := httptest.NewRecorder()  
    handler := http.HandlerFunc(addHandler)  

    handler.ServeHTTP(rr, req)  

    if status := rr.Code; status != http.StatusOK {  
        t.Errorf("handler returned wrong status code: got %v want %v",  
            status, http.StatusOK)  
    }  
}
