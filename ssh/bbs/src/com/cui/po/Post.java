package com.cui.po;

import java.io.Serializable;
import java.util.*;

public class Post implements Serializable {

private Integer id;

public Post(Integer id) {
	this.id = id;
}

private String name;
private String content;
private Student sid;
private Board bid;
private String publishTime;
private Set replies = new HashSet();

public Post() {
}

public Integer getId() {
	return id;
}

public void setId(Integer id) {
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

public String getPublishTime() {
	return publishTime;
}

public void setPublishTime(String publishTime) {
	this.publishTime = publishTime;
}

public Admin getAid() {
	return aid;
}

public void setAid(Admin aid) {
	this.aid = aid;
}

public Integer getCount() {
	return count;
}

public void setCount(Integer count) {
	this.count = count;
}

private Admin aid;
private Integer count;


public Set getReplies() {
	return replies;
}

public void setReplies(Set replies) {
	this.replies = replies;
}
}
