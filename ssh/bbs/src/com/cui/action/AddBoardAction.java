package com.cui.action;

import com.cui.po.Admin;
import com.cui.po.Board;
import com.cui.service.BoardLoadService;
import com.cui.util.SessionManager;
import com.opensymphony.xwork2.ActionSupport;
import org.apache.commons.io.FileUtils;
import org.apache.struts2.ServletActionContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.io.File;

@Component
public class AddBoardAction extends ActionSupport {
private String sessionId;
private Integer parentId;
private String boardName;
private String boardDescription;
private File icon;
@Autowired
private BoardLoadService boardLoadService;

public String execute() {
	if (parentId == null || boardDescription == null || sessionId == null || boardName == null || SessionManager.IsExist(sessionId)
				|| SessionManager.IsInitIPAddr(sessionId, ServletActionContext.getRequest().getRemoteAddr())) {
		return ERROR;
	} else {
		try {
			//文件名位置使用md5进行随机编码 以防止文件名重复
			if (icon != null) {
				File icons = new File(destPath, iconFileName);
				FileUtils.copyFile(icon, icons);
			}
			Admin admin = (Admin) SessionManager.Get(sessionId).getObject();
			Board board = new Board();
			if (parentId != - 1) {
				board.setParentId(new Board(parentId));
			}
			board.setAid(admin);
			board.setDescription(boardDescription);
			board.setName(boardName);
			board.setBoardImg(iconFileName);
			boardLoadService.SaveOrUpdate(board);
		} catch (Exception ex) {
			ex.printStackTrace();
			return ERROR;
		}
	}
	return SUCCESS;
}


private String iconContentType;
private String iconFileName;

public File getIcon() {
	return icon;
}

public void setIcon(File icon) {
	this.icon = icon;
}

public String getIconContentType() {
	return iconContentType;
}

public void setIconContentType(String iconContentType) {
	this.iconContentType = iconContentType;
}

public String getIconFileName() {
	return iconFileName;
}

public void setIconFileName(String iconFileName) {
	this.iconFileName = iconFileName;
}

public String getDestPath() {
	return destPath;
}

public void setDestPath(String destPath) {
	this.destPath = destPath;
}

private String destPath;


public String getSessionId() {
	return sessionId;
}

public void setSessionId(String sessionId) {
	this.sessionId = sessionId;
}

public Integer getParentId() {
	return parentId;
}

public void setParentId(Integer parentId) {
	this.parentId = parentId;
}

public String getBoardName() {
	return boardName;
}

public void setBoardName(String boardName) {
	this.boardName = boardName;
}

public String getBoardDescription() {
	return boardDescription;
}

public void setBoardDescription(String boardDescription) {
	this.boardDescription = boardDescription;
}


public BoardLoadService getBoardLoadService() {
	return boardLoadService;
}

public void setBoardLoadService(BoardLoadService boardLoadService) {
	this.boardLoadService = boardLoadService;
}
}
