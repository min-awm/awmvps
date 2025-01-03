package terminal

import (
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Init(app *fiber.App) {
	setXTERM()

	app.Get("/terminal", websocket.New(func(c *websocket.Conn) {
		defer c.Close()

		cmd := exec.Command("bash")
		ptmx, err := pty.Start(cmd)
		if err != nil {
			return
		}

		defer func() {
			if cmd.Process != nil {
				cmd.Process.Kill()
			}

			ptmx.Close()
			cmd.Wait()
		}()

		go func() {
			buf := make([]byte, 1024)
			for {
				n, err := ptmx.Read(buf)
				if err != nil {
					if err == io.EOF {
						break
					}
					break
				}

				if err := c.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
					break
				}
			}
		}()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				break
			}
			_, err = ptmx.Write(msg)
			if err != nil {
				break
			}
		}
	}))
}

func setXTERM() {
	err := os.Setenv("TERM", "xterm")
	if err != nil {
		log.Error("Error setting TERM:", err)
	}
}
