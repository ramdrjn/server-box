package serverbox

type Statistics struct {
	enabled bool
}

func InitializeStatistics(sbc *SbContext) (statistics *Statistics, err error) {
	if (sbc.Conf.Stats.Enabled == false) {
		Log.Debugln("statistics not configured")
		return nil, nil
	}

	statistics = &Statistics{true}
	return statistics, nil
}
