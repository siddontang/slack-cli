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
		err = fmt.Errorf("%s has not been supported", tp)
	case "channels":
		err = fmt.Errorf("%s has not been supported", tp)
	case "chat":
		v, err = s.handleChat(action, params)
	case "emoji":
		v, err = s.handleEmoji(action, params)
	case "files":
		err = fmt.Errorf("%s has not been supported", tp)
	case "groups":
		err = fmt.Errorf("%s has not been supported", tp)
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
		return nil, fmt.Errorf("invalid api type %s", cmds[0])
	}

	return v, err
}

func getIntParam(params map[string]string, key string, defValue int) (int, error) {
	v, ok := params[key]
	if !ok {
		return defValue, nil
	} else {
		return strconv.Atoi(v)
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
		postParam.LinkNames, err = getIntParam(params, "link_names", slack.DEFAULT_MESSAGE_LINK_NAMES)
		if err != nil {
			return nil, err
		}

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

		if historyParms.Count, err = getIntParam(params, "count", slack.DEFAULT_HISTORY_COUNT); err != nil {
			return nil, err
		}

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
	if highlight, err := getIntParam(params, "highlight", 0); err != nil {
		return nil, err
	} else {
		searchParams.Highlight = (highlight == 1)
	}

	if searchParams.Count, err = getIntParam(params, "count", slack.DEFAULT_SEARCH_COUNT); err != nil {
		return nil, err
	}

	if searchParams.Page, err = getIntParam(params, "page", slack.DEFAULT_SEARCH_PAGE); err != nil {
		return nil, err
	}

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

		if starsParams.Count, err = getIntParam(params, "count", slack.DEFAULT_STARS_COUNT); err != nil {
			return nil, err
		}
		if starsParams.Page, err = getIntParam(params, "page", slack.DEFAULT_STARS_PAGE); err != nil {
			return nil, err
		}

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
