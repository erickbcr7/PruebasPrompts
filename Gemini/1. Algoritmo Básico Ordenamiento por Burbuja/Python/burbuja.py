def bubble_sort(lista):
    n = len(lista)
    for i in range(n - 1):
        for j in range(0, n - i - 1):
            if lista[j] > lista[j + 1]:
                lista[j], lista[j + 1] = lista[j + 1], lista[j]
    return lista

entrada = input("Ingrese una lista de números enteros separados por espacios: ")
numeros = [int(x) for x in entrada.split()]
lista_ordenada = bubble_sort(numeros)
print("Lista ordenada:", lista_ordenada)