package main

import (
	"encoding/json"
	"fmt"
)

func getEmployeesFromDB() []byte {
	var (
		employee  Employee
		employees []Employee
	)
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for rows.Next() {
		rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone)
		employees = append(employees, employee)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(employees)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func getEmployeeFromDB(empID string) []byte {
	var (
		employee Employee
	)
	rows, err := db.Query("SELECT * FROM employee WHERE id=?", empID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for rows.Next() {
		rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(employee)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func insertEmployeeInDB(empDetails Employee) bool {
	stmt, err := db.Prepare("INSERT into employee SET Name=?,Email=?,Phone=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(empDetails.Name, empDetails.Email, empDetails.Phone)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func deleteEmployeeFromDB(empID string) bool {
	stmt, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(empID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func updateEmployeeInDB(empDetails Employee) bool {
	stmt, err := db.Prepare("UPDATE employee SET Name=?,Email=?,Phone=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(empDetails.Name, empDetails.Email, empDetails.Phone, empDetails.ID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}
