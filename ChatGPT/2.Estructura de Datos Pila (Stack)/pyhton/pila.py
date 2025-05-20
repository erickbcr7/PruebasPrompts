class Stack:
    def __init__(self):
        self.items = []
    def push(self, item):
        self.items.append(item)
    def pop(self):
        return self.items.pop() if not self.is_empty() else None
    def is_empty(self):
        return len(self.items) == 0

if __name__ == "__main__":
    s = Stack()
    while True:
        print("\nOpciones: 1) Push 2) Pop 3) Ver pila 4) Salir")
        op = input("Elija opción: ")
        if op == "1":
            valor = input("Valor a apilar: ")
            s.push(valor)
            print("Valor apilado.")
        elif op == "2":
            print("Valor desapilado:", s.pop())
        elif op == "3":
            print("Contenido pila:", s.items)
        elif op == "4":
            break
