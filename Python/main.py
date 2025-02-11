import time

def mcd(a, b):
    while b:
        a, b = b, a % b
    return a

def mcm(a, b):
    return abs(a * b) // mcd(a, b)

def mcm_varios(*args):
    resultado = 1
    for num in args:
        resultado = mcm(resultado, num)
    return resultado

start_time = time.time()
for i in range(300000):
    mcm_varios(12321, 5674, 123, 821)
end_time = time.time()
print(f"{(end_time - start_time) * 1000}")
