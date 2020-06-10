#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int removeDuplicates(vector<int> &nums)
{
}

int main()
{
  vector<int> ns = {0, 1, 2, 3, 4, 5, 6};
  for (int i=0; i<ns.size(); i++) {
    if (i == 4) {
      ns.push_back(ns[i]);
      ns.erase(ns.begin() + i);
    }
    //cout << ns[i] << endl;
  }
  for (int i=0; i<ns.size(); i++) {
    cout << ns[i] << endl;
  }
}