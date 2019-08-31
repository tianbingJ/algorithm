package leetcode

import "math"

/**
有个坑:twitterId的大小不能保证先后关系,添加一个时间递增戳字段，同时用map保存当前twitter和时间戳的关系。

用一个map记录用户关注的人， 一个map记录用户发布的twitter。

这里有个优化的地方是，如何快速获取用户发布了哪些twitter:
首先可以通过followerMap某个用户关注了哪些人，即一个列表follows[]int
	同时也知道一个user发布的twitter列表，那么问题变成：
	一个用户关注了k个人，user1, ..., userK， K个人twitter列表
	user1  l1 = [twitter11, twitter12 ....., twitter1End]
    user2  l2 = [twitter21, twitter22, ...., twitter2End]
    .....
    userk  lk = [twitterk1, twitterk2, ..., twitter3End]
    最新的twitter一定在l1到lk里的某个列表twitter最右侧。
    取出这条twitter之后，需要把这个位置保存下来，假设某个列表中这次取的下表twitter2End， 下次需要考虑的位置是twitter2End - 1。
    这个类似丑数系列的思想，用steps数据保存k个列表下一个需要考虑的位置。

其他操作都比较简单。
*/
type Twitter struct {
	//user_id -> 关注的人
	followerMap map[int][]int

	//user_id -> twitterMap
	twitterMap map[int][]int
	timestamp  int
	//index timestamp
	timestampMap map[int]int
}

func Constructor1() Twitter {
	return Twitter{
		followerMap:  make(map[int][]int),
		twitterMap:   make(map[int][]int),
		timestampMap: make(map[int]int),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.timestampMap[tweetId] = this.timestamp
	if v, ok := this.twitterMap[userId]; ok {
		this.twitterMap[userId] = append(v, tweetId)
	} else {
		this.twitterMap[userId] = []int{tweetId}
	}
	this.timestamp++
	return
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	followees := this.followerMap[userId]
	if followees == nil {
		followees = []int{userId}
	} else {
		followees = append(followees, userId)
	}
	//userId关注的followee的twitter列表的位置，从大到小开始查找
	steps := make([]int, len(followees))
	for i := 0; i < len(followees); i++ {
		user := followees[i]
		if twitters, ok := this.twitterMap[user]; ok {
			steps[i] = len(twitters) - 1
		} else {
			steps[i] = -1
		}
	}

	result := make([]int, 0)
	for i := 0; i < 10; i++ {
		maxTimeStamp := math.MinInt32
		maxTwitter := math.MinInt32
		index := -1
		for j := 0; j < len(followees); j++ {
			if steps[j] == -1 {
				continue
			}
			user := followees[j]
			twitters := this.twitterMap[user];
			if steps[j] >= len(twitters) {
				continue
			}
			twitterId := this.twitterMap[user][steps[j]]
			if this.timestampMap[twitterId] > maxTimeStamp {
				index = j
				maxTimeStamp = this.timestampMap[twitterId]
				maxTwitter = twitterId
			}
		}
		if index == -1 {
			break
		}
		result = append(result, maxTwitter)
		steps[index] --
	}

	followees = followees[0 : len(followees)-1]
	return result
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if v, ok := this.followerMap[followerId]; ok {
		for i := 0; i < len(v); i++ {
			if followeeId == v[i] {
				return
			}
		}
		this.followerMap[followerId] = append(v, followeeId)
		return
	}
	this.followerMap[followerId] = []int{followeeId}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followees := this.followerMap[followerId];
	index := -1
	n := len(followees)
	for i := 0; i < n; i++ {
		if followeeId == followees[i] {
			index = i
		}
	}
	if index == 0 && len(followees) == 1 {
		followees = followees[0:0]
	} else if index != -1 {
		followees[index] = followees[n-1]
		followees = followees[0 : n-1]
	}
	this.followerMap[followerId] = followees
}
