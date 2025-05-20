def bubble_sort(arr)
  n = arr.length
  (0...n).each do |i|
    (0...(n-i-1)).each do |j|
      if arr[j] > arr[j+1]
        arr[j], arr[j+1] = arr[j+1], arr[j]
      end
    end
  end
  arr
end

puts "Ingrese números separados por espacios:"
entrada = gets.chomp
numeros = entrada.split.map(&:to_i)
puts "Lista ordenada: #{bubble_sort(numeros)}"
