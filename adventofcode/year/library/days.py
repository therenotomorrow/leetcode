import dataclasses


@dataclasses.dataclass
class Day:
    indata: str

    def first_star(self) -> int:
        raise NotImplementedError()

    def second_star(self) -> int:
        raise NotImplementedError()
