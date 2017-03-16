package user_shell

import "os"
import "runtime"

func GetUserShell() (string) {
    switch runtime.GOOS {
        case "windows":
            if os.Getenv("COMSPEC") != "" {
                return os.Getenv("COMSPEC")
            }
            return "/cmd.exe"
        case "darwin":
            if os.Getenv("SHELL") != "" {
                return os.Getenv("SHELL")
            }
            return "/bin/bash"
        default:
            if os.Getenv("SHELL") != "" {
                return os.Getenv("SHELL")
            }
            return "/bin/sh"
    }
    return ""
}
