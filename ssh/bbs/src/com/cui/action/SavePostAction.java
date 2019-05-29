package com.cui.action;

import com.alibaba.druid.support.json.JSONUtils;
import com.cui.po.Admin;
import com.cui.po.Board;
import com.cui.po.Post;
import com.cui.po.Student;
import com.cui.service.PostLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

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
	Map<String, Object> json = new HashMap<>();
	if (board == null || title == null || sessionId == null || content == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		json.put("code", 200);
		json.put("msg", "参数非法");
		json.put("error", 1);
	} else {
		Post post = new Post();
		post.setBid(new Board(board));
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj != null) {
			if (oj instanceof Admin) {
				post.setAid((Admin) oj);
			} else {
				post.setSid((Student) oj);
			}
			post.setName(title);
			post.setContent(content);
			post.setPublishTime(new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").format(new Date()));
			post.setId(null);
			post.setCount(0);
			if (postLoadService.saveOrUpdate(post)) {
				json.put("code", 200);
				json.put("msg", "");
				json.put("error", 0);
			} else {
				json.put("code", 200);
				json.put("msg", "发布失败,请联系管理员");
				json.put("error", 1);
			}
		} else {
			json.put("code", 200);
			json.put("msg", "禁止未登录用户发帖");
			json.put("error", 1);
		}
		
	}
	HttpServletResponse response = ServletActionContext.getResponse();
	response.setContentType("application/json;charset=UTF-8");
	String data = JSONUtils.toJSONString(json);
	try {
		response.getWriter().println(data);
	} catch (IOException e) {
		e.printStackTrace();
	}
	return null;
}
}
