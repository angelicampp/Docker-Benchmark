def mcd(a, b)
    while b != 0
        a, b = b, a % b
    end
    a
end

def mcm(a, b)
    (a * b).abs / mcd(a, b)
end

def mcm_varios(*args)
    args.reduce(1) { |acc, num| mcm(acc, num) }
end

start_time = Time.now
for i in 0...300000
  mcm_varios(12321, 5674, 123, 821)
end
end_time = Time.now
puts "#{(end_time - start_time) * 1000}"
