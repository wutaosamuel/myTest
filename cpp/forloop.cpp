#include <iostream>

int main()
{
  int loop_index = 0;
  for (int i = 0; i <= 20; i += loop_index)
  {
    //std::cout << i << std::endl;
    for (int j = 1; j <= 4; j++)
    {
      std::cout << loop_index << std::endl;
      loop_index += j;
    }
  }

  return 0;
}