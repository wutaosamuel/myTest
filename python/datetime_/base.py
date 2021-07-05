from datetime import datetime


def main():
    birthday = datetime(year=1993, month=10, day=20)
    print(birthday)

    now = datetime.now()
    print(now)

    if now == now:
        print("now == now")


if __name__ == "__main__":
    main()
