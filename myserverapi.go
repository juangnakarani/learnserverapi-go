package main

 import (
         "context"
         "fmt"
         "net/http"
 )
 var keyValue = struct{}{}
 
 func WithParamHandler(p string, h http.Handler) http.Handler {
         return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                 //v := r.URL.Query().Get(p)
                 v := "admin"
                 h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), keyValue, v)))
         })
 }
 
 func MustAdminHandler(h http.Handler) http.Handler {
         return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 
                 if v, ok := r.Context().Value(keyValue).(string); ok && v == "admin" {
                         h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), keyValue, v)))
                         return
                 }
 
                 http.Error(w, "You must have administrative permissions", http.StatusForbidden)
 
         })
 }
 func AdminPage(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("Yes you are admin"))
 }

 func Cars(w http.ResponseWriter, r *http.Request) {
         w.Header().Set("Content-Type", "application/json")

         jsonStr := `{"data":[{ "name":"Ford", "models":[ "Fiesta", "Focus", "Mustang" ] },
         { "name":"BMW", "models":[ "320", "X3", "X5" ] },
         { "name":"Fiat", "models":[ "500", "Panda" ] }]}`

         w.Write([]byte(jsonStr))
 }

 func main() {
         port := ":7070"
         //mux := http.NewServeMux()
         
         http.Handle("/car", http.HandlerFunc(Cars))
        //  mux.HandleFunc("/admin", AdminPage)
        http.Handle("/admin", WithParamHandler("token", MustAdminHandler(http.HandlerFunc(AdminPage))))
         fmt.Printf("server run on %s\n", port)
         fmt.Println("/car for test json data")
         fmt.Println("/admin for test admin page")
         http.ListenAndServe(port, nil)
        
 }