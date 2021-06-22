from decimal import Decimal

a = 1.23
a_decimal = Decimal(a).quantize(
    Decimal("0.00000001"), rounding="ROUND_HALF_UP")
a_float = float(a_decimal)
print(a_float)

b = 1.26789
b_decimal = Decimal(b).quantize(Decimal("0.1"), rounding="ROUND_HALF_UP")
b_float = float(b_decimal)
print()
print("round up:", b_float)
b_decimal = Decimal(b).quantize(Decimal("0.001"), rounding="ROUND_DOWN")
b_float = float(b_decimal)
print("round down:", b_float)

dict = {
    "a decimal": a_decimal,
    "b decimal": b_decimal,
}
print(str(dict))
dict = {
    "a float": a_float,
    "b float": b_float,
}
print(str(dict))
