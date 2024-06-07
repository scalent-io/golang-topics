package main

import "fmt"

func main() {

	var ageArray []int
	var employeeNameArray []string

	employeesInformationStruct := []struct {
		Name        string
		Age         int
		CompanyName string
	}{
		{Name: "ScalentEmployee1", Age: 25, CompanyName: "ScalentInfotechPrivateLimited"},
		{Name: "ScalentEmployee2", Age: 24, CompanyName: "ScalentInfotechPrivateLimited"},
		{Name: "ScalentEmployee3", Age: 29, CompanyName: "ScalentInfotechPrivateLimited"},
	}

	for _, employeInformation := range employeesInformationStruct {

		//Basic if statement usage
		if employeInformation.Name != "" {
			employeeNameArray = append(employeeNameArray, employeInformation.Name)
		}
		if employeInformation.Age != 0 {
			ageArray = append(ageArray, employeInformation.Age)
		}
	}

	fmt.Println(employeeNameArray)
	fmt.Println(ageArray)

}
