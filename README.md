slack-cli is a command tool to communicate with [slack](https://slack.com).

## Why?

I just want to study and know [slack API](https://api.slack.com) easily.

## Use

```
shell>slack-cli -token=your_token
slack>users.list
slack>users.info user=U023BECGF
slack>stars.list user=U023BECGF page=1 count=100
```

## todo

+ add help description for commands
+ more test

## Limitation

Use Redis linenoise lib, so Windows may be not supported. :-)