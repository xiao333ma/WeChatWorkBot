package utils

import (
	"WeChatWorkRobot/model"
	"strings"
)

func GreenString (s string) string  {
	return `<font color="info">` + s + `</font>`
}
func GrayString (s string) string  {
	return `<font color="comment">` + s + `</font>`
}
func OrangeString (s string) string  {
	return `<font color="warning">` + s + `</font>`
}

func Title(level int, s string) string  {
	t := ""
	for i := 0; i < level ; i++  {
		t += "#"
	}
	return  t + " " + s
}

func Blod(s string) string {
	return "**" + s + "**"
}

func Link(linkTitle string, url string) string  {
	return "[" + linkTitle + "](" + url + ")"
}

func Code(s string) string {
	return "`" + s + "`"
}

func Ref(s string) string  {
	return "> " + s
}

func Newline() string  {
	return "\n"
}

func WhiteSpace () string  {
	return " "
}

func CommitMessage(hookModel model.GitLabHookModel) string  {
	line := ""
	for i := len(hookModel.Commits) - 1; i >= 0; i-- {
		s := hookModel.Commits[i].Message
		arr := strings.Split(s, "\n")
		for _, value := range arr {
			line += Ref(value) + Newline()
		}
	}
	return line + Newline()
}