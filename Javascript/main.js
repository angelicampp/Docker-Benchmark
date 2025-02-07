function mcd(a, b) {
  while (b !== 0) {
    [a, b] = [b, a % b];
  }
  return a;
}

function mcm(a, b) {
  return Math.abs(a * b) / mcd(a, b);
}

function mcmVarios(...args) {
  return args.reduce((acc, num) => mcm(acc, num), 1);
}

const startTime = performance.now();
mcmVarios(12321, 5674, 123, 821);
const endTime = performance.now();
console.log(`${endTime - startTime}ms`);
