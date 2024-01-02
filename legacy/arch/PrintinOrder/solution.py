from typing import Callable

TP = Callable[[], None]
TM = Callable[['Foo', TP], None]


class Foo:
    def __init__(self):
        self.priority = 1
        self.done = [False] * 3

    @staticmethod
    def wait(priority: int) -> Callable[[TM], TM]:
        def decorator(method: TM) -> TM:
            def wrapper(self: 'Foo', func: TP) -> None:
                nonlocal priority

                while self.priority != priority:
                    pass

                method(self, func)
                self.priority += 1
                self.done[priority-1] = True

            return wrapper

        return decorator

    @wait(1)
    def first(self, printFirst: TP) -> None:
        printFirst()

    @wait(2)
    def second(self, printSecond: TP) -> None:
        printSecond()

    @wait(3)
    def third(self, printThird: TP) -> None:
        printThird()
