public class MCM {
    public static int mcd(int a, int b) {
        while (b != 0) {
            int temp = b;
            b = a % b;
            a = temp;
        }
        return a;
    }

    public static int mcm(int a, int b) {
        return Math.abs(a * b) / mcd(a, b);
    }

    public static int mcmVarios(int... args) {
        int resultado = 1;
        for (int num : args) {
            resultado = mcm(resultado, num);
        }
        return resultado;
    }

    public static void main(String[] args) {
        long startTime = System.nanoTime();
        mcmVarios(12321, 5674, 123, 821);
        long endTime = System.nanoTime();
        System.out.println((endTime - startTime) / 1_000_000.0);
    }
}
