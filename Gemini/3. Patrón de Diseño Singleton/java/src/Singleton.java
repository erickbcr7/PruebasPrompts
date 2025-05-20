public class Singleton {
    private static Singleton instance;
    private String valor;

    private Singleton(String valor) {
        this.valor = valor;
    }

    public static Singleton getInstance(String valor) {
        if (instance == null) {
            instance = new Singleton(valor);
        }
        return instance;
    }

    public String getValor() {
        return valor;
    }

    public void setValor(String valor) {
        this.valor = valor;
    }

    public static void main(String[] args) {
        Singleton instancia1 = Singleton.getInstance("Primer valor");
        System.out.println("Valor de instancia1: " + instancia1.getValor());

        Singleton instancia2 = Singleton.getInstance("Segundo valor");
        System.out.println("Valor de instancia2: " + instancia2.getValor());
        System.out.println("¿Instancia1 es la misma que instancia2? " + (instancia1 == instancia2));
        System.out.println("Valor de instancia1 después de obtener instancia2: " + instancia1.getValor());

        instancia2.setValor("Valor modificado");
        System.out.println("Valor de instancia1 después de modificar instancia2: " + instancia1.getValor());
    }
}