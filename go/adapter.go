package main
import "fmt"

type NewFormat struct {
	Name       string
	ID         int
	Department string
	Pay        float64
}

func getLegacyData() map[string]interface{} {
	legacyData := map[string]interface{}{
		"employee_name": "John Doe",
		"employee_id":   12345,
		"department_code": "XYZ",
		"salary": 50000.0,
	}
	return legacyData
}

func adaptLegacyDataToNewFormat(legacyData map[string]interface{}) NewFormat {
	return NewFormat{
		Name:       legacyData["employee_name"].(string),
		ID:         legacyData["employee_id"].(int),
		Department: legacyData["department_code"].(string),
		Pay:        legacyData["salary"].(float64),
	}
}

func main() {
	legacyData := getLegacyData()
	newFormatData := adaptLegacyDataToNewFormat(legacyData)
	fmt.Printf("Nuevo formato de datos: %+v\n", newFormatData)

}