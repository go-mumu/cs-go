/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-04
 * File: gen_models.go
 * Desc: generate model
 */

package cmd

import (
	"github.com/go-mumu/cs-go/service/container"
	"github.com/go-mumu/cs-go/service/dal"
	"github.com/spf13/cobra"
)

// Usage: go run main.go gen-model -d defMysql
// var Args string
//
// func init() {
// 	   rootCmd.AddCommand(genModel)
// 	   genModel.Flags().StringVarP(&Args, "defMysql", "d", "defMysql", "")
// 	   _ = genModel.MarkFlagRequired("defMysql")
// }

// go run main.go gen-model -d defMysql
var genModel = &cobra.Command{
	Use:   "gen-model",
	Short: "generate models",
	Long:  "use gorm gen, generate models",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, cleanfunc, err := container.InitApp()
		defer cleanfunc()

		if err != nil {
			return err
		}

		dal.GenDefModels(app.DefMysql.DB)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(genModel)
}
