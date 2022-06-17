import os


def runPWD():
    return os.system("pwd")


def runLocal():
    local = "./"
    return os.path.abspath(local)


if __name__ == "__main__":
    print("They are the same, outputting path with executing location")
    print("Execute pwd: ")
    runPWD()
    print("os.path.abspath with ./ ")
    print(runLocal())
