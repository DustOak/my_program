package com.cui.service;


import com.cui.dao.DaoOperating;
import com.cui.po.Admin;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class AdminLoadService implements UserLoad<Admin> {
public DaoOperating getDaoOperating() {
	return daoOperating;
}

public void setDaoOperating(DaoOperating daoOperating) {
	this.daoOperating = daoOperating;
}

@Autowired
DaoOperating daoOperating;

@Override
public Admin GetUser(int id) {
	return (Admin) daoOperating.Get(new Admin(), id);
}

@Override
public Admin CheckUsernameAndPassword(String username, String password) {
	String hql = "from Admin a where  a.account='" + username + "' and  a.password='" + password + "'";
	return daoOperating.Finds(hql).size() == 0 ? null : (Admin) daoOperating.Finds(hql).get(0);
}
}
