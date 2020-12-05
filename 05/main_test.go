package main

import "testing"

type TestCase struct {
	Iid           string
	rid, cid, sid int
}

var testCases = []TestCase{
	TestCase{"BFFFBBFRRR", 70, 7, 567},
	TestCase{"FFFBBBFRRR", 14, 7, 119},
	TestCase{"BBFFBBFRLL", 102, 4, 820},
}

func TestDecodeSeatIid(t *testing.T) {
	for _, tc := range testCases {
		rid, cid, sid := DecodeItinerary(tc.Iid)
		if rid != tc.rid {
			t.Errorf("expected row %d, got %d", tc.rid, rid)
		}
		if cid != tc.cid {
			t.Errorf("expected col %d, got %d", tc.cid, cid)
		}
		if sid != tc.sid {
			t.Errorf("expected sid %d, got %d", tc.sid, sid)
		}
	}
}

func TestMaxInt(t *testing.T) {
	nums := []int{1, 12, 22, 4, 42, 9, 1, 10, 23}
	max := MaxInt(nums)
	if max != 42 {
		t.Errorf("expected 42, got %d", max)
	}
}
