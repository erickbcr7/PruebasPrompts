class Singleton:
    _instance = None

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super(Singleton, cls).__new__(cls, *args, **kwargs)
            # Inicialización de la instancia (opcional)
            cls._instance.valor = None
        return cls._instance

    def asignar_valor(self, valor):
        self.valor = valor

    def obtener_valor(self):
        return self.valor

# Prueba
instancia1 = Singleton()
instancia1.asignar_valor("Primer valor")
print(f"Valor de instancia1: {instancia1.obtener_valor()}")

instancia2 = Singleton()
instancia2.asignar_valor("Segundo valor")
print(f"Valor de instancia2: {instancia2.obtener_valor()}")
print(f"¿Instancia1 es la misma que instancia2? {instancia1 is instancia2}")
print(f"Valor de instancia1 después de modificar instancia2: {instancia1.obtener_valor()}")