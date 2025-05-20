class Singleton {
  constructor(valor) {
    if (!Singleton.instance) {
      Singleton.instance = this;
      this.valor = valor;
    }
    return Singleton.instance;
  }

  getValor() {
    return this.valor;
  }

  setValor(valor) {
    this.valor = valor;
  }
}

// Prueba
const instancia1 = new Singleton("Primer valor");
console.log("Valor de instancia1:", instancia1.getValor());

const instancia2 = new Singleton("Segundo valor");
console.log("Valor de instancia2:", instancia2.getValor());
console.log("¿Instancia1 es la misma que instancia2?", instancia1 === instancia2);
console.log("Valor de instancia1 después de obtener instancia2:", instancia1.getValor());

instancia2.setValor("Valor modificado");
console.log("Valor de instancia1 después de modificar instancia2:", instancia1.getValor());