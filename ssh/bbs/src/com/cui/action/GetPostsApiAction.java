package com.cui.action;

import com.cui.po.Post;
import com.cui.service.PostLoad;
import com.cui.service.PostLoadService;
import com.opensymphony.xwork2.ActionSupport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class GetPostsApiAction extends ActionSupport {
private int board;
private int start;
private int length;
private String data;
private List<Post> posts;
@Autowired
private PostLoadService postLoadService;

public int getStart() {
	return start;
}

public void setStart(int start) {
	this.start = start;
}

public int getLength() {
	return length;
}

public void setLength(int length) {
	this.length = length;
}

public String execute() {
	posts = postLoadService.pageAllPost(board, start, length);
	for (int j = 0; j < posts.size(); j++) {
		System.out.print(posts.get(j).getBid().getName() + " ");
		System.out.println(posts.get(j).getName() + " ");
	}
	return SUCCESS;
}

public String getData() {
	return data;
}

public void setData(String data) {
	this.data = data;
}

public int getBoard() {
	return board;
}

public void setBoard(int board) {
	this.board = board;
}


public PostLoadService getPostLoadService() {
	return postLoadService;
}

public void setPostLoadService(PostLoadService postLoadService) {
	this.postLoadService = postLoadService;
}

public List<Post> getPosts() {
	return posts;
}

public void setPosts(List<Post> posts) {
	this.posts = posts;
}
}
