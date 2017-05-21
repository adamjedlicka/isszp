package doc

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/skratchdot/open-golang/open"
)

var (
	port = "6060"
	pkg  = "gitlab.fit.cvut.cz/isszp/isszp/"
)

func Run() {
	var err error

	if flag.Arg(1) != "" {
		port = flag.Arg(1)
	}

	go func() {
		time.Sleep(1 * time.Second)

		address := fmt.Sprintf("http://localhost:%s/pkg/%s", port, pkg)
		err = open.Run(address)
		if err != nil {
			err = open.Start(address)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	godoc := ""
	if runtime.GOOS == "widnows" {
		godoc = "godoc_windows"
	} else if runtime.GOOS == "linux" {
		godoc = "godoc_linux"
	} else if runtime.GOOS == "darwin" {
		godoc = "godoc_macos"
	}

	cmd := exec.Command(filepath.Join(dir, godoc), "-http", "localhost:"+port)
	err = cmd.Run()
	if err != nil {
		cmd := exec.Command("godoc", "-http", "localhost:"+port)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
