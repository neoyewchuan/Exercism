class Leap {
  bool leapYear(int yr) {
    return (yr % 4 == 0 && yr % 100 != 0) || (yr % 400 == 0);
  }
}
