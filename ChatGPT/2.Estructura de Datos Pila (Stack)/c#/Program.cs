using System;
using System.Collections.Generic;

class Stack {
    private List<int> items = new List<int>();
    public void Push(int item) => items.Add(item);
    public int? Pop() {
        if (IsEmpty()) return null;
        int val = items[items.Count - 1];
        items.RemoveAt(items.Count - 1);
        return val;
    }
    public bool IsEmpty() => items.Count == 0;
    public List<int> Items => items;
}

class Program {
    static void Main() {
        var s = new Stack();
        while (true) {
            Console.WriteLine("\nOpciones: 1) Push 2) Pop 3) Ver pila 4) Salir");
            var op = Console.ReadLine();
            if (op == "1") {
                Console.Write("Valor a apilar: ");
                int v = int.Parse(Console.ReadLine());
                s.Push(v);
                Console.WriteLine("Valor apilado.");
            } else if (op == "2") {
                Console.WriteLine("Valor desapilado: " + s.Pop());
            } else if (op == "3") {
                Console.WriteLine("Contenido pila: " + string.Join(",", s.Items));
            } else if (op == "4") {
                break;
            }
        }
    }
}
