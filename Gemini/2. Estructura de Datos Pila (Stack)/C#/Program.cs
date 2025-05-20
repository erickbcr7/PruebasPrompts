using System;
using System.Collections.Generic;

public class Pila<T>
{
    private Stack<T> _stack = new Stack<T>();

    public void Push(T item)
    {
        _stack.Push(item);
    }

    public T Pop()
    {
        if (_stack.Count > 0)
        {
            return _stack.Pop();
        }
        else
        {
            return default(T); // O podrías lanzar una excepción
        }
    }

    public bool IsEmpty()
    {
        return _stack.Count == 0;
    }
}

public class PilaEjemplo
{
    public static void Main(string[] args)
    {
        Pila<string> pila = new Pila<string>();

        while (true)
        {
            Console.Write("Ingrese una acción (push <elemento>, pop, verificar, salir): ");
            string linea = Console.ReadLine();
            string[] accion = linea.Split(' ');
            string comando = accion[0].ToLower();

            switch (comando)
            {
                case "push":
                    if (accion.Length > 1)
                    {
                        string elemento = accion[1];
                        pila.Push(elemento);
                        Console.WriteLine($"Se agregó '{elemento}' a la pila.");
                    }
                    else
                    {
                        Console.WriteLine("Por favor, especifique el elemento a agregar.");
                    }
                    break;
                case "pop":
                    string elementoPop = pila.Pop();
                    if (elementoPop != null)
                    {
                        Console.WriteLine($"Se extrajo '{elementoPop}' de la pila.");
                    }
                    else
                    {
                        Console.WriteLine("La pila está vacía.");
                    }
                    break;
                case "verificar":
                    if (pila.IsEmpty())
                    {
                        Console.WriteLine("La pila está vacía.");
                    }
                    else
                    {
                        Console.WriteLine("La pila no está vacía.");
                    }
                    break;
                case "salir":
                    return;
                default:
                    Console.WriteLine("Comando no válido.");
                    break;
            }
        }
    }
}