package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func renderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/index.html")
}

func getEmployees(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	jsonResponse := getEmployeesFromDB()

	if jsonResponse == nil {
		returnErrorResponse(response, request, httpError)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func getEmployee(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	empID := mux.Vars(request)["id"]
	if empID == "" {
		httpError.Message = "Employee id can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		jsonResponse := getEmployeeFromDB(empID)
		if jsonResponse == nil {
			returnErrorResponse(response, request, httpError)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}
	}
}

func insertEmployee(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	var empDetails Employee
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&empDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if empDetails.Name == "" {
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if empDetails.Email == "" {
			httpError.Message = "Email  can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if empDetails.Phone == "" {
			httpError.Message = "Phone can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isInserted := insertEmployeeInDB(empDetails)
			if isInserted {
				getEmployees(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}

func deleteEmployee(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	empID := mux.Vars(request)["id"]
	if empID == "" {
		httpError.Message = "Employee id can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		isdeleted := deleteEmployeeFromDB(empID)
		if isdeleted {
			getEmployees(response, request)
		} else {
			returnErrorResponse(response, request, httpError)
		}
	}
}

func updateEmployee(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	var empDetails Employee
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&empDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if empDetails.Name == "" {
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if empDetails.ID == 0 {
			httpError.Message = "user Id can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if empDetails.Email == "" {
			httpError.Message = "Email can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if empDetails.Phone == "" {
			httpError.Message = "Phone can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isUpdated := updateEmployeeInDB(empDetails)
			if isUpdated {
				getEmployees(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage ErrorResponse) {
	httpResponse := &ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
