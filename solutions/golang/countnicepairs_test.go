package golang

import "testing"

func TestCountNicePairs(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{42, 11, 1, 97}}, want: 2},
		{name: "smoke 2", args: args{nums: []int{13, 10, 35, 24, 76}}, want: 4},
		{
			name: "test 59: wrong answer",
			args: args{nums: []int{
				74888847, 394615546, 133272331, 251696003, 209613900, 12499421, 6046406, 159141802, 274060323,
				483333235, 371201227, 193606391, 382411271, 213774310, 316562611, 252897303, 487120838, 261433215,
				443362491, 371865161, 458190703, 362302113, 465747415, 275555423, 382252133, 365908616, 325421521,
				350845041, 331161133, 272655122, 473592372, 368969714, 340582041, 21922912, 311897260, 171706225,
				184595331, 283746232, 338131833, 196150745, 211340100, 326966611, 294332341, 275080422, 134505431,
				266757512, 375282424, 336877780, 337360731, 458322905, 394190547, 383636234, 5875785, 246360630,
				383101234, 164988512, 477453827, 160645114, 467948613, 184949332, 457714742, 459574803, 256413706,
				471140023, 166232511, 459858804, 268847712, 285020433, 152716100, 360867114, 210161012, 378433927,
			}},
			want: 239,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countNicePairs(test.args.nums); got != test.want {
				t.Errorf("countNicePairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
