import typing

import pytest

from .add_poly import PolyNode, PolyNodeT, add_poly


class TestAddPoly:
    @pytest.mark.parametrize(
        'poly1, poly2, want',
        [
            (
                PolyNode(coefficient=1, power=1),
                PolyNode(coefficient=1, power=0),
                PolyNode(
                    coefficient=1,
                    power=1,
                    next_node=PolyNode(coefficient=1, power=0),
                ),
            ),
            (
                PolyNode(
                    coefficient=2,
                    power=2,
                    next_node=PolyNode(
                        coefficient=4,
                        power=1,
                        next_node=PolyNode(coefficient=3, power=0),
                    ),
                ),
                PolyNode(
                    coefficient=3,
                    power=2,
                    next_node=PolyNode(
                        coefficient=-4,
                        power=1,
                        next_node=PolyNode(coefficient=-1, power=0),
                    ),
                ),
                PolyNode(
                    coefficient=5,
                    power=2,
                    next_node=PolyNode(coefficient=2, power=0),
                ),
            ),
            (
                PolyNode(coefficient=1, power=2),
                PolyNode(coefficient=-1, power=2),
                None,
            ),
            (None, None, None),
            (
                PolyNode(
                    coefficient=9,
                    power=10,
                    next_node=PolyNode(
                        coefficient=-8,
                        power=9,
                        next_node=PolyNode(
                            coefficient=1,
                            power=5,
                            next_node=PolyNode(coefficient=-7, power=3),
                        ),
                    ),
                ),
                PolyNode(
                    coefficient=-4,
                    power=4,
                    next_node=PolyNode(coefficient=7, power=3),
                ),
                PolyNode(
                    coefficient=9,
                    power=10,
                    next_node=PolyNode(
                        coefficient=-8,
                        power=9,
                        next_node=PolyNode(
                            coefficient=1,
                            power=5,
                            next_node=PolyNode(coefficient=-4, power=4),
                        ),
                    ),
                ),
            ),
        ],
        ids=(
            'smoke 1',
            'smoke 2',
            'smoke 3',
            'test 3: runtime',
            'test 45: wrong answer',
        ),
    )
    def tests(
        self,
        poly1: PolyNodeT,
        poly2: PolyNodeT,
        want: PolyNodeT,
    ) -> None:
        got = add_poly(poly1, poly2)

        while got:
            want = typing.cast(PolyNode, want)

            assert got.coefficient == want.coefficient, f'add_poly() = {got}, want = {want}'
            assert got.power == want.power, f'add_poly() = {got}, want = {want}'

            got, want = got.next, want.next
