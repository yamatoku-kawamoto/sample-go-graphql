package graph

import "gographql/graph/model"

var todos = []*model.Todo{
	{ID: "0", Text: "mail check", User: users[0]},
	{ID: "1", Text: "MTG", User: users[0]},
}

var users = []*model.User{
	{ID: "0", Name: "Alice"},
}
