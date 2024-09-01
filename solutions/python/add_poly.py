import typing


class PolyNode:
    def __init__(
        self,
        coefficient: int = 0,
        power: int = 0,
        next_node: 'PolyNodeT' = None,
    ):
        self.coefficient = coefficient
        self.power = power
        self.next = next_node


PolyNodeT = typing.Optional[PolyNode]


def add_poly(poly1: PolyNodeT, poly2: PolyNodeT) -> PolyNodeT:
    if not poly1 and not poly2:
        return None

    dummy = PolyNode()
    curr = dummy
    curr1, curr2 = poly1, poly2

    while curr1 and curr2:
        if curr1.power == curr2.power:
            coefficient = curr1.coefficient + curr2.coefficient
            power = curr1.power
            curr1, curr2 = curr1.next, curr2.next

        elif curr1.power > curr2.power:
            coefficient = curr1.coefficient
            power = curr1.power
            curr1 = curr1.next

        else:
            coefficient = curr2.coefficient
            power = curr2.power
            curr2 = curr2.next

        if coefficient:
            curr.next = PolyNode(coefficient=coefficient, power=power)
            curr = curr.next

    curr.next = curr1 if curr1 else curr2

    return dummy.next
