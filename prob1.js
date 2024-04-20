// 1: Adapter Pattern
// Base function
function getLegacyData() {
    // Simulate fetching data from a legacy system
    const legacyData = {
      "employee_name": "John Doe",
      "employee_id": 12345,
      "department_code": "XYZ",
      "salary": 50000
    };
    return legacyData;
  }

//-----------------------------------

// Return Dep name given Dep code
function depCodeToName(code) {
    if (code === "XYZ") {
        return "DepName01";
    } else {
        return "DepNameXX";
    }
}  

// Adapter function
function legacyDataAdapter(legacyData) {
    return {
        "name": legacyData["employee_name"],
        "id": legacyData["employee_id"],
        "department": depCodeToName(legacyData["department_code"]),
        "pay": legacyData["salary"]
    };
}

const legacyData = getLegacyData();
const adaptedData = legacyDataAdapter(legacyData);
console.log(adaptedData);