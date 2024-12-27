import functools
import random
import threading
import time
import unittest

from solution import Foo


def printer(_: str) -> None:
    time.sleep(random.randint(1, 3))


class TestFoo(unittest.TestCase):
    def test_order(self) -> None:
        foo = Foo()
        threads = [
            threading.Thread(
                target=foo.third,
                args=(functools.partial(printer, 'third'),),
            ),
            threading.Thread(
                target=foo.second,
                args=(functools.partial(printer, 'second'),),
            ),
            threading.Thread(
                target=foo.first,
                args=(functools.partial(printer, 'first'),),
            ),
        ]

        for th1 in threads:
            th1.start()

        for th2 in threads:
            th2.join()

        self.assertEqual(foo.priority, 4)  # after third will be 4
        self.assertEqual(foo.done, [True, True, True])
