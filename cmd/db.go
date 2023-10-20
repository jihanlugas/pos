package cmd

import (
	"github.com/jihanlugas/pos/cryption"
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/utils"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"time"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Run server",
	Long: `With this command you can
	up : create database table
	down :  drop database table
	seed :  insert data table
	`,
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Up table",
	Long:  "Up table",
	Run: func(cmd *cobra.Command, args []string) {
		up()
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Down table",
	Long:  "remove public schema, create public schema, restore the default grants",
	Run: func(cmd *cobra.Command, args []string) {
		down()
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed data table",
	Long:  "Seed data table",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Down, up, seed table",
	Long:  "Down, up, seed table",
	Run: func(cmd *cobra.Command, args []string) {
		down()
		up()
		seed()
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(upCmd)
	dbCmd.AddCommand(downCmd)
	dbCmd.AddCommand(resetCmd)
	dbCmd.AddCommand(seedCmd)
}

func up() {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	// table

	err = conn.Migrator().AutoMigrate(&model.Log{})
	if err != nil {
		panic(err)
	}
	err = conn.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Company{})
	if err != nil {
		panic(err)
	}

	err = conn.Migrator().AutoMigrate(&model.Usercompany{})
	if err != nil {
		panic(err)
	}

	// view
	vUser := conn.Model(&model.User{}).
		Select("users.*, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join users u1 on u1.id = users.create_by").
		Joins("left join users u2 on u2.id = users.update_by").
		Joins("left join users u3 on u3.id = users.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_USER, gorm.ViewOption{
		Replace: true,
		Query:   vUser,
	})
	if err != nil {
		panic(err)
	}

	vCompany := conn.Model(&model.Company{}).
		Select("companies.*, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join users u1 on u1.id = companies.create_by").
		Joins("left join users u2 on u2.id = companies.update_by").
		Joins("left join users u3 on u3.id = companies.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_COMPANY, gorm.ViewOption{
		Replace: true,
		Query:   vCompany,
	})
	if err != nil {
		panic(err)
	}

	vUsercompany := conn.Model(&model.Usercompany{}).
		Select("usercompanies.*, companies.name as company_name, users.fullname as user_name, u1.fullname as create_name, u2.fullname as update_name, u3.fullname as delete_name").
		Joins("left join companies companies on companies.id = usercompanies.company_id").
		Joins("left join users users on users.id = usercompanies.user_id").
		Joins("left join users u1 on u1.id = usercompanies.create_by").
		Joins("left join users u2 on u2.id = usercompanies.update_by").
		Joins("left join users u3 on u3.id = usercompanies.delete_by")

	err = conn.Migrator().CreateView(model.VIEW_USERCOMPANY, gorm.ViewOption{
		Replace: true,
		Query:   vUsercompany,
	})
	if err != nil {
		panic(err)
	}

}

// remove public schema
func down() {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	err = conn.Exec("DROP SCHEMA public CASCADE").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("CREATE SCHEMA public").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("GRANT ALL ON SCHEMA public TO postgres").Error
	if err != nil {
		panic(err)
	}

	err = conn.Exec("GRANT ALL ON SCHEMA public TO public").Error
	if err != nil {
		panic(err)
	}

	//// view
	//err = conn.Migrator().DropView(model.VIEW_USER)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// table
	//err = conn.Migrator().DropTable(&model.User{})

	if err != nil {
		panic(err)
	}
}

func seed() {
	now := time.Now()
	password, err := cryption.EncryptAES64("123456")
	if err != nil {
		panic(err)
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	userID := utils.GetUniqueID()
	companyID := utils.GetUniqueID()

	users := []model.User{
		{
			ID:          userID,
			RoleID:      utils.GetUniqueID(),
			Email:       "jihanlugas2@gmail.com",
			Username:    "jihanlugas",
			NoHp:        "6287770333043",
			Fullname:    "Jihan Lugas",
			Passwd:      password,
			PassVersion: 1,
			Active:      true,
			LastLoginDt: nil,
			PhotoID:     "",
			CreateBy:    userID,
			CreateDt:    now,
			UpdateBy:    userID,
			UpdateDt:    now,
		},
	}
	tx.Create(&users)

	companies := []model.Company{
		{
			ID:          companyID,
			Name:        "Jihan Company",
			Description: "This is my company",
			Address:     "Jl. Kehidupan",
			CreateBy:    userID,
			CreateDt:    now,
			UpdateBy:    userID,
			UpdateDt:    now,
			DeleteBy:    "",
			DeleteDt:    nil,
		},
	}
	tx.Create(&companies)

	usercompanies := []model.Usercompany{
		{
			UserID:           userID,
			CompanyID:        companyID,
			IsDefaultCompany: true,
			IsCreator:        true,
			CreateBy:         userID,
			CreateDt:         now,
			UpdateBy:         userID,
			UpdateDt:         now,
			DeleteBy:         "",
			DeleteDt:         nil,
		},
	}
	tx.Create(&usercompanies)

	err = tx.Commit().Error
	if err != nil {
		panic(err)
	}

}
