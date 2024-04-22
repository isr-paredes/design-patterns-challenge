package main
import (
    "fmt"
    "strings"
)

// Component interface
type Employee interface {
    Display(depth int)
}

// Leaf class
type Individual struct {
    Name string
}

func (i *Individual) Display(depth int) {
    fmt.Printf("%s%s\n", strings.Repeat("-", depth), i.Name)
}

// Composite class
type Department struct {
    Name     string
    Employees []Employee
}

func (d *Department) Display(depth int) {
    fmt.Printf("%s%s\n", strings.Repeat("-", depth), d.Name)
    for _, employee := range d.Employees {
        employee.Display(depth + 2)
    }
}


func main() {
    // Hierarchy
    ceo := &Individual{Name: "CEO"}
    cto := &Individual{Name: "CTO"}
    cfo := &Individual{Name: "CFO"}

    marketing := &Department{Name: "Marketing", Employees: []Employee{ceo}}
    finance := &Department{Name: "Finance", Employees: []Employee{cto, cfo}}

    root := &Department{Name: "Company", Employees: []Employee{marketing, finance}}

    root.Display(0)
}