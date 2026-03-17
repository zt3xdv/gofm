/*
Copyright © 2026 Zoom theoldzoom@proton.me

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

	"github.com/theOldZoom/gofm/internal/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var recentCmd = &cobra.Command{
	Use:   "recent [username]",
	Short: "Show recently played tracks",

	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			limit = 10
		}
		if len(args) == 1 {
			username = args[0]
		}
		if username == "" {
			fmt.Println("No username provided. Pass one explicitly or run setup first.")
			return
		}

		tracks, err := api.GetRecentTracks(username, limit)
		if err != nil {
			fmt.Println("Failed to get recent tracks:", err)
			return
		}
		for i, track := range tracks {
			fmt.Printf("%d. %s - %s\n", i+1, track.Artist.Name, track.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(recentCmd)
	recentCmd.Flags().IntP("limit", "l", 10, "Number of tracks to show")
}
