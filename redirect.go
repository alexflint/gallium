package gallium

import (
	"os"
	"syscall"
)

// RedirectStdoutStderr overwrites the stdout and stderr streams with a file
// descriptor that writes to the given path. This is done with os.Dup2,
// meaning that even C functions that write to stdout or stderr will be
// redirected to the file.
func RedirectStdoutStderr(path string) (*os.File, error) {
	return redirect(path, os.Stdout.Fd(), os.Stderr.Fd())
}

// RedirectStdout overwrites the stdout streams with a file
// descriptor that writes to the given path. This is done with os.Dup2,
// meaning that even C functions that write to stdout will be
// redirected to the file.
func RedirectStdout(path string) (*os.File, error) {
	return redirect(path, os.Stdout.Fd())
}

// RedirectStderr overwrites the stderr streams with a file
// descriptor that writes to the given path. This is done with os.Dup2,
// meaning that even C functions that write to stderr will be
// redirected to the file.
func RedirectStderr(path string) (*os.File, error) {
	return redirect(path, os.Stderr.Fd())
}

func redirect(path string, fds ...uintptr) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	for _, fd := range fds {
		err = syscall.Dup2(int(f.Fd()), int(fd))
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}
