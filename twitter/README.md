# twitter apiの試し場
## 申請手順(2020/1/13申請)
https://digitalnavi.net/internet/3072/

## 申請時の文言
### How will you use the Twitter API or Twitter data?
```
I would like to output data to my web application in the future using the results of analyzing the twitter timeline.
Specifically, I want to use the twitter API to analyze tweets about disasters.
I want to create a web app that will be useful in times of disaster.
Thank you.
```

### Are you planning to analyze Twitter data?
```
From tweets at the time of the disaster, tweets related to the disaster are extracted and analyzed.
Specifically, morphological analysis is used.
```

### Will your app use Tweet, Retweet, like, follow, or Direct Message functionality?
```
At the moment, I want to use tweets and retweets as analysis functions, but I also want to use other functions for analysis.
```

### Do you plan to display Tweets or aggregate data about Twitter content outside of Twitter?
```
I want to output the result of summarizing data analyzed from tweets on my web application.
Specifically, we plan to aggregate and output data obtained by analyzing weather and geographic information contained in tweets.
```

## twurl
https://github.com/twitter/twurl<br>
`twurl authorize --consumer-key key --consumer-secret secret`<br>
※key/secretは`https://developer.twitter.com/` でログイン後の「Apps > disaster-analyzer」のとこから取得<br>

https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-user_timeline<br>
`twurl -H "api.twitter.com" "/1.1/statuses/user_timeline.json" | jq . > data/timeline.json`<br>
