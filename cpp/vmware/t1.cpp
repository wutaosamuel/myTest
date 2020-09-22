#include <iostream>
#include <math.h>

using namespace std;

int main()
{
  float n, r;
  int m;
  cin >> n >> m >> r;

  float result[m][2];
  for (int i = 1; i <= m; i++)
  {
    float a = float(i) * r;
    float c = floor(a / (4 * n));
    if (a >= 4 * n)
      a = a - 4 * n * c;
    if (0 <= a && a <= n)
    {
      result[i - 1][0] = a;
      result[i - 1][1] = 0;
    }
    if (n < a && a <= 2 * n)
    {
      result[i - 1][0] = n;
      result[i - 1][1] = a - n;
    }
    if (2 * n < a && a <= 3 * n)
    {
      result[i - 1][0] = n - (a - 2 * n);
      result[i - 1][1] = n;
    }
    if (3 * n < a && a <= 4 * n)
    {
      result[i - 1][0] = 0;
      result[i - 1][1] = n - (a - 3 * n);
    }
  }

  for (int i = 0; i < m; i++)
  {
    printf("%2.2f %2.2f\n", result[i][0], result[i][1]);
  }

  return 0;
}