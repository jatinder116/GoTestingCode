package main

import (
    "fmt"
    "encoding/json"
    "student"
    "github.com/gorilla/mux"
    //  "gopkg.in/validator.v2"
    "gopkg.in/go-playground/validator.v9"
    "log"
    "net/http"
    "io/ioutil"
    "strings"
    "strconv"
)

type Student = student.Student

var Students[] Student

// Globally declared validate paramater
var validate = validator.New()


func homePage(writer http.ResponseWriter, req * http.Request) {
    fmt.Fprintf(writer, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}


//============= Add Student =================================
func AddStudent(writer http.ResponseWriter, req * http.Request) {
    reqBody, _:= ioutil.ReadAll(req.Body)
    var student Student

    //======== unmarshel ==============================
    if err:= json.Unmarshal(reqBody, &student);
    err != nil {
        ErrorResponse(http.StatusUnprocessableEntity, "invalid json", writer)
        return
    }

    //======================== Vaidate Json Inputs ======================
	if ok, errors:= validateJsonParams( &student);
	!ok {
        ValidationResponse(errors, writer)
        return
    }

    Students = append(Students, student)

    response:= make(map[string] interface {})
    response["status"] = 1
    response["message"] = "Success"
    response["data"] = Students
    data, err:= json.Marshal(response)
    if err != nil {
            writer.WriteHeader(http.StatusInternalServerError)
            writer.Write([] byte("error"))
        }
        //=============== Send header, status code and output to writer==================================
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    writer.Write(data)
}



//===================== Update student ==============================
func UpdateStudent(writer http.ResponseWriter, req * http.Request) {
    checkStudentExist:= 0;
    params:= mux.Vars(req)
    key:= params["id"]
    Id,err:= strconv.Atoi(key)
    if err != nil {
        ErrorResponse(http.StatusUnprocessableEntity, "not converted", writer)
        return
    }

    //============= Loop for students to find student for update ================
    for i,_:= range Students {
        if Students[i].Id == Id {
            Students = append(Students[:i], Students[i + 1:]...)
            var student Student
            reqBody, _:= ioutil.ReadAll(req.Body)
                //======== unmarshel ==============================
            if err:= json.Unmarshal(reqBody, &student);
            err != nil {
                ErrorResponse(http.StatusUnprocessableEntity, "invalid json", writer)
                return
            }
            student.Id = Id
            Students = append(Students, student)

            response:= make(map[string] interface {})
            response["status"] = 1
            response["message"] = "Success"
            response["data"] = Students
            data, err:= json.Marshal(response)
            if err != nil {
                    writer.WriteHeader(http.StatusInternalServerError)
                    writer.Write([] byte("error"))
                }
                // ===============Send header, status code and output to writer==========================
            writer.Header().Set("Content-Type", "application/json")
            writer.WriteHeader(http.StatusOK)
            writer.Write(data)
            checkStudentExist = 1;
        }
    }
    if checkStudentExist == 0 {
        ErrorResponse(http.StatusUnprocessableEntity, "Student Not Found", writer)
        return
    }
}



//==================== Delete Student ========================================
func DeleteStudent(writer http.ResponseWriter, req * http.Request) {
    checkStudentExist:= 0;
    params:= mux.Vars(req)
    key:= params["id"]
    Id,err:= strconv.Atoi(key)
    if err != nil {
        // ==========Handle error for string converted to int or not============
        ErrorResponse(http.StatusUnprocessableEntity, "not converted", writer)
        return
    }

    //==========Loop for students to delete student===================
    for i,_:= range Students {
        if Students[i].Id == Id {
            fmt.Println(Students[i])
            Students = append(Students[:i], Students[i + 1:]...)
            response:= make(map[string] interface {})
            response["status"] = 1
            response["message"] = "Success"
            response["data"] = Students
            data, err:= json.Marshal(response)
            if err != nil {
                    writer.WriteHeader(http.StatusInternalServerError)
                    writer.Write([] byte("error"))
                }
                //================ Send header, status code and output to writer======================
            writer.Header().Set("Content-Type", "application/json")
            writer.WriteHeader(http.StatusOK)
            writer.Write(data)
            checkStudentExist = 1;
            return
        }
    }
    if checkStudentExist == 0 {
        ErrorResponse(http.StatusUnprocessableEntity, "Student Not Found", writer)
        return
    }
}


//================= Get All Student ===========================
func GetAllStudents(writer http.ResponseWriter, req * http.Request) {
    response:= make(map[string] interface {})
    response["status"] = 1
    response["message"] = "Success"
    response["data"] = Students
    data,err:= json.Marshal(response)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write([] byte("error"))
    }
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    writer.Write(data)
}



//=================== Get Student by Id ===============================
func GetStudentByID(writer http.ResponseWriter, req * http.Request) {
    checkStudentExist:= 0;
    params:= mux.Vars(req)
    key:= params["id"]
    Id,err:= strconv.Atoi(key)
    if err != nil {
        // ==========Handle error for string converted to int or not============
        ErrorResponse(http.StatusUnprocessableEntity, "not converted", writer)
        return
    }
    //==============Loop for students to get single student ===============================
    for i,student:= range Students {
        if Students[i].Id == Id {
            response:= make(map[string] interface {})
            response["status"] = 1
            response["message"] = "Success"
            response["data"] = student
            data,err:= json.Marshal(response)
            if err != nil {
                writer.WriteHeader(http.StatusInternalServerError)
                writer.Write([] byte("error"))
            }
            //================= Send header, status code and output to writer =======================================
            writer.Header().Set("Content-Type", "application/json")
            writer.WriteHeader(http.StatusOK)
            writer.Write(data)
            checkStudentExist = 1;
        }
    }
    if checkStudentExist == 0 {
        ErrorResponse(http.StatusUnprocessableEntity, "Student Not Found", writer)
        return
    }
}


//============ Error Response function ================================
func ErrorResponse(statusCode int, errorMessage string, writer http.ResponseWriter) {
    response:= make(map[string] interface {})
    response["status"] = 0
    response["message"] = errorMessage
    data,err:= json.Marshal(response)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write([] byte("error"))
    }
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(statusCode)
    writer.Write(data)
}

//=================== validation Response function ====================
func ValidationResponse(params map[string][] string, writer http.ResponseWriter) {
    response:= make(map[string] interface {})
    response["status"] = 0
    response["message"] = "validation error"
    response["errors"] = params
    data,err:= json.Marshal(response)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write([] byte("error"))
    }

    // ========================== Send header, status code and output to writer ===========================================
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusUnprocessableEntity)
    writer.Write(data)
}


//================= validate Json Params ==============================
func validateJsonParams(dataSet interface {})(bool, map[string][] string) {
    if err:= validate.Struct(dataSet);
    err != nil {
        errors:= make(map[string][] string)
        for _,err:= range err.(validator.ValidationErrors) {
            var name = strings.ToLower(err.StructField())
            switch err.Tag() {
                case "required":
                    errors[name] = append(errors[name], "The " + name + " is required")
                    break
                case "email":
                    errors[name] = append(errors[name], "The " + name + " should be a valid email")
                    break
                default:
                    errors[name] = append(errors[name], "The " + name + " is invalid")
                    break
            }
        }
        return false,
        errors
    }
    return true, nil
}



//============ Hnadling all routes ==================================
func handleRequests() {
    //============ creates a new instance of a mux router ===================
    Router:= mux.NewRouter().StrictSlash(true)
        // replace http.HandleFunc with myRouter.HandleFunc
    Router.HandleFunc("/", homePage)
    Router.HandleFunc("/addStudent", AddStudent).Methods("POST")
    Router.HandleFunc("/updateStudent/{id}", UpdateStudent).Methods("PUT")
    Router.HandleFunc("/deleteStudent/{id}", DeleteStudent).Methods("DELETE")
    Router.HandleFunc("/allStudent", GetAllStudents).Methods("GET")
    Router.HandleFunc("/getSingleStudent/{id}", GetStudentByID).Methods("GET")

    log.Fatal(http.ListenAndServe(":10000", Router))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")

    // ======Static student data ==============
    Students = [] Student {
        Student {
            Id: 1,
            Name: "Alicia",
            Email: "alicia@gmail.com",
            City: "USA",
            Gender: "Female",
        },
        Student {
            Id: 2,
            Name: "Mike",
            Email: "mike@gmail.com",
            City: "Canada",
            Gender: "Male",
        },
        Student {
            Id: 3,
            Name: "Milina",
            Email: "milina@gmail.com",
            City: "Russia",
            Gender: "Female",
        },
    }

    handleRequests()
}