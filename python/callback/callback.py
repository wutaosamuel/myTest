
def my_callback(a, b=-1):
    print("my callback:", a+b)


# type Callback func(x, y int) int
def caller(a, b=-1, Callback=None):
    if Callback:
        Callback(a, b)
    if not Callback:
        print("work callback is none")


def Main():
    print("Test with callback:")
    caller(1, 2, my_callback)
    print()
    print("Test none callback:")
    caller(1, 2)


if __name__ == "__main__":
    Main()
