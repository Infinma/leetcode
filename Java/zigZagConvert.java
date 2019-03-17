class Solution {
  public String convert(String s, int numRows) {
      if (numRows <= 1 || numRows >= s.length()) {
          return s;
      }
      String[] words = new String[numRows];
      boolean increment = true;
      int row = 0;
      int i = 0;
      for (int w = 0; w < words.length; w++) {
          words[w] = "";
      }
      while (i < s.length()) {
          words[row] += s.charAt(i);
          i += 1;
          if (increment) {
              row += 1;
          } else {
              row -= 1;
          }
          if (row == numRows - 1 && increment) {
              increment = false;
          } else if (row == 0 && !increment) {
              increment = true;
          }
      }
      String sol = "";
      for (int w = 0; w < words.length; w++) {
          sol += words[w];
      }
      return sol;
  }
}
