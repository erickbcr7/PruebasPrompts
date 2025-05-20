class Pila
  def initialize
    @items = []
  end

  def push(item)
    @items << item
  end

  def pop
    if !is_empty?
      @items.pop
    else
      "La pila está vacía"
    end
  end

  def is_empty?
    @items.empty?
  end
end

pila = Pila.new

loop do
  print "Ingrese una acción (push <elemento>, pop, verificar, salir): "
  accion = gets.chomp.split
  comando = accion[0].downcase

  case comando
  when "push"
    if accion.length > 1
      elemento = accion[1]
      pila.push(elemento)
      puts "Se agregó '#{elemento}' a la pila."
    else
      puts "Por favor, especifique el elemento a agregar."
    end
  when "pop"
    elemento = pila.pop
    puts "Se extrajo '#{elemento}' de la pila."
  when "verificar"
    if pila.is_empty?
      puts "La pila está vacía
      ."
    else
      puts "La pila no está vacía."
    end
  when "salir"
    break
  else
    puts "Comando no válido."
  end
en