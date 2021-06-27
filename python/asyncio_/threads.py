import time
import asyncio
import concurrent.futures as cf


def myPrint(i):
    print("start {}th".format(i))
    time.sleep(1)
    print("finish @", i)


async def threadsPrint():
    with cf.ThreadPoolExecutor(max_workers=10) as executor:
        loop = asyncio.get_event_loop()
        futures = (
            loop.run_in_executor(
                executor, myPrint, i) for i in range(10)
        )
        for result in await asyncio.gather(*futures):
            pass


def main():
    loop = asyncio.get_event_loop()
    try:
        loop.run_until_complete(threadsPrint())
    finally:
        loop.close()


if __name__ == "__main__":
    main()
