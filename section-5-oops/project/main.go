package main

import "fmt"

// Payable Processor

type Payable interface {
	fmt.Stringer           // We'll also rely on fmt.Stringer, so our types should implement String()
	CalculatePay() float64 // calculate monthly pay
}

type SalariedEmp struct {
	Name         string
	AnnualSalary float64
}

type HourlyEmp struct {
	Name        string
	HourRate    float64
	HoursWorked float64
}

type CommissionEmp struct {
	Name          string
	BaseSalary    float64
	CommisionRate float64
	SalesAmt      float64
}

func (e SalariedEmp) CalculatePay() float64 {
	return e.AnnualSalary / 12.0
}
func (he HourlyEmp) CalculatePay() float64 {
	return he.HourRate * he.HoursWorked
}

func (ce CommissionEmp) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommisionRate * ce.SalesAmt)
}

func (se SalariedEmp) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

func (he HourlyEmp) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", he.Name, he.HourRate, he.HoursWorked)
}

func (ce CommissionEmp) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.2f, CommRate: %.2f%%, Sales: $%.2f)",
		ce.Name, ce.BaseSalary, ce.CommisionRate*100, ce.SalesAmt)
}

func PrintEmpSummary[T fmt.Stringer](emp T) {
	fmt.Printf("  - Processing: %s\n", emp) // Relies on String() method
}

func ProcessPayroll(emp []Payable) {
	fmt.Println("--------------------------------")
	fmt.Println("Processing Payment")
	totalPayroll := 0.0
	for _, emp := range emp {
		PrintEmpSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("   Monthly Pay : %.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("Total Monthly Payroll is %.2f\n", totalPayroll)
	fmt.Println("--------------------------------")
}
func main() {
	fmt.Println("Welcome to Payroll Processor")
	ujjwal := SalariedEmp{
		Name:         "Ujjwal",
		AnnualSalary: 120000.00,
	}
	pro := HourlyEmp{
		Name:        "Pro",
		HourRate:    12.00,
		HoursWorked: 10.00,
	}
	baranwal := CommissionEmp{
		Name:          "Baranwal",
		CommisionRate: 10.00,
		BaseSalary:    11.00,
		SalesAmt:      12.00,
	}
	payrollList := []Payable{
		ujjwal, pro, baranwal,
	}
	ProcessPayroll(payrollList)

}
