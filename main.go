package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glacier"
)

type Archive struct {
	VaultARN      string
	InventoryDate string
	ArchiveList   []struct {
		ArchiveId          string
		ArchiveDescription string
		CreationDate       string
		Size               int
		SHA256TreeHash     string
	}
}

var (
	AccountId = "043039367084"
	VaultName = "Backups"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})

	if err != nil {
		panic(err)
	}

	svc := glacier.New(sess)

	file, err := os.Open("output.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var archive Archive

	json.Unmarshal(bytes, &archive)

	alist := archive.ArchiveList

	sort.Slice(alist[:], func(i, j int) bool {
		return alist[i].Size > alist[j].Size
	})

	for _, a := range alist {
		fmt.Println(a.Size)

		input := &glacier.DeleteArchiveInput{
			AccountId: &AccountId,
			ArchiveId: &a.ArchiveId,
			VaultName: &VaultName,
		}

		result, err := svc.DeleteArchive(input)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(result)
		}

		time.Sleep(50 * time.Millisecond)
	}
}
