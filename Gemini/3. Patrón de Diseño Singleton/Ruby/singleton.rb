require 'singleton'

class MiSingleton
  include Singleton
  attr_accessor :valor

  def initialize(valor=nil)
    @valor = valor
  end
end

# Prueba
instancia1 = MiSingleton.instance
instancia1.valor = "Primer valor"
puts "Valor de instancia1: #{instancia1.valor}"

instancia2 = MiSingleton.instance
instancia2.valor = "Segundo valor"
puts "Valor de instancia2: #{instancia2.valor}"
puts "¿Instancia1 es la misma que instancia2? #{instancia1.equal?(instancia2)}"
puts "Valor de instancia1 después de modificar instancia2: #{instancia1.valor}"