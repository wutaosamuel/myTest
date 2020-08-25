#include <iostream>

using namespace std;

int partition(int *arr, int low, int high)
{
  int p = *(arr + low);
  while (low < high)
  {
    while (low < high && *(arr+high) >= p) 
    {
      high--;
    }
    *(arr+low) = *(arr+high);
    while (low < high && *(arr+low) <= p)
    {
      low++;
    }
    *(arr+high) = *(arr+low);
  }
  *(arr+low) = p;

  return low;
}

void quicksort(int *arr, int low, int high)
{
  if (low < high) {
    int p = partition(arr, low, high);
    quicksort(arr, low, p - 1);
    quicksort(arr, p + 1, high);
  }
}

int main()
{
  int arr[] = {9,5,8,7,2,1,4,6,3,0};
  quicksort(&arr[0], 0, 9);
  for (int i=0; i < 10; i++) {
    cout << arr[i] << endl;
  }
  return 1;
}