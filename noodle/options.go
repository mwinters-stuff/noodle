package noodle

type NoodleOptions struct {
	Config string `short:"c" long:"config" description:"Noodle Configuration"`
	Debug  bool   `short:"d" long:"debug" description:"Debug Information"`
}

type AllNoodleOptions struct {
	NoodleOptions NoodleOptions
}
