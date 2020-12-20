package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	code := `
		var posts = {posts};
		var res = [];
		var i = 0;
		while (i < posts.length) {
			var post = {};
			post.post_id = posts[i].id;
			post.text = posts[i].text;
			var likes = API.likes.getList({
							 "type": "post",
							 "owner_id": {user_id},
							 "item_id": posts[i].id,
							 "filter": "likes",
							 "extended": 1
			});
			if (likes) {
				post.likes = likes;
			} else {
				post.likes = { "count": 0, "items": [] };
			}
			if (posts[i].attachments) {
				post.attachments = posts[i].attachments;
			}            
			var comments = API.wall.getComments({
							 "owner_id": {user_id},
							 "post_id": posts[i].id,
							 "fields": "first_name,last_name",
							 "extended": 1
			});
			if (comments) {
				post.comments = comments;
			} else {
				post.comments = { "count": 0, "items": [] };
			}
			res.push(post);
			i = i + 1;
		}
		return res;`
	r := strings.NewReplacer("{posts}", "HUY", "{user_id}", "USID", )
	fmt.Println(r.Replace(code))

	type T struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	jStr := []byte(`{"a": 12}`)
	var m T
	_ = json.Unmarshal(jStr, &m)
	fmt.Printf("%#v\n", m)
}