package com.cui.dao;

import org.hibernate.HibernateException;
import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.query.Query;
import org.springframework.orm.hibernate4.HibernateCallback;
import org.springframework.orm.hibernate4.HibernateTemplate;

import java.util.List;

public class DaoOperating implements Dao {

public DaoOperating() {
}

public DaoOperating(SessionFactory sf) {
	sessionFactory = sf;
	hibernateTemplate = new HibernateTemplate(sessionFactory);
}

private static HibernateTemplate hibernateTemplate = null;
private static SessionFactory sessionFactory = null;

public static Object Get(Object oj, int id) {
	return getHibernateTemplate().get(oj.getClass(), id);
}

public static List Gets(Object oj) {
	return getHibernateTemplate().findByExample(oj);
}

//public static Object Query(String sql, Object... value) {
//	return getHibernateTemplate().execute(new HibernateCallback<List>() {
//		@Override
//		public List doInHibernate(Session session) throws HibernateException {
//			Query query = session.createQuery(sql);
//			for (int i = 0; i < value.length; i++) {
//				query.setParameter(i, value[i]);
//			}
//			return query.list();
//		}
//	});
//}

public static List Finds(String fields, Object... value) {
	return getHibernateTemplate().execute(new HibernateCallback<List>() {
		@Override
		public List doInHibernate(Session session) throws HibernateException {
			Query query = session.createQuery(fields);
			for (int i = 0; i < value.length; i++) {
				query.setParameter(i, value[i]);
			}
			return query.list();
		}
	});
}

public static boolean Save(Object oj) {
	return (boolean) getHibernateTemplate().save(oj);
}

public static void Update(Object oj) {
	getHibernateTemplate().update(oj);
}

public static boolean Delete(Object oj) {
	try {
		getHibernateTemplate().delete(oj);
		return true;
	} catch (Exception ex) {
		ex.printStackTrace();
		return false;
	}
}

public static List LimitQuery(String hql, int start, int length) {
	return getHibernateTemplate().execute(new HibernateCallback<List>() {
		@Override
		public List doInHibernate(Session session) throws HibernateException {
			Query query = session.createQuery(hql);
			query.setFirstResult(start);
			query.setMaxResults(length);
			return query.list();
		}
	});
}

public static int Count(Object oj) {
	return getHibernateTemplate().findByExample(oj).size();
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
