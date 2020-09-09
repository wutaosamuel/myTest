extern void swap_with_zero(int *array, int len, int n);

class Solution
{
public:
  void sort(int *array, int len)
  {
    int zero_p;
    for(int i=0; i < len; i++)
    {
      if (*(array +i) == 0)
      {
        zero_p = i;
        break;
      }
    }
  }
};