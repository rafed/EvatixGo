package parse

import "github.com/gocarina/gocsv"

// CliStreamerRecord holds streaming info
type CliStreamerRecord struct {
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

// CliRunnerRecord holds how many times it will run
type CliRunnerRecord struct {
	Run string `csv:"Run"`
	CliStreamerRecord
}

// CliStreamerRecordCsv returns streamer runner of a csv
func (cliRunnerRecord CliRunnerRecord) CliStreamerRecordCsv() string {
	cliStreamerRecords := []CliStreamerRecord{cliRunnerRecord.CliStreamerRecord}

	out, err := gocsv.MarshalString(cliStreamerRecords)

	if err != nil {
		panic(err)
	}

	return out
}

// // CliStreamerRecord returns a streamer object
// func (c CliRunnerRecord) GetCliStreamerRecord() CliStreamerRecord {
// 	return CliStreamerRecord{
// 		Title:       CliRunnerRecord.Title,
// 		Message1:    CliRunnerRecord.Message1,
// 		Message2:    CliRunnerRecord.Message2,
// 		StreamDelay: CliRunnerRecord.StreamDelay,
// 		RunTimes:    CliRunnerRecord.RunTimes,
// 	}
// }
