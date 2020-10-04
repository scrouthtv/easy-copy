package main;

import "fmt";
import "os";

import "github.com/spf13/viper"

var verbose bool;
var onExistingFile uint8 = 2;
// 0 skip
// 1 overwrite
// 2 ask
var followSymlinks uint8 = 1;
// 0 ignore symlinks
// 1 follow symlinks, copying them as links
// 2 fully dereference

func verboseFlags() {
	fmt.Printf(FGColors.Green)
	fmt.Println(" Verbose:", verbose);
	fmt.Println(" Overwrite Mode:", onExistingFile);
	fmt.Print(" Follow symlinks:", followSymlinks);
	fmt.Println(Textstyle.Reset);
}

func readConfig() {
	viper.SetDefault("verbose", false);
	viper.SetDefault("overwrite", 2);
	viper.SetDefault("followSymlinks", 1);

	viper.AddConfigPath("$XDG_CONFIG_DIR/");
	viper.AddConfigPath("$XDG_CONFIG_DIR/ec/");
	viper.AddConfigPath("$HOME/.config/");
	viper.AddConfigPath("$HOME/.config/ec/");
	viper.SetConfigName("ec");
	viper.SetConfigType("toml");
	err := viper.ReadInConfig();
	if err != nil {
		warnConfig(err);
	}
	viper.WriteConfig();

	verbose = viper.GetBool("verbose");
	onExistingFile = uint8(viper.GetInt("overwrite"));
	followSymlinks = uint8(viper.GetInt("followSymlinks"));
	if onExistingFile > 2 {
		warnBadConfig("overwrite", onExistingFile, "0, 1, 2");
		onExistingFile = 2;
	}
	if followSymlinks > 2 {
		warnBadConfig("followSymlinks", followSymlinks, "0, 1, 2");
		followSymlinks = 2;
	}
}

func parseFlag(prefix string, flag string) {
	switch(flag) {
	case "h", "help":
		printHelp();
		os.Exit(0);
	case "v", "version":
		printVersion();
		os.Exit(0);
	case "V", "verbose":
		verbose = true;
		verboseVerboseEnabled();
		break;
	case "f", "force":
		onExistingFile = 1;
		break;
	case "i", "interactive":
		onExistingFile = 2;
		break;
	case "n", "no-clobber": //case "no-overwrite":
		onExistingFile = 0;
		break;
	default:
		errUnknownOption(prefix + flag);
	}
}
