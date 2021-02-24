package tiktokdl

func init() {
	if !IsCommandAvailable("wget") {
		panic("Please install wget")
	}
}
