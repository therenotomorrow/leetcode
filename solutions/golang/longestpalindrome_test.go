package golang

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abccccdd"}, want: 7},
		{name: "smoke 2", args: args{s: "a"}, want: 1},
		{name: "test 49: wrong answer", args: args{
			s: "" +
				"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbat" +
				"tlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheir" +
				"livesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedic" +
				"atewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecratedit" +
				"faraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneve" +
				"rforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughthe" +
				"rehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromth" +
				"esehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatwehere" +
				"highlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandth" +
				"atgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
		}, want: 983},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := longestPalindrome(test.args.s); got != test.want {
				t.Errorf("longestPalindrome() = %v, want = %v", got, test.want)
			}
		})
	}
}
