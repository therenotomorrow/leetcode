package golang

import "testing"

func TestGetDecimalValue(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}},
			want: 5,
		},
		{
			name: "smoke 2",
			args: args{head: &ListNode{Val: 0, Next: nil}},
			want: 0,
		},
		{
			name: "test 73: wrong answer",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 1,
										Next: &ListNode{
											Val: 1,
											Next: &ListNode{
												Val: 1,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val: 1,
														Next: &ListNode{
															Val: 1,
															Next: &ListNode{
																Val: 0,
																Next: &ListNode{
																	Val: 0,
																	Next: &ListNode{
																		Val: 0,
																		Next: &ListNode{
																			Val: 0,
																			Next: &ListNode{
																				Val: 0,
																				Next: &ListNode{
																					Val: 0,
																					Next: &ListNode{
																						Val: 0,
																						Next: &ListNode{
																							Val: 0,
																							Next: &ListNode{
																								Val:  1,
																								Next: nil,
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: 685569,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getDecimalValue(test.args.head); got != test.want {
				t.Errorf("getDecimalValue() = %v, want = %v", got, test.want)
			}
		})
	}
}
