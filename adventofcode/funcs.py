import io
import typing


def cleaned_line(infile: io.TextIOBase) -> typing.Iterable[str]:
    yield from filter(bool, map(str.strip, infile))
