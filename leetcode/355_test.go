package leetcode

import (
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {

	twitter := Constructor1()
	twitter.PostTweet(1, 5)
	result := twitter.GetNewsFeed(1)
	if !reflect.DeepEqual(result, []int{5}) {
		t.Errorf("expected [5], result: %v", result)
	}
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	result = twitter.GetNewsFeed(1)
	if !reflect.DeepEqual(result, []int{6, 5}) {
		t.Errorf("expected [5, 6], result: %v", result)
	}
	twitter.Unfollow(1, 2)
	result = twitter.GetNewsFeed(1)
	if !reflect.DeepEqual(result, []int{5}) {
		t.Errorf("expected [5], result: %v", result)
	}
	twitter.Follow(1, 3)
	twitter.Follow(1, 4)
	twitter.PostTweet(3,7)
	twitter.PostTweet(3,8)
	twitter.PostTweet(4,9)
	twitter.PostTweet(3,10)
	result = twitter.GetNewsFeed(1)
	if !reflect.DeepEqual(result, []int{10, 9, 8, 7, 5}) {
		t.Errorf("wrong result :%v\n", result)
	}
}
