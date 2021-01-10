package main

import (
	"log"

	"github.com/vnkrtv/go-vk-tracker/pkg/service"
)

const (
	cfgPath    = "config/config.json"
	usersPath  = "config/users.json"
	groupsPath = "config/groups.json"
)

func main() {
	cfg, err := service.GetConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	vkLoader, err := service.NewVKLoaderFromCfg(cfg)
	if err != nil {
		log.Fatal(err)
	}

	usersIDs, err := service.GetUsersIDs(usersPath)
	if err != nil {
		log.Fatal(err)
	}
	groupsScreenNames, err := service.GetGroupsScreenNames(groupsPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := vkLoader.InitDB(); err != nil {
		log.Println(err)
	}
	if err := vkLoader.AddTrackingUsers(usersIDs); err != nil {
		log.Fatal(err)
	}
	log.Println()
	if err := vkLoader.AddTrackingGroups(groupsScreenNames); err != nil {
		log.Fatal(err)
	}
	log.Println()

	for {
		if err := vkLoader.LoadUsersInfo(); err != nil {
			log.Println(err)
		} else {
			log.Printf("Loaded all tracking users info")
		}
		log.Println()
		log.Println()
		if err := vkLoader.LoadGroupsInfo(); err != nil {
			log.Println(err)
		} else {
			log.Printf("Loaded all tracking groups info\n")
		}
		log.Println()
		log.Println()
		vkLoader.Sleep(10 * 1000)
	}
}
