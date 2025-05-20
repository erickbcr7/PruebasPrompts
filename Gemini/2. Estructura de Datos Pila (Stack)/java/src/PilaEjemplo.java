import java.util.Scanner;
import java.util.Stack;

public class PilaEjemplo {
    public static void main(String[] args) {
        Stack<String> pila = new Stack<>();
        Scanner scanner = new Scanner(System.in);

        while (true) {
            System.out.print("Ingrese una acción (push <elemento>, pop, verificar, salir): ");
            String linea = scanner.nextLine();
            String[] accion = linea.split(" ");
            String comando = accion[0].toLowerCase();

            switch (comando) {
                case "push":
                    if (accion.length > 1) {
                        String elemento = accion[1];
                        pila.push(elemento);
                        System.out.println("Se agregó '" + elemento + "' a la pila.");
                    } else {
                        System.out.println("Por favor, especifique el elemento a agregar.");
                    }
                    break;
                case "pop":
                    if (!pila.isEmpty()) {
                        String elemento = pila.pop();
                        System.out.println("Se extrajo '" + elemento + "' de la pila.");
                    } else {
                        System.out.println("La pila está vacía.");
                    }
                    break;
                case "verificar":
                    if (pila.isEmpty()) {
                        System.out.println("La pila está vacía.");
                    } else {
                        System.out.println("La pila no está vacía.");
                    }
                    break;
                case "salir":
                    scanner.close();
                    return;
                default:
                    System.out.println("Comando no válido.");
            }
        }
    }
}