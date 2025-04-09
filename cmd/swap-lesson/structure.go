package main

type node struct {
	IsDir bool
	Files []string
}

var GameDirStructure = map[string]node{
	"game": {
		IsDir: false,
		Files: nil,
	},
	"game.set": {
		IsDir: false,
		Files: nil,
	},
	"tr": {
		IsDir: true,
		Files: nil,
	},
	"sounds": {
		IsDir: true,
		Files: nil,
	},
	"logos": {
		IsDir: true,
		Files: nil,
	},
	"information.txt": {
		IsDir: false,
		Files: nil,
	},
	"version": {
		IsDir: false,
		Files: nil,
	},
}
