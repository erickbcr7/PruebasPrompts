class Stack
  def initialize
    @items = []
  end
  def push(item)
    @items.push(item)
  end
  def pop
    @items.empty? ? nil : @items.pop
  end
  def empty?
    @items.empty?
  end
  def items
    @items
  end
end

s = Stack.new
loop do
  puts "\nOpciones: 1) Push 2) Pop 3) Ver pila 4) Salir"
  op = gets.chomp
  case op
  when "1"
    print "Valor a apilar: "
    v = gets.chomp
    s.push(v)
    puts "Valor apilado."
  when "2"
    puts "Valor desapilado: #{s.pop}"
  when "3"
    puts "Contenido pila: #{s.items.inspect}"
  when "4"
    break
  end
end
