class Stack {
    constructor() { this.items = []; }
    push(item) { this.items.push(item); }
    pop() { return this.items.length === 0 ? null : this.items.pop(); }
    isEmpty() { return this.items.length === 0; }
}

const readline = require('readline').createInterface({
    input: process.stdin, output: process.stdout
});
const s = new Stack();

function menu() {
    console.log("\nOpciones: 1) Push 2) Pop 3) Ver pila 4) Salir");
    readline.question("Elija opción: ", op => {
        if (op === "1") {
            readline.question("Valor a apilar: ", v => {
                s.push(v);
                console.log("Valor apilado.");
                menu();
            });
        } else if (op === "2") {
            console.log("Valor desapilado:", s.pop());
            menu();
        } else if (op === "3") {
            console.log("Contenido pila:", s.items);
            menu();
        } else if (op === "4") {
            readline.close();
        } else {
            menu();
        }
    });
}
menu();
