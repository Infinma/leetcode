class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.length;
        int n = nums2.length;
        if (m > n) {
            return findMedianSortedArrays(nums2, nums1);
        }
        int iMin = 0;
        int iMax = m;
        int median = -1;
        int i = -1;
        int j = -1;
        while (iMin <= iMax) {
            i = (iMin + iMax) / 2;
            j = (m + n + 1) / 2 - i;
            
            if (i < m && j > 0 && nums2[j-1] > nums1[i]) {
                iMin = i + 1;
            } else if (i > 0 && j < n && nums1[i-1] > nums2[j]) {
                iMax = i - 1;
            } else {
                if (i == 0) {
                    median = nums2[j-1];
                } else if (j == 0) {
                    median = nums1[i-1];
                } else {
                    if (nums1[i-1] > nums2[j-1]) {
                        median = nums1[i-1];
                    } else {
                        median = nums2[j-1];
                    }
                }
                break;
            }
        }
        if ((m + n) % 2 == 1) {
            return (double)median;
        } else if (i == m) {
            return (median + nums2[j]) / 2.0;
        } else if (j == n) {
            return (median + nums1[i]) / 2.0;
        } else {
            if (nums1[i] < nums2[j]) {
                return (median + nums1[i]) / 2.0;
            } else {
                return (median + nums2[j]) / 2.0;
            }
        }
    }
}
