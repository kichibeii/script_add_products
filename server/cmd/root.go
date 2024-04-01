package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/script_add_products/server/commons"
	"github.com/script_add_products/server/config"
	"github.com/script_add_products/server/domain/repositories"
	"github.com/script_add_products/server/domain/thirdparties"
	services "github.com/script_add_products/server/service"
	"github.com/spf13/cobra"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "script-sync-product",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	cobra.OnInitialize()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() {
	ctx := context.Background()
	cfg := config.GetConfig()

	fmt.Println("connect to database")

	// connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.AdditionalParameters)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	opt := commons.Options{
		Config:   cfg,
		Database: db,
	}

	fmt.Println("initiate repository")
	repo := initRepository(opt)

	fmt.Println("initiate thirdparty")
	thirdParty := initThirdparty(opt)

	// run process sync
	fmt.Println("initiate usecases")
	syncProduct := services.NewSyncProductService(opt, repo, thirdParty)

	fmt.Println("run process sync")
	err = syncProduct.Run(ctx)
	if err != nil {
		fmt.Println("[ERROR LOG] SYNC PRODUCT", err)
	}

	fmt.Println("process finish")
}

func initThirdparty(opt commons.Options) thirdparties.IThirdParty {
	thirdparty := thirdparties.NewThirdParty(opt)

	return thirdparty
}

func initRepository(opt commons.Options) repositories.IRepository {
	repo := repositories.NewRepository(opt)

	return repo
}
