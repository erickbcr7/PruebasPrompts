class Pila {
  constructor() {
    this.items = [];
  }

  push(item) {
    this.items.push(item);
  }

  pop() {
    if (!this.isEmpty()) {
      return this.items.pop();
    } else {
      return "La pila está vacía";
    }
  }

  isEmpty() {
    return this.items.length === 0;
  }
}

const pila = new Pila();

function interactuarPila() {
  const accion = prompt("Ingrese una acción (push <elemento>, pop, verificar, salir):").split(" ");
  const comando = accion[0].toLowerCase();

  if (comando === "push") {
    if (accion.length > 1) {
      const elemento = accion[1];
      pila.push(elemento);
      console.log(`Se agregó '${elemento}' a la pila.`);
    } else {
      console.log("Por favor, especifique el elemento a agregar.");
    }
  } else if (comando === "pop") {
    const elemento = pila.pop();
    console.log(`Se extrajo '${elemento}' de la pila.`);
  } else if (comando === "verificar") {
    if (pila.isEmpty()) {
      console.log("La pila está vacía.");
    } else {
      console.log("La pila no está vacía.");
    }
  } else if (comando === "salir") {
    return;
  } else {
    console.log("Comando no válido.");
  }
  interactuarPila();
}

interactuarPila();