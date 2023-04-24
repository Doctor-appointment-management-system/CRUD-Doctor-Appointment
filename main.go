package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Doctor struct {
	ID                               int
	Name                             string
	Gender                           string
	Address                          string
	City                             string
	Phone                            string
	Specialisation                   string
	Opening_time                     string
	Closing_time                     string
	Availability_time                string
	Availability                     string
	Available_for_home_visit         string
	Available_for_online_consultancy string
	Fees                             int
}

func Err(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func dbCreation() {

	//connecting to mysql

	db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/")
	Err(err)
	defer db.Close()

	// database creation

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS dasNew")
	Err(err)
}

func db_connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/dasNew")

	if err != nil {
		return nil, err
	}
	return db, nil
}

func sql_Doctor_tabel_creation() {
	db, err := db_connection()
	Err(err)
	// sql table creation

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Doctor(ID INT NOT NULL AUTO_INCREMENT, Name VARCHAR(30),Gender VARCHAR(10),Address VARCHAR(50), City VARCHAR(20),Phone VARCHAR(15),Specialisation VARCHAR(20),Opening_time VARCHAR(10),Closing_time VARCHAR(10),Availability_time VARCHAR(30),Availability VARCHAR(10),Available_for_home_visit VARCHAR(4),Available_for_online_consultancy VARCHAR(4),Fees INT ,PRIMARY KEY (ID) );")
	Err(err)
	fmt.Println("Docter Table Created")
}

func sql_Patient_tabel_creation() {
	db, err := db_connection()
	Err(err)
	// sql table creation

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS PATIENTS(ID INT,Name VARCHAR(20),Gender VARCHAR(8),Address VARCHAR(255),City VARCHAR(20),State VARCHAR(20),Mobile_no VARCHAR(15))")
	Err(err)
}

func Add_docter() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := db_connection()
		Err(err)
		var data Doctor
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf(`INSERT INTO Doctor (Name,Gender,Address,City,Phone,Specialisation,Opening_time,Closing_time,Availability_time,Availability,Available_for_home_visit,Available_for_online_consultancy,Fees) VALUES ( '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',%d)`, data.Name, data.Gender, data.Address, data.City, data.Phone, data.Specialisation, data.Opening_time, data.Closing_time, data.Availability_time, data.Availability, data.Available_for_home_visit, data.Available_for_online_consultancy, data.Fees)
		fmt.Println(query_data)
		//insert data
		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		c.JSON(http.StatusOK, gin.H{"message": "Doctor added successfully"})

	}
}

func Get_my_profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := db_connection()
		Err(err)
		var mob Doctor
		err = c.BindJSON(&mob)
		if err != nil {
			return
		}
		get_detail := fmt.Sprintf("SELECT * FROM Doctor WHERE Phone = '%s'", mob.Phone)
		detail, err := db.Query(get_detail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer detail.Close()

		var output interface{}
		for detail.Next() {

			var ID int
			var Name string
			var Gender string
			var Address string
			var City string
			var Phone string
			var Specialisation string
			var Opening_time string
			var Closing_time string
			var Availability_Time string
			var Availability string
			var Available_for_home_visit string
			var Available_for_online_consultancy string
			var Fees float64
			err = detail.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Opening_time, &Closing_time, &Availability_Time, &Availability, &Available_for_home_visit, &Available_for_online_consultancy, &Fees)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s' '%s'  '%s' '%s''%s' %f", ID, Name, Gender, Address, City, Phone, Specialisation, Opening_time, Closing_time, Availability_Time, Availability, Available_for_home_visit, Available_for_online_consultancy, Fees)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Doctor details": output})

		}

	}
}

func Update_docter() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := db_connection()
		Err(err)

		var data Doctor
		var updateColumns []string
		var args []interface{}

		err = c.BindJSON(&data)

		if err != nil {

			return

		}
		fmt.Println(data)
		if data.Address != "" {
			updateColumns = append(updateColumns, "Address = ?")
			args = append(args, data.Address)
		}

		if data.City != "" {
			updateColumns = append(updateColumns, "City = ?")
			args = append(args, data.City)
		}

		if data.Phone != "" {
			updateColumns = append(updateColumns, "Phone = ?")
			args = append(args, data.Phone)
		}

		if data.Specialisation != "" {
			updateColumns = append(updateColumns, "Specialisation = ?")
			args = append(args, data.Specialisation)
		}

		if data.Opening_time != "" {
			updateColumns = append(updateColumns, "Opening_time = ?")
			args = append(args, data.Opening_time)
		}

		if data.Closing_time != "" {
			updateColumns = append(updateColumns, "Closing_time = ?")
			args = append(args, data.Closing_time)
		}

		if data.Availability_time != "" {
			updateColumns = append(updateColumns, "Availability_time = ?")
			args = append(args, data.Availability_time)
		}

		if data.Availability != "" {
			updateColumns = append(updateColumns, "Availability = ?")
			args = append(args, data.Availability)
		}

		if data.Available_for_home_visit != "" {
			updateColumns = append(updateColumns, "Available_for_home_visit = ?")
			args = append(args, data.Available_for_home_visit)
		}

		if data.Available_for_online_consultancy != "" {
			updateColumns = append(updateColumns, "Available_for_online_consultancy = ?")
			args = append(args, data.Available_for_online_consultancy)
		}

		if data.Fees != 0 {
			updateColumns = append(updateColumns, "Fees = ?")
			args = append(args, data.Fees)
		}

		if len(updateColumns) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No update data provided"})
			return
		}

		updateQuery := "UPDATE Doctor SET " + strings.Join(updateColumns, ", ") + " WHERE id = ?"
		args = append(args, data.ID)
		fmt.Println(updateQuery)
		stmt, err := db.Prepare(updateQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()
		if _, err := stmt.Exec(args...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, data)

		c.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully"})

	}
}

func Delete_docter() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := db_connection()
		Err(err)

		var data Doctor

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		// _, err = db.Exec("DELETE FROM Dost WHERE id = 10")

		delete_query := fmt.Sprintf("DELETE FROM Doctor WHERE ID = %d", data.ID)

		delete, err := db.Query(delete_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Doctor Deleted successfully"})

	}
}

type Patients struct {
	ID                      int
	Name                    string
	Age                     int
	Gender                  string
	Address                 string
	City                    string
	Phone                   string
	Disease                 string
	Selected_specialisation string
	Patient_history         string
}

func Add_patient() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := db_connection()
		Err(err)
		defer db.Close()

		var patient Patients

		err = c.BindJSON(&patient)

		Err(err)

		c.IndentedJSON(http.StatusCreated, patient)

		sql_query := fmt.Sprintf(`INSERT INTO patient(Name,Age,Gender,Address,City,Phone,Disease,Selected_Specialisation,Patient_history) VALUES('%s',%d,'%s','%s','%s','%s','%s','%s','%s')`, patient.Name, patient.Age, patient.Gender, patient.Address, patient.City, patient.Phone, patient.Disease, patient.Selected_specialisation, patient.Patient_history)

		insert, err := db.Query(sql_query)
		Err(err)
		defer insert.Close()
		fmt.Println("Yeah! Patient is Added to the database")
	}
}

func Get_patient() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := db_connection()
		Err(err)

		query, err := db.Query("SELECT * FROM Patients")
		Err(err)
		defer query.Close()

		var output interface{}
		for query.Next() {
			var ID int
			var Name string
			var Age int
			var Gender string
			var Address string
			var City string
			var Phone string
			var Disease string
			var Selected_specialisation string
			var Patient_history string

			err = query.Scan(&ID, &Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Selected_specialisation, &Patient_history)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  %s  %d  %s  %s  %s  %s  %s  %s  %s ", ID, Name, Age, Gender, Address, City, Phone, Disease, Selected_specialisation, Patient_history)
			c.IndentedJSON(http.StatusOK, output)
		}
	}
}

func Update_patient() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := db_connection()
		Err(err)
		fmt.Println("database Connected Successfully")
		defer db.Close()
		var patient Patients
		err = c.BindJSON(&patient)
		Err(err)
		c.IndentedJSON(http.StatusCreated, patient)
		update_query := fmt.Sprintf("UPDATE Patients SET Name='%s',Age=%d,Gender='%s',Address='%s',City='%s',Phone='%s', Diseases='%s',Selected_specialisation='%s',Patient_history='%s', WHERE Id=%d", patient.Name, patient.Age, patient.Gender, patient.Address, patient.City, patient.Phone, patient.Disease, patient.Selected_specialisation, patient.Patient_history, patient.ID)
		fmt.Println(update_query)
		_, err = db.Exec(update_query)
		if err != nil {
			panic(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

func Delete_patient() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := db_connection()
		Err(err)

		var patient Patients
		err = c.BindJSON(&patient)

		query := fmt.Sprintf(" DELETE FROM PATIENTS WHERE Mobile_no='%s'", patient.Phone)
		delete, err := db.Query(query)
		Err(err)
		defer delete.Close()
		c.JSON(http.StatusOK, gin.H{"message": "Yeh! Patient is removed succesfully from database"})
	}
}

func main() {

	dbCreation()
	sql_Doctor_tabel_creation()
	sql_Patient_tabel_creation()
	router := gin.Default()
	router.POST("/", Add_docter())
	router.GET("/", Get_my_profile())
	router.PUT("/", Update_docter())
	router.DELETE("/", Delete_docter())

	router.POST("patient/add_patients", Add_patient())
	router.GET("patient/get_patients", Get_patient())
	router.PUT("patient/update_patients", Update_patient())
	router.DELETE("patient/delete_patients", Delete_patient())

	router.Run("localhost:8080")
}
