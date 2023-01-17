package noodle

type NoodleOptions struct {
	Config string `short:"c" long:"config" description:"Noodle Configuration"`
	Debug  bool   `short:"d" long:"debug" description:"Debug Information"`
	Drop   bool   `long:"drop" description:"Drop Database"`
}

type AllNoodleOptions struct {
	NoodleOptions NoodleOptions
}
