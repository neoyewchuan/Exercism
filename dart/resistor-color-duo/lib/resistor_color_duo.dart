class ResistorColorDuo {
  int value(List<String> colors) {
    String tResult = "";
    colors.forEach((element) {
      switch (element.toLowerCase()) {
        case 'black':
          tResult += "0";
          break;
        case 'brown':
          tResult += "1";
          break;
        case 'red':
          tResult += "2";
          break;
        case 'orange':
          tResult += "3";
          break;
        case 'yellow':
          tResult += "4";
          break;
        case 'green':
          tResult += "5";
          break;
        case 'blue':
          tResult += "6";
          break;
        case 'violet':
          tResult += "7";
          break;
        case 'grey':
          tResult += "8";
          break;
        case 'white':
          tResult += '9';
          break;
      }
    });
    // line 42 covert the string into number first then back to string
    // this will remove the first '0' if the first resistor color is black
    // as there is no specification of how should be handled for such scenario
    // this code assume that the result must always be two digit number
    tResult = int.parse(tResult).toString();
    return int.parse(tResult.substring(0, 2));
  }
}
