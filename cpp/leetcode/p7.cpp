#include<iostream>

using namespace std;

int reverse(int x) {
        unsigned int result;
        int index = 0;
        int tmp[31];
        bool positive = false;
        if(x < 0)
        {
            positive = false;
            result = -x;
        } else {
            positive = true;
            result = x;
        }
        for (int i=0; i < 31; i++)
        {
            tmp[index] = result % 10;
            result /= 10; 
            if (result == 0)
                break;
            index++;
        }
        for (int j=0; j<=index; j++)
            result = result * 10 + tmp[j];
        return positive ? result : -result;
    }

int main() {
    int a = 1534236469;
    cout << reverse(a) << endl;
}