package main

import (
	"fmt"
	"os/exec"
)

func main() {
	checkPlayStore := exec.Command("adb", "shell", "pm", "path", "com.android.vending")

	out, err := checkPlayStore.Output()

	// If the Play Store is not installed (no output and exit status 1)
	if err != nil && string(out) == "" {
		installPlayStore := exec.Command("adb", "shell", "pm", "install-existing", "com.android.vending")
		installChrome := exec.Command("adb", "shell", "pm", "install-existing", "com.android.chrome")

		out, err := installPlayStore.Output()

		if err != nil {
			fmt.Println(string(out))
			fmt.Println(err)
		}

		out, err = installChrome.Output()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println(err)
		}
	} else {
		removePlayStore := exec.Command("adb", "shell", "pm", "uninstall", "--user", "0", "com.android.vending")
		removeChrome := exec.Command("adb", "shell", "pm", "uninstall", "--user", "0", "com.android.chrome")

		out, err := removePlayStore.Output()

		if err != nil {
			fmt.Println(string(out))
			fmt.Println(err)
		}

		out, err = removeChrome.Output()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println(err)
		}
	}
}
