package com.cui.action;


import com.cui.dao.DaoOperating;
import com.cui.po.Admin;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionContext;
import com.opensymphony.xwork2.ActionSupport;

import java.util.List;

public class HomeAction extends ActionSupport {
private Admin admin;
private String sessionId = null;

public String execute() {
	if (sessionId != null) {
		if (! sessionId.replace(" ", "").equals("")) {
			admin = (Admin) SessionManager.Get(sessionId).getObject();
			List<Object> list = DaoOperating.Gets(admin);
			if (! list.isEmpty()) {
				admin = (Admin) list.get(0);
				return SUCCESS;
			} else {
				return INPUT;
			}
		}
	}
	return INPUT;
}

public Admin getAdmin() {
	return admin;
}

public void setAdmin(Admin admin) {
	this.admin = admin;
}

public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}
}
