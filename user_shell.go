package user_shell

import "os"
import "runtime"

func GetUserShell() (string) {
    if runtime.GOOS == "windows" {
        if os.Getenv("COMSPEC") != "" {
            return os.Getenv("COMSPEC")
        } else {
            return "/cmd.exe"
        }
    } else if runtime.GOOS == "darwin" {
        if os.Getenv("SHELL") != "" {
            return os.Getenv("SHELL")
        } else {
            return "/bin/bash"
        }
    }
    if os.Getenv("SHELL") != "" {
        return os.Getenv("SHELL")
    } else {
        return "/bin/sh"
    }
    return "";
}
