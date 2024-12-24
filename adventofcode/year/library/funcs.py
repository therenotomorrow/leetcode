import typing


def split(indata: str) -> typing.Iterable[str]:
    yield from filter(bool, map(str.strip, indata.split('\n')))


def parts(indata: str) -> tuple[str, str]:
    comps = indata.split('\n\n')

    return comps[0], comps[1]
