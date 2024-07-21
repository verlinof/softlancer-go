package database

import (
	"log"

	"github.com/verlinof/softlancer-go/config/db_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var errConnection error

	dsnMysql := db_config.DB_USER + ":" + db_config.DB_PASSWORD + "@tcp(" + db_config.DB_HOST + ":" + db_config.DB_PORT + ")/" + db_config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, errConnection = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{}) 

	if(errConnection != nil) {
		panic(errConnection)
	}

	log.Println("Database connected!")
}

// // migrateCmd represents the migrate command
// var migrateCmd = &cobra.Command{
// 	Use:   "migrate",
// 	Short: "Run database migrations",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//Load .env file
// 		err := godotenv.Load()
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		//Connect to Database
// 		ConnectDatabase()

// 		// Migrate the schema
// 		err = DB.AutoMigrate(&models.User{})
// 		if err != nil {
// 			log.Fatal("failed to migrate database:", err)
// 		}

// 		log.Println("Database migration completed")
// 	},
// }

// func init() {
// 	migrateCmd.Root().AddCommand(migrateCmd)
// }