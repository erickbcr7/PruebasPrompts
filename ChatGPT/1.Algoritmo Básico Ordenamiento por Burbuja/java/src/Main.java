import java.util.*;
public class Main {
    public static int[] bubbleSort(int[] arr) {
        int n = arr.length;
        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < n - i - 1; j++) {
                if (arr[j] > arr[j + 1]) {
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                }
            }
        }
        return arr;
    }
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Ingrese números separados por espacios:");
        String linea = sc.nextLine();
        String[] partes = linea.trim().split("\\s+");
        int[] arr = Arrays.stream(partes).mapToInt(Integer::parseInt).toArray();
        System.out.println("Lista ordenada: " + Arrays.toString(bubbleSort(arr)));
    }
}
