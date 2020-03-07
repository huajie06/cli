package cmdparser

type Command struct {
	Use   string
	Short string
	Long  string
	Flags []Flag
	Help  string
}

type Flag struct {
	Use   string
	Short string
	Long  string
	Help  string
}

func NewCmd() *Command {
	return &Command{
		Use:   "",
		Short: "A CLI app",

		Flags: []Flag{{
			Use:   "help",
			Short: "Please use --help for more info.",
			Long:  "",
		}},
	}
}

func (c *Command) AddFlag(f Flag) {
	c.Flags = append(c.Flags, f)
	return
}
