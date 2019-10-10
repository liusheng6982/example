package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	githppt "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Project struct {
	GitUrl string `json:"http_url_to_repo"`
	Name   string `json:"name"`
}

// Basic example of how to clone a repository using clone options.
func main() {
	if len(os.Args) < 5 {
		println("请输入项目名称！")
		println("Usage:  back_pom userName password private-token branch-name")
		return
	}
	httpGet(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
}

func httpGet(userName, password, privateToken, branchName string) {
	gitUrl := "http://git.taocaimall.com:2469/api/v4/projects?private_token=%s&per_page=1000"
	resp, err := http.Get(fmt.Sprintf(gitUrl, privateToken))
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	projects := make([]Project, 1000)
	json.Unmarshal(body, &projects)

	for _, v := range projects {
		//f v.Name == "o2o-order" {
		sigle(v, branchName, userName, password)
		//}
	}

}

func sigle(project Project, branchName, userName, password string) {
	fmt.Println(project.GitUrl)
	fmt.Println(userName + "    " + password)
	workPath := "./" + project.Name
	r, err := git.PlainClone(workPath, false, &git.CloneOptions{
		Auth: &githppt.BasicAuth{
			Username: userName,
			Password: password,
		},
		URL: project.GitUrl,
	})
	if err != nil {
		fmt.Println("PlainClone:" + err.Error())
		r, err = git.PlainOpen(workPath)
		if err != nil {
			fmt.Println("PlainOpen:" + err.Error())
			removeDir(workPath)
			return
		}
	}
	w, err := r.Worktree()
	if err != nil {
		fmt.Println("Worktree:" + err.Error())
		removeDir(workPath)
		return
	}

	ref, err := r.Head()
	if err != nil {
		fmt.Println("head:" + err.Error())
		removeDir(workPath)
		return
	}
	fmt.Println(ref.Hash())

	referenceName := plumbing.ReferenceName(branchName)
	fmt.Println("branch:=" + referenceName)
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewRemoteReferenceName("origin/", branchName),
	})
	if err != nil {
		fmt.Println("Checkout:" + err.Error())
		removeDir(workPath)
		return
	}
	//r.Head()

	err = os.RemoveAll(fmt.Sprintf("%s%s", workPath, "/.git"))
	if err != nil {
		fmt.Println("RemoveAll .git:" + err.Error())
		println(err.Error())
	}

	removeOtherFile(workPath, false)

}

func removeDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Println("RemoveAll dir:" + err.Error())
		println(err.Error())
	}
}

func removeOtherFile(path string, delDir bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		println("removeOtherFile  ioutil.ReadDir:" + path + "    " + err.Error())
		return
	}
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			newFileName := path + "/" + fileName
			removeOtherFile(newFileName, true)
			if delDir {
				os.Remove(newFileName)
			}
		} else {
			if strings.Index(file.Name(), "pom.xml") == -1 {
				os.Remove(path + "/" + fileName)
			}
		}
	}

}
