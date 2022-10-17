package main

import (
	"fmt"
	"log"
	"os"
	g "project4/input"
	"project4/student"
	"strings"
	"text/tabwriter"
)
func main(){
    // Number of students is determined
	fmt.Print("Enter the number of Students to be registered: ")
    num, err := g.Student(g.InputUser2())
	if err != nil{
		log.Fatal("An error has occured: ", err)
	}
	
//created a hash map database to hold all informations
	registeredStudent := map[string]student.Student{}

	value := student.Student{} // created a student.student value type 
	var name string // created a string value type to hold the keys for hashmap

	//created a for loop to accept multiple data
	for{
	fmt.Print("Enter FirstName: ") 
	_, err := fmt.Scanln(&name)
	if err != nil{
		log.Fatal(err)
	}
	name  = strings.TrimSpace(name)


	// firstname and lastname
	fmt.Print("Enter LastName: ")
	err = value.SetName(g.InputUser(), name)
    if err != nil{
		log.Fatal(err)
	}
    //gender
	fmt.Print("Enter Student gender: ")
	err = value.SetGender(g.InputUser())
    if err != nil{
	log.Fatal(err)
	}


	//date of birth
	fmt.Println("Enter your date of birth")
	fmt.Print("Day: ")
	err = value.SetDay(g.InputUser2())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Print("Month: ")
	err = value.SetMonth(g.InputUser2())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Print("Year: ")
	err = value.SetYear(g.InputUser2())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%v-%v-%v\n",value.Day(), value.Month(), value.Year())
	

	//email
	fmt.Print("Enter a Valid Email Address: ")
	err = value.SetEmail(g.InputUser())
    if err != nil{
		log.Fatal(err)
	}
	
	//department
	fmt.Print("Enter Student Department: ")
	err = value.SetDept(g.InputUser())
    if err != nil{
		log.Fatal(err)
	}

    // matric Number
	fmt.Print("Enter Student Matric Number: ")
	err = value.SetRegNo(g.InputUser())
    if err != nil{
		log.Fatal(err)
	}

	fmt.Println("")

	fmt.Println("Enter the details for another student.")

	registeredStudent[name] = value// each key equal a specofied value
	// fmt.Println(name, registeredStudent[name])
	if len(registeredStudent) == num{
		break
	}
	}
  


	//Prints all the data 
	var values student.Student
	// for _, values = range registeredStudent{
	// 	fmt.Println(values.Name1(), values.Name(), values.Gender(), values.Date, values.Email(), values.Dept(), values.RegNo())
	// }
    //println in a table format
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w,"FirstName\tSurname\tGender\tDateOfBirth\tEmail\tDepartment\tMatric-No\t")
	for _, values = range registeredStudent{
		fmt.Fprint(w, values.Name1(), "\t", values.Name(), "\t", values.Gender(), "\t", values.Date, "\t", values.Email(), "\t", values.Dept(), "\t", values.RegNo(), "\t")
	// fmt.Print("\n")
	}
	

	w.Flush()

	fmt.Println("")




	//Enabling two deletion of student
	fmt.Println("Would you like to delete a student name?")

	fmt.Print("Please enter yes or no: ")
	answer := g.InputUser()

	var deleteName string
	
	if answer == "yes"{
	fmt.Print("Enter a Student name: ")
	_, err = fmt.Scanln(&deleteName) 
	if err!= nil{
		log.Fatal(err)
	}
	// registeredStudent[deleteName] = value
	delete(registeredStudent, deleteName)
	}else if answer == "no"{
		fmt.Println("No data was deleted")
	}else{
		fmt.Println("Enter yes or no!")
	}



    fmt.Println("")


	
    fmt.Println("Do you want to delete another data?")
	fmt.Print("Please enter yes or no: ")
    answer2 := g.InputUser()
	var deleteName2 string
	if answer2 == "yes"{
		fmt.Print("Enter a Student name: ")
		_, err = fmt.Scanln(&deleteName2) 
		if err!= nil{
			log.Fatal(err)
		}
		// registeredStudent[deleteName2] = value
		delete(registeredStudent, deleteName2)
	}else if answer2 == "no"{
		fmt.Println("No data was deleted")
	}

	w = tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w,"FirstName\tSurname\tGender\tDateOfBirth\tEmail\tDepartment\tMatric-No")
	for _, values = range registeredStudent{
	fmt.Fprint(w, values.Name1(), "\t", values.Name(), "\t", values.Gender(), "\t", values.Date, "\t", values.Email(), "\t", values.Dept(), "\t", values.RegNo())
	}

	w.Flush()
	fmt.Println("")
	// for _, values = range registeredStudent{
	// 	fmt.Println(values.Name1(), values.Name(), values.Gender(), values.Date, values.Email(), values.Dept(), values.RegNo())
	// }

	fmt.Println("Exit")

}