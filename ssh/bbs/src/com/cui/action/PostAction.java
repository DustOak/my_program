package com.cui.action;


import com.cui.po.Post;
import com.cui.service.PostLoadService;
import com.opensymphony.xwork2.ActionSupport;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


@Component
public class PostAction extends ActionSupport {
public Integer getBoard() {
	return board;
}

public void setBoard(Integer board) {
	this.board = board;
}

public Integer getPost() {
	return post;
}

public void setPost(Integer post) {
	this.post = post;
}

private Integer board;
private Integer post;
@Autowired
private PostLoadService postLoadService;
private Post postData;

public PostLoadService getPostLoadService() {
	return postLoadService;
}

public void setPostLoadService(PostLoadService postLoadService) {
	this.postLoadService = postLoadService;
}


public String execute() {
	postData = postLoadService.loadPost(post);
	if (postData != null) {
		return SUCCESS;
	}
	return ERROR;
}

public Post getPostData() {
	return postData;
}

public void setPostData(Post postData) {
	this.postData = postData;
}
}
