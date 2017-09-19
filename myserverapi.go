package main

 import (
         "fmt"
         "net/http"
 )

 func Cars(w http.ResponseWriter, r *http.Request) {
         w.Header().Set("Content-Type", "application/json")

         jsonStr := `{"data":[{ "name":"Ford", "models":[ "Fiesta", "Focus", "Mustang" ] },
         { "name":"BMW", "models":[ "320", "X3", "X5" ] },
         { "name":"Fiat", "models":[ "500", "Panda" ] }]}`

         w.Write([]byte(jsonStr))
 }

 func main() {

         mux := http.NewServeMux()
         
         mux.HandleFunc("/car", Cars)
         fmt.Println("server run on 7070")
         fmt.Println("/car for test json data")
         http.ListenAndServe(":7070", mux)
        
 }