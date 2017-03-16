package user_shell

import "runtime"
import "strings"
import "testing"

func TestGetUserShell(t *testing.T) {
    if runtime.GOOS == "windows" {
        shell := GetUserShell()
        if strings.HasPrefix(shell, "/bin/") {
            t.Error("Expected shell to be a Windows path, got ", shell)
        }
    } else {
        shell := GetUserShell()
        if strings.HasPrefix(shell, "C:\\") {
            t.Error("Expected shell to be a Unix path, got ", shell)
        }
    }
}
