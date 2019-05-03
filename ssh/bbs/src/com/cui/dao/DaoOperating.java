package com.cui.dao;

import org.hibernate.SessionFactory;
import org.springframework.orm.hibernate4.HibernateTemplate;

import java.util.List;

public class DaoOperating implements Dao {

public DaoOperating() {
}

public DaoOperating(SessionFactory sf ) {
	sessionFactory=sf;
	hibernateTemplate=new HibernateTemplate(sessionFactory);
}

private static HibernateTemplate hibernateTemplate = null;
private static  SessionFactory sessionFactory = null;

public static Object Get(Object oj, int id) {
	return getHibernateTemplate().get(oj.getClass(),id);
}

public static List<Object> Gets(Object oj) {
	return getHibernateTemplate().findByExample(oj);
}

public static List<Object> Finds(String fields, Object...value) {
	return (List<Object>) getHibernateTemplate().find(fields, value);
}

public static boolean Save(Object oj) {
	return (boolean) getHibernateTemplate().save(oj);
}

public static void Update(Object oj) {
	getHibernateTemplate().update(oj);
}

public static void Delete(Object oj) {
	getHibernateTemplate().delete(oj);
}

public static HibernateTemplate getHibernateTemplate() {
	return hibernateTemplate;
}

public void setHibernateTemplate(HibernateTemplate hibernateTemplate) {
	this.hibernateTemplate = hibernateTemplate;
}

public SessionFactory getSessionFactory() {
	return sessionFactory;
}

public void setSessionFactory(SessionFactory sessionFactory) {
	this.sessionFactory = sessionFactory;
}
}
