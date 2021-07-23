String reverse(String s) {
  // Put your code here
  String reverse = '';
  if (s.length == 0) {
    return s;
  }
  for (int i = s.length - 1; i >= 0; i--) {
    reverse += s[i];
  }
  return reverse;
}
