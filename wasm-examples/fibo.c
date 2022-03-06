// To compare V8 vs wasm number crunching
int fibo(int n, int a, int b) {
  return n < 1 ? a: fibo(n - 1, a + b, a);
}

int fibonacci(int n) {
  return fibo(n, 0, 1);
}
