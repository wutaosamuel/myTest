#include <stdio.h>

struct TreeNode 
{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
  // NOT WORK ON C  
  // TreeNode(int x): val(x), left(NULL), right(NULL) {};
};