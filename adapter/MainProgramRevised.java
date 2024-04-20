public class LegacyData {
  private String employee_name;
  private int employee_id;
  private String department_code;
  private double salary;

  public LegacyData(String employee_name, int employee_id, String department_code, double salary) {
    this.employee_name = employee_name;
    this.employee_id = employee_id;
    this.department_code = department_code;
    this.salary = salary;
  }
  public String getEmployee_name() {
    return employee_name;
  }

  public void setEmployee_name(String employee_name) {
    this.employee_name = employee_name;
  }

  public int getEmployee_id() {
    return employee_id;
  }

  public void setEmployee_id(int employee_id) {
    this.employee_id = employee_id;
  }

  public String getDepartment_code() {
    return department_code;
  }

  public void setDepartment_code(String department_code) {
    this.department_code = department_code;
  }

  public double getSalary() {
    return salary;
  }

  public void setSalary(double salary) {
    this.salary = salary;
  }
}



