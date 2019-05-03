package com.cui.util;

public class Session {
private Object object;
private long destroyTime;
private String ipAddr;

public Session(Object object, String ipAddr) {
	this.object = object;
	this.destroyTime = System.currentTimeMillis() + 3600 * 1000;
	this.ipAddr = ipAddr;
}

public long getDestroyTime() {
	return destroyTime;
}

public void setDestroyTime(long destroyTime) {
	this.destroyTime = destroyTime;
}

public Object getObject() {
	return object;
}

public void setObject(Object object) {
	this.object = object;
}

public String getIpAddr() {
	return ipAddr;
}

public void setIpAddr(String ipAddr) {
	this.ipAddr = ipAddr;
}
}
