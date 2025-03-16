package cmds

import (
	"encoding/json"
	"net/http"

	"github.com/teamdunno/bottle/bot"
)

type GithubUser struct {
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	Url       string `json:"html_url"`
	ApiUrl    string `json:"url"`
}

type GithubRepo struct {
	Name        string     `json:"name"`
	FullName    string     `json:"full_name"`
	Description string     `json:"description"`
	Language    string     `json:"language"`
	Owner       GithubUser `json:"owner"`

	Stars    int `json:"stargazers_count"`
	Forks    int `json:"forks_count"`
	Watchers int `json:"watchers_count"`
}

func init() {
	registry.AddCommand("repo", repo)
	registry.SetHelp("repo", "show a github repo (github)")
}

func repo(ctx bot.Context) {
	resp, err := http.Get("https://api.github.com/repos/" + ctx.Args[0])
	if err != nil {
		ctx.Send("There was an error! " + err.Error())
		ctx.Send("Are you sure this repo exists, " + ctx.User + "?")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ctx.Send("There was an error! " + resp.Status)
		return
	}

	var repo GithubRepo
	err = json.NewDecoder(resp.Body).Decode(&repo)
	if err != nil {
		ctx.Send("There was an error! " + err.Error())
		return
	}

	ctx.Send(repo.FullName)
	ctx.Send(repo.Description)
	ctx.Sendf("Stars: %d, Forks: %d, Watchers: %d", repo.Stars, repo.Forks, repo.Watchers)
	ctx.Send(repo.Language)
}
