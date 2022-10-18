package student


import "fmt"


type Student struct{
	firstName string
	lastname string
	gender string
	Date
	email string
	department string
	matricNumber string
	
	
}
 
//Set and get method for firstname and lastname
func (s *Student) SetName(first string, last string)error{
	 //merged both firstname and lastname togther for simplicity 
	if first == "" && last == ""{
		return fmt.Errorf("error occured! you haven't entered a name") // setting an error message if user input no value
	}
	s.firstName = first // assigning firstname to name2 variable
	s.lastname = last // assigning lastname to name variable
	return nil
}
func (s *Student)Name1()string{
	return s.firstName //returns firstname 
}
func (s *Student)Name()string{
	return s.lastname //returns lastname
}

// repeated the above error check method in the following lines of code
//Set and get method for gender
func (s *Student) SetGender(gender string)error{
	if gender == ""{
		return fmt.Errorf("error occured! you haven't entered a name")
	}
	s.gender = gender
	return nil
}
func (s *Student)Gender()string{
	return s.gender
}

//Set and get method for Email
func (s *Student) SetEmail(email string)error{
	if email == ""{
		return fmt.Errorf("error occured! you haven't entered an email")
	}
	s.email = email
	return nil
}
func (s *Student)Email()string{
	return s.email
}

//Set and get method for department
func (s *Student) SetDept(department string)error{

	if department == ""{
		return fmt.Errorf("error occured! you haven't entered an Department")
	}
	s.department = department
	return nil
}
func (s *Student)Dept()string{
	return s.department
}

//Set and get method for MatricNumber
func (s *Student) SetRegNo(id string)error{
	if id == ""{
		return fmt.Errorf("error occured! you haven't entered an email")
	}
	s.matricNumber = id
	return nil
}
func (s *Student)RegNo()string{
	return s.matricNumber
}

