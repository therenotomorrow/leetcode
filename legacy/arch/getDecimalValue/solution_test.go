package getDecimalValue

import "testing"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
					},
				},
			}},
			want: 5,
		},
		{
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			}},
			want: 11,
		},
		{
			args: args{head: &ListNode{Val: 0}},
			want: 0,
		},
		{
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
																								Val: 1,
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
