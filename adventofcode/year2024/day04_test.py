import pytest

from adventofcode.library import testit
from adventofcode.year2024 import day04

example = """
....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
"""
answer = """
MMMSAMXAMMXSXXMAXAXXXSASAMXSAMXSAMXMAXSAMXSXSMSASAMXXXSASXSMSSSSXMAMSAXMSSXXASASMASXAXXMSMSSMMMSAMXMMMMXMXSXXMSXMAXXMAMXMSSXXMAMMMSMMMMMSXAX
XAAXAMXMASXMASMSSSXSASASAMXMXMAMXMAMXASXAXMAXAMXSAMSMMSAMASAAAAAMXMASMMMMMMSMMASAXAMMMSMSMAAAASMMXSMSMSASMAASXMASMMSXMASASASXMASXMASAAASMMAM
SMMSAMAMMAMXAMAMAAAMAMXSMMAXAMMSXSAMXASMSAMAMAMASXMASASAMAMXMMMMAXMAMAAAAAAAMMMMXMSAAXMASMMSMMSAMXSXAASASAMXMASXMAASAMAXMXMAXSASASASXSXXXAMA
XAAMAXMASASMSSXMMMMMAMAMASMSMSAAAXASMAMAMMXMXAMXMAMASXSAMXMASXSXSXMASXMSSSSMMAASAMXSMSMAMAMXMXXAMASMMXMMMXMASXMAMMMSAMXSXSMMMSASAMASAMASMSAS
MMMMMMMASASAMXMXMSMMAMAXAMMAAMMSSMAMMSMSXSAMXSSMSSMASMSXMAMXMASAMAMXMMAXAXXXSSSSMXMASMMASAMMMMSAMASXXAAAAASAMAXXMAXSXMAMMXAAAMMMXMAMAMAMAMMM
SASASAAMMAMXMASXMAMSASXSSSSMSMMAAMAMAXAXMMASAMAMAXMSMAXMSASMMMMXSAMAMMMMMMXMXMAXXSAMXXSASASXAMSAMXSXSXSMSXMMSSMSSSXSASMSMSSMSSXAXMXSSMAMSMXX
MASASASMMMXAMXAXSMMSMSAAAMXAMXXMMSAMXSXXASAMASAMAMMAMMMASXXAMMAAMXSXSASASASMAMXMXMAMMXMASAMMAXMASXMAXAAAMAXAAAAAAXAMAMXAAAXMAXMASXAAXAXMMXSX
MAMAMMXXAMSXSAMMXMASMMMMMMMMMMMXXMAXMAMSMMXSMMXMSSSMXAAMMXSSMMMSXMXMXMSASASMASAMASAMXAMMMMMASMMASMMAMSMMMAMSSMMMSMMMSMSMMMSMXSMSMMMMMMAMXASA
MXSXMAMSAXXMXAMASMASXSXMAMAMAMXMASMMXAMSAMXSAMXXAAXXMSSXXXXMAXAAMMXSMAMMMXMXAMASASMSSMMAAAXSMAMASAMAMXAXMXXXAXAXAAXXAAMXSAXAXMXAXAMMSMAXMAMA
MAMAMAXSXXAMSSMAXMMSAMASXSASXXSAMXAMMSXSAMMMMMAMMSMSAMXAMSAMXMASASAMMMMAXASMSSMMASXAASASXMXMMMMASXXXXXSMSXMSXMSSSSMSMSMAXMASASXMSXMAASMMMAMM
MMSSMMMXSXXXAXMSMMAMASAMMMMSAAXXXXSMAXAMAMAAAAAAAXAXXMMMMSAMSSMAXAAMXMSMSAXMXAXMSMMMSMMMASASASMXSMMSAAMAMMMSXMAXMAXXXAMXSMXMXMAXAAMSMMSXSMSS
XAAXAXMAMSAMMXMXAMXMAMASXAAMMMMXMAXMMMMXMMSSXSASMMAMXSAMMMAMAAXMASMMXXAMMXMMSSMMAMAMXAXSMMXSAMASXMASMSMXMAXMASMMSAMSSMSAMXAMASMMMSMXAASMMAAA
XMXSMMMAMAXASASMSMSMXSAMMMSSMXMXMAXASMSAXMAXAMAMXSMMASMSMSMMSMMXMXMAMMSSMMSAAMXSAXAMMAMXXSAMAMXSAMXSXXXMSMMAMXMAXXAAAASMMSASXSXXMMAMMXMAMMMM
MMMMMASXMXSXSASAAAXAAMASXMAMAAMAMMSXSASMSMMSSXXAAMAMXSXAXMAAMXSAMASAXAMAAAAMMMASMSSSMAXXAMXSAMXSAMAXMMAXAAMSMSMSSMSMMMMAMSXXMMMMAMAMXAXXMXMS
XAAXMASAMMXAMXMXMMMMMSXMMSAXSASXSXMMMAMXMAAMMSAMXSXMASXMSSMMSASASASASXSSMMSMXMASXAAMMSMMMSMXASASAMMSASMMMSMAAXAMXAAXAAMAMMMMXAAXXMSAXMSXMAMA
SSSSMMXAMXMXMAMXSMXMAMAXXXAXMXSASAAXMAMAXMMSASMMMAMXAMAXAAAAMMSMMAMAMMAXSMAMXMAXMMSMAMXSAAASXMASAMMXAXXAXAMMSMXMMMMXXXSXMAAAMXSSSXXAASAMMSSM
MAMXMXMMMMAXMMXAXXXMASAMMMSMSAMASMMMSAMMXSAMASXSXSXMASMMSSMMSMSAMXMMMMSMMMMSXMAXSMMMXMAXMAXMMMAMAMSMMMSXSMXSAXAASASXSMAMSSMSSXAXXAMXMMASXAAX
SAMASMSAAASMMXMSSSSSMXAAXAAXAXMMMAMASXMSAMASAXAXXXSMAAAAAAXXAAXMMSXMAAAASXMMMMMMSAXAAMMMSSSXSMMXMMSASAMXAMXAAXMAMXMAAMXMAXMAMMMMMMMXXMAMMSMM
SAMAXAMMSMAAMMSMAMXAMSSMMSSSMSSXSXMXSMAMSSMMMMSMMMXMASMMMAMSMSMXAMASMSSSMAAMMMSASXMMSSMAMAAASMXAAMSAMASMSSSMMMMSMMMSMMSMMXSASASAMASXXXASMMXS
SMMSXSXXAXSAMXAMAMSAMXMAMXMAXAXXMAMAXMAAXMXAAAAAAXMAXMXXMXMXAMMMASXMAMMXXSMMAAMASXXAAMMAMMMMMASXSMXAMAXAAAXXSAMXASAAASXXXXMASASXSASAMMMSASAM
XMAXSMXMAXMAMSASAMXASMSSSSSSMMSSMSMMXSMSMMSXMSSSMSSMSSSXMAXMMMMXAMMMSMAAXXXSMSSMMMMMSXSXSAAAMXMAAAXXMXMMMMMMSASMMMSXSAMSMXMMMAMAMMSXMASMXMAM
MMXXAXAXXASAMSAMXMSMMAMAAAAXAAAAAMAXXXXMAMASAMXXXAAXAXMAXAXSAMAMMSMAMMMMSMXMAMXMAASAMMSMSXSSSMMSMMSMMXMAAMSXSXMASMMXMAMXMAXAMAMAMMSASXSAMSAM
SSMSSMMMMMXMMSAMSAMXMAMMSMMMMMSSMSAMMXMSSMMAMMMMMXSMSSSSMASMAMXSAAMASAAXMMAMSMXMSMMAMAXASAXAMAMXAXXAAMSSSMSAMMSSMAMMSMSMSMSXSASASMXASAMXAMAM
XAAAAAAAXXSSXSAMXMAXSXXAAMMAMXAXXXAMASXMASMXMASASXMAXAAMMXMMXMASMXMSSXMSSMXSAAXMMXSAMXSSSMMMMSMSMMSSMXXAMXMXMASAMAMAAXAAMAAASASASMMMMMMSSSSM
SMMXXSMXSAMXAMXMAMXMAXMXSXXASAMXMSSMAXAXSAMSSMXMXAMXMMMXMSMXSAMXXSXXXXXAAXSAMSMSAMSXMMXAXXAAAAAAAAAAXMMAXMSAMXSAMAMSSMMSMMMXMXMMMXAAAXMAAAXM
AXMSAAAAMAMMXMAMASASXSSMAMMMXXXAMAAMAMSMMAXAAAMSSMMMXAAMMXAASMSMAMMSXSMSSMSAMXMMAMMASMMMMXMMMSMSMMSSMXSXMAXXXASAMXMAMASMMAXXMSAMXSSMSSMXMXMX
AMAAMMMSMSMMASMXXSASAAXMAMXSSMMSSSSMMSMASXMSMMMMAXAXSSMSAMMMMAAAAAAXAAXXAMXAMXSSSMSAMXAMASXSAMXAAAXXXXXAMSSSMASMMSSMSMMASMSMMASMXXXAMXMAXSXM
SXMMXMAXAMXSMAMSAMXMMMMMASAMAAMAAAXMXAMXMXAXXSAMXMMXAAAXMAMAMSMSSSSMSMSSSMSAMXAAAAMAXSXMMMAMAMSXMMSMSMSAMAAXMAMAMAAASXSMMMAAMAMMXAMXMAMMXXAM
AAXXSXAMMMXSMAAMAMASAXMAXMAMSMMMMMMSSMMMAMSMASMSAMXMMMMMXMASMMAAAXAAXXAAAXXAMMMSMMMXMAXSMSXMAMAMSAMAXAAXMMSMMMSAMMSMMAXXASMSMXSXMXMSSMSSMSAM
MSAMAMASAMXXMXXMSMASASXSXSXMXXSAAAAXXSXMASMMAMASXSXXAASXSAMXAMMMSMMSMMMXMMXMXSXXMAXXAMMXAAAXXMSXMAMXMMMXSXMAXAMAMAAXMSMSMMAXXMMAMXAAXSAAASMM
XMAXMMAXAMMASXSSMMXSMMMMASAMXASMXMMSASXSXSXMASAMMMMMXXSAMXSMMXAMAXXAAAAASMMSAMAMSXSAMXAMSMSAAAXXMSMAAAMMMASMMSSSMSMXMAAAXMASXAXAMAMMSMMSMMMS
MSMMXMMMXMAAMAXMASAMXAXMAMAAMMMSXSAMXMASAMMSXSAMAXAASMMXMAXASMXSASXSSMSSXAXMAMSMMASXXMXMMAXMMMMMMXAXSAXAMXMMAXAMAMAASMSMXMMSXMMAMXXXMAAXMXXX
XAAXSASAAXSSMSMSXMXSMMXSASXMXAAXXMAXAMXMAMASMSASMSXXAAMSSMSXMAMMMMMAMMMMMSMMAMXAMAMMXMMXMAMXSAMAMMAMXMSMMAXMAMSMAMSMMAAAAXAMXMAXXMASXMMSMSSS
SSXMSAASXXAMAXAXMMMAXAAXXSASAMXMXXMXSMMSSMMXAXAMAAXASXMAAXMMMAMASAMAMAASMMAMAMSXMAMSASASMXMASXSAXXXMXMAXSXSXAMXMMXXMMSMSMSASMMSXAXAXAASXAAAX
MMSAMSMMMAMMAMSAAXAAAMXSASAXASAMASMAXAMAXMSMSMMMMMMMMAMSSMMASMSASASMSSMSASMMAXMMXMAMAXMXMAAMSMSMSMMMXSAMMAMXMASAMXMSAAAAAAAXMAMSMMSXMMMAMMSM
XAXMXAXMAMAMXSXMMAMAAAAMAMMMMSASAMMASAMMSXAAAAMXMMAXSAMXMASASXMAXAMMAMXMAMASXXMXAXMMSMMSSMSXMASXAAAMMMMAAAAXMASXMASMSMSMSMXMMSXMMAXAXXXAMXAA
MMSSXMXXSXSXMAXASXSSSSMMXAXMXXXMASMASXMMMMSSMSAMXMMAXAMASMMMMXMAMXMMSXSASMMXAMSXMSXAXAAAAAXMSMMXSXMMXAXMSSMMSASAMXMXXXXXAXSXXXASMAMMMXMAMXMX
SAAMAXMAMXXAMASXMAMAAAXAXSXMASMSMMXXXAXAAAMAMXAXMSMSSXMXSMAASMMASXAXAMMAMAXMMMAASAMMSMMSSMMAAAAAMXSXSASXAAAXMAMXAMXAXMMSASAXXMAMMMSSSXSMSASX
MMSXMMAASMSSMAXAMSMMSMMSXMAXAAAAAXXSAMSMMMSAMXAMAAAXAASAMASMAASASMMMSSMXMXMXSMSXMAMXAAAMAAMSSMMMSASASAAMSSMMMSMSMSMXMAAXMMAMXMSMSXAAAAAASAMA
SXAMMSSXSAAXMASXMMAXAXAXASXMXMSMMSASXXAAXXMASXSSSMMMMMMASAMXSAMXSAXAMXMASMMXXAMMSXMSSSMXSAMXMXAAMAMAMMMAXAXMXXASASASMSMMSMXASXXAXMXSMMMMMAMX
AXMAXAXMMMMMMXSAASMMMMSSMMASXAXASMAMMSSSMSSMMMXAAXXAAXSXMASAMASMSXMXXAXXAAXAMXMAMAMAAAAXXSXAMSSMMXMXMAXMMAAAAMSMAMAXAXXAAMXAAAMMSMMXXAXAXAMX
MMXMMSXMAXXXMXSMXMSAMSMAMMASMAMMMMAMXAAAAMAMSAMSMMSSXMAXMXMMSAMMMAMMSSMSSMMAXSMXSSMMSMMMAMMMXAAASXMMSXSAAASMXSAMXMMMAMMSASMAMAXAAXMASMSSMMSM
XAMXAAMSAMSAMAMXXAMXMAAAXMASAMMAMSAMMMSSMSAMXSAXMAXMMMAMXMSXMMSASXSAAAAXAASMMXAXSAMXXAAMAMAXMSMMMASXAAMASAMAMAMXAAMMSXXMXXXMAMMSSMSAMAAXMMAX
ASAMMSXSAXSAMASXMSMSXMXMSMASAMXAMSASXMAXAMXMMMXMAMXAAXSMMAAMSAXASAMMXMSMSMMAMMMMXAMSMSSSMSSSXXAASAMMMSMAMXMMMAMSSSXAXMASMMXMAMAMXAMMSMMSXSAS
AMSXSXAMMMMXXAMAAAMXASAMXMASAMSSXSMMMMAMMMMMAMASMMMMMMXAASMMMAMMMMMSXAXAXMXAMXAAMMMSAMAAMAXMASXMMAMXXXMAXXXMMMXXXAMASASAAMMSAXSAMXMAAAXAAMAM
XMAMSMXXMASXMSSMSMXSAMASAMMSAMAMMMMAXMXSAAMSASAXAAXXAAMSMMXXSASAAMAMMMMXMMSMMSMMSAAMXMSMMMMXXAMXSAMXMMMSMXAMAMSMXXXMAXXMSMAMMSXAAXMMSMMMSMAM
XSXSXASXMASXAAAAAAAMAMSSXSASAMXSASXSXSASXSMSASASMMSAMSXMAMAMSASMXMASAMAXXMASMAAASMSMAMMMASMMMAXXMXXAAXAXXXXXASASMXMMSXSAXMSMMXXAMXXAMASMXXAM
MXMAXSMAXAXMMMMMMMMXXMXMMMASAMXSASAAAMXXAXASAMAMXAAXXXAMAMAMXMMXXSASMXSAASASMSMMSXAMMMAXAAAMSXMXMASMXMSSMXMSMSAMXAMAXAMMMMAASMMSSSMMXAMAMSSS
AAAMMMSXMSAMXSAMXMSMSXSXAMMMXSAMAMMMMMSMAMAMMMAMMMMMMSMMAMASMSAXXMASAAMSMMASAXMXXXMASXSMSXSMSAMXXMASAMXSAAMAMMXMSXMAMAMXAMSSMAAAAXASMSMSXAXX
MASXAAXXXAAMAMXMASAMSAMSXSAXMMMSMXAXAASMAMXMSXMSMSAMAAXSMSXXAMSSXSXSMMMAXAAMAMXMASMSSXAAAAMASXXAXMASXMASXXSASXSMMXMSSSSSSMXAAMMMSMMMAXAXMSSS
SMXXMSSMXSXMAMASMMXSMAMAMMMSMSAAMMMXMASAMAMMSAAAXMAMSMMAASAMSMXMASAMASMMSMXMAMAXMAXAXMMMMSMAMSXMSXAMAMXSMMXXSMSAASMAXAAAXXSXMXMXXMSMXMAMXAAA
ASASXMAMMMASASAMMMSXXSMMSAAMMMSXSASASAMAMMSASMMMSSMMXXMSMMAMMAAMMMAMAMXMAXASXSSMSXMSSMMAAXMAMMXSAMMSXMXMXSSXMASMMSSMMMMMMMMAMAMXMAAMAMXXMMSM
MMXSASXMASXMMAXXAMXAMMAXSMSSMAMMSASAMAXAMXMMSXSMXAAAXAXMASXMMSSSXSXMASASASMXAAMMMAMAAASMSXSAASMMMMAAAXXMASMAMAMAAMAXAXSASAAAMASAMSMXXXAMAMAX
XMASAMXMXXAAXAMXXMMXMSMMXXMAMASAMAMMSSMASMSXMASMMMMMSSMXAMMAMAMMXMASASXSAXXMMMMAMAMMSMMAMAXXMXAAAMXSSMAASXXAMXMMMSMMMMMASMSXSASXXAMMSXMAXXAX
SMXSMXXMASXMMXSAMXSMMAASXMSXMAMAMXMMAAXSMMMAMAMAAXAXXXAMSMMMMAMXXMAAAMMMMMMXSASMSSSXAAMXMAMSSSMMMMMAMMSMMMSMSXXSXSXMASMXMXXMAMMXMASASASMSMSM
SSMMMMXMAMAAAAMASXAAMMXMAMAMMMMAMXMMSSMXAMMAMASXMSSSMMXMAAAMXSMSMXAMSMSAAAXASMSAAAXXXMMMMXMXAAXAAXMAMAXAXXAAAMAMAMXXAXXSSXAMMSSXSXMASXMASAMS
MAMAAXXMMSSMMXMAMXMSMMMSXMXSASMSSSXAAAASAMXSSMXAMXMAXAASXSMSAMAAASMXMASMSSMMSAMMMMMMXMAXSAMXSMMMMXXAMXMMMSMSMSSMAMMMSSMAAAMSAAAXAMMMMAMAMXMA
SAMSSSMSAAAASAASXSMAXAAAMAXMASAAAXMMSSMMXMAXAXSXMAMSMSMXXAAMAMAMMAXAMAMXMAXAMXMMXMAXAMAXSASMXMXSASXSSSSXAXAAXMASXMAXXAAMSMMMMSSMMXSASMMSSMXM
SXXXAAAMMSSMMAXXAASASMSMMMSMSMMMMMMMAMMSXSXSAMXMSMSAAXAXSMSMSMAXSMSAMMXSMASMAAXMASXSMSSMSXMAXMAMMSAAAAXMMSMSMSMMMXSMXXXXAAXXAXAAXMMAXSAXMAXX
SMSSSMMMAMAAXSSMSMMXSMMAXMAMAAXXXSXMASXXAXAMXMXMASXMMMXMAAXAAXXXSXSXMSAMXAAXSSSMXSAAAAXAMMMSMMASAMMMMMMMXAAAAAXMAXXAMMSSMSSMMSSMMAMSMMMSMSSX
SAAAXXXMASXMMXAMAXMAMASXMSASMSMMASAMSSMMSMMMAXAMXXMAAAMSMMMMMSMASAMAAMMMMMSMMAMMMSAMXMMSMMAXASMSXSSXMXXXXMSMSMSMSSMXSAAAXAMAXAAAMMMXAAXAAASX
MMMSMMXSAXAXSSSMSXSAMAMAMXMMXXAXAMMMMMAAMAMSMSMSASXSMSMXAAAXAXMAMAMXMASMAMAMMAMMAXXXXSAMAMAXAMASXXMAMMMMSMAAMXSXAAMAMXSXMMSSMSXMMASMSMXSMMSS
XMAXAMAMMSMMMAXAXAMXSASMMSMSMSXMMSXXAMSMSAMXMAMMASXAXXXSSMSMMSMMSXMXXSXSASXSMMSXSSMXXMASMMMSMMAMSMSMMASAAMXMMMMMSXMXSAMMMXAXXMMSSMSAAMXMAMXX
SMAXAMAXXAXAMXMMMSMASXSXAASAAASXXAXSMXAAXXMMSMMMMMXSAMXAMAXXSAMMXAXSMXXSMSAAXMAMMAMSSSMMMAXAXMXMXXAASAMXSXSAAXXAAAXXMASAMMXSAMXAAXMXMXAXAMAS
XMSXMMAXSMSXMSASAAMMMXXMMXSMSMAMMMMMMMMSMAXMAAASASXMASXXMMMMSASXSMMMASMSMMXMMMAMSAMAAXMAXXMXMAMXSMSMMASMXASMSMSXSXMAMAMAMXMSAMMSSMSMSMSSSMXS
AXMASMMXSAMXMSAMSMMAAMMSSMMXMXMAAXAAASAMMMMSMSMSASMSXMASMMXAMXMAAXAXSXAXAAMXAXMXSXMMSMSAMSASXMMXXAXMXMMAMAMAMASAAMSXMASXXMASAMAAAASAAAAAAMMM
MSMAMASAMMMAXMAMAMXSAMMAAASXMXXSSSMXMSAMASAXXXAMAMMMASAXAAMXSMMSMSMMMAMSMMSXMXSAMMMSAMXAAXAMAMSSMMMAAXMXMSMXMAMMMMMASAXXAMXSAMMSMMMSMSMSMMAA
SAMXXXMAXXSMSSXMAXMXAMMSMMMAMMMMMMMSMXMXAMXSSMXSAMASMMSXMMMAAMAMAMXSMSMAXMAAMSMMMAXMASMMMMMSAMAAAMASXXAMXAMXMASXMMSXMASXSMXSAMXXASXXXXXXASMS
AXMASMSSMAXAMASXSSMXAMXAMSSMMAAAAAAMXAXMXSXMXMASASXXMAXMSAMMSMASAMSAAASXMMMSMAAAXSMSXXMAAAMAAMSSMSXMMSSMSMSSMAMAMMMSMAMXASAMXMXMASAMMSAMXMAX
SASASAAAMXMASAMXMAMSMMMMSMAAMSSSXMSSSMMSAAAMAMMSAMAAMASMSASAAMXMAXXMSMSXSAAMMMMMMMAAMMSSSSMXXMAXAXAXAAAAMXMXMASXMAASMSXSAMMSAMMXAXASAAXMXMSX
MAMASMMMMXSAMXSXSAMXXAAMAMMMMMMMAXAAAAAMSSMSMSAMMMSXMMSASAMXSSMSSMMXXAXASMXMAXAXAMSMSAMAMMMMXMXMMSMMMSSMSASXSASMSMSMAAXMXAMXASXMMSAMXMXSAAAA
MAMAMXSAAXMXMAXAAAMXSSSSXSMSXSASMMSMSMMXMMAXXMMSXAMAMAMXMXASAAXAMAAAMXMMMSSSMMMMSMXMAXMAMAAMMSSSXSAAAAXXSASXMASAXMAXXMXMMMSMMMMAMMMMSXAMMMMM
MXMAXSMMSSXMASMSMSSMMMAMAMAMAMMMSAXXAMXAAMAMXMXXMMSAMAXAXXXMASMSXXMASMSXMAMAMSMAAAAXAXMAMMXSAAMAAMXSMMSAMAMXMASMSSSSXXXXAAAAAAXSMASAMMSMSXSA
MSSMSSMXMAXXAAAMAMXAAAAMAMXMASXAMXSAMXSXSMSMMSSMMAMXXXXSMXXXXMAXMMXXSASXMAXAMAMSSSMSMSSSMSAMMSMMXMAXAXMAMAMXMASAXAAXMXMSMSSSMXMAXXMAMXXAXAAM
AAASAMXASMMSXMXMXMSSMSXXAXXSAMMXSXSAMXMAXMMAAAMAXSSSXMAMAXSMMMXMASAXMMMMMMSSSMXAXAAXAAAAXMASXAAXXMAXAXXXMASASMMMMMMMMAMMMAMXMAMXMSSSMAMXMMMX
XSXMXMSASAMMXMMSAAAAXAASXMMMSXSXMASAMXMASAXMMSXSMXAAAMSMSMSAMXMSAMXXMSMAMAMMAMMSSMMMMMSMXSMMMSSMXMMMMSMSSMSASXASXMXXMAXAXXMASXSMMAAXASAMXMSM
MMAMXMMAXAAMXSASMMMSMMMMSAAAMXSAMASAMXMASMMXSAXMAMMMMMMAMXSSMXMMMSAMSAMSMAMSAMMAXXAAXXAXAXASAMXMXAAASAMAAAMAMXMMMSXMSSSMSSMXAMMAMMSMMMXXXMMS
ASAMAXMMSSMSAMXSXMSXAXMASMMMXASXMASASMMXSAAMAMAMSASASMXXSMMAMXSMSMMMSASXSXXSASMMXSSSSMSXMSMMASAMSMMMSASMMMMMMXXAXMAMAXAMAXMMSASAMXAAASMSMMAX
XSASMSXAAMXMASAMAXMSSMMXSXSXMXMMMASAMXSASMMSMXMSMMSASXAAXASAMAMXMASASXMAAMXSAMXXXXAAAXXAMAXMASXMASMASMMAXXAXMASMXSAMASMMSSXAXXSASMSSMSAAXMAS
ASAMXXMMSSXSAMASXMMAMXXXMASAMXAXMAMMMAMASAAMXMXAMXMMMMMMSMMSMASMSXMASAMXMMAMAMXXAXMXMMSXMASMXMASASMAMXSMMXMMSMMMAMAMMSAAXXMMSASAMXAAAMXMSMAS
XXAMXAAAXXXXXMXMXAMASMSMXMSAMMMMMAMAMASASMMSASMXSAXAAAAAXAAXSMSASXSASXMAXMSSMMSMMXSASXMASXSXASAMXSMXSASASMSAMAAMMSSMSXMASXSAMMMXMMSMMMAXXMAS
ASMMMSSSMXMAMMSAMMMAXMAMXXXAMMXAXAXMAAAASAASAXXAMMSSSSSSMMMMMXSMMAMXSASXSXMAMMAMXXSASASXMMMXXXMAXSAXMASASAMXMXMSXAMXMAMXSAMASXMASAXMASMSXMAS
MMAAAXAMAAMASAMASXMSXSASXMXXMASASXSXSXSASXMMAMSSSMAAAAMAMXXSXASAMXMASMMSMXSXMXMMSAMXMAXMAAAXMAMSAMXXSASMMMMSMXMSMXMAMAMAMXMAMASAMXMMAXMAXMMS
XSSMMMAMSMSAXXXAMXAAASASMXAASMSAAMAMXAMMMMXMAMAAAMMMMMMASAAMMMSXMAMASXXMXMMXMAMASAMSMAMSSMMSAMXMMMMMMASXAXMASAAXAASXSXMASAMAMXMASASMSSMXMAMX
AXAMSSMMAXMMMSMXSMMMXMAMAASMMAMXMSSSMMMAAAMMSMMXMMXASXSASMAXAMXXXASXXMXMAMMSSSMASXMMMXMAMMMMAMAXAAAAMXMXXXSASMSMMMXMAXSASMSSSXSAMMSAAAMSSSXS
SSSMAAAMXMSAAMMMXASMAMXMSMMAMXMSXMAAAASXMSXMAAXSMXMXXAMXSXAMMMMMSASXAMSSSSXAAAMAMMMAMXMAMSAMASXSSSSSXMASXXMAMAMXXXXMXMMASXAXMMMAXXMMMSMAAXAX
MAMMSSMXSASMXSAMSAMXAXXMMSMXMASXAMXMMMMAXAASMMMAMSMSMAMASMMMSAAXMMMMMMAAXAMMSMMASAMAXXMAMXAXXSAAAAAMXSASXSMSMMXSAMXSSSXAMMSSMMSMMMMXXMMMMMMM
MXMAXAAAMAMAMSAMMAMSXSASAAXAMXSAXMXSASXSMSXMASAMXSAMXXMASAAAXMSMSXMASXMSMMAXXXSXSMSSSSMSXSSMAMMMMMMMAMAMAMAAXXAXAMXXAMMMMXAAAAAASMSAMMMSMAAM
MAMSSMMMMAMXMMASXXMAASAMSMSMMAMAMMXSASAMXMASXMSXAMMMMMMXMMMMMXMAXXSASMMMXASXMXMASAAAAAAAMAMMMXSAAAAMAMMMMMSMSMXSMMSMSMAMMMSSMMXMSASASAAASMMS
SAMAAXXXSMSAXMAMMAMXMMMMAXAXAASAMXAMXMAMXSXMASMXMXAAAAMMSXAASXMMMMMAXAAAAMMAMXMAMMMMMMMMXMXMXASXSSSSSSMSAAAAXXMAMXXAAMMXSAXXMMSMMAMXMMSMSMAM
SSSXSMSXMASMXMASMXMASAXSSSMSSXMASXXSASAMXMASMMAASMSSSXSAAMSXSAMSAMMSMSMMSXSAMAMASMSXSXMSAMAMMMSAMXAAAAASMSMSXXAMXAMSMSAAMXMAXXAMMMMAMXXAMMAX
SASAAAXSMASAXSAMAASXXXMXAAAAMMSSMMMMASAXAMAMAMMMMAAMMAMASXMAMXMMAXXXAXAAMMSXSXSAXMXAXAASASASAAMAMMSSSMMMAXMMMXSAMXXXAMMSSSMSSSMMAAXAXAMAMSMS
SAMSMMMXMAMMXMASXMMXSSSMSMMMSAXAAAAMMMMMSMASAMASMMMSSXMXXXMAMAMXSMSMMMMMSAMMMMMXSMMSMMMSXMASMXMAMXAAAXXMMMSAAAXMAMMMSMXMAMAMXAXSSSSXSASXMAAX
MAMXXXXAMMSSMSAMAXASAMXAXMXXMASMSMSSMASXXXXXAXASAXSAAASXMXSMSMSAAAXAAAAXMXMAAAXAAXAAAXAMAMXMMMMAMXMSMMSASASMSSXMMASAMMAMXSSMSAMAAAAASXMASMSM
SSMSSSSXSMAAXMMSAMXMSXMSMMSMMAAAMAMXMAMMSMMSSMASXMMMMMMASMSAAAMMSXSXMSSSXAXSXSSSMSXXSMSSMMXSAMXSMAXAXXXAMXMMAMASAMMASMXXAMAMSASMMMMMMASAMMMA
AAAAAAAAMMMSASXMMSXAXMMMAASAMXSMMASMMXSAAAMAXMAMXXMAXASMMAMMMSMAMMMAXAMAXMMMAMXMASXMMAXAXAAMMSAXMASXSSMXMSXMASAMAXMXMAXMAXAMSXMAXXMXSXMMXXXX
SMMMMMMSMXMXASAMAMMSMAXSMMSAMAMXSASASMSXSSMASMAMMMSMMAMASXMSXXMAMASMMMXMMMAMAMAMAMAXSXSXMMMSAMMSXAXMAXAAMMAMAMASXMSASMSSMSXXXMXMMMSAMXMSMSMM
MMMMXXXMAMAMAMXMASXXSSMAMMSAMXSAMASXXAMAMXMAMMASAXAMXSSMMAMXMXMASXMMASAMAXAMSSMMASXMMMSASXAMASASMSMMSMSSSSSMASAMAXMASXMAMAMSAMSAMXMAXAMAMAAS
XMASXMMMSMAMSMXSXSAMAMXMSAMXSXMAMXMXMSMMMAMSMMXMMMASAAAXSXMAMAMXSAASASAMSSMMMAASAMXAAMXAMMMSAMMMAMXAAAXAAAAXXMAMMMMMMMMAMAXSAMXMASXSMMSASXSM
MMASAAAAAMAMAMMAMMMMMMMXMMSAMXMMSAAMAAAASXMXASMSXSAMMSMMMMSXSAMMMXMMAXAMMAXAMXMMAMSSMSMSMMXSXMXMAMSXMSMMMMMMSSSMMXSAASMMSMXSXMASXSAMXXAAXMAM
XMXMMMMXSMMSSXMAAAXSXMMAAXMASMAASMSASXXMMMXSSMAAMAXAXAXAXMAMMASMSAMXSSSXSAMSSMXSAMAMMSAMASAXMMAMAXMSMXMXSAMXXAXXMASXMMAAAXASMSAMMMMMMMMSMMSS
XXMMSMSAMAMXMASXSMMSAMSSSSSMMMMXXXMMAMXMSMAMAMXMMMSXXSMXMASXSXSMSASAAXMAMMMXAXASAMXSASASAMMSASXSXSMAMAMAMAMSMSMSMASMSSXMMMASAMXMAXMASAAMAXAS
MMXAAAMASXMSSMMAXMASAMAAXAASXMSMMMSSXSAAAMSSXMAMAXMAXMXSXXXAAAXASAMMXAMMMAMSAMXSXMSMAMAMXXAMXXXMASXMSASASAMAAAAAMMSMASAAXMAMMMSSSMSASMXSAMMS
AAMSSMMMMMXXAMXXMMMMAMMSMSAMXSAAASMSASMMMXAAMSMXMAMXAMAMXMMMMMMMMAXMXMAMSAMMSMAMAXAMXMSMMMSSMSXMXMAXMXMASMSMSMMMSAXMASMMXMAMXMAAMXMASAAMXXXS
MXMMAMXXMASMSMMAAASXMMXXAMAMXSMSMSAMMMXMMMMXXAAASXSMSMAMAMAMAMXXSMMSAMMASMSAASMXSAMXMAXAAAMAMXAMASAMXXMASAAAAAASMXSMMSASXSASAMMSMMXAMMMSMMAS
XMASMMAMXAXXMAXXSMXAXSMSMSASAMXAAMAMXXMASXMMXMMMXAXAAXMMAXSXSSSMSXASAXAMXAMAMAXAXMMAMMSXMSSXMSXMASASAMXAMXMSMSMSAMXAAXASASASXSAXAXMMSAMXXXAX
SXMXXXXAAMSMSSMMMASXMSAAASAMASXMMMAMXXMMAAAAMSSSMAMSMSMSMSMAMAMASMMSAMSXMXMMMXMASXSXSAAASAXAXMASMSAMAXMASXMAMXMMMMXXSMAMAMASMMMSMMSMAMSMMMSS
SAMXSMSMASXAAMASMAMXAMMMMMAMMMMAASMMMSXXMAMXSAAAXSMXXXAAXAMXMAMMMAXMAMXASXMXSXAAAAXAMASXMAMMMSAMMSXMAMMSMMSMSMMAAAMMMAMXXMAMAAAAXAAXMXSAMAAX
SAMAXAASXAMXMXXMMSMMXMAXAMAMXAMMMSAAAXAMSMXSMMSMMAAXSSSSXMMMXMMXSSMSSMSXMASXMASMMSMSMMMXMAMMAMXSMSMMXSXAXAAXXASMMSAAXSMXXMXSSMSSMSSSMAMAMSSM
SAMXMSMSXMSSSMMSAMXSSSSSSSMXMMXSXSXMMSXMASMMMMMAXMMMMAAMXSAMAXAXXXXAXXSASMMAXAXAAXXXAMSAMASMASAXAMMSASXMMSSSSXMXMAXMSAMXSAMXAAXXAAXXMASXMAXX
SXSAAMAMXXAAAAAMASMMAAXAMXXAASASXSASAXAMXXMASASAMXAAMMMMASASASMMMMMMMASAMXMXMASMMMMSMMMAMAAXAMMMMMAMMMAXAAMAMAXMAXSMSASMMASXSMMMXAMXSASAMAMX
XASMSMAMSMMSMMMSAMAMSMMSMASXSMASASAMMASMXMSAMMSAMSSSSSSMXSAMXMAAAXAASAMAMXXAMXXAAMXSXAMAMXSMMSXSXMXSAXAMMXXAMXSAMXSASASASAMAMXMAXMMMMAXAMSSS
MMMXAMAMAAXXAXXMXSXMXAAXMASMAMSMMMAMMAXXXXMASXSMMXAXAAMSAMXSASMSMXSAMXMAMSMXSXSSMSAMXMXXMMAXMAXMASMSXSSXSSSSXMAXASMAMAMMMAXAMSMXMAAAMMMSMMAM
XAASXMMXSSMSSMSMMXXMSMMSMAXMAMXAXXXMMSMMMMMXMAXAXMMMMMMMXSAMMSAAAXXMMSXMASXXAAXAAMASMMXSAMMMXSXSAMAXAXAAMXAMXAXMMAMXMMMXSAMMMSAASXSMSAAXASMS
XXXSMASAMAAAXMAAMXXXAMXSMMSSMSSMMSSMAMAXMASMSXSXMMXXXSAMAMXSSMMMMMXMASAMAMMMMSMMMSXMAAASAMXSAXAMXSSMMMMAMMMMMSMSASXMXXXAMASXAMMMSAAASMMSAMAS
SMMMMAMASMMMSXSMMAMSASAXMAMAAAXAAAXMAMXMAMSAAXSAMMAMXXAXSAMXXAAMMSMMASXMMXAAAAAXAMASMMMSXMXMASXMASAMAAXSSXSAAAASAMXAMMMXSAMMAMAAMXMMMMXMAMSM
XMAXMXSAMMSAMXAMMSXSAMASMXSMMXSMMSSMSMSAXXMMMMSAMAAXMSMMXAAXMSMSAAASXSXSSXSMSSSMMSAMAMAMMAMMMMXMXXAXSSSMAASMSSMMAMXMASMMMASMAXMSSXSAAXXSSMMM
XSMSAMMMMAMASAMXAXMXAMXXAAXAMAXXAMAAAAXMSXSASASAMXSSMAAMSMMXMAXMMSMMAXMAMAAAAAAAAMMSAMAXMAMXMMMSMSSMMXMMMMMMMMMXMASMSAMASAXXMSSMMMSSMXAXAASM
XAAMXXAAMMSMXMSMMSXSASAXXXMAMASMMSMMMSMAAAXAMASMMXMAMSSMSAASXMMAXMMMAMXSMMMMMSMMMMXSASMMSAXASAAMMAAAMAMXAAXAMAAASMMMXASMSAMSMMMAAAMAMSSSXMMS
MMXMMSSMMXAMAXXXXAMSAMMMSSSSXMXAXXXXMMXMMSMAMAMAMMSAMXAASXMMAMSMSMASAMAXAMXXMXASXMXSMMAASASAMMSSMMMMMSXMXMSMSMSXXAAXXAMXMAMXAASMMSMAMSASXMAM
ASXMXXMASMMSXSMSMSXMMMMAMMAXXXSAMMSXMMAXAXMXSMMXXAMAXSMMMAMSMMAAXMASAMXSAMMSMSXMAMMXAMMMMAAASXMXAMMSAXASMMAAMAMMSSMSSXSAMMMSMMSAMXMMMMAMSMXS
SMAXMXMAMAAXAXAAAMMSAMSSSMAMXXMASAAASXMMMSMAMAMMMXSSMMXMMMAAMSMSMMXSAMAMXMAAMMMSMMASAMXXMSMMMAAMMMAMASAMASXXMAMMAAMAAXSMSMAMXAXAXXXXAMAMAMMM
MSMMSSMASMMMMMSMSMAXAMAAMMAXAMSXMMXXMAAAAAMASAMAXMAMAMASAMSMXAMAAXMSAMMSAMSMSAAAAXXXAMMXMAMXSMMMSMMSASASXMMMSASMSXMXSMSMAMASMASMMSXXASXSASAS
XAXXXAXASXMXSAAXXXMSSMMSMSXSXMAMMXMSMSMMXMXASASXSAMSXMASAMAXSASMSMASAAMMMAXASMSMSMMMSSSSSXSAAXMXAAXMASAMMMAXXAMXAASXXASMSSMMXMXXAAMSMMASAMAS
SSMMSMMXMAXAMSXMMMXAMAXAAXAAMAAMSAMXAMXSASMXSAMAXMXSAMAMXAMMXMMAMXMSMMXMXMMAMXAMMAXXXAXAAASAMXXSXSMMXMMMASMSXMMMSMMAMSMAMMMXXXAMMMXASMMMSMAM
MAAAMXMAXMMMMMSMAAMMSSMMMMMMAMMMXASMAMAMAXXXMAMMMSAMAMAMAMXXMSMAMAXXXAAXAAXMXMMXXAMSMAMXMMMSXMXAAXXMAXAMXXAAAASAXAMXMAMXMAAXMMSSMMSASXMAXXAS
MSMMSASMXMAMAAAMSSXMAMXAAAXAMMSMXAMMXMMMSMMMMAMAAMASXMXXSXSXAAMMSMMMSMMXSSMSAASMMAMMAAAMXMXXMSMMXMASASXSMMSMMMMASMMAMMXSMMMAXAAAAMMAMAMXMSAS
XMAMSASXASASMSMMAMXMASXSSSSSSMAAMSSMMMSAMAMMXAXXMXMMMMMMAAMMSMSMAMMMSAXXAAAAMMAASMMXMASASXMMAXXAASXMMMMAMAAAAXMAMAMAMXXAAASMSMSSMMMMSSMSMMMX
MMAMMMMXMSASAXXMASMMMSAAAAXAAXMMMAAMXXMAMSASXSSSSXSAAAXMMMMMAAAMAMAASAMASMMSXMSAMAMXMAMXXAAMSMAMXMMMXSSSMSXSMSMMSSSMSSMMSMAAXMAMMMXMMMAAXAXA
MSSMMAAAMMMMAMASAMAAAMMMMMMSMMXSMMSMMMSAMXMXAAAAAASXSAXAAAAXMSMXSMMMSAMMAAAXXXMASXMXMAXAMMMMAAASAAXSAXMAXXAXXMAXAAAMAXAXAXMMMMASXMAAAMXMSMSM
MAXASMSXSAMSASMMASMMMSAXXXXMASMXMMMMAAASAXXSMMMMMMMXMXSXSSSSMXMXXASAMASXSMMMAXAAMXMAMXSSMMSSXMXSMSMMXSMMMMMMASMMMSMMXMSSMXXMAXXSAXSMMXSXXXAM
MASXMAXXSAMSASXSASXMASXXXMAMSMMASAMXMAXMXSASASXMXSMMMMSXAMAMXAMMMAMASXMAMSAMXMMMSASAMXAAAAMAMSMMMMAMXMAAMASMXMXAMMXSAAXAXSSSXSAMMMXAMXSMMXXX
MAMAMMMAMSMMAMAMXSAMASMSSMSMXASAXAXSXSXSAMASAMMSAAAXXAMMMMSXXSMAMMMXMAXAMAMXMXAAMXXAMMXMMMSMSAAAMXAMMSSMMASMAMSMSMASMSMAMXAAMMMSXAXAMAXAXSSS
MASMMAMXAMMMSMXMXSAMXSAMMAAAMAMMXSMMXMAMXMAMAMAMSSMMMSSMMAXMAMSSSXMXSSMXSASAMSMSSMMSXSXMMXSXSMSMSSMSXMASMASMMMAMAMXMXXMAMMMMSAAXMSXMMSSSMAAX
SMAASXSMMXAMMXAMXSXMXMAMMSMXMXMXAMAMMMAMSMAMMMMXAMAAAMAAMXMAXXAXAMMAMAAMAAMMAAAXMAAMASMXSAMMSAMAXXXMASAMMXMAMSXMASXSMXSAXXAXMMMSAAASAAAXMMMM
XMXMAAXXMASMSSMSAMSXMSAMXMASMMXMASAMXXAMMXXMAAXMASXMSMXMMASXMMSMMAMAMMMMMXMXSMSMMSXMAMAAMAMAMMMMMMAMXMMSMASAMMMMAMAXAXSXSSSSMASAMMXMASAMXXXX
XMMXMXMMSAXMXAXXMSMXAXMSXMSMMAXMASXMAMXXSASXSMXSAMASMMSAMAXAAAMASMSSSXSXXAAMXMAXAAAXMMMMSXMXSAAASMMSASMMMASASASMSMMMSAMXXXAAMXMAMMAAAXMSMSMM
XMAMSXSAMMSSMMMSMAMMSMAMAMAAXSXMASMMAMXAMMSAAAAMASMMAASAMSSSMMSAMAAAAASASAXSASXMSSSMSAXXMMSMMMXSAAXMASAMMMMAMAMAXMASMAMAMMSMMMSAMXAMXMAAAAAA
XMASAXMASAAXAAAASASAAMAMASMSMMAMASAMASMAMXMMMMMSAMXMMMSXMAAAMAMAMMMMMMMAMAMSASMMAAAASMSAMAAAAASAMXMASMMMAXMSMSMMMXMASAAAMXAMXAMMSAMXSMSMSMSX
SSMSAXSMMMSSSMSMSMSXSSMSASMMMSMMMSXMAXXAMMSAMXXMMSXXSAMXMMSMMASMMSAMXXMMMXMMAMMMMSMMMXSAMSSSMSSSXAASXXXSXSAXAMXSSXMASXSMMSASMXSASXMASMXXMAMA
"""


class TestDay:
    @pytest.mark.parametrize(testit.names, [(example, 18), (answer, 2618)], ids=testit.ids())
    def test_first_star(self, indata: str, want: int) -> None:
        testit.first_star(day04.Day(indata), want)

    @pytest.mark.parametrize(testit.names, [(example, 12), (answer, 2011)], ids=testit.ids())
    def test_second_star(self, indata: str, want: int) -> None:
        testit.second_star(day04.Day(indata), want)
