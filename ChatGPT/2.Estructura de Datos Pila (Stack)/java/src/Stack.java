import java.util.ArrayList;
import java.util.Scanner;

public class Stack {
    private ArrayList<Integer> items = new ArrayList<>();
    public void push(int item) { items.add(item); }
    public Integer pop() {
        if (isEmpty()) return null;
        return items.remove(items.size() - 1);
    }
    public boolean isEmpty() { return items.isEmpty(); }
    public ArrayList<Integer> getItems() { return items; }

    public static void main(String[] args) {
        Stack s = new Stack();
        Scanner sc = new Scanner(System.in);
        while (true) {
            System.out.println("\nOpciones: 1) Push 2) Pop 3) Ver pila 4) Salir");
            String op = sc.nextLine();
            if (op.equals("1")) {
                System.out.print("Valor a apilar: ");
                int v = Integer.parseInt(sc.nextLine());
                s.push(v);
                System.out.println("Valor apilado.");
            } else if (op.equals("2")) {
                System.out.println("Valor desapilado: " + s.pop());
            } else if (op.equals("3")) {
                System.out.println("Contenido pila: " + s.getItems());
            } else if (op.equals("4")) {
                break;
            }
        }
    }
}

