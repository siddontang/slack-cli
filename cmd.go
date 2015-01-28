package main

var helpCommands = [][]string{
	[]string{"chat.delete", "ts channel", ""},
	[]string{"chat.postMessage", "channel text [username] [parse] [link_names] [attachments] [unfurl_links] [unfurl_media] [icon_url] [icon_emoji]",
		"attachments is a json format string"},
	[]string{"chat.update", "ts channel text", ""},
	[]string{"emoji.list", "", ""},
	[]string{"im.close", "channel", ""},
	[]string{"im.history", "channel [latest] [oldest] [count]", "latest is a timestamp, default is now, oldest default is 0"},
	[]string{"im.list", "", ""},
	[]string{"im.mark", "channel ts", "ts is a timestamp"},
	[]string{"im.open", "user", ""},
	[]string{"search.all", "query [sort] [sort_dir] [highlight] [count] [page]", "sort is score or timestamp, default is score, sort_dir is asc or desc, default is desc, pass 1 to enable highlight"},
	[]string{"search.files", "query [sort] [sort_dir] [highlight] [count] [page]", "sort is score or timestamp, default is score, sort_dir is asc or desc, default is desc, pass 1 to enable highlight"},
	[]string{"search.messages", "query [sort] [sort_dir] [highlight] [count] [page]", "sort is score or timestamp, default is score, sort_dir is asc or desc, default is desc, pass 1 to enable highlight"},
	[]string{"stars.list", "[user] [count] [page]", "default user is your token user, default count is 100 and page is 1"},
	[]string{"users.getPresence", "user", ""},
	[]string{"users.info", "user", ""},
	[]string{"users.list", "", ""},
	[]string{"users.setActive", "", ""},
	[]string{"users.setPresence", "presence", "presence is auto or away"},
}
