#include <iostream>

using namespace std;

int main()
{
  int num = 4321;
  
  for(int i=0; i <3; i++)
  {
    cout << num << endl;
    num = num << 1;
  }
}