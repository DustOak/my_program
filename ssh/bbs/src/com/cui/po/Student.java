package com.cui.po;

import java.io.Serializable;

public class Student implements Serializable {
private int id;
private String stuNum;
private String realName;
private String nickName;
private String password;
private String qq;
private String email;
private String major;
private String className;
private String photoPath;

public int getId() {
	return id;
}

public void setId(int id) {
	this.id = id;
}

public String getStuNum() {
	return stuNum;
}

public void setStuNum(String stuNum) {
	this.stuNum = stuNum;
}

public String getRealName() {
	return realName;
}

public void setRealName(String realName) {
	this.realName = realName;
}

public String getNickName() {
	return nickName;
}

public void setNickName(String nickName) {
	this.nickName = nickName;
}

public String getPassword() {
	return password;
}

public void setPassword(String password) {
	this.password = password;
}

public String getQq() {
	return qq;
}

public void setQq(String qq) {
	this.qq = qq;
}

public String getEmail() {
	return email;
}

public void setEmail(String email) {
	this.email = email;
}

public String getMajor() {
	return major;
}

public void setMajor(String major) {
	this.major = major;
}

public String getClassName() {
	return className;
}

public void setClassName(String className) {
	this.className = className;
}

public String getPhotoPath() {
	return photoPath;
}

public void setPhotoPath(String photoPath) {
	this.photoPath = photoPath;
}
}
