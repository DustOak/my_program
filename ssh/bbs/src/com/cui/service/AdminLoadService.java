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
	String hql = "from Admin  where account=" + username + "and password=" + password;
	return (Admin) DaoOperating.Finds(hql);
}
}
