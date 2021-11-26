package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func SplitToString(a []int) string {
	if len(a) == 0 {
		return ""
	}

	var sep string
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}
func SumIntArray(a []int) int {
	if len(a) == 0 {
		return 0
	}

	var sep int
	for i := range a {
		sep += a[i]
	}
	return sep
}
func SumMatrix(a [][]string) int {
	var sep int
	for i, h := range a {
		for j, cell := range h {
			fmt.Print(cell, " ")
			new_int, err := strconv.Atoi(a[i][j])
			if err != nil {
				fmt.Sprintf("error %s", err.Error())
			}
			sep += new_int
		}

	}
	return sep
}

func MultiplyMatrix(a [][]string) int {
	sep := 1
	for i, h := range a {
		for j, cell := range h {
			fmt.Print(cell, " ")
			new_int, err := strconv.Atoi(a[i][j])
			if err != nil {
				fmt.Sprintf("error %s", err.Error())
			}
			sep *= new_int
		}

	}
	return sep
}

func read_from_file(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

func flatten_from_file(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var res string
	var response string
	var flattened_response string
	for _, row := range records {
		res = fmt.Sprintln(res, strings.Join(row, ","))
		response = strings.Replace(res, "\n", ",", -1)
		flattened_response = strings.Replace(response, " ", "", -1)
	}
	fmt.Fprint(w, flattened_response)
}
func sum_from_file(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response int
	response = SumMatrix(records)
	fmt.Fprint(w, response)
}

func multiply_from_file(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response int
	response = MultiplyMatrix(records)
	fmt.Fprint(w, response)
}
func invert_from_excel(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for i := 0; i < len(records); i++ {
		for j := 0; j < i; j++ {
			records[i][j], records[j][i] = records[j][i], records[i][j]
		}
	}
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(w, response)
}
func main() {
	http.HandleFunc("/echo", read_from_file)
	http.HandleFunc("/flatten", flatten_from_file)
	http.HandleFunc("/sum", sum_from_file)
	http.HandleFunc("/multiply", multiply_from_file)
	http.HandleFunc("/invert", invert_from_excel)
	http.ListenAndServe(":8080", nil)
}

func echo() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
		fmt.Fprint(w, response)
	})
	http.ListenAndServe(":8080", nil)
}
