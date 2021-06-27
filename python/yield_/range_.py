def foo(num):
    print("starting")
    # yield num+100
    # yield num+200
    while num < 10:
        num += 1
        yield num

    while num < 50:
        num += 5
        yield num


def main():
    i = 0

    for n in foo(0):
        print(i, ":", n)
        i += 1


if __name__ == "__main__":
    main()
