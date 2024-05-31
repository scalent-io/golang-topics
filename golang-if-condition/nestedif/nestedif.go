package main

import "fmt"

func main() {

	customerID := 30202405
	customerIncomeType := "SALARIED"

	//nested if used
	if customerIncomeType == "SALARIED" {
		currentBalance := fetchCurrentAccountBalance()

		if currentBalance > 20000 {
			fmt.Println("customer is eligible to avail loan and credit card schemes")
			callCustomerOnMobileNumber(customerID)
		}
	}

}

func fetchCurrentAccountBalance() float64 {
	return 24000.00
}

func callCustomerOnMobileNumber(ID int) {
	fmt.Printf("Call customer with ID %v for marketing \n", ID)
}
