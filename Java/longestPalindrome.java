/**
 * For each character in the string, we move the indices outwards
 * and compare the characters until there is a mismatch.
 * 
 * j (in palindrome func) is for even palindromes (Ex. "abba")
 */

class Solution {
  public String longestPalindrome(String s) {
      String sol = "";
      String temp = "";
      for (int i = 0; i < s.length(); i++) {
          temp = palindrome(s, i, i);
          if (temp.length() > sol.length()) {
              sol = temp;
          }
          temp = palindrome(s, i, i+1);
          if (temp.length() > sol.length()) {
              sol = temp;
          }
      }
      return sol;
  }

  public String palindrome(String s, int i, int j) {
      int l = i;
      int r = j;
      while (l >= 0 && r < s.length() && s.charAt(l) == s.charAt(r)) {
          l -= 1;
          r += 1;
      }
      return s.substring(l+1, r);
  }
}
