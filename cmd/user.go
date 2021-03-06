/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"GoGinChat/config"
	"GoGinChat/models"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Command to add new user to db",
	Long: `Command to add new user to db using a bash script`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add user to db")
		flag.String("username", "user", "Add username in DB")
		flag.String("email", "emai", "Add email to DB")
		flag.Parse()

		if _, err := os.Stat("addUser.sh"); os.IsNotExist(err){
			panic(err)
		}else{
			output, err := exec.Command("/bin/sh", "$GOPATH/GoGinChat/addUser.sh").Output()
			if err!=nil {
				panic(err)
			}else{
				if err:= config.ReadConfig(); err!=nil{
					panic(err)
				}
				fmt.Println("Read config successful for DB")
				var dbConn models.DBConnection
				viper.SetDefault("database.dbname", "goeat")
				err := viper.Unmarshal(&dbConn)
				if err!=nil{
					fmt.Println("Error unmarshaling the config")
					panic(err)
				}

				//Adding bash script for default user add in DB
				cmd, err := exec.Command("/bin/sh", "$GOPATH/cmd/addUser.sh").Output()
				if err!=nil {
					log.Fatal(err)
					panic(err)
				}
				fmt.Printf("User added %s : %s", cmd, output)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
