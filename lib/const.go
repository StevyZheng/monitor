package lib

var (
	ConstCpuTempThresholdMax    int
	ConstMemoryTempThresholdMax int
	ErrMap                      map[string]string
)

func ConstInit() {
	ConstCpuTempThresholdMax = 90
	ConstMemoryTempThresholdMax = 40
	ErrMap = map[string]string{
		"io error": "io error",
	}
}
