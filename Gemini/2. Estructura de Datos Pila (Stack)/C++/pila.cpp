#include <iostream>
#include <stack>
#include <string>

int main() {
    std::stack<std::string> pila;
    std::string comando;

    while (true) {
        std::cout << "Ingrese una acción (push <elemento>, pop, verificar, salir): ";
        std::cin >> comando;

        if (comando == "push") {
            std::string elemento;
            std::cin >> elemento;
            pila.push(elemento);
            std::cout << "Se agregó '" << elemento << "' a la pila." << std::endl;
        } else if (comando == "pop") {
            if (!pila.empty()) {
                std::cout << "Se extrajo '" << pila.top() << "' de la pila." << std::endl;
                pila.pop();
            } else {
                std::cout << "La pila está vacía." << std::endl;
            }
        } else if (comando == "verificar") {
            if (pila.empty()) {
                std::cout << "La pila está vacía." << std::endl;
            } else {
                std::cout << "La pila no está vacía." << std::endl;
            }
        } else if (comando == "salir") {
            break;
        } else {
            std::cout << "Comando no válido." << std::endl;
        }
    }

    return 0;
}