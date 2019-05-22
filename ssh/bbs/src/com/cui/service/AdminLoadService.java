package com.cui.service;


import com.cui.dao.DaoOperating;
import com.cui.po.Admin;


public class AdminLoadService implements UserLoad<Admin> {
@Override
public Admin GetUser(int id) {
	return (Admin) DaoOperating.Get(new Admin(), id);
}

@Override
public Admin CheckUsernameAndPassword(String username, String password) {
	String hql = "from Admin a where  a.account='" + username + "' and  a.password='" + password + "'";
	return DaoOperating.Finds(hql).size() == 0 ? null : (Admin) DaoOperating.Finds(hql).get(0);
}
}
