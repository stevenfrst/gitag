package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"gitlab-tag-hook/internal"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	appPath := "."
	r, err := git.PlainOpen(appPath)

	if err != nil {
		os.Exit(1)
	}
	hashCommit, err := r.Head()
	worktree, err := r.CommitObject(hashCommit.Hash())
	commitMsg := worktree.Message
	//log.Println(commitMsg)
	regexFix, _ := regexp.Compile(`([Ff]ix|[Ff]ixing|[Ff]ixed):`)                                                                              // Fix: fix: Fixed:
	regexMinor, _ := regexp.Compile(`([Mm]inor:|[Nn]ew [Ff]eature:|[Mm]enambah:|[Aa]dd:|[Rr]emove:|[Mm]odify:|[Mm]engganti:|[Mm]enghapus:).*`) // Minor: New Feature: Menambah:
	regexMajor, _ := regexp.Compile(`([Vv]ersi [Bb]aru|[Nn]ew [Vv]ersion|[Mm]igrate|[Mm]ajor):.*`)                                             // Versi Baru: New Version: Migrate:

	if err != nil {
		os.Exit(1)
	}
	//fmt.Println("hash", hashCommit)
	tags, err := internal.GetTags(r)

	s := strings.Split(tags, ".")

	PatchVer, _ := strconv.Atoi(s[2])
	MinorVer, _ := strconv.Atoi(s[1])
	MajorVer, _ := strconv.Atoi(s[0])

	var newTags string
	if regexFix.MatchString(commitMsg) {
		PatchVer++
		newTags = internal.ConvSemVersion(strconv.Itoa(MajorVer), strconv.Itoa(MinorVer), strconv.Itoa(PatchVer))
	} else if regexMinor.MatchString(commitMsg) {
		PatchVer = 0
		MinorVer++
		newTags = internal.ConvSemVersion(strconv.Itoa(MajorVer), strconv.Itoa(MinorVer), strconv.Itoa(PatchVer))
	} else if regexMajor.MatchString(commitMsg) {
		MajorVer++
		MinorVer = 0
		PatchVer = 0
		newTags = internal.ConvSemVersion(strconv.Itoa(MajorVer), strconv.Itoa(MinorVer), strconv.Itoa(PatchVer))
	} else {
		newTags = tags
	}

	fmt.Println(newTags)
	os.Exit(0)
}
