package com.cui.po;

import java.io.Serializable;
import java.util.Date;

public class Post implements Serializable {

private int id;
private String name;
private String content;
private Student sid;
private Board bid;
private Date publishTime;

public int getId() {
	return id;
}

public void setId(int id) {
	this.id = id;
}

public String getName() {
	return name;
}

public void setName(String name) {
	this.name = name;
}

public String getContent() {
	return content;
}

public void setContent(String content) {
	this.content = content;
}

public Student getSid() {
	return sid;
}

public void setSid(Student sid) {
	this.sid = sid;
}

public Board getBid() {
	return bid;
}

public void setBid(Board bid) {
	this.bid = bid;
}

public Admin getAid() {
	return aid;
}

public void setAid(Admin aid) {
	this.aid = aid;
}

public int getCount() {
	return count;
}

public void setCount(int count) {
	this.count = count;
}

private Admin aid;
private int count;

public Date getPublishTime() {
	return publishTime;
}

public void setPublishTime(Date publishTime) {
	this.publishTime = publishTime;
}
}
