package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/schollz/progressbar/v3"
)

func main() {
	checkPlayStore := exec.Command("adb", "shell", "pm", "path", "com.android.vending")
	out, err := checkPlayStore.Output()

	// If the Play Store is not installed (no output and exit status 1)
	if err != nil && string(out) == "" {
		// Install apps
		bar := progressbar.NewOptions(2,
			progressbar.OptionSetDescription("Restoring"),
			progressbar.OptionSetWidth(30),
			progressbar.OptionShowCount(),
		)

		installPlayStore := exec.Command("adb", "shell", "pm", "install-existing", "com.android.vending")

		out, err := installPlayStore.Output()
		if err != nil {
			fmt.Println()
			fmt.Println(string(out))
			fmt.Println(err)
			return
		}
		bar.Add(1)

		installChrome := exec.Command("adb", "shell", "pm", "install-existing", "com.android.chrome")
		out, err = installChrome.Output()
		if err != nil {
			fmt.Println()
			fmt.Println(string(out))
			fmt.Println(err)
			return
		}
		bar.Add(1)

		fmt.Println()
	} else {
		// Apps are already installed
		fmt.Println("Apps Restored")
	}

	// Wait for user to press Enter
	fmt.Print("Press Enter to remove apps...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// Remove apps
	bar := progressbar.NewOptions(2,
		progressbar.OptionSetDescription("Removing"),
		progressbar.OptionSetWidth(30),
		progressbar.OptionShowCount(),
	)

	removePlayStore := exec.Command("adb", "shell", "pm", "uninstall", "--user", "0", "com.android.vending")

	out, err = removePlayStore.Output()
	if err != nil {
		fmt.Println()
		fmt.Println(string(out))
		fmt.Println(err)
	}
	bar.Add(1)

	removeChrome := exec.Command("adb", "shell", "pm", "uninstall", "--user", "0", "com.android.chrome")
	out, err = removeChrome.Output()
	if err != nil {
		fmt.Println()
		fmt.Println(string(out))
		fmt.Println(err)
	}
	bar.Add(1)

	fmt.Println()
}
