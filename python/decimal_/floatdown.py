
a = 1.23
a_float = round(a - 0.000000005, 8)
print(a_float)

b = 1.26789
b_float = round(b-0.05, 1)
print()
print("round up:", b_float)
b_float = round(b-0.0005, 3)
print("round down:", b_float)