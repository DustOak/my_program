package com.cui.action;

import com.cui.service.PostLoadService;
import com.opensymphony.xwork2.ActionSupport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class SavePostAction extends ActionSupport {
private Integer board;
@Autowired
private PostLoadService postLoadService;
private String title;

public String getTitle() {
	return title;
}

public void setTitle(String title) {
	this.title = title;
}

public String getContent() {
	return content;
}

public void setContent(String content) {
	this.content = content;
}

private String content;

public Integer getBoard() {
	return board;
}

public void setBoard(Integer board) {
	this.board = board;
}

public PostLoadService getPostLoadService() {
	return postLoadService;
}

public void setPostLoadService(PostLoadService postLoadService) {
	this.postLoadService = postLoadService;
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

private String sessionId;

public String execute() {
	System.out.println(board);
	System.out.println(sessionId);
	System.out.println(title);
	System.out.println(content);
	return SUCCESS;
}
}
