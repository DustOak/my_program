package com.cui.service;

import com.cui.dao.DaoOperating;

import com.cui.po.Admin;
import com.cui.po.Student;

public class StudentLoadService implements UserLoad<Student> {
@Override
public Student GetUser(int id) {
	return (Student) DaoOperating.Get(new Student(), id);
}

@Override
public Student CheckUsernameAndPassword(String username, String password) {
	String hql = "from Admin  where stuNum=" + username + "and password=" + password;
	return (Student) DaoOperating.Finds(hql);
}
}
