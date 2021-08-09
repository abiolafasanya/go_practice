package main

import "strings"

func name(n string) (string, string) {
	s := strings.ToTitle(n)
	 fn := strings.Split(s, " ")

	 var fullname []string
	for _, v := range fn{
		fullname = append(fullname, v)
	}

	if len(fullname) > 1{
		return fullname[0], fullname[1]
	}

	return fullname[0], "_"
}