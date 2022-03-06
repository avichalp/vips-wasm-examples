int add(int a[], int len) {
  int sum = 0;

  for (int i = 0; i < len; i++) {
    sum += a[i];
  }

  return sum;
}
