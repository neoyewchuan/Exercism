int score(String word) {
  int theScore = 0;
  for (int i = 0; i < word.length; i++) {
    theScore += letterValue(word[i]);
  }

  return theScore;
}

int letterValue(String s) {
  String sUpper = s.toUpperCase();
  if (sUpper.codeUnitAt(0) >= 65 && sUpper.codeUnitAt(0) <= 90) {
    switch (sUpper) {
      case 'Q':
      case 'Z':
        return 10;
      case 'J':
      case 'X':
        return 8;
      case 'K':
        return 5;
      case 'F':
      case 'H':
      case 'V':
      case 'W':
      case 'Y':
        return 4;
      case 'B':
      case 'C':
      case 'M':
      case 'P':
        return 3;
      case 'D':
      case 'G':
        return 2;
      default:
        return 1;
    }
  }
  return 0;
}
