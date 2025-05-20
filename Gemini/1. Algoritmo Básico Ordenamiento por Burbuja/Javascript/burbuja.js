function bubbleSort(arr) {
  const n = arr.length;
  for (let i = 0; i < n - 1; i++) {
    for (let j = 0; j < n - i - 1; j++) {
      if (arr[j] > arr[j + 1]) {
        [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
      }
    }
  }
  return arr;
}

const entrada = prompt("Ingrese una lista de números enteros separados por espacios:");
const numeros = entrada.split(" ").map(Number);
const listaOrdenada = bubbleSort(numeros);
console.log("Lista ordenada:", listaOrdenada);