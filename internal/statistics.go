package serverbox

type Statistics struct {
	enabled bool
}

func InitializeStatistics(sbc *SbContext) (statistics *Statistics, err error) {
	err = nil
	statistics = new(Statistics)
	return statistics, err
}
