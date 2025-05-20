import java.util.Arrays;
import java.util.Scanner;

public class Burbuja {
    public static void bubbleSort(int[] arr) {
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
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Ingrese una lista de números enteros separados por espacios: ");
        String entrada = scanner.nextLine();
        String[] numerosStr = entrada.split(" ");
        int[] numeros = new int[numerosStr.length];
        for (int i = 0; i < numerosStr.length; i++) {
            numeros[i] = Integer.parseInt(numerosStr[i]);
        }

        bubbleSort(numeros);
        System.out.println("Lista ordenada: " + Arrays.toString(numeros));
        scanner.close();
    }
}