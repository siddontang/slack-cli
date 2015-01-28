package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"strconv"
	"strings"
)

func extractParams(args []string) map[string]string {
	params := make(map[string]string)
	for _, arg := range args {
		seps := strings.SplitN(arg, "=", 2)
		if len(seps) == 2 {
			params[seps[0]] = strings.Trim(seps[1], "\"'")
		} else {
			params[seps[0]] = ""
		}
	}
	return params
}

func (s *Slack) handle(cmd string, args []string) (interface{}, error) {
	cmds := strings.Split(cmd, ".")
	if len(cmds) != 2 {
		return nil, fmt.Errorf("cmd must be type.action format, not %s", cmd)
	}

	var v interface{}
	var err error

	tp := strings.ToLower(cmds[0])
	action := strings.ToLower(cmds[1])
	params := extractParams(args)

	switch tp {
	case "api":
		err = fmt.Errorf("%s has not been supported", tp)
	case "auth":
		v, err = s.s.AuthTest()
	case "channels":
		v, err = s.handleChannels(action, params)
	case "chat":
		v, err = s.handleChat(action, params)
	case "emoji":
		v, err = s.handleEmoji(action, params)
	case "files":
		v, err = s.handleFiles(action, params)
	case "groups":
		v, err = s.handleGroups(action, params)
	case "im":
		v, err = s.handleIM(action, params)
	case "oauth":
		err = fmt.Errorf("%s has not been supported", tp)
	case "rtm":
		err = fmt.Errorf("%s has not been supported", tp)
	case "search":
		v, err = s.handleSearch(action, params)
	case "stars":
		v, err = s.handleStars(action, params)
	case "users":
		v, err = s.handleUsers(action, params)
	default:
		return nil, fmt.Errorf("unsupported api type %s", cmds[0])
	}

	return v, err
}

func getIntParam(params map[string]string, key string, defValue int) int {
	v, ok := params[key]
	if !ok {
		return defValue
	} else {
		vv, err := strconv.Atoi(v)
		if err != nil {
			return defValue
		} else {
			return vv
		}
	}
}

func getStringParam(params map[string]string, key string, defValue string) string {
	v, ok := params[key]
	if !ok {
		return defValue
	} else {
		return v
	}
}

func getBoolParam(params map[string]string, key string, defValue bool) bool {
	v, ok := params[key]
	if !ok {
		return defValue
	} else if v == "true" {
		return true
	} else {
		return false
	}
}

func (s *Slack) handleChannels(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error

	switch action {
	case "archive":
		err = s.s.ArchiveChannel(params["channel"])
	case "create":
		ch, err := s.s.CreateChannel(params["name"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"channel": ch,
		}
	case "history":
		historyParam := slack.HistoryParameters{}
		historyParam.Latest = getStringParam(params, "latest", slack.DEFAULT_HISTORY_LATEST)
		historyParam.Oldest = getStringParam(params, "latest", slack.DEFAULT_HISTORY_OLDEST)
		historyParam.Count = getIntParam(params, "count", slack.DEFAULT_HISTORY_COUNT)
		v, err = s.s.GetChannelHistory(params["channel"], historyParam)
	case "info":
		ch, err := s.s.GetChannelInfo(params["channel"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"channel": ch,
		}
	case "invite":
		ch, err := s.s.InviteUserToChannel(params["channel"], params["user"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"channel": ch,
		}
	case "join":
		ch, err := s.s.JoinChannel(params["name"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"channel": ch,
		}
	case "kick":
		err = s.s.KickUserFromChannel(params["channel"], params["user"])
	case "leave":
		_, err = s.s.LeaveChannel(params["channel"])
	case "list":
		exclude := getIntParam(params, "exclude_archived", 0)
		chs, err := s.s.GetChannels(exclude == 1)
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"channels": chs,
		}
	case "mark":
		err = s.s.SetChannelReadMark(params["channel"], params["ts"])
	case "rename":
		ch, err := s.s.RenameChannel(params["channel"], params["name"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"channel": ch,
		}
	case "setpurpose":
		purpose, err := s.s.SetChannelPurpose(params["channel"], params["purpose"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"purpose": purpose,
		}
	case "settopic":
		topic, err := s.s.SetChannelTopic(params["channel"], params["topic"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"topic": topic,
		}
	case "unarchive":
		err = s.s.UnarchiveChannel(params["channel"])
	default:
		return nil, fmt.Errorf("invalid files action %s", action)
	}

	return v, err
}

func (s *Slack) handleGroups(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error

	switch action {
	case "archive":
		err = s.s.ArchiveGroup(params["channel"])
	case "close":
		noop, closed, err := s.s.CloseGroup(params["channel"])
		if err != nil {
			return nil, err
		}
		v = map[string]bool{
			"no_op":          noop,
			"already_closed": closed,
		}
	case "create":
		group, err := s.s.CreateGroup(params["name"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"group": group,
		}
	case "createchild":
		group, err := s.s.CreateChildGroup(params["channel"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"group": group,
		}
	case "history":
		historyParam := slack.HistoryParameters{}
		historyParam.Latest = getStringParam(params, "latest", slack.DEFAULT_HISTORY_LATEST)
		historyParam.Oldest = getStringParam(params, "latest", slack.DEFAULT_HISTORY_OLDEST)
		historyParam.Count = getIntParam(params, "count", slack.DEFAULT_HISTORY_COUNT)
		v, err = s.s.GetGroupHistory(params["channel"], historyParam)
	case "invite":
		group, in, err := s.s.InviteUserToGroup(params["channel"], params["user"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"already_in_group": in,
			"group":            group,
		}
	case "kick":
		err = s.s.KickUserFromGroup(params["channel"], params["user"])
	case "leave":
		err = s.s.LeaveGroup(params["channel"])
	case "list":
		exclude := getIntParam(params, "exclude_archived", 0)
		groups, err := s.s.GetGroups(exclude == 1)
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"groups": groups,
		}
	case "open":
		noop, opened, err := s.s.OpenGroup(params["channel"])
		if err != nil {
			return nil, err
		}
		v = map[string]bool{
			"no_op":        noop,
			"already_open": opened,
		}
	case "mark":
		err = s.s.SetGroupReadMark(params["channel"], params["ts"])
	case "rename":
		group, err := s.s.RenameGroup(params["channel"], params["name"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"group": group,
		}
	case "setpurpose":
		purpose, err := s.s.SetGroupPurpose(params["channel"], params["purpose"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"purpose": purpose,
		}
	case "settopic":
		topic, err := s.s.SetGroupTopic(params["channel"], params["topic"])
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"topic": topic,
		}
	case "unarchive":
		err = s.s.UnarchiveGroup(params["channel"])
	default:
		return nil, fmt.Errorf("invalid groups action %s", action)
	}

	return v, err
}

func (s *Slack) handleFiles(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error

	switch action {
	case "info":
		count := getIntParam(params, "count", slack.DEFAULT_FILES_COUNT)
		page := getIntParam(params, "page", slack.DEFAULT_FILES_PAGE)
		files, comments, pages, err := s.s.GetFileInfo(params["file"], count, page)
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"file":     files,
			"comments": comments,
			"paging":   pages,
		}
	case "list":
		listParam := slack.GetFilesParameters{}
		listParam.UserId = getStringParam(params, "user", slack.DEFAULT_FILES_USERID)
		listParam.TimestampFrom = slack.JSONTime(getIntParam(params, "ts_from", slack.DEFAULT_FILES_TS_FROM))
		listParam.TimestampTo = slack.JSONTime(getIntParam(params, "ts_to", slack.DEFAULT_FILES_TS_TO))
		listParam.Types = getStringParam(params, "types", slack.DEFAULT_FILES_TYPES)
		listParam.Count = getIntParam(params, "count", slack.DEFAULT_FILES_COUNT)
		listParam.Page = getIntParam(params, "page", slack.DEFAULT_FILES_PAGE)

		files, pages, err := s.s.GetFiles(listParam)
		if err != nil {
			return nil, err
		}

		v = map[string]interface{}{
			"files":  files,
			"paging": pages,
		}

	case "upload":
		uploadParams := slack.FileUploadParameters{}
		uploadParams.File = getStringParam(params, "file", "")
		uploadParams.Content = getStringParam(params, "content", "")
		uploadParams.Filetype = getStringParam(params, "filetype", "")
		uploadParams.Filename = getStringParam(params, "filename", "")
		uploadParams.Title = getStringParam(params, "title", "")
		uploadParams.InitialComment = getStringParam(params, "initial_comment", "")
		channels := getStringParam(params, "channels", "")
		uploadParams.Channels = strings.Split(channels, ",")
		file, err := s.s.UploadFile(uploadParams)
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"file": file,
		}
	default:
		return nil, fmt.Errorf("invalid files action %s", action)
	}
	return v, err
}

func (s *Slack) handleChat(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error

	switch action {
	case "delete":
		ch, ts, err := s.s.DeleteMessage(params["channel"], params["ts"])
		if err != nil {
			return nil, err
		}
		v = map[string]string{
			"channel": ch,
			"ts":      ts,
		}
	case "postmessage":
		postParam := slack.PostMessageParameters{}
		postParam.Username = getStringParam(params, "username", slack.DEFAULT_MESSAGE_USERNAME)
		postParam.Parse = getStringParam(params, "parse", slack.DEFAULT_MESSAGE_PARSE)
		postParam.LinkNames = getIntParam(params, "link_names", slack.DEFAULT_MESSAGE_LINK_NAMES)

		postParam.UnfurlLinks = getBoolParam(params, "unfurl_links", slack.DEFAULT_MESSAGE_UNFURL_LINKS)
		postParam.UnfurlMedia = getBoolParam(params, "unfurl_media", slack.DEFAULT_MESSAGE_UNFURL_MEDIA)
		postParam.IconURL = getStringParam(params, "icon_url", slack.DEFAULT_MESSAGE_ICON_URL)
		postParam.IconEmoji = getStringParam(params, "icon_emoji", slack.DEFAULT_MESSAGE_ICON_EMOJI)

		if err = json.Unmarshal([]byte(getStringParam(params, "attachments", "[]")), &postParam.Attachments); err != nil {
			return nil, err
		}

		ch, ts, err := s.s.PostMessage(params["channel"], params["text"], postParam)
		if err != nil {
			return nil, err
		}

		v = map[string]string{
			"channel": ch,
			"ts":      ts,
		}
	case "update":
		ch, ts, text, err := s.s.UpdateMessage(params["channel"], params["ts"], params["text"])
		if err != nil {
			return nil, err
		}
		v = map[string]string{
			"channel": ch,
			"ts":      ts,
			"text":    text,
		}
	default:
		return nil, fmt.Errorf("invalid chat action %s", action)
	}

	return v, err
}

func (s *Slack) handleEmoji(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error
	switch action {
	case "list":
		m, err := s.s.GetEmoji()
		if err != nil {
			return nil, err
		}

		v = map[string]interface{}{
			"emoji": m,
		}

	default:
		return nil, fmt.Errorf("invalid emoji action %s", action)
	}
	return v, err
}

func (s *Slack) handleIM(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error

	switch action {
	case "close":
		noop, closed, err := s.s.CloseIMChannel(params["channel"])
		if err != nil {
			return nil, err
		}

		v = map[string]bool{
			"no_op":          noop,
			"already_closed": closed,
		}
	case "history":
		historyParms := slack.HistoryParameters{}
		historyParms.Latest = getStringParam(params, "latest", slack.DEFAULT_HISTORY_LATEST)

		historyParms.Oldest = getStringParam(params, "oldest", slack.DEFAULT_HISTORY_OLDEST)

		historyParms.Count = getIntParam(params, "count", slack.DEFAULT_HISTORY_COUNT)

		v, err = s.s.GetIMHistory(params["channel"], historyParms)

	case "list":
		ims, err := s.s.GetIMChannels()
		if err != nil {
			return nil, err
		}

		v = map[string]interface{}{
			"ims": ims,
		}
	case "mark":
		err = s.s.MarkIMChannel(params["channel"], params["ts"])
	case "open":
		noop, opened, ch, err := s.s.OpenIMChannel(params["user"])
		if err != nil {
			return nil, err
		}

		v = map[string]interface{}{
			"no_op":        noop,
			"already_open": opened,
			"channel": map[string]string{
				"id": ch,
			},
		}

	default:
		return nil, fmt.Errorf("invalid im action %s", action)
	}

	return v, err
}

func (s *Slack) handleSearch(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error

	searchParams := slack.SearchParameters{}
	searchParams.Sort = getStringParam(params, "sort", slack.DEFAULT_SEARCH_SORT)
	searchParams.SortDirection = getStringParam(params, "sort_dir", slack.DEFAULT_SEARCH_SORT_DIR)
	highlight := getIntParam(params, "highlight", 0)
	searchParams.Highlight = (highlight == 1)

	searchParams.Count = getIntParam(params, "count", slack.DEFAULT_SEARCH_COUNT)

	searchParams.Page = getIntParam(params, "page", slack.DEFAULT_SEARCH_PAGE)

	query := params["query"]

	switch action {
	case "all":
		m, f, err := s.s.Search(query, searchParams)
		if err != nil {
			return nil, err
		}

		v = map[string]interface{}{
			"query":    query,
			"messages": m,
			"files":    f,
		}

	case "files":
		f, err := s.s.SearchFiles(query, searchParams)
		if err != nil {
			return nil, err
		}

		v = map[string]interface{}{
			"query": query,
			"files": f,
		}
	case "messages":
		m, err := s.s.SearchMessages(query, searchParams)
		if err != nil {
			return nil, err
		}

		v = map[string]interface{}{
			"query":    query,
			"messages": m,
		}

	default:
		return nil, fmt.Errorf("invalid search action %s", action)
	}

	return v, err
}

func (s *Slack) handleStars(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error
	switch action {
	case "list":
		starsParams := slack.StarsParameters{}
		starsParams.User = params["user"]

		starsParams.Count = getIntParam(params, "count", slack.DEFAULT_STARS_COUNT)
		starsParams.Page = getIntParam(params, "page", slack.DEFAULT_STARS_PAGE)

		items, paging, err := s.s.GetStarred(starsParams)
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"items":  items,
			"paging": paging,
		}
	default:
		return nil, fmt.Errorf("invalid stars action %s", action)
	}

	return v, err
}

func (s *Slack) handleUsers(action string, params map[string]string) (interface{}, error) {
	var v interface{}
	var err error
	switch action {
	case "getpresence":
		v, err = s.s.GetUserPresence(params["user"])
	case "info":
		v, err = s.s.GetUserInfo(params["user"])
	case "list":
		v, err = s.s.GetUsers()
	case "setactive":
		err = s.s.SetUserAsActive()
	case "setpresence":
		err = s.s.SetUserPresence(params["presence"])
	default:
		return nil, fmt.Errorf("invalid users action %s", action)
	}

	return v, err
}
