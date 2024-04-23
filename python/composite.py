
from abc import ABC, abstractmethod

class Employee(ABC):
    @abstractmethod
    def get_name(self):
        pass

    @abstractmethod
    def get_salary(self):
        pass

    @abstractmethod
    def get_position(self):
        pass

    @abstractmethod
    def get_subordinates(self):
        pass

    @abstractmethod
    def add_subordinate(self, employee):
        pass

    @abstractmethod
    def remove_subordinate(self, employee):
        pass


class Department(Employee):
    def __init__(self, name):
        self._name = name
        self._subordinates = []

    def get_name(self):
        return self._name

    def get_salary(self):
        return sum(sub.get_salary() for sub in self._subordinates)

    def get_position(self):
        return f"Department: {self._name}"

    def get_subordinates(self):
        return self._subordinates

    def add_subordinate(self, employee):
        self._subordinates.append(employee)

    def remove_subordinate(self, employee):
        self._subordinates.remove(employee)

class EmployeeLeaf(Employee):
    def __init__(self, name, salary, position):
        self._name = name
        self._salary = salary
        self._position = position

    def get_name(self):
        return self._name

    def get_salary(self):
        return self._salary

    def get_position(self):
        return self._position

    def get_subordinates(self):
        return []

    def add_subordinate(self, employee):
        pass

    def remove_subordinate(self, employee):
        pass

def main():
    john = EmployeeLeaf("John Doe", 50000, "Software Engineer")
    jane = EmployeeLeaf("Jane Doe", 60000, "Product Manager")

    development_department = Department("Development")
    development_department.add_subordinate(john)
    development_department.add_subordinate(jane)

    marketing_department = Department("Marketing")

    company = Department("Company")
    company.add_subordinate(development_department)
    company.add_subordinate(marketing_department)

    print(company.get_name())
    for subordinate in company.get_subordinates():
        print(f" {subordinate.get_name()}")

if __name__ == "__main__":
    main()
