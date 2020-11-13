package cutil

import (
	"time"
)

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func Min32(a, b int32) int32 {
	if a <= b {
		return a
	} else {
		return b
	}
}

func Min64(a, b int64) int64 {
	if a <= b {
		return a
	} else {
		return b
	}
}

func MinDuration(a, b time.Duration) time.Duration {
	if a <= b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Max32(a, b int32) int32 {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Max64(a, b int64) int64 {
	if a >= b {
		return a
	} else {
		return b
	}
}

func MaxDuration(a, b time.Duration) time.Duration {
	if a >= b {
		return a
	} else {
		return b
	}
}
