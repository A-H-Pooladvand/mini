package main

import (
	"kraken/internal/bootstrap"
)

func main() {
	bootstrap.Boot()

	// Todo: Implement console command to run database migration files.
	// Todo: Implement easy-to-use Mysql access OR repository pattern.

	// io := style.
	// 	NewConsoleStyler().
	// 	AddInputArgument(
	// 		argument.
	// 			New("name", argument.REQUIRED),
	// 	).
	// 	AddInputOption(
	// 		option.
	// 			New("foo", option.NONE).
	// 			SetShortcut("f"),
	// 	).
	// 	ParseInput().
	// 	ValidateInput()
	//
	// // enable stylish errors
	// defer io.HandleRuntimeException()
	//
	// //
	// // Do what you want with console and args !
	// //
	//
	// io.Title("Starting console")
	//
	// io.TextArray([]string{
	// 	"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	// 	"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	// 	"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	// 	"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	// })
	//
	// io.Note(
	// 	fmt.Sprintf(
	// 		"name argument value '%s'",
	// 		io.GetInput().GetArgument("name"),
	// 	),
	// )
	//
	// if option.DEFINED == io.GetInput().GetOption("foo") {
	// 	io.Success("foo option is set")
	// }
	//
	// panic("this error will be stylish!")

	// dsn := "root:1@tcp(127.0.0.1:3306)/news?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// news := models.News{
	// 	Title: "Hello",
	// 	Body:  "World",
	// }
	// db.Create(&news)
}
