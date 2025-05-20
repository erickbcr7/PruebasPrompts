const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function bubbleSort(arr) {
    let n = arr.length;
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < n - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
            }
        }
    }
    return arr;
}

rl.question('Ingrese números separados por espacios: ', (input) => {
    const numeros = input.trim().split(/\s+/).map(Number);
    console.log('Lista ordenada:', bubbleSort(numeros));
    rl.close();
});
