package com.cui.interceptor;

import com.cui.util.Session;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionContext;
import com.opensymphony.xwork2.ActionInvocation;
import com.opensymphony.xwork2.interceptor.AbstractInterceptor;
import com.opensymphony.xwork2.interceptor.Interceptor;
import com.sun.security.ntlm.Server;
import org.apache.struts2.ServletActionContext;
import org.apache.struts2.interceptor.SessionAware;
import org.hibernate.SessionFactory;

import javax.servlet.http.Cookie;
import java.util.Map;

public class SessionCheck implements Interceptor {


@Override
public void destroy() {

}

@Override
public void init() {

}

@Override
public String intercept(ActionInvocation actionInvocation) throws Exception {
	Cookie[] cookies = ServletActionContext.getRequest().getCookies();
	if (cookies != null) {
		for (int i = 0; i < cookies.length; i++) {
			if (cookies[i].getName().equals("sessionId")) {
				String sessionId = cookies[i].getValue();
				if (sessionId != null || ! sessionId.equals("")) {
					if (! SessionManager.IsExist(sessionId) && ! SessionManager.IsExpired(sessionId)
								&& ! SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
						ServletActionContext.getRequest().setAttribute("sessionId", sessionId);
						actionInvocation.invoke();
					}
				}
			}
		}
	}
	return "input";
}
}