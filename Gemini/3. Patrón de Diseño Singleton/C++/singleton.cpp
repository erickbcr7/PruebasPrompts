#include <iostream>
#include <string>
#include <mutex>

class Singleton {
private:
    static Singleton* instance;
    static std::mutex mutex;
    std::string valor;

    Singleton(const std::string& valor) : valor(valor) {}

public:
    Singleton(const Singleton&) = delete;
    Singleton& operator=(const Singleton&) = delete;

    static Singleton* getInstance(const std::string& valorInicial = "") {
        std::lock_guard<std::mutex> lock(mutex);
        if (instance == nullptr) {
            instance = new Singleton(valorInicial);
        }
        return instance;
    }

    std::string getValor() const {
        return valor;
    }

    void setValor(const std::string& nuevoValor) {
        valor = nuevoValor;
    }
};

Singleton* Singleton::instance = nullptr;
std::mutex Singleton::mutex;

int main() {
    Singleton* instancia1 = Singleton::getInstance("Primer valor");
    std::cout << "Valor de instancia1: " << instancia1->getValor() << std::endl;

    Singleton* instancia2 = Singleton::getInstance("Segundo valor");
    std::cout << "Valor de instancia2: " << instancia2->getValor() << std::endl;
    std::cout << "¿Instancia1 es la misma que instancia2? " << (instancia1 == instancia2) << std::endl;
    std::cout << "Valor de instancia1 después de obtener instancia2: " << instancia1->getValor() << std::endl;

    instancia2->setValor("Valor modificado");
    std::cout << "Valor de instancia1 después de modificar instancia2: " << instancia1->getValor() << std::endl;

    // No es necesario eliminar la instancia aquí en un ejemplo simple,
    // pero en aplicaciones más complejas se debe manejar su ciclo de vida.
    return 0;
}