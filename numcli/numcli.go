package numcli

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func flagUsage() {
	usageText := `example is an example cli tool.

Usage:
example command [arguments]
The commands are:
convhex    convert Number to Hex
convbinary convert Number to binary
Use "Example [command] --help" for more infomation about a command`

	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func Convert(mode string, num []string) {

	flag.Usage = flagUsage
	convHexCmd := flag.NewFlagSet("convhex", flag.ExitOnError)
	convBinaryCmd := flag.NewFlagSet("convbinary", flag.ExitOnError)

	// コマンドが指定されていない場合
	if mode == "" {
		flag.Usage()
		return
	}

	switch mode {
	case "convhex":
		i := convHexCmd.Int64("i", 0, "Convert number to hex")
		convHexCmd.Parse(num)
		fmt.Println(strconv.FormatInt(*i, 16))
	case "convbinary":
		i := convBinaryCmd.Uint64("i", 0, "convert number to binary")
		convBinaryCmd.Parse(num)
		fmt.Println(strconv.FormatUint(*i, 2))
	default:
		flag.Usage()
	}
}