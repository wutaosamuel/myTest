#include <iostream>
#include <vector>

using namespace std;

struct node {
  int val;
  node(int x): val(x) {}
};

int main()
{
  vector<node> v{node(1), node(2)};

  node *n = &v[1];
  cout << "before: " << v[1].val << endl;
  cout << "before pointer: " << n->val << endl;
  n->val = 10;
  cout << "" << endl;
  cout << "after: " << v[1].val << endl;
  cout << "after pointer: " << n->val << endl;
  return 1;
}