import 'dart:convert';
import 'dart:developer';

class Diamond {
  // Put your code here
  final String alpha = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';

  List<String> rows(String theAlpha) {
    String endAlpha = theAlpha.substring(0, 1).toUpperCase();
    String strAlpha = 'A';
    int alphaCount = endAlpha.codeUnitAt(0) - strAlpha.codeUnitAt(0) + 1;
    List<String> result = [];
    int ttlRowCol = (alphaCount * 2) - 1; // no. of rows = alphaCount * 2 - 1

    for (int rowPtr = 0; rowPtr < ttlRowCol; rowPtr++) {
      // loop through the row
      String temp = '';
      for (int colPtr = 0; colPtr < ttlRowCol; colPtr++) {
        // loop through the column
        if (rowPtr < alphaCount &&
            (colPtr - rowPtr == (alphaCount - 1) ||
                colPtr + rowPtr == (alphaCount - 1))) {
          temp += alpha[rowPtr];
        } else if (rowPtr >= alphaCount &&
            (rowPtr - colPtr == (alphaCount - 1) ||
                rowPtr + colPtr == (alphaCount - 1) * 3)) {
          temp += alpha[alphaCount - 1 - (rowPtr - alphaCount + 1)];
        } else {
          temp += ' ';
        }
      }
      result.add(temp);
    }
    return result;
  }
}
