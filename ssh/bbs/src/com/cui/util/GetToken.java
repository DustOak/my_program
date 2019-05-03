package com.cui.util;

import org.apache.struts2.ServletActionContext;

public class GetToken {
	public static String GetToken(String name){
		return (String) ServletActionContext.getRequest().getAttribute(name);
	}
	public  static  void RemoveToken(String name ){
		ServletActionContext.getRequest().removeAttribute(name);
	}
}
