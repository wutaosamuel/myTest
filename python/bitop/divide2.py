
def devideByInteger():
    a = 20
    print(a, "devide: ", a >> 1)
    print("works")


def devideByFloat():
    a = 20.2224
    #print(a, "devide: ", a << 1)
    print(a, "float number not works")


if __name__ == "__main__":
    devideByInteger()
    devideByFloat()
