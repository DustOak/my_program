package com.cui.action;


import com.alibaba.druid.support.json.JSONUtils;
import com.cui.po.Admin;

import com.cui.po.Student;
import com.cui.util.Encryption;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.struts2.ServletActionContext;


import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

public class CheckPasswordAction extends ActionSupport {
private String sessionId;

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public String getOldPwd() {
	return oldPwd;
}

public void setOldPwd(String oldPwd) {
	this.oldPwd = oldPwd;
}

private String oldPwd;

public String execute() {
	Map<String, Object> data = new HashMap<>();
	if (sessionId == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		data.put("code", 200);
		data.put("msg", "参数非法");
		data.put("error", 1);
	} else {
		Object oj = SessionManager.Get(sessionId).getObject();
		if (oj instanceof Admin) {
			if (Encryption.getMd5(oldPwd).equals(((Admin) oj).getPassword())) {
				data.put("code", 200);
				data.put("msg", "");
				data.put("error", 0);
			}
		} else {
			if (Encryption.getMd5(oldPwd).equals(((Student) oj).getPassword())) {
				data.put("code", 200);
				data.put("msg", "");
				data.put("error", 0);
			}
		}
	}
	String json = JSONUtils.toJSONString(data);
	HttpServletResponse response = ServletActionContext.getResponse();
	response.setContentType("application/json;charset=UTF-8");
	try {
		response.getWriter().println(json);
	} catch (IOException e) {
		e.printStackTrace();
	}
	return null;
}
}
