package com.cui.action;

import com.alibaba.druid.support.json.JSONUtils;
import com.cui.po.*;
import com.cui.service.ReplyLoadService;
import com.cui.util.Session;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;

import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

public class SaveReplyAction extends ActionSupport {
private Integer board;
private Integer post;
private String content;
private String sessionId;

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

public String getContent() {
	return content;
}

public void setContent(String content) {
	this.content = content;
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public ReplyLoadService getReplyLoadService() {
	return replyLoadService;
}

public void setReplyLoadService(ReplyLoadService replyLoadService) {
	this.replyLoadService = replyLoadService;
}

private ReplyLoadService replyLoadService;

public String execute() {
	Map<String, Object> json = new HashMap<>();
	if (board == null || post == null || sessionId == null || content == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		json.put("code", 200);
		json.put("msg", "参数非法");
		json.put("error", 1);
	} else {
		Reply reply = new Reply();
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj != null) {
			if (oj instanceof Admin) {
				reply.setAid((Admin) oj);
			} else {
				reply.setSid((Student) oj);
			}
			reply.setPid(new Post(post));
			reply.setContent(content);
			reply.setPublishTime(new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").format(new Date()));
			reply.setId(null);
			if (replyLoadService.saveOrUpdate(reply)) {
				json.put("code", 200);
				json.put("msg", "");
				json.put("error", 0);
			} else {
				json.put("code", 200);
				json.put("msg", "回帖失败,请联系管理员");
				json.put("error", 1);
			}
		} else {
			json.put("code", 200);
			json.put("msg", "未登录用户无法回帖");
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
