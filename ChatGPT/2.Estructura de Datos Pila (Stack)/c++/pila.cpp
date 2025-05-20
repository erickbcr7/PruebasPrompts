#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Stack {
    vector<int> items;
public:
    void push(int item) { items.push_back(item); }
    int pop() {
        if (is_empty()) return -1;
        int val = items.back();
        items.pop_back();
        return val;
    }
    bool is_empty() { return items.empty(); }
    vector<int>& get_items() { return items; }
};

int main() {
    Stack s;
    string op;
    while (true) {
        cout << "\nOpciones: 1) Push 2) Pop 3) Ver pila 4) Salir\n";
        getline(cin, op);
        if (op == "1") {
            cout << "Valor a apilar: ";
            string input; getline(cin, input);
            int v = stoi(input);
            s.push(v);
            cout << "Valor apilado.\n";
        } else if (op == "2") {
            cout << "Valor desapilado: " << s.pop() << "\n";
        } else if (op == "3") {
            cout << "Contenido pila: ";
            for (int i : s.get_items()) cout << i << " ";
            cout << endl;
        } else if (op == "4") {
            break;
        }
    }
}
