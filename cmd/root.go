/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"math/rand"
	"net/http"
	"net/url"

	echo "github.com/labstack/echo"
	distuv "gonum.org/v1/gonum/stat/distuv"

	middleware "github.com/labstack/echo/middleware"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var poissonGenerator distuv.Poisson

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rest-in-chaos",
	Short: "add unreliability to any HTTP service",
	Long: `rest-in-chaos is an HTTP proxy that randomly fails any request you make to its server, and lets the remaining ones go.
  `,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("URL argument required. Run rest-in-chaos <<YOUR URL>>")
		}

		if u, err := url.Parse(args[0]); err != nil || u.Scheme == "" || u.Host == "" {
			fmt.Println("WARNING. Just so you know, the URL you've entered does not look right. We're going forward anyway.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		url, _ := url.Parse(args[0])

		poissonGenerator = distuv.Poisson{
			Lambda: 10000,
		}

		e := echo.New()
		e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
			Skipper: RequestSkipper,
			Balancer: middleware.NewRandomBalancer([]*middleware.ProxyTarget{
				{
					URL: url,
				},
			}),
		}))
		e.GET("/", func(c echo.Context) error {
			return c.JSON(GetRandomErrorCode(), nil)
		})
		e.HideBanner = true
		fmt.Println(`
     ____  _________________   ____         ________                    
    / __ \/ ____/ ___/_  __/  /  _/___     / ____/ /_  ____ _____  _____
   / /_/ / __/  \__ \ / /     / // __ \   / /   / __ \/ __ '/ __ \/ ___/
  / _, _/ /___ ___/ // /    _/ // / / /  / /___/ / / / /_/ / /_/ (__  ) 
 /_/ |_/_____//____//_/    /___/_/ /_/   \____/_/ /_/\__,_/\____/____/  
                                                                        
    `)
		fmt.Println("Adding unreliability to ", url)
		e.Logger.Fatal(e.Start(":24267"))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RequestSkipper(echo.Context) bool {
	if int(poissonGenerator.Rand())%2 == 0 {
		fmt.Println(time.Now().Format("15:04:05"), "Received request. Returning Error.")
		return true
	}
	fmt.Println(time.Now().Format("15:04:05"), "Received request. Forwarded.")
	return false
}

func GetRandomErrorCode() int {
	failureCodes := []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusForbidden, http.StatusGatewayTimeout, http.StatusRequestTimeout}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	return failureCodes[r.Intn(len(failureCodes))]
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rest-in-chaos.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".rest-in-chaos" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".rest-in-chaos")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
