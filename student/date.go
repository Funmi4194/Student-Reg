package student

import (
	"fmt"
	"time"
)


type Date struct{
	year int
	month int
	day int 
}

//creating a setter method to set fields values and store
//in the getter method
func (d *Date) SetYear(year int)error {
	var now time.Time = time.Now()
	currentYear := now.Year()
	if year < 1 ||  year > currentYear{ // checking to see if user inputs invalid year 
		return fmt.Errorf("invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int)error{
	if month < 1 || month > 12{ // checking to see if user inputs invalid month
        return fmt.Errorf("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int)error{
	if day < 1 || day > 31{ // checking to see if user inputs invalid day
		return fmt.Errorf("invalid day")
	}
	d.day = day
	return nil
}



//creating a getter method to get values from get method

/* the return value for each methods declared above is to 
be able to return individual day,month and year*/
func (d *Date)Day()int{ 
	return d.day
}
func (d *Date)Month()int{
	return d.month
}

func (d *Date) Year()int{
	return d.year
}




