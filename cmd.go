package main

var helpCommands = [][]string{
	[]string{"auth.test", "", ""},

	[]string{"channels.archive", "channel", ""},
	[]string{"channels.create", "name", ""},
	[]string{"channels.history", "channel [latest] [oldest] [count]", ""},
	[]string{"channels.info", "channel", ""},
	[]string{"channels.invite", "channel user", ""},
	[]string{"channels.join", "name", ""},
	[]string{"channels.kick", "channel user", ""},
	[]string{"channels.leave", "channel", ""},
	[]string{"channels.list", "[exclude_archived]", ""},
	[]string{"channels.mark", "channel ts", ""},
	[]string{"channels.rename", "channel name", ""},
	[]string{"channels.setPurpose", "channel purpose", ""},
	[]string{"channels.setTopic", "channel topic", ""},
	[]string{"channels.unarchive", "channel", ""},

	[]string{"groups.archive", "channel", ""},
	[]string{"groups.close", "channel", ""},
	[]string{"groups.create", "name", ""},
	[]string{"groups.createChild", "channel", ""},
	[]string{"groups.history", "channel [latest] [oldest] [count]", ""},
	[]string{"groups.invite", "channel user", ""},
	[]string{"groups.kick", "channel user", ""},
	[]string{"groups.leave", "channel", ""},
	[]string{"groups.list", "[exclude_archived]", ""},
	[]string{"groups.mark", "channel ts", ""},
	[]string{"groups.open", "channel", ""},
	[]string{"groups.rename", "channel name", ""},
	[]string{"groups.setPurpose", "channel purpose", ""},
	[]string{"groups.setTopic", "channel topic", ""},
	[]string{"groups.unarchive", "channel", ""},

	[]string{"files.info", "file [count] [page] [count]", ""},
	[]string{"files.list", "[user] [ts_from] [ts_to] [types] [count] [page]", ""},
	[]string{"files.upload", "[file] [content] [filetype] [filename] [title] [initial_comment] [channels]", ""},

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
