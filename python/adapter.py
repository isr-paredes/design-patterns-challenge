def getLegacyData():
    legacyData = {
        "employee_name": "John Doe",
        "employee_id": 12345,
        "department_code": "XYZ",
        "salary": 50000
    }
    return legacyData

class LegacyToNewFormatAdapter:
    def __init__(self, legacy_data):
        self.legacy_data = legacy_data

    def convert(self):
        return {
            "name": self.legacy_data["employee_name"],
            "id": self.legacy_data["employee_id"],
            "department": self.legacy_data["department_code"],
            "pay": self.legacy_data["salary"]
        }

def main():
    legacy_data = getLegacyData()
    adapter = LegacyToNewFormatAdapter(legacy_data)
    new_format_data = adapter.convert()
    print(new_format_data)


main()
