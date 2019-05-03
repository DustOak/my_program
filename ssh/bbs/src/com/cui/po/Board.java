package com.cui.po;

import java.io.Serializable;

public class Board implements Serializable {
private int id;
private String name;
private String description;
private int parentId;
private Admin aid;
private String boardImg;

public String getBoardImg() {
	return boardImg;
}

public void setBoardImg(String boardImg) {
	this.boardImg = boardImg;
}

public Admin getAid() {
	return aid;
}

public void setAid(Admin aid) {
	this.aid = aid;
}

public int getParentId() {
	return parentId;
}

public void setParentId(int parentId) {
	this.parentId = parentId;
}

public String getDescription() {
	return description;
}

public void setDescription(String description) {
	this.description = description;
}

public String getName() {
	return name;
}

public void setName(String name) {
	this.name = name;
}

public int getId() {
	return id;
}

public void setId(int id) {
	this.id = id;
}
}
