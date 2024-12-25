import dataclasses
import typing


@dataclasses.dataclass
class Day:
    indata: str

    def parse(self) -> typing.Any:
        raise NotImplementedError()

    def first_star(self) -> int:
        raise NotImplementedError()

    def second_star(self) -> int:
        raise NotImplementedError()
