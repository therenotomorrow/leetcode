import pytest

from adventofcode.library import testit
from adventofcode.year2024 import day12

example = """
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
"""
answer = """
OORRQQQQQDUDDDDDDDDDDDDDDDDDDDYYDDDDPGGGGGGGGGGGNNNNNNNTTTTTTTTTTYYYHKKKKKKKKKKXXXXXXXXXXXXYEEEEEEEEEEEEVVAAAAAAAAAAAAAAAATERREVEEEEEEEEEEEE
OOORQBQQQDDDDDDDDDDDDDDDDDDDDDDDDDDDDGGGGGGGGGGGGNNNNNNTTTTTTTTTTTTYHKKKKKKKKKKXXXXXXXXXXXEEEEEEEEEEEEEEEEEEAAAAAAAAAAAAAAREEEEEEEEEEEEEEEEE
OOOOOOOQODDDDDDDDDDDDDDDDDDDDDDDDDDGGGGGGGGGGGGGGNNNNNTTTTTTTTTTTTTHHHKKKKKKKKKKXXXXXXXXXEEEEEEEEEEEEEEEEEEEAAAAAAAAAAAAAARRREEEEEEEEEEEEEEE
OOOOOOOOOVVDDDDDDDDDDDDDDDDDDYYYYDDGGGGGGGGGGGGGNNNNNNTTTTTTTTTTHHHHHHKKKHHHHKKKXXXXXXXXXXXXEEEEEEEEEEEEEEEEAAAAAAAAAAAAAARREEEEEEEEEEEEEEEE
OOOOOOOOVVVDDDDDDDDDDDDDDDDDDYYYYYYYGGGGGGGCCCCNNNNNNNNTTUTTTTTTHHHHHKKKKHHHHHXXXXXXXXXXXXXXEREEEEEEEEEEEEEEEAAAAAAAAAAARRRRRREERRREEEEEEEEE
OOOOOOOODDDDDDDDDDDDDDDLDDDDDBYYYYYYYYYGGCCCCCCCNCCCNNNTDUUUTTTTRHHHHHHHHHHHHHXAAAAAXXXXXSSXXEEEEEEEEEENEEEEAAAAAAAAAAAAAARRRRRRRRRREEEEEEEE
OOOOOOOWWWWWWWDDDDDDDLLLDFFDBBYYYYYYYYYJCCCCCCCCCCCCNNNIIUUUTTRRRRRHHHHHHHHHHHAAAAAAAXXXSSSSSRREEEEEEEENEEAAAAKKKAAAAAAAAARRRRRRRRRREEEEEEEE
OOOOOOOOWWWWWDDDDDDLLLLLFFCBBYYYYYYYAAACCCCCCCCCCCCCIINIIUUUKTRRRRRSHXXXHHHHHUAAAAAXXXXXSSSSSREEREEEEEPAAAAKKAKKKAAAAAAARRRRRRRRRRREEEEEEEEE
OOOOOOWWWWWWWDDDDDDLLLLLFFCCBYYYYYYYACCCCCCCCCCCCCCIIIIIIIKKKKRRRRRRHXXXHHHXAAAAAAAAXXXSFSSSSSQQEEEECEEEVKKKKKKKKKKAAAAARRRRRRRRRRREEEEEEEEJ
OOOOOWWOOWWWWWDDDDOLILLLFLLCCYYYYYYYACCCCCCCCCCCCCCCIIIIIIIKKRRRRRRNRXXXXXXXXAAAAAAAXSSSSSVVSSQQEEEQCEEVVKKKKKKKKKXAAAAARRRRRRRRRFREEEEEEEEE
OOOOOOOOJJWWWJIIIDLLLLLLLLLCCCYYYYYYACCCCCCCCCCCCCCIIIIIIIIKKRRRRRRRRRXRXXXXXAAAAAJJJSSSSSVQQQQQQQEQQVVVVKKKKKKKKKKEKRRRRRRRRRRREEEEEEEEEEEE
OOOOOOOJJJJJJJJIIIIILLLLLLLCCCYYYCCCCUFFCCCCCCCCCCCIIIIIIIIIKRRRRRRRRRRRXXXXXAAAAAASSSSSSSVVVQQQQQQQQVVVVVKKKKKKKKKKKARRRRRRRRRPEEEEEEEEKKKK
OOJJOOOJJJJJJJIIIIIILLLLLLLLCCCCCCCCUUFUCCCCCCCCPCCIIIIIIIIRRRRRRRRRRXXXXXXXXAARRRSSSSSSSSSSQQQQQQQYQVHVVVVKKKKKKKKKKKKKKRRRRRRPPDDKEKEEKKKK
OOOJJJJJJJJJJJIIEEEILLLLLLLLLLCCCCCCUUFUCCCCCPCCPPPPIIIIITTTRRRRRRRRRXXXXXXJJJJRRRSSSSSSSSSSEQQQQQQQQVHHHHHKKKKKKKKKTKKKTRRRPPRDDDDKKKKEEKKK
OOJJJJJJJJJJJJEEEEEEELLLLLLLLLCCCCCUUUUUCCCCCPPPPKIIIIIIIIRRRRRRRRRRRRRRXXJJJJJRRRASSSSSSSSSEQQQQQHHVVHHHHHKKKKKKKKTTKTTTTPPPPDDDDDKKKKKKKKK
JJJJJJJJJJJJJEEPEEEEPLLLLLLLLLLCCCCCCUCUUSSSKKPKKKIIIIIIIIIJJRRRRRRRRRRXXXJJJJJRRRAASSSSSSSEEQQPQQHHHHHHHHKKKKKKKKTTTJTTTTTPPACDAADKKKKKKKKK
JJJJJJJJJJJJJEEEEEEEPLLLLLLLLLLCCCCCCCCCKKKKKKKKKKIIIIIIBIIIRRLRRRRRXXXXXXJJJRRRRRRRRRSHHHPPPPPPPHHHHHHHHKKKKKKKKKTTTTTTTTTPPAADAAAKKKKKKKKK
JJJJJJJJJJJJEEEEEEEEEXLLLLLLCCCCCCCCCCCKKKKKKKKKKKKKIIIIBIIILRLLRRRRXXXXXJJJRRRRRRRRRRSSHHPPPPPPPPHHHHHHHHKKKKKKKKTTTTTTTTTPPAAAAAAKKKKKKKKK
JJJJJJJJJJEEEEEEEEEEXXLLLLCCCCCCGCCCCCCKKKKKKKKKKKKKKKIIBBILLLLLRRRRPPPPXXJJRRRRRRRRRRSSNNNPPPPPPPPHHHHHHKKKKKKKKKKTTTTTTTTAAAAAAAAKKKKKKKKK
JJJJJJJJJJJJEEEEEEEEEXXXXCCCCCCCGCEECCCKKKKKKKKKKKKKKKIBBBBBBBRRRRRRRPPPPXXXRRRRRRRRRRSNNNNPPPPPPPPHHHHHKKKKKKKKKKKTTTTTTTTTAAAAAAAKKKKKKKKZ
JJJJJJJJJJJFEEEEEEEEXXXXXXXCCCCCGEEEECCGGHKKKKKKKKKKKKKBBBBBPPPPPPPPPPPPPPQQRRRRRRRRRRSNNNCCCPPPPPPIILHKKKKKKKKKKTTTTTTTTTTTAAAAAAAJKKKKKKKK
JJPJJJJJJJJEEEEEEEEEXXXXXXXCCCCCGGEEECGGGHKKKKKKKKKKKKKBBBBBBPPPPPPPPPPPPPQQRRRRRRRRRRVVNNNCCPPPIIIIIIOOKKKKFFKKFTTTTTTTTTTTAAAAAAKKKKKKKKKK
JPPPPPPPPPJEEGGEEEEEXXXXXCCCCCCGGGGGGGGGGGPKKKKKKKKKKKKKBBPPPPPPPPPPPPPPPBQQRRRRRRRRRRVVNNCCPPPPPIIIIIIIKYKKFFFFFTTTTTTTFTVVVVAAVAKVKKKKKKKK
PPMPPPPPPPGGGGGGGGGXXXXXXXXCCCCNNGGGGGGGLLPKKKKKKKKKKKKKBPPPPPPPPPPPPPPXXQQQRRRRRRRRRRAZNCCCCPPPCIIIIIIIIIFFZFFFFTTTTTTTTVVVVVVVVVVVKKKKKKKK
YPPPPPPPGGGGGGGGGGGXXXXXXNNCCCNNNNNGGGGGLAKKKLKKKKKKKKKKBPPPPPPPPPPPPPPXXQQQRRRRQZZZZZZZZCCCPPPCCIIIIIIIIPFFFFFFFFTTTTTTVVVVVVVVVVVVKKKKKKKK
YPPPPPPPPGGGGGGLLLLLLXXNNNNNNNNNNNNGGGLGLLLKKLKLKOKKKKKKBPPPPPPPPPPPPPXXXQQQRRRRQZZZZZZZZCCCCCPCIIIIIIIIIFPFFFFFFFMTTTTUVVUVVVVVVVVVKKKKKKKK
PPPPPPPPPGGGOGGGGGLLLXXXNNNNNNNNNNNBGGLLLLLKKLLLKOKKKKKPPPPPPPPPPPPPPPXXXXXQQQQQQQQZZZZZCCCICCCCCIIILLLIRFFFFFFFFFFFFUUUUUUVVVVVVVVVVKFKKKKK
PPPPPPPPPPOOOOGGGGLLLLXXNNNNNNNNNNNBGULLLLLLLLLKKKKKKKKPPPPPPPPPPPPPPPPXXXXXXQQQQNZZZZZZZCIICCCCCIYILRLRRRRFFFFFFFFFFUUUUUUVWWVVVWVVVKKKKOOO
PPPPPPPPPPOOOOGGGLLLLLLLNNLNNNNNNNNBNUULLLLLLLLLLLSKKSKKKPPPPFPFPPPPXXXXXXXXQQQQQNNNZCCZEIIICCICCRRLLRRRRRFFFFFFFFFFFUUUUUUUWWWWWWWWVVVKKXOX
PPPQPPPPPPOOOOGGGLLLLMMLLLLLLLNNNNNNNNUULLLLLLLLLSSSSSKKKPPPPFPFNPPPXXXXXXXXXQQNNNNFCCCIIIIIIIIICRRRRRRRRFFFFFFFFUUUUUUUUUUWWWWWWWWVVVVVVXXX
PPPPPPPPCWOOOOCGGGLLLLLLLLLLNNNNNNNNNNULLLLLLLLLLLSSSKKSKKPFPFPFFFXXXTXTTTXXXXQFFFFFFCIIIIIIIIIICRRRRRRRRFFFFFYFFUHHHUUUUUUUWWWWWWWXVVVVVXXX
PPKKPPPPCCCOCCCCCGLLLLLLLLLLLNNNNNNNJCLLLLLLLLLLLSSSSKSSKKKFFFFUFFQQQTXTTTTTTXGFFFFFFCCIIIIIIIITRRRRRRRRRRFFYFYYYYHHHHHHHHHUWWWWWWWXXXXXVXXX
CPKKCCPPCCCCCCCCCGLLLLLLLLLLLNJNNNNNJCLLLLLLLLLLLSSSSSSSSSFFFFSFFFFQQTTTTTTGGGGGFFFFFCCIIIIIIIITRRRRRRRRRRFYYYYYYYHHHHHHHHHUUWWWWLLXXXXXXXXX
CCCCCCPCCCCCCCCCCLLLLLSKLLLLJJJJJNNNJJXLLLLLLLLLSSSSSSSSSKKFFFFFFFFTTTTTTTTTTGUGUUFFFFCTTTTTTTTTTRRRRRRRRRRYYYYYYYHHHHHHHHHUUWWWWLLXXXXXXXXX
CCCCCCPCCCCCCCCAALSSSSSKLLJJJJJJJJJJJJJLDLLLLLLLFFSFSSSSSFFFFFFFFFKTTTTTTTTTTUUUUFFFFFCTTTTTTTTTTRRRRRRRRYYYYYYYYYHHHHHHHHHUUUWDWLLLXXXXXXXX
CCCCCCCCCCCCCCCCCLSSSSSKKSJJJJJJJJJJJJDDDDOLOCLFFFFFSSESFFFFFFFFFFFFTTTTTTTTHHUUUFFBBBBBBBBBBTTTTRRRRRRYYYYYYYYYYYHHHHHHHHHUUUXULLLLXXXXXXXX
CCCCCCCCCCCCCCCCSSSSSSSKSSSJJJJJJJJJJJJDDOOOOOVOFFFFSSFFFFFFFFFFFFFTTTTTTTTTTHUUUFFBBBBBBBBBBTSSSRRRRRRBYYYYYYYYYYHHHHHHHHHUUUUULLLLXXXXXXXX
CCCCCCCCCCCCCCCSSSSSSSSSSSSJJJJJJJJJJJDDDOOOOOOOOFYYYYYYYYYYFFFFFFFOOTTTTTTTTHUUUUFBBBBBBBBBBTSSSSSRRRSYYYYYYYYYYYHHHHHHHHHUUUUUDLLLLXXXXXXX
CCFCCCCCCCCCCCCSSSSSSSSSSSSSJJJJJJJJJJDDDDOOOOOOOOYYYYYYYYYYYFFFRRRFFFGTTTTTHHHHBBBBBBBBBBBBBYTTSSSSRSSYYYYYYYYYYYHHHHHHHHHUUUUSDLLLLXXXXXXX
CFFFCFFFCCCCCCCSSSSSSSSSSSSJJWWJWJJJJDDDDOOOOOOOOOYYYYYYYYYYYOOFFRFFFFGTSSTTMHHMBBBBBBBBBBBBBTTSSAASSSSYYYYYYYYYYYHHHHHHHHHUUUSSDLDDDDXXXXXX
FFFFCFFFFCCCCCCCSSSSSSSSSSZWWWJJWWJJDDDDDDDOOOOOOOYYYYYYYYYYYOOFFRRFFFFFFSSSMMMMBBBBBBBBTFTTTTTTAAASSSYYYYYYYYYYVVYYAAAADUUUUUUDDDDPDDIIXXXX
FFFFFFFFFCCCCSSCCSSSSSSSSWZWWWJJWWWWWDDDDDOOOOOOOYYYYYYYYYYYYOOFFRFFFFFFFSSSMMMBBBBBBBBBFFTTTAAAAAASAAAAFYYYYYYYVVVVAAAADUUUDDDDDDDDDDDIIXEX
FFFFFFFFFFFCVSVSSSVSSSSSQWWWWTTJWWWWWDDDDDOOOOOOOOYYYYYYYYYYYOOOFRFFFFFFFFSSMMMBBBBBBBBBQQTATAAAAAAAAAAAFFFFYYYYVVVVAAAAVUTUDDDDDDDDDDDKKEEE
FFFFFFFFFFFFVVVVVVVVVSSQQWWWWWWWWWWWWWWDDDDOOOOOOOYYYYYYYYYYYOOOOOFFFFFFFLMMMMMBBBBBBBBBQQOAAAAAAAAAAAAAFFFYYYYYVVVVVAVVVUBBBBBDDDDDDDDKZIIE
FFFFFFFFFFVVVVVVVVVVVVWWWWWWWWWWWWWWWWDDDDEEOOOOOEEYYYYYYYYYYOOOOFFFFFFFFMMMMMMBBBBBBBBBOOOAAAAAAAAAAAAAAFYYYYXYVVVVVVVVVVBBBBDDDDDDDDDKKKII
FFFWWFWFFVVVVVVVVVVVBWWWWWWWWWWWWWWWWWDDDDEEEEEYYYYYYYYYYYYYYOOOFFFFFFFFMMMMMMMMBBBBBBBBBBBAAAAAAAAAAAAAAYYYYYYVVVVVVVVVVVVVBBBBDDDDDDKKKKII
FFWWWWWWWWWWVVVVVVVBBWWWWWWWWWWWWWWWWWDDDDEEEEEYYYYYYYYYYYHHTHOOFFFFFFFFMMMMMMMMBBBBBBBBBBBAAAAAAAAAAAVYYYYYYYAVVVVVVVVVVVVVBBBBDDDDDDKKKKII
FFWWWWWWWWWWVVVVVVVVVWWWWFFWFFWWWWWWWWWXXDEEEEEYYYYYYYYYHHHHHHHFFFFFFFFFMMMMMMMMBBBBBBBBBBBAAAAAAAAAAAVYYYAAYAAVVVVVVVVVVVVVVBBDDDDDDKKKKKII
FFFWWWWWWWWWVVVVVVVVQQQQWFFFFWWWWWWWXWXXXDEEEEEEEEEYYYEEHHHHHHAFFFFFFFFFLMMMMMKKBBBBBBBBBBBAAAAAAAAAAAVVYYYAAAAJVJVVVVVVVVVVVDDDDDDIIKKKKKII
FFFFWWWWWWWWWVWWWWVVQQQQQQFFFFFWWWWXXXXXXXXEEEEEEELYYYHHHHHHHHHGHHFFLFFFLLLLMMKKBBBBBBBBBBBAAAAAAAAAAAAAYYYYAAJJJJJCVVVVVVVOODDDDDDIIIIIKKII
FFFFFWWWWWWWWWWWWWQQQQQQQQFFFFWWWWWXXXXXEXEEEEEEEELYYYHHHHHHHHHGHHHFLLLLLLLLLKKKKKKKKBBBBBBAAAAAAAAAAAAYYYYYAAJJJJJJVVJVVVVOOODDDDIIIIIIIIII
FFFFFWWWWWWWWWWWQQQQQQQQQQFFFFWJWPXXXXXXEEEEEEEEEEEYYYHHHHHHHHHHHLLLLLLLLLLLLKKKKKKKKBBBBBBAAAAAAAAAAAYYYYYYAAJJJJJJJJJOOOOOOODDDIIIIIIIIIII
FFFFFWWWWWWWWWWWQQQQQQQQQQFFFFPPPPXXXXXXEEEEEEEEEEEYYYHHHHHHHHHHHLLLLLLLLLLLLLAKKKKKKBBBBBBQQAYYYAAYAYYYYYYYYAAAJJJJJJJJJOOOOODDOIOIIIIIIIII
FFFWWWWWWWWWWWWWQQQQQSQQQFFPPPPPPPPXXXXXXXEEEEEEHHHHHHHHHHHHHHHHHLLLLLLLLLLLLLLUKKKKKBBBBBBQQQYYYYAYYYYYYYYYCYTTTJJJJJJJJOOOOOOOOOOIIIIIIIII
FFFFFWWWWWWWWWWWWWSQVQQQQQFFFPPPPPXXXXXXXXEEEEEEHHHHHHHHHHHHHHHLLLLLLLLLLLLLLLUUUUKKKBBBBBBQQQYYYYYYYYYYYYYTTTTTTJJJJJJJOOOOOOOOOIIIIIIIIIII
FFFFFFWWWWGWWWWWWWQQKKQQQFFFFFPPPMXXXSXXSSSSEEEEEHHHHHHHHHHHHHHHHLLLLLLLLLLLLLUUKKKKKKKKBBBQQQQQYYYYYYYYYYYTTTTTTTTTJJJJJOOOOOOOOOIIIIIIIIII
FFFFFFWWWWGWWWWWWKKKKKKQFFFFFFPSSSSXXSXSSSSCEECEEHHHHKHHHHHHNHKHHLLLLLLLLLLUUUUKKKMKKKKQBBBQQQQYYYYYYYYYYYYTTTTTTTTTJJJJJJOOOOOOOOIIIIIIIIII
FFFFFFFWWWWWWWWWKKKKKKKKKKFFFFFSSSSSSSSSSSCCQCCCCHHHHKKKKJKKKKKEEELLLLLLLLUUUUUQKQQQQQQQBBBQQQQYYYYYYYYYYYYTTTTTTTTTJJJJJJJOJJZZOOIIIIIIIIII
FFFFFFFSSWWWKKKKKKKKKKKKKFFFFFSSSSSSSSSSSCCCCCCCCHHKKKKKKKKKKKKKKEELLLDLLLLUUUQQQQQQQQQQBBBQQQQQYYYYYYYYYJJTTTTTTTTTJJJJJJJJJJZZZZZIIIIIIIII
FFFFFTFSSSSSSKKKKKKKKKKDKFFFFFSSSSSSSSSSSSCCCCCCCCCCCKKKKKKKKKKKKEELLDDDZLZUUUQQQQQQQQQQBBBAQQQQYYYYYYYYYJTTTTTTTTTTJJJJJJJJJJJZZHFIIIIIIIII
FFFFFFFQSSSSSKKKKKKKKKKKFFFFFSSSSSSSSSSSSSCCCCCCCCCCKKKKKKKKKKKKKEEEEZDDZZZUQQQQQQQQQQQQBBBAAAAQYYYYYYYYYJTTTTTTTTTTJJJJLJJJVVVZZHFFFIIIIIII
FFFFFFFQQSSSSKKKKKKKKKFFFFFFFFSSSSSSSSSSSSCCCCCCCCCNKKKKKKKKKKKKKEEEEZZZZZUUQQQQQQQQQQQQBBBAAAAAYYYYYYYNJJTTTTTTTTTTJJJJJVJVVFFFFFFFYYIIIIIC
FFFFFFFQQQQSSSSKKKKKKKFFFFFFFFFFSSSSSSSSSSSCCCCCCCNNKKKKKKKKKKKKKKKEEZZZZZUZJQQQQQQQQQQQBBBAAAAAAYAAAYJJJJTTTTTTTTTTJJJJJVJVVFFFFFYYYYIIIIIC
FFFFFFFQQQQSSSSSSKKKKKKFFFFFFFSSSSSSSSSSSSCCCNCCCNNNKKKKKKKKKKKKKKKEZZZZZZZZZQQQQQQQQQQQBBBAAAAAAAAAAYFFJJTTTTTTTTTTVVVVJVVVFFFFYYYYYYIIICCC
FFFFFQQQQQQQQSSSKKKKKKFFFFFFFFFFFSSSSSSSSSSCCNCCCCNNKKKKKKKKKKKKKKKKKZZZZZZZZZAAQQQAAAQYAAAAAAAAAAAAAYFFFJTTTTTTTTTTJVVVJVVVFFFFFYYYYYYYYCCC
FFQQQQQQQQQQQSSSKEKKKKKFFFFFFFFFCSSSSSSSSSSCCCPCPMNNKKKKKKKKKKKKKLKNNZZZZZZZZZAAAAAAAQQYAAAAAAAAAAAAAFFJJJTTTTTTTSSHHVVVVVVFFFFFYYYYYYYCCCCC
QQQQQQQQQQQQQVVVKKKKVKKFFFFFFFFFCYSSSSSSSSSSPPPPPMMMMKKKKKKKKKKKLLZZZZZZZZZZZZAAAAAAAAAAAAAAAAAAAAAAAFFJJJTTTTTTTHHHHVVVVVFFFFFFFYYYYYYYCCCC
QQQQQQQQQQQQVVVVVVVVVVVVVVFFFFFFYYSSSSSSSSSSPPPPPMMMMMMMKKKKKKILLLLZZZZZZZZZZAAAAAAAAAAAAAAAAAAAAAAFFFJJJJTTTTTTTHHHHVVLVVVFFFFFFYYYYYYYCYCC
QQQQQQQQQQQQQQVVVVVVVVVVVTFFFFFFYYYVSYYYYYYPPPPMMMMMMMMMKKIIIIIJLLLZZZZZZZZZZZAAAAAAAAAAAAAAAAAAAAAAAAAJJJTTTTTTTJYHMLLLLLLLSFOFPPYYYYYYYYOC
QQQQQQQQQQQQQQVVVVVVVVVVVAFFFFFYYYYYYYYYYYYYPPPMMMMMMMMMMIIIIIIIGGLLZZZZZZZZZFFFFAAAAAAAAAAAOOAAAAAAAAYYJJYYYYJJJJYLLLLLLLLLSSCCCPPPYYYVVVOO
QQQQQQQQQQQQQQVVVVVVVVUVVAFFSFFYYYYYYYYYYYPPPPPPMMMMMMMMIIIIIIIIILLLAZZZZZZZZFFFFFFFAAAAAAAOOOOAAAAAAYYYJJYYYYYJJYYYYLLLLLLLSCVCECPPPOVVVOOM
QQQQQQQQQQQQQXVVVVVVVVUVAAAAAFRYYYYYYYYYYYPPPPPPMMMMMMMMIIIIIIIIIALAAZZZZZFSFFFFFFFAAAAEEAAOOOOOOAAAAAAAWWYYYYYYYYYLLLLLLLLSSCCCCCCPPOOOOOOO
QQQQQQQQQQQXXXXVVVVVVVVXAAAAAYYYYYYYYYYYYPPPPPPPMMMMMMMMIIIIIIIIIALAAAAZFZFFFFFFFFFFAANEENROOOOOOOAAKAAAYYYYYYYYYYYYLLLLLLGGCCCCCCCPPOOOOOOO
QQQQQQQQQXQXXXVVVVVVVVXXAAAAAYYYYYYYYYYYPPPPPPPPPPMMMMMMMIIIIIIIIAAAAAFFFFFFFFFFFFFFAANNENNNLOOOOOOKKYYAYYYYYYYYYYYYYYLLLLGGGCCCCCCPOOOOOOOO
XXQQQQQQQXXXXVVVVVVVVVXXXAUAJJYYYYYYYYYYYPPPJPPPPPMMMMMMMIIIIIIAAAAAAAAFFFFFFFFFFFFFNNNEENNNNOOOORRKKYYYYYYYYYYYYYYYYYLLLLGGGGSCCCCOOOOOOOOO
XXXQQQXQQXXXXVVVVVEEEEXXXAUJJJJYYJYYYYYYPPPJJJJPPPMMMMMMMMIIIIIAARRAAAFFFFFFFFFFFFFFNNNNNNNNOROOOORRKKYYYYYYYYYYYYYYYYYLLAGGGASSSCOOOOOOOOOO
XXXQXXXXXXXXXXVVVVVVEUUUUXUJJUJJJJJJYYYYYJJJJJJJPPMMMMMMMIIAAIIAARRRAAAFFFFFFFFFFFFNNNNNNNNOOOOOOORRRKKYYYYYYYYYYYYYYLLLLAAAAAAOOHHOOOOOOOOO
XXXXXXXXXXXXXDZZZDVVEUUUUUUUUUJJJJJJJYYYYJJJJJJJPPPPMMSMMMAAAAAAARRAAAFFFFFFFFFFFFNNNNNNNNNOOOOZZORRRRRRRNNNNYYYYYYYLLLLAAAAAAAOOOOOOOOOOOOO
XXXXXXXXXDDDDDZZZDDDEEUUUUUUUJJJJJJJYYYYJJJJJJJJPPPPMMMAAAAAARRRRRRRRAAAFFFFFFFFFBBNNNNNNNNONBRZRRRRRRRRRRRNYYYYDYYDLLLAGAAAAAAAAAAOOOOOOOOO
XXXXXXXXXDXDDDDZZDDDEEEUUUUUUUJJJJJJJJJYYJJJJJJJPPPPAAAAAAAAAAARRRRRRRARRRFFFFFBBBBNNNNNNNNNNRRRRRRRRRRRRRRRRDDDDDDDAAAAAAAAAAAAAAABOOOOOOOO
XXXXXXXXXXXDDDDDZDDDEEEUUUUUUUJJUJJJJJJJJJJJJJPPPPPPAPAAAAAAAARRRRRRRRRRKRRFFFFXBBXXNNNNNNNNRRRRRRRRRRRRRRRRRDDDDDDDBBAAAAAAAAAAAABBOOOOOOOX
XXXXXXXXXXDDDDDDDDDDEEDZUUUUUUJJUJJJJJPJJJJJJJJPQPPPPPAAAAAAAAHRRRRRRRRRKRRFFFXXXXXXXNNNNNNNNNLRRRRRRRRRRRRDDDDDDDADDBBBAAAAAAAAAAAOOHOOLLOO
XXXXXXXXXXDDDDDDDDDDDDDDUUUUUUUUUJJPJPPJJJJJJJPPQQQPPOOOAAAARRRRRRRRRRRRRRRRXXXXXXXXXXXXXXXXXXRRRRRRRRRRRRDDDDDDDDDBBBBBBBAAAAAAAAAAOOOLLLZO
XXXXXXXXXXDDDDDDDDDDDDDDEEEEUJUUJJJPPDPPJJJPJPPQQQQPPOOOOAATRRRRRRRRRRRRRRRRXXXRXXXXXXXXXXXXXXRRRRRRRRRRRRLRDDDDDDBBBBDDBAAAAAAAAAAAALLLLLZZ
XXXXXSXSXXDDDDDDDDDDDDDDEEEEEEEEEEEEPPPPPPPPPPPQQQQQQWOOOAAARTRRRRRRRRRRRRRRRRRRXXXXXXXXXXXXXXBRRRRRRRRRRRLRRDDDDDBBBBDDAAAAAAAAAUAARLLLLLZZ
XXXSSSSSSSWDDDDDDDDDDDDDEEEEEEEEEEEEPPPPPPPPPPPQQQQQQQQOORAAARRRRRRRRRRRRRRRRRRRXXXXXXXXXXXXXXRRRRRRRRRRWRRRKDDBDBBDBDDDDDAAAAAAAAAAALLLLZZZ
XXXSSSSSSSDDDDDDDDDDDDDDEEEEEEEEEEEEEEEEEPPPPPQQQQQQQQQQORRAARRRYRRRRRHHHHHRRRRRXXXXXXXXXXXXXXTRRRRRRRXRWRRVVVBBBBBDDDDDDDDDAAAAAAABLLLLZZZZ
XXXXSSSSSSSSDDDDDDDDDDDDEEEEEEEEEEEEEEEEEPPTTTQQQQQQQQQQORRAAAARYRRRHHHHHHHHHXRRXXXXXXXXXXXXXXTVRVRRVRVVWVVVVYYYBBBDDDDDDDDDDAAAAAABLLLZZZZZ
EEEEESESSSSSDDDYDDDDDDDDDDEEEEEEEEEEEEEEEPTTTTQQQQQQQQQOORRAAAARRRRRHHHHHHHHHHRRRXXXXXRXXXXXXXTVVVVVVVVVVVVVVYYYBBBDDDDDDDDDDADAJAAJJZZZZZZZ
EEWEESESSSSSSCDDDDDDDYZZDDJEEEEEEEEEPPPPPPHTTTTQQQQQQQQOOORAAAARRRPHHHHHHHHHHHRRRRRRRRXXXXXXXXVVVVVVVVVVVVVVYYYYBBBDDDDDDDDDDDDJJJJJJZZZZZZZ
EEEEEEEESSSSSCCDDQYYYYYZZZZEEEEEEEEEPPPPPHPTTTTMQKQQQIQQOOOAAAARPHHHHHHHHHHHHHRRRRRRRRXXXXXXXXVVVVVVVVVIIVIIYBBYBBBBBDDDDDDDDDDDJJJJJJJZZZZZ
EEEEEEEESSSSSSCDDDIYYYZZZZZEEEEEEEEEPPPPPPPPTTTMQRRRXIBOOOOOOOORPPHHHHHHHHHHHHRRRRRRRRXXXXXXXXVVVVVVVVIIIIIIIBBBBBBBBOODDDDDDDJJJJJJJJZZZZZZ
EEEEEEEESJSSSJJJJDDYYYZLZZZEEEEEEEEEPPPPPPPXTTTTQRRIIIOOOOOPPPOOPHHHHHHHHHHHHHRRRRRRRRXXXXKKVVVVVVVVIIIIIIIIBBBBOBBBBOOODDDDDDJJJJJJJZZZZZPP
EEEEEEEEEJSSJJJJJJHYYPZZZZTEEEEEEEEEPPPPPPPXRRBRRRRRIIIIOODTTPPPPHHHHHHHHHHHHHRRRRRRRRXXXXXKKVVVVVVVVIIIIIIIIIBOOOOOOOOOODDDGDJJJJJJJZZZZPPP
EEEEEEEEEJJJJJJJJJYYYYZZZZZTTTTEEEPPCCCRRPPPRRRRRRRRIIIITTTTCTTPPHHHHHHHHHUUMMUUUKKKKXXXXXKKVVVVVVVVVVIIIIIIIIICOOOOOODDDDDDGRRRJJJJJJPPPPPP
PEEEEAEEEJJJJJJJJJYYTTZZZZZZTTTEEEPPCCRRRPRRRRRRRRRRIIITTTTTTTTPPHHHHHHHHHUUUUUUUUUKXXXXXKKKKVVVVVVIIIIIIIIIIIIIOOOOOOODDDDGGARRJJJJJJPPPPPP
PEEEEAAXXJJXJJJJJJJJTZZZZZZZZZTEEEECCCCARRRRRRRRRRRRIIITTTTTTTTTPHHHHHHHHHUDDUUUUUKKKXXXKKKKVVVVVVVIIIIIIIIIIIIOOOOOOOOOODDGGARRJRRJPPPPPPPP
PPEEEEEXIXXXJJJJNNNJJZZZZZZZEEEEEEEECCCARRRRRRRRRRRRIIIITTTTTTTTPPPTTTHHHHDDDUUUUUUKKKKKKHHHHHHHVVVIIIIIIIIIIIOOOOOOOOOOODDGGGRRRRRJPPPPPPPP
PEEEEEXXXXXXJNNNNNNZZZZZZZZEEEEEEEEEAACARRRRRRRRRRNNIIIIITTTTTTPPPPPTTTHHDDDDUUUUUKKKKKKKKHHHHHHHHVIIIGICIIIOOOOOOOOOOOOODDCRRRRRRRJPPPPPPPP
PXXXXXXXXXYYNNNNNNNZZZZZZZEEEEEEEEEEAAAAAARRRRRRRNNNIIIIITTTTTTPPPPTTTTHHDDDDDUUUKKKKKKKKHHHHHHHWWWWWCCCCCIIIOOOOOOOOOOODDDDRRRRRRRRRPPPPPPP
PXXXXXXXXXYYYNNNNNNNZZZZZZEEEEEEEEEAAAAAARRRRRRRNNNNIIIIITTTTTTPPPPTTTTHHHDDDDUUUKKKKKKKKKHHHHHHHCCCCCCCCCIOOOOOOOOOOOOOODDDRRRRRRRRRLPPPPPP
PXXXXXXXXYYXNNNNNNNNZZZZZZZEEEEAEEAAAAAAAAARRRRRNNNNNIIIAATTCTTPPPTTTTTHHDDDDDKKKKDKKKKKKHHHHHHHHCCCCCCCCCCCCOOOOOOOOOOOOOOOOORRRRRLLLLPPPPP
PXXXXXXXXXXXXNNNNNNNZZZZZZZEEEAAEAAAAAAAARRRRRRRNNNNNIIIAAAACCYCCCCCCTTTDDDDDDDKKDDDKKKHHHHHHHHHHICCCCCCCCCCCCCOOOOOOOOOOOOOOORRRRRRRPPPPPPP
XXXXXXXXXXXXXNNNNNNNZZZZZZZZZAAAAAAEEEAAAARRRRRRNNNNNNAAAAACCCCCCCCCCTTTDDDDDKKKKDDDKKKKHHHHHHHHHICCCCUCCCCCCCCCOOOOOOOOOOOOOOORRRRROPPPPPPP
XKXXXXXXXXXXXXNNNNNNCCZQZZRRRAAAAEAEEEAAARRRRRRNNNNNNNAAAAACCCCCCCCCCTTFDDDDDDKKKDDDDKHHHHHHHHHHHHQCCCUUCCCCCCCCOOOOOOOQOOOOOOORROOOOPPPPPPP
KKKKXXXXXXXXXXNNNNLCCCCCZCRRRRAAAEEEEEARRRRRRRRNNHNNNNAAAAAACCCCCCCCTTFFFDDDDKKKKDDDDDDDHHHHHHHHHHHRUUUUUCCCCCCCCOOOJOOOOOOOOOOOOOOOPPPPPPPP
KKKKXXXXXXXDXNNNNNLCCCCCCCRRRREEEEEEEEERRRRRKRRRRRNNNNAOAAACCCCCCCCCTTFFFFDDDDDDDDDDDDDHHHHHHHHHHHHRRUUUUUUCCCCCCCOOJJOOOMOOOOOOOOOOOPPPPPPP
KKKKKXXXXXXXXCCLLLLLCCCCCCIRRRREEEEEEEEEKKKKKRRRRKAAAAAOOOOOCCFFCCCCTTFFZZVVDVVDDDDDDDDHHHHHHHHHHHHHRJUUUCCCCNNCCCCCCJOOMMMOOOOOOOOOOPPPPPPP
KKKKKXXXXXXAXXCCLCLCCCCCCIIRRRRREEEEEEEEIKKKKKKKRKAAOOOOOOOOCCOFCCCCZTTZZZVVVVVVDDDDDDHHHHHHHHHHHHJJJJUUUUCCNUNNNCCCCJJMMMMOOOOOOWWOOPYYPPPP
KKKKKKXXXXAAXXCCLCCCCCCCCCIQRRRREEEEEEIIIKKKKKKKKKAAOOOOOOOOOCOFCZZCZZZZZVVVVDDDDDDDDDDHHHHHHHWWEHJJJJJUUNNNNNNNNNGCJJMMMMMMMMWWWWWWWPPPPWPW
KKKKKKKKKKKAAXCCCCCCCCCCCCIQQRRREEEEEEEIIIKKKKKKKAAOOOOOOOOOOOOOCZZZZZZZZZVVVVDDDDDDDDDDDHHHEEEEEJJJJJJJNNNNNNNNNNNJJJMMMMMMMMWWWWWWWWWWWWWW
KKKKKKKKKKCCCCCCCCCCCCCCCCCQHHHHEEEEEEIIIKKKKKKKAAKOOOOOOOOOOOOOOZZZZZZZZZZVVDDDDDDDDDDDDHHHHEEEJJJJJJJJYYYYFNNNNNNNJJMMMMMMMMWWWWWDWWTTTTTW
KKKKKKKKCCCCCCCCCCCCCCCCCCCQQQHHHEEEEEIKKKKKKKKKKKKKOOOOOOOOOOOOOZZZZZZZZZZVVDDDDDDDDDDDDDDDEEEJJJJJJJJJYYYMIINNNNJJJMMMMMMMMWWWWWDDWWWWWTTT
KKKKKKVKCCCCCCCCCCCCCCKKCCMQQQHHHEEEEEKIKKKKKKKKKKKKOOOOOOOOOOOOOZZZZZZZZZVVVVDDDDDDDDLLDDDEEEEEJJJJJJJJYYYYIIINNNNNIIMMMMMMMWWWWWDDDDTTTTTT
KKKKKVVVPPCBCBBBCBBBBBBBBPQQQQHHHHKKEKKIKKKKKKKKKKKKOOOOOOOOOOOOOZZZZZZZZZVVVVDDDDDDDDLLDDDDEEEEEEEEJYJJYYYYIIINNNNNIIIIMMMMWWWWWWWDDCTTTTTT
KKKKKVVVVVVBBBBBBBBBBBBBBPQQQQQHHKKKKKKKKKKKKKKKKKKKKOOOOOOOOOOAOOZZZZZZZZZVVVDDDDLLLLLLDDDDEEEEEEEEYYJJYYYYYIIIINNIIIIIMMMMMTWWWWWWTTTTTTTT
KKKKVVVVVVBBBBBBBBBBBBBBBBQQQQQQHHKKKKKKKKKKKKKKKKKKOOOOOOOOOOJOOOZZZZZIZZIIIVDDDLLLLLLEEEEEEEEEEEEYYYYJYYYYIIIIIIIIIIITMTMMMTWWWWWTTTTTTTTT
KKKKVVVVVVVBBBBBBBBBBBBBBQQQQQQHHHKKKKKKKKKWWKKKKKKKOOOOOOOOOOOOZZZZZZZIIIIIIIDDDLLLLLLLLEEEEEEEEEEYYYYYYYYYYIIIIIIIQIITTTTTTTWWWTTTTTTTTTTT
KKKKVVVVVVVBBBBBBBBBBBBBBBQQQQQQHHKKKKKKKWWWWWKKYKKOOOOPOOOOOOOOOOZZZZZIIIIIIIIDDLLLLLLLLLEEEEEEEEEYYYYYYYYYIIIIIIITTTTTTTTTTTTTWTTTTTTTTTTT
KKKKKVVVVXVVLBBBBBBBBBBBBBBQQQQLLLKKKKKKWWWWWWKKYOOOOOOPPPPOOOOOOZZZZZZIIIIIIIIILLLLLLLLLLLEYYEYYYYYYYYYYYYIIIIIIITTTTTTTTTTTTTTTTTTTTTTTTTT
KKKKKKKKVLLLLBBBBBBBBBBBBBQQTTQLLLKKKKKKWWWWIIIOOOOOOOPPPPPPPPPOOZZZZZZIIIIIIIIIIILLLLLLLLDDDYYYYYYYYYYYYYTPPIIIIITTTTTTTTTTTTTRRTTTTTTTTTTT
KKKKKKLLLLLLLBBBBBBBBBBBBBBBTTTLLLLKZKRKOOWIIIIIOOOOOPPPPPPPPOOOOOZZZZZIIIIIIIIIIILLLLLLLLDDDFYYYYYYYYYYYYTPPNIIINTETATTTTTTTTRRRRRRTTTTTTTT
KKKKKKKLRLLLLJBBBBBBBBBBBBGLLLLLLLLLLVKKOOOIIIIOOOOOOOOOPPPPPPPOPZZZZZZIIIIIIIIIILLLLLLLLLLDDFYYYYYYYYYYYYPPPNNNNNTTTTTTTTTTTTRRRRRRLLTTTTTT
KKKKKKKLLLLBBBBBBBBBBBBBBGGLLLLLLLLLLVKKOOOIOOOOOOOOOOPPPPPPPPPPPZZZZZIIIIIIIIIIILLLLLLLLLDDDDKKYYYYYYYYYYPPNNNNNNNNTTTCCTCTPRRRRRRRLPTTTTTV
KKKKKKLLLELLBLLLBBBBEBBBBBBLLLLLLLLLLVVOOYOOOOOOOOOOOPPPPPPPPPPZZZZZZZZIIIIIIIIIILLLLULLLLLDDDKKYYYYYYYYYPPPNNNNNNNNCCCCCCCPPRRRRRRLLITIIVVV
KKKKKLLLLLLLLLLLLLCBBBBBCCBLLLLLLLLLLVYYYYOOOOOOOOOOOPPPPPPPPPPZZZZZZZZIIIIIIIIIZLJLULLLLLGGDDKKKYYYYYYPPPPPNONNNNNNCCCCCCCPPPRPPRRLLIIIIVVV
KKKKKKKLLLLLLLLLLLCCCCCBCCLLLLLLLLLLLVYYYYYYOYOOOOOOOOWWPPPPPPPPZZZZZZZIIIIIIIZIZLJLULLLLLGGGDDDPPYYYYYYPPPNNNNNNNNNCCCCCCCCPPPPPRRRRRIVVVVV
KKKKKKKLDLLLLLLLLYCCCCCBCLLLLLLLLLLVVVYYYYYYYYYOOOOOOOWWPWPWPPHPPZPZZZZIIZZZIIZIZZEEEQQQQQGGGDDDGPPPYYPPPPNJNNNNNNNNNCCCCCCCCCPPPRRRRRVVVVVV
KKKKKKDNDLLLLLLLLYCCCCCCCCCPPDDLLLLVVVYYYYYYYYOOOOOOOOWWWWWWHHHPPPPHHHZZZZZZZIZZZEEEEEQQQQQGGDGGGPPPPPPPPPNNNNNNNNNNCCCCCCCCCPPPPRRRVRVVVVVV
KKKKKKDDDDDDLLLLYYYYCCCCCCCPPDPSSPVVVVYYYYYYYYOWOOOWWWWWWWWHHHHHHHHHHZZZZZZZZIZZZEEEEEQQQQQQGGGNNPNNPPPPPPPNNVNNCCCNCCCCICCCINPPPRRVVVVVVVVV
DKKKKKDDDDDDTLLYYYYYYCCCCCCPPPPPSPPVVVYYYYYYYYWWOOWWWWWWWWHHHHHHHHHHFZZZZZZZZZZZEEEEQQQQQQQQQGNNNNNNPPPPPPPPVVVWCCCNCCCIIIIIINPPPRRIIIVVVVVV
DDDDDDDDDDDTTLLYYYYYYCYCCCPPPPPPPPVVVVYYYYYYYYWWWWWWWWWWWWHHHHHHHHHHFZFZZZZZZZZZEEEEQQQQQQQQQGNNNNNNNNPPPLPPPVVWWCNNCCCIIIIIIPPPPRIIYVVVVVVV
DDDDDDDDDDDTTLLTTYYYYYYYCPPPPPPPVPPVVYYYYYYYYYYWWWWWWWWWWHHHHHHHHHHFFFFFZZZZZEEEEEEEQQQQQBGGNNNNNNNNNNNPPWWWWWWWWCCCCHHIIIIIIIIPIIIIIVVVVVVV
DDDDDDDDDDDDTTTTTTYTTTYYCCPPPPPPVVVVYYGYYYYYYYYWWWWWWWWWWWHHHHHHHHFFFFFFZZZZZEEEEEEEEQQBBBGGGNNNNNNNNNNNNNWWWWWWWCCCCCHHIIIIIIIIIIIIIIVVVVVV
KKDDDDDDDDDDTTTTTTTTTTTTCCCPPPPPVVVVVYYYYYYYYYWWWWWWWWWWWHHHHHHHHHFFFFFFFZZZZEEEEEEEQQQBBBBBGBBNNNNNNNNNNNWWWWWWWWWWHHHHHIIIIIIIIIIILLLVVLVL
KKDDDDDDDDTTTTTTTTTTTTTTCCCPPPPVVVVVYYYYYYYYYYWWWWWWWWWWWWHHHHHHHFFFFFFFFFFZZEEEEEEBQQQBBBBBBBBBNNNNNNNNNNWWWWWWWWWMHHHHHIIIIIIIIIIILLLLLLVL
KKKDDDDDHDTTTTTTTTTTTTTCCCPPPPPPPPVYYYYYYYYYYYWWWWWWWWWWWWHHHHHHHHQQFFFFFFFZZEEZEEBBQQQQBBBBBBBBBBNNNNNNNNWWWWWWWWWWHHHHHIIIIIIIIIIILLLLLLVL
KKKDDDDDDTTTTTTTTTTTTTTCCPPPPPPPPCYYYYYYYYYYRYWWWWWWWWWWWWHHHHHHHHQFFFFFFFFFZZZZZERBBBBBBBBBBBBBBBBBNNNWWWWWWWWWWWWWWHHNIIIIIIIIIIIICLLLLLLL
KKDDDDDDDTTTTTTTTTTTTTTTCCCPPPPPPCYLLYYLYYYRRWWWWWWWWWWWHHHHHHHHHHHHFGFFZZFFZZZEEEEEBBBBBBBBBBBBBBNNNNWWWWWWWWWWWWWNNNHNIIIIIIIIIIIICCLLLLLL
KKKKDDDDTTTTTTTTTTTTTTTCCCCCPPPPPYYYLLLLLYYRRRRRWWWTTWWWWQQHHHHHHHHHHHHFZZZZZZZEQEEEBXBBBBBBBBBBBBNNNNNWWWWWWWWWWWWWNNNNIIIIIIIIIIIICCLLLLLL
"""


class TestDay:
    @pytest.mark.parametrize(testit.names, [(example, 1930), (answer, 1396298)], ids=testit.ids())
    def test_first_star(self, indata: str, want: int) -> None:
        testit.first_star(day12.Day(indata), want)

    @pytest.mark.parametrize(testit.names, [(example, 1206), (answer, 853588)], ids=testit.ids())
    def test_second_star(self, indata: str, want: int) -> None:
        testit.second_star(day12.Day(indata), want)
