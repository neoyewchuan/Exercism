class CollatzConjecture {
  int steps(int num) {
    int numSteps = 0;
    if (num > 0) {
      while (num != 1) {
        if (num % 2 == 0) {
          num ~/= 2;
        } else {
          num *= 3;
          num += 1;
        }
        numSteps++;
      }
    }
    return numSteps;
  }
}
