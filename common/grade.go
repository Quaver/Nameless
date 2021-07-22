package common

type Grade string

const (
	GradeNone string = "None"
	GradeA           = "A"
	GradeB           = "B"
	GradeC           = "C"
	GradeD           = "D"
	GradeF           = "F"
	GradeS           = "S"
	GradeSS          = "SS"
	GradeX           = "X"
)

func GetGradeFromAccuracy(acc float32, failed bool) Grade {
	if failed {
		return GradeF
	}

	if acc >= 100 {
		return GradeX
	}
	if acc >= 99 {
		return GradeSS
	}
	if acc >= 95 {
		return GradeS
	}
	if acc >= 90 {
		return GradeA
	}
	if acc >= 80 {
		return GradeB
	}
	if acc >= 70 {
		return GradeC
	}

	return GradeD
}
