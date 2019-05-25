package com.cui.action;


import com.cui.po.Admin;
import com.cui.po.Post;
import com.cui.po.Student;
import com.cui.service.PostLoadService;
import com.cui.util.SessionManager;
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
private String sessionId;
@Autowired
private PostLoadService postLoadService;
private Post postData;

public Student getStudent() {
	return student;
}

public void setStudent(Student student) {
	this.student = student;
}

public Admin getAdmin() {
	return admin;
}

public void setAdmin(Admin admin) {
	this.admin = admin;
}

private Student student;
private Admin admin;

public PostLoadService getPostLoadService() {
	return postLoadService;
}

public void setPostLoadService(PostLoadService postLoadService) {
	this.postLoadService = postLoadService;
}


public String execute() {
	postData = postLoadService.loadPost(post);
	if (postData == null) {
		return ERROR;
	}
	if (sessionId != null && ! SessionManager.IsExist(sessionId)) {
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj != null) {
			if (oj instanceof Admin) {
				admin = (Admin) oj;
			} else {
				student = (Student) oj;
			}
		}
		if (oj != null) {
			return "USER_LOGIN";
		} else {
			return "TOURIST";
		}
		
	} else {
		return "TOURIST";
	}
	
}

public Post getPostData() {
	return postData;
}

public void setPostData(Post postData) {
	this.postData = postData;
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}
}
