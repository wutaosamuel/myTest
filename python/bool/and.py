
def withAnd():
    isTrue = True
    a = 1

    print("Init: isTrue is true")
    print("or False = True --> pass")
    isTrue = isTrue or False
    print(isTrue)
    print("and (a==1) is True --> pass")
    isTrue = isTrue and (a == 1)
    print(isTrue)
    print("and (a==2) is False --> pass")
    isTrue = isTrue and (a == 2)
    print(isTrue)


if __name__ == "__main__":
    withAnd()
