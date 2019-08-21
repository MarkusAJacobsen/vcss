package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"vcss/statuspage"
)

// Version Control Systems Status - A tool for fetching operational status from different VC providers

const GithubStatusEndpointSummary = "https://kctbh9vrtdwd.statuspage.io/api/v2/summary.json"
const BitBucketStatusEndpointSummary = "https://bqlf8qjztdtr.statuspage.io/api/v2/summary.json"

type Target string

const (
	Github    Target = "gh"
	Bitbucket Target = "bb"
)

func (t *Target) String() string {
	return string(*t)
}

func (t *Target) Set(value string) error {
	if len(*t) > 0 {
		return errors.New("target flag already set")
	}
	*t = Target(value)
	return nil
}

var target Target

func init() {
	usage := "which VCS to target - accepts gh (github), bb (bitbucket)"
	flag.Var(&target, "target", usage)
	flag.Var(&target, "t", usage)
}

func main() {
	printProgramLog()

	flag.Parse()

	var te string
	switch target {
	case "gh":
		te = GithubStatusEndpointSummary
		break
	case "bb":
		te = BitBucketStatusEndpointSummary
		break
	default:
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(0)
	}

	getVCStatusAndPrint(te)
}

func getVCStatusAndPrint(ep string) {
	stat, err := getStatus(ep)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	res := statuspage.Res{}
	err = json.Unmarshal(stat, &res)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	printGithubStatus(res)
}

func getStatus(ep string) (res []byte, err error) {
	req, err := http.Get(ep)
	if err != nil {
		return
	}
	defer req.Body.Close()

	res, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	return
}

func printGithubStatus(stat statuspage.Res) {
	fmt.Printf("Target VCS: \t%s\n", stat.Page.Name)
	fmt.Printf("Updated at: \t%s\n", stat.Page.UpdatedAt)
	fmt.Printf("Overall status: %s\n", stat.Status.Description)
	fmt.Println()

	for _, comp := range stat.Components {
		fmt.Printf("Module: \t%s\n", comp.Name)
		fmt.Printf("Updated at: \t%s\n", comp.UpdatedAt)
		fmt.Printf("Status: \t%s\n", comp.Status)
		fmt.Println()
	}
}

func printProgramLog() {
	fmt.Printf("\t _  _  ___  ___  ___\n")
	fmt.Printf("\t( \\/ )/ __)/ __)/ __)\n")
	fmt.Printf("\t \\  /( (__ \\__ \\\\__ \\\n")
	fmt.Printf("\t  \\/  \\___)(___/(___/\n")
	fmt.Println()
}
