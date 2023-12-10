package colors

type ColorsStruct struct {
	Primary   string
	Secondary string
	Error     string
	Warning   string
}

var Colors ColorsStruct

func init() {
	Colors.Primary = "\033[36m"
	Colors.Secondary = "\033[34m"
	Colors.Error = "\033[31m"
	Colors.Warning = "\033[33m"
}
