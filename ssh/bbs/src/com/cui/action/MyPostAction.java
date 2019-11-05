package com.cui.action;

import com.cui.po.Admin;
import com.cui.po.Post;
import com.cui.po.Student;
import com.cui.service.PostLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;

import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class MyPostAction extends ActionSupport {

private String sessionId;
private List<Post> posts;
private Student student;
private Admin admin;

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

public String execute() {
	if (sessionId == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj instanceof Admin) {
			admin = (Admin) oj;
			posts = postLoadService.getUserPosts(oj);
		} else {
			student = (Student) oj;
			posts = postLoadService.getUserPosts(oj);
		}
		if (posts == null) {
			return ERROR;
		}
		return SUCCESS;
	}
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public List<Post> getPosts() {
	return posts;
}

public void setPosts(List<Post> posts) {
	this.posts = posts;
}

public PostLoadService getPostLoadService() {
	return postLoadService;
}

public void setPostLoadService(PostLoadService postLoadService) {
	this.postLoadService = postLoadService;
}

@Autowired
private PostLoadService postLoadService;

}
