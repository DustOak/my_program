package com.cui.po;

import java.io.Serializable;

public class Admin implements Serializable {
private Integer id;
private String account;
private String password;
private Integer qx;
private String nickname;
private String name;

public Admin(Integer id, String account, String password, Integer qx, String nickname, String name, String photoPath) {
	this.id = id;
	this.account = account;
	this.password = password;
	this.qx = qx;
	this.nickname = nickname;
	this.name = name;
	this.photoPath = photoPath;
}

private String photoPath;



public Admin() {
}

public String getAccount() {
	return account;
}

public void setAccount(String account) {
	this.account = account;
}

public String getPassword() {
	return password;
}

public void setPassword(String password) {
	this.password = password;
}


public String getNickname() {
	return nickname;
}

public void setNickname(String nickname) {
	this.nickname = nickname;
}

public String getPhotoPath() {
	return photoPath;
}

public void setPhotoPath(String photoPath) {
	this.photoPath = photoPath;
}

public Integer getId() {
	return id;
}

public void setId(Integer id) {
	this.id = id;
}

public Integer getPermission() {
	return qx;
}

public void setPermission(Integer permission) {
	this.qx = permission;
}

public String getName() {
	return name;
}

public void setName(String name) {
	this.name = name;
}
}
