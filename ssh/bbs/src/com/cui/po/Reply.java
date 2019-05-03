package com.cui.po;

import java.io.Serializable;
import java.util.Date;

public class Reply implements Serializable {
private int id;
private String content;
private Date publishTime;
private Post pid;
private Admin aid;
private Student sid;

public int getId() {
	return id;
}

public void setId(int id) {
	this.id = id;
}

public String getContent() {
	return content;
}

public void setContent(String content) {
	this.content = content;
}

public Date getPublishTime() {
	return publishTime;
}

public void setPublishTime(Date publishTime) {
	this.publishTime = publishTime;
}

public Post getPid() {
	return pid;
}

public void setPid(Post pid) {
	this.pid = pid;
}

public Admin getAid() {
	return aid;
}

public void setAid(Admin aid) {
	this.aid = aid;
}

public Student getSid() {
	return sid;
}

public void setSid(Student sid) {
	this.sid = sid;
}
}
