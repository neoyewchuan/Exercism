class MatchingBrackets {
  // Put your code here
  bool isPaired(String s) {
    List<String> brackets = [];

    for (int i = 0; i < s.length; i++) {
      if (s[i] == '(' || s[i] == '[' || s[i] == '{') {
        brackets.add(s[i]);
      } else if (s[i] == ')' && brackets.length > 0 && brackets.last == '(') {
        brackets.removeLast();
      } else if (s[i] == ']' && brackets.length > 0 && brackets.last == '[') {
        brackets.removeLast();
      } else if (s[i] == '}' && brackets.length > 0 && brackets.last == '{') {
        brackets.removeLast();
      }
    }
    return (brackets.length == 0);
  }
}
