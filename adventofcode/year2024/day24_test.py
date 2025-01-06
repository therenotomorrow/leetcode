import pytest

from adventofcode.library import testit
from adventofcode.year2024 import day24

example = """
x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj
"""
answer = """
x00: 1
x01: 0
x02: 0
x03: 1
x04: 1
x05: 1
x06: 0
x07: 0
x08: 0
x09: 1
x10: 0
x11: 0
x12: 0
x13: 0
x14: 0
x15: 1
x16: 1
x17: 0
x18: 1
x19: 0
x20: 0
x21: 0
x22: 0
x23: 1
x24: 1
x25: 0
x26: 0
x27: 0
x28: 1
x29: 0
x30: 1
x31: 0
x32: 0
x33: 0
x34: 1
x35: 1
x36: 1
x37: 0
x38: 1
x39: 1
x40: 0
x41: 0
x42: 0
x43: 1
x44: 1
y00: 1
y01: 1
y02: 1
y03: 1
y04: 0
y05: 1
y06: 0
y07: 1
y08: 0
y09: 1
y10: 1
y11: 1
y12: 1
y13: 1
y14: 1
y15: 0
y16: 0
y17: 0
y18: 1
y19: 1
y20: 1
y21: 1
y22: 1
y23: 1
y24: 0
y25: 1
y26: 1
y27: 0
y28: 1
y29: 0
y30: 1
y31: 0
y32: 1
y33: 0
y34: 1
y35: 1
y36: 0
y37: 1
y38: 0
y39: 0
y40: 0
y41: 0
y42: 0
y43: 0
y44: 1

rds AND wpc -> cmj
vbn XOR mkk -> z13
y04 AND x04 -> bbw
ftn OR swv -> bmw
dkj AND gqs -> vbm
x38 XOR y38 -> bhv
jhh XOR vws -> z04
y38 AND x38 -> tbv
whm AND rbn -> tpt
fkf XOR whs -> z31
x37 XOR y37 -> bmt
y43 AND x43 -> jbc
qkj XOR fdg -> z15
tsq XOR pmr -> z21
x11 XOR y11 -> jtg
x35 XOR y35 -> hqk
whs AND fkf -> vpr
y16 XOR x16 -> sbt
y15 AND x15 -> gmt
tdp AND qqs -> sph
bsn OR bdc -> mkk
y12 AND x12 -> bdc
qmj OR rmv -> tdp
fdg AND qkj -> fjw
wfd XOR cwj -> z36
rgm OR vgp -> cvt
vws AND jhh -> shg
qgd OR ggn -> cvr
x28 AND y28 -> rtf
y42 XOR x42 -> jgh
jsv OR kbr -> cwj
hdp OR rtr -> gms
gdm OR kqg -> wpc
x00 XOR y00 -> z00
x17 AND y17 -> qqc
hrt AND sbt -> dfq
y07 AND x07 -> nsg
y18 AND x18 -> z18
mjm AND njj -> bch
y39 AND x39 -> qnt
x16 AND y16 -> hsb
grq XOR vfq -> z28
rbn XOR whm -> z40
x25 AND y25 -> rwp
x41 AND y41 -> hdp
x42 AND y42 -> ftn
rcm AND rhd -> qvh
wsq XOR cvr -> z19
rwp OR bvm -> gdq
kgd OR kqf -> z10
x40 AND y40 -> gkj
x09 XOR y09 -> rds
jhm AND fqt -> rtr
y07 XOR x07 -> jfr
x35 AND y35 -> kbr
sqr XOR mwq -> mwk
hmc AND brk -> pjk
y23 XOR x23 -> gwq
tdd OR nsg -> cct
y14 XOR x14 -> njj
wtm XOR phh -> z06
wws AND cdh -> dff
thq XOR mfd -> z17
y13 XOR x13 -> vbn
y24 XOR x24 -> hsw
x12 XOR y12 -> jnk
vvj OR hph -> jsp
swf XOR hqk -> z35
jjg OR pjp -> tqk
vbm OR wss -> whs
x22 XOR y22 -> kdn
y11 AND x11 -> pjp
x08 XOR y08 -> cjn
std XOR jfr -> z07
tjj OR vpp -> jhh
pjk OR vpj -> rhd
vkd XOR vtd -> z34
cjn AND cct -> kqg
hnw OR nth -> phh
x30 AND y30 -> wss
ptr OR dnq -> wrn
wfd AND cwj -> vvj
wpd AND ndm -> jcq
jmh XOR mrs -> z24
x29 XOR y29 -> wws
x33 AND y33 -> wvp
y02 AND x02 -> wwc
nfh XOR nqq -> qgd
y30 XOR x30 -> gqs
jcq OR qnt -> rbn
wws XOR cdh -> z29
bmn OR cmj -> mwq
y06 AND x06 -> mgc
wkt OR ddd -> swf
phh AND wtm -> cjd
wvp OR gqp -> vtd
jhm XOR fqt -> z41
x20 XOR y20 -> vcn
gms AND jgh -> swv
y23 AND x23 -> sdd
ndb OR qqc -> nfh
png XOR gdq -> z26
nfh AND nqq -> ggn
x40 XOR y40 -> whm
y33 XOR x33 -> wwp
sph OR hpb -> vfq
vjs OR sfm -> cpv
y43 XOR x43 -> bng
cct XOR cjn -> z08
jmh AND mrs -> bmv
x01 XOR y01 -> brk
tsq AND pmr -> vjs
cwt AND mcd -> pcq
jfr AND std -> tdd
sbt XOR hrt -> z16
y20 AND x20 -> twp
y05 AND x05 -> hnw
y27 XOR x27 -> qqs
bmv OR hsw -> mtn
y10 XOR x10 -> sqr
bmt XOR jsp -> z37
sdd OR hcp -> mrs
shg OR bbw -> ckj
y04 XOR x04 -> vws
jsp AND bmt -> mwp
png AND gdq -> rmv
mwk XOR jtg -> z11
twp OR rqv -> tsq
x19 AND y19 -> dtr
x18 XOR y18 -> nqq
wpc XOR rds -> z09
kdn XOR cpv -> z22
x31 AND y31 -> vjp
rtf OR cbh -> cdh
dkj XOR gqs -> z30
y34 XOR x34 -> vkd
mgc OR cjd -> std
tff AND mtn -> bvm
gwq AND wrn -> hcp
vcn AND spr -> rqv
y44 XOR x44 -> mcd
y01 AND x01 -> vpj
vbn AND mkk -> tsd
tqk AND jnk -> bsn
x27 AND y27 -> hpb
bhv AND kqb -> nrc
spr XOR vcn -> z20
ckj XOR tbw -> z05
y22 AND x22 -> ptr
ckj AND tbw -> nth
cpv AND kdn -> dnq
fjw OR gmt -> hrt
qvh OR wwc -> sgb
gwq XOR wrn -> z23
jgh XOR gms -> z42
y32 AND x32 -> vgp
mwq AND sqr -> kgd
gkj OR tpt -> fqt
x34 AND y34 -> wkt
x14 AND y14 -> pwr
y31 XOR x31 -> fkf
mwk AND jtg -> jjg
y05 XOR x05 -> tbw
pwr OR bch -> fdg
x21 AND y21 -> sfm
vfq AND grq -> cbh
x41 XOR y41 -> jhm
y36 XOR x36 -> wfd
mcr AND bgf -> rgm
psb XOR sgb -> z03
bmw AND bng -> trs
x02 XOR y02 -> rcm
jnk XOR tqk -> z12
dfq OR hsb -> mfd
vtd AND vkd -> ddd
bhv XOR kqb -> z38
y37 AND x37 -> ftj
fmp OR pcq -> z45
brk XOR hmc -> z01
dff OR hcb -> dkj
cvt AND wwp -> z33
x13 AND y13 -> vwv
qqs XOR tdp -> z27
x26 XOR y26 -> png
x15 XOR y15 -> qkj
x17 XOR y17 -> thq
trs OR jbc -> cwt
y36 AND x36 -> hph
x10 AND y10 -> kqf
x06 XOR y06 -> wtm
x32 XOR y32 -> mcr
mtn XOR tff -> z25
y28 XOR x28 -> grq
mwp OR ftj -> kqb
x03 XOR y03 -> psb
x25 XOR y25 -> tff
njj XOR mjm -> z14
y03 AND x03 -> vpp
x00 AND y00 -> hmc
mfd AND thq -> ndb
x26 AND y26 -> qmj
x29 AND y29 -> hcb
y19 XOR x19 -> wsq
bgf XOR mcr -> z32
vpr OR vjp -> bgf
rhd XOR rcm -> z02
bmw XOR bng -> z43
nrc OR tbv -> ndm
y24 AND x24 -> jmh
x44 AND y44 -> fmp
wwp XOR cvt -> gqp
cvr AND wsq -> ghw
x39 XOR y39 -> wpd
x09 AND y09 -> bmn
hqk AND swf -> jsv
y08 AND x08 -> gdm
sgb AND psb -> tjj
vwv OR tsd -> mjm
y21 XOR x21 -> pmr
cwt XOR mcd -> z44
ghw OR dtr -> spr
wpd XOR ndm -> z39
"""


class TestDay:
    @pytest.mark.parametrize(testit.names, [(example, 2024), (answer, 45121475050728)], ids=testit.ids())
    def test_first_star(self, indata: str, want: int) -> None:
        testit.first_star(day24.Day(indata), want)

    @pytest.mark.parametrize(
        testit.names,
        [
            (example, 'ffh,hwm,kjc,mjb,ntg,rvg,tgd,wpb,z02,z03,z05,z06,z07,z08,z10,z11'),
            (answer, 'gqp,hsw,jmh,mwk,qgd,z10,z18,z33'),
        ],
        ids=testit.ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        testit.second_star(day24.Day(indata), want)
