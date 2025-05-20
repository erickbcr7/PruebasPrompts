#include <iostream>
#include <vector>
#include <sstream>
#include <algorithm>

std::vector<int> bubbleSort(std::vector<int> arr) {
    int n = arr.size();
    for (int i = 0; i < n - 1; i++) {
        for (int j = 0; j < n - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                std::swap(arr[j], arr[j + 1]);
            }
        }
    }
    return arr;
}

int main() {
    std::cout << "Ingrese una lista de números enteros separados por espacios: ";
    std::string entrada;
    std::getline(std::cin, entrada);
    std::stringstream ss(entrada);
    int numero;
    std::vector<int> numeros;
    while (ss >> numero) {
        numeros.push_back(numero);
    }

    std::vector<int> listaOrdenada = bubbleSort(numeros);
    std::cout << "Lista ordenada: ";
    for (int num : listaOrdenada) {
        std::cout << num << " ";
    }
    std::cout << std::endl;
    return 0;
}