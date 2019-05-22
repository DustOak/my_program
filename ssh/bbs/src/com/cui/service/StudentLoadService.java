package com.cui.service;

import com.cui.dao.DaoOperating;


import com.cui.po.Student;

import java.util.List;

public class StudentLoadService implements UserLoad<Student> {
@Override
public Student GetUser(int id) {
	return (Student) DaoOperating.Get(new Student(), id);
}

@Override
public Student CheckUsernameAndPassword(String username, String password) {
	String hql = "from Student  s where s.stuNum='" + username + "' and  s.password='" + password + "'";
	return DaoOperating.Finds(hql).size() == 0 ? null : (Student) DaoOperating.Finds(hql).get(0);
}
}
