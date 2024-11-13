package main

import (
	"flag"
	"log"
	"tasklist_REST_API/internal/app/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", api.ConfigFilePath, "path to config file in .toml format")
}

func main() {

	flag.Parse()
	log.Println("It works")

	config := api.DefaultConfig()
	_, err := toml.DecodeFile(configPath, config) // Десериализиуете содержимое .toml файла
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
	}

	server := api.New(config)

	log.Fatal(server.Start())
}

/*
Реализовать REST API сервис для задач и использовать в качестве хранилища - sqlite db:
Можно использовать стандартный Mux, можно Gorilla/mux

---- ГОТОВО
// POST   /tasks/              :  создаёт задачу и возвращает её ID
// GET    /tasks/<taskid>      :  возвращает одну задачу по её ID
// GET    /tasks/              :  возвращает все задачи
// GET    /tags/<tagname>      :  возвращает список задач с заданным тегом

----- В ПРОЦЕССЕ


----- НЕ ГОТОВО

// DELETE /tasks/<taskid>      :  удаляет задачу по ID
// DELETE /tasks/              :  удаляет все задачи
// GET    /due/<yy>/<mm>/<dd> :  возвращает список задач, запланированных на указанную дату


type task struct {
	id int
	Text string
	Tags []string
	Due time.Time // deadline date
}
*/
