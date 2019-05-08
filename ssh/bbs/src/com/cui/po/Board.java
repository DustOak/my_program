package com.cui.po;

import java.io.Serializable;

public class Board implements Serializable {
private Integer id;

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

public String getDescription() {
	return description;
}

public void setDescription(String description) {
	this.description = description;
}

public Admin getAid() {
	return aid;
}

public void setAid(Admin aid) {
	this.aid = aid;
}

public String getBoardImg() {
	return boardImg;
}

public void setBoardImg(String boardImg) {
	this.boardImg = boardImg;
}

public Board getParentId() {
	return parentId;
}

public void setParentId(Board parentId) {
	this.parentId = parentId;
}

private String name;
private String description;
private Admin aid;
private String boardImg;
private Board parentId;


}
