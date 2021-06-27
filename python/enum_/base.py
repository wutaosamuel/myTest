from enum import Enum


class Status(Enum):
    SELL = "sell"
    BULL = "bull"
    DONE = "done"


def main():
    print(Status.SELL)
    print(Status("done"))
    print("Status.SELL name:value: " +
          Status.SELL.name + ":" + Status.SELL.value)
    print("Test equal -> pass")
    if Status.SELL == Status("sell"):
        print("Status.SELL == Status(\"sell\")")
    print("Test Name equal -> pass")
    if Status.SELL.name == "SELL":
        print("Status.SELL.name == \"SELL\"")


if __name__ == "__main__":
    main()
