package enums

var ActivityTypes = map[int16]string{
	1:  "Walking",
	2:  "Yoga",
	3:  "Stretching",
	4:  "Cycling",
	5:  "Swimming",
	6:  "Dancing",
	7:  "Hiking",
	8:  "Running",
	9:  "HIIT",
	10: "JumpRope",
}

var Calories = map[string]int{
	"Walking":    4,
	"Yoga":       4,
	"Stretching": 4,
	"Cycling":    8,
	"Swimming":   8,
	"Dancing":    8,
	"Hiking":     10,
	"Running":    10,
	"HIIT":       10,
	"JumpRope":   10,
}
