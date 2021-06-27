def foo():
    print("starting...")
    while True:
        res = yield 4
        print("res: ", res)


def main():
    g = foo()
    next(g)
    g.send(7)
    g.send(2)
    # print(next(g))
    # print(g.send(7))
    # print(g.send(2))


if __name__ == "__main__":
    main()
