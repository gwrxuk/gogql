package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gwrxuk/graphql/graph/generated"
	"gwrxuk/graphql/graph/model"
	"strconv"

	gql "github.com/machinebox/graphql"
)

func (r *queryResolver) LastProject(ctx context.Context, n int) ([]*model.Projects, error) {
	client := gql.NewClient("https://gitlab.com/api/graphql")

	req := gql.NewRequest(`
	  query last_projects($n: Int) {
	  projects(last:$n) {
	    nodes {
	      name
	      description
	      forksCount
	    }
	  }
	}
	`)

	req.Var("n", n)
	var res ResponseStruct
	if err := client.Run(context.Background(), req, &res); err != nil {
		panic(err)
	}

	var all_projects []*model.Projects

	for i, s := range res.Projects.Nodes {
		fmt.Println(i, strconv.Itoa(s.ForksCount))
		var project model.Projects
		project.Info = s.Name + ", " + strconv.Itoa(s.ForksCount)
		all_projects = append(all_projects, &project)
	}
	return all_projects, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type ResponseStruct struct {
	Projects ProjectsStruct `json:"projects"`
}
type ProjectsStruct struct {
	Nodes []NodeStruct `json:"nodes"`
}
type NodeStruct struct {
	Name        string `json:"name"`
	Description string `json:"description, omitempty"`
	ForksCount  int    `json:"forksCount"`
}
