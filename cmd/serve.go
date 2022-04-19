/*
Copyright © 2022 Alex Bechanko alexbechanko@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const defaultPort = 8080

var port int
var rootPath string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves static files from a provided configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		// check if root path is empty, if it is then default to current working directory
		if rootPath == "" {
			logrus.Info("No path argument, using current working directory")
			var err error
			rootPath, err = os.Getwd()
			if err != nil {
				logrus.WithError(err).Error("Failed to get working directory")
				return
			}
		}
		logrus.WithField("port", port).Info("Starting file server")
		http.ListenAndServe(
			fmt.Sprintf(":%d", port),
			http.FileServer(http.Dir(rootPath)),
		)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntVar(&port, "port", defaultPort, "Port to serve files from")
	serveCmd.Flags().StringVar(&rootPath, "path", "", "File path to serve")
}
