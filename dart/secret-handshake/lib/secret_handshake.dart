class SecretHandshake {
  List<String> commands(int theCommand) {
    // first convert the number to binary
    String commandBinary = theCommand.toRadixString(2);
    List<String> theSecret = [];
    for (int i = commandBinary.length - 1; i >= 0; i--) {
      if (commandBinary[i] == '1') {
        switch (commandBinary.length - i) {
          case 1:
            theSecret.add('wink');
            break;
          case 2:
            theSecret.add('double wink');
            break;
          case 3:
            theSecret.add('close your eyes');
            break;
          case 4:
            theSecret.add('jump');
            break;
          case 5:
            theSecret.reversed;
        }
      }
    }
    return theSecret;
  }
}
