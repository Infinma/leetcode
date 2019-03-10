/**
 * Uses a sliding window to obtain O(n) time complexity
 * charList = <char, latest index of char>
 * 
 * If a matching char is found that is past the index of `start`,
 * start is updated so that it is the first char after the matching char
 */

class Solution {
    public int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> charList = new HashMap<>();
        int maxLen = 0;
        int start = 0;
        int end = 0;
        while (end < s.length()) {
            if (charList.containsKey(s.charAt(end))) {
                if (charList.get(s.charAt(end)) >= start) {
                    if (end - start > maxLen) {
                        maxLen = end - start;
                    }
                    start = charList.get(s.charAt(end)) + 1;
                }
            }
            charList.put(s.charAt(end), end);
            end += 1;
        }
        if (end - start > maxLen) {
            return end - start;
        } else {
            return maxLen;
        }
    }
}
