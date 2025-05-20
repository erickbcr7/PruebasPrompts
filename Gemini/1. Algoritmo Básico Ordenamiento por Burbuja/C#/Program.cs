using System;
using System.Linq;

public class Burbuja
{
    public static void BubbleSort(int[] arr)
    {
        int n = arr.Length;
        for (int i = 0; i < n - 1; i++)
        {
            for (int j = 0; j < n - i - 1; j++)
            {
                if (arr[j] > arr[j + 1])
                {
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                }
            }
        }
    }

    public static void Main(string[] args)
    {
        Console.Write("Ingrese una lista de números enteros separados por espacios: ");
        string entrada = Console.ReadLine();
        int[] numeros = entrada.Split(' ').Select(int.Parse).ToArray();

        BubbleSort(numeros);
        Console.WriteLine("Lista ordenada: " + string.Join(", ", numeros));
    }
}