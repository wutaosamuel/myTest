import asyncio

from asyncio.base_events import _run_until_complete_cb


async def asyncPrint(i):
    print("start {}th".format(i))
    await asyncio.sleep(1)
    print("finish @", i)


def main():
    loop = asyncio.get_event_loop()
    myFuncList = (asyncPrint(i) for i in range(10))
    try:
        loop.run_until_complete(asyncio.gather(*myFuncList))
    finally:
        loop.close()


def mainFuture():
    loop = asyncio.get_event_loop()
    myFuncList = [asyncio.ensure_future(asyncPrint(i)) for i in range(10)]
    try:
        loop.run_until_complete(asyncio.wait(myFuncList))
    finally:
        loop.close()


if __name__ == "__main__":
    mainFuture()
