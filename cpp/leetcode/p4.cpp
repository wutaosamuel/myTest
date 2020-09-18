#include<iostream>
#include<vector>

using namespace std;

double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int i = 0, j = 0;
        int result1 = 0, result2 = 0;
        int totalSize, location;
        bool isDouble;
        totalSize = nums1.size() + nums2.size();
        isDouble = (totalSize % 2 == 0) ? true : false;
        location = isDouble ? totalSize / 2 - 1 : totalSize / 2;
        if (totalSize == 1)
            return (nums1.size() == 1) ? (double)nums1[0] : (double)nums2[0];
        if (nums1.size() == 0)
            return isDouble ? (double)(nums2[location] + nums2[location+1])/2 : double(nums2[location]);
        if (nums2.size() == 0)
            return isDouble ? (double)(nums1[location] + nums1[location+1])/2 : double(nums1[location]);
        
        for (int index=0; index <= location+1; index++)
        {
            cout << i << " " << j << endl;
            if (result1 >= result2)
                result2 = result1;
            if (i >= nums1.size())
            {
                if (result1 <= nums2[j])
                    result1 = nums2[j];
                j++;
                continue;
            }
            if (j >= nums2.size())
            {
                if (result1 <= nums1[i])
                    result1 = nums1[i];
                i++;
                continue;
            }
            if (nums1[i] > nums2[j]) {
                result1 = nums2[j];
                j++;
            } else {
                result1 = nums1[i];
                i++;
            }
        }
        cout << result1 << " " << result2 << endl;
        if (isDouble)
            return ((double)result1 + (double)result2) / (double)2;
        return (double)result1;
    }

int main() {
  vector<int> n1 = {1, 2, 3, 4, 5, 6, 7};
  vector<int> n2 = {1};
  double v = findMedianSortedArrays(n1, n2);
  cout << v << endl;
}