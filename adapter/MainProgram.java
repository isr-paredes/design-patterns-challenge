import java.util.HashMap;
import java.util.Map;

public class LegacyDataFetcher {

    public static Map<String, Object> getLegacyData() {
        Map<String, Object> legacyData = new HashMap<>();
        legacyData.put("employee_name", "John Doe");
        legacyData.put("employee_id", 12345);
        legacyData.put("department_code", "XYZ");
        legacyData.put("salary", 50000);
        return legacyData;
    }

    public static void main(String[] args) {
        Map<String, Object> data = getLegacyData();
        System.out.println(data);
    }
}
