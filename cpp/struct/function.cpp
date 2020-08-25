#include <iostream>

using namespace std;

struct TreeNode
{
  int val;
  // can work on 'struct TreeNode *left;'
  TreeNode *left;
  TreeNode *right;
  TreeNode(int x) : val(x), left(NULL), right(NULL){};
};

int main()
{
  TreeNode *first, *tree;
  first = new TreeNode(10);
  tree = first;
  for(int i = 0; i < 3; i++) {
    tree->right = new TreeNode(i);
    tree = tree->right;
  }
  for(int i = 0; i < 4; i++) {
    cout << "i: " << first->val << endl;
    first = first->right;
  }
  return 1;
}