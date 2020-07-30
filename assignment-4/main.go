package main
import (
    "fmt"
    "student"
    
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

type Student = student.Student

func homePage(writer http.ResponseWriter, req * http.Request) {
    fmt.Fprintf(writer, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func JSONHeaders(handler http.HandlerFunc) http.HandlerFunc {
    return func(writer http.ResponseWriter, req *http.Request) {
        writer.Header().Set("Content-Type", "application/json")
        handler(writer, req)
    }
}

//============ Hnadling all routes ==================================
func handleRequests() {
    //============ creates a new instance of a mux router ===================
    Router:= mux.NewRouter().StrictSlash(true)
        // replace http.HandleFunc with myRouter.HandleFunc
    Router.HandleFunc("/", homePage)
    Router.HandleFunc("/addStudent", JSONHeaders(student.AddStudent)).Methods("POST")
    Router.HandleFunc("/updateStudent/{id}", JSONHeaders(student.UpdateStudent)).Methods("PUT")
    Router.HandleFunc("/deleteStudent/{id}", JSONHeaders(student.DeleteStudent)).Methods("DELETE")
    Router.HandleFunc("/allStudent", JSONHeaders(student.GetAllStudents)).Methods("GET")
    Router.HandleFunc("/getSingleStudent/{id}", JSONHeaders(student.GetStudentByID)).Methods("GET")

    log.Fatal(http.ListenAndServe(":10000", Router))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}