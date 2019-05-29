package com.cui.service;

import com.cui.dao.DaoOperating;
import com.cui.po.Reply;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class ReplyLoadService implements ReplyLoad {
public DaoOperating getDaoOperating() {
	return daoOperating;
}

public void setDaoOperating(DaoOperating daoOperating) {
	this.daoOperating = daoOperating;
}

@Autowired
private DaoOperating daoOperating;

@Override
public boolean saveOrUpdate(Reply reply) {
	try {
		if (reply.getId() == null) {
			daoOperating.Save(reply);
			return true;
		} else {
			daoOperating.Update(reply);
			return true;
		}
	} catch (Exception ex) {
		ex.printStackTrace();
		return false;
	}
}

@Override
public void deleteReply(Reply reply) {
	try {
		daoOperating.Delete(reply);
	} catch (Exception ex) {
		ex.printStackTrace();
	}
}

@Override
public List<Reply> getReplies(Integer id) {
	String hql = "from Reply where pid=" + id;
	return daoOperating.Finds(hql);
}
}
