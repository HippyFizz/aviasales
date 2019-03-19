package aviasales

import (
	"github.com/urfave/cli"
	"os"
)

var (
	app = cli.NewApp()
)

func main() {
	app.Name = "aviasales"
	app.Usage = "fight the loneliness!"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf",
			Value: os.ExpandEnv("$GOPATH/config/config.yaml"),
			Usage: "configuration yaml file",
		},
	}

	//app.Commands = append(app.Commands,
	//	commands.Web,
	//	commands.Worker,
	//)

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
