package com.cui.util;


import java.util.*;

public class SessionManager {
private static Map<String, Session> sessionMap;
private static long destroyTime;

//销毁时间 毫秒
public static void Init(long destroyTime) {
	System.out.println("开始执行session管理功能");
	System.out.println("初始化session列表");
	sessionMap = new HashMap<String, Session>();
	SessionManager.destroyTime = destroyTime;
	System.out.println("初始化销毁时间为:" + destroyTime / 1000 + "秒");
	System.out.println("开始定时多线程执行清理过期session任务");
	Timer timer = new Timer();
	timer.schedule(new TimerTask() {
		@Override
		public void run() {
			ClearExpiredSession();
		}
	}, destroyTime, destroyTime);
}

public static synchronized void Put(String key, Session session) {
	sessionMap.put(key, session);
}

public static synchronized void Remove(String key) {
	sessionMap.remove(key);
}

//不存在返回true 否则false
public static boolean IsExist(String key) {
	return ! sessionMap.containsKey(key);
}

public static Session Get(String key) {
	return sessionMap.get(key);
}

//是否过期是 true 否 false
public static boolean IsExpired(String key) {
	Session session = Get(key);
	System.out.println("销毁时间" + session.getDestroyTime() + "当前时间" + System.currentTimeMillis());
	if (session != null && System.currentTimeMillis() < session.getDestroyTime()) {
		return false;
	}
	return true;
}

//是否更换ip 是 true 否 false
public static boolean IsInitIPAddr(String key, String ipAddr) {
	Session session = Get(key);
	System.out.println("创建iP" + session.getIpAddr() + "当前访问iP" + ipAddr);
	if (session != null && session.getIpAddr().equals(ipAddr)) {
		return false;
	}
	return true;
}

public static synchronized void ClearExpiredSession() {
	System.out.println("开始清理过期Session");
	if (sessionMap != null) {
		for (Iterator<Map.Entry<String, Session>> it = sessionMap.entrySet().iterator(); it.hasNext(); ) {
			Map.Entry<String, Session> map = it.next();
			if (map.getValue().getDestroyTime() <= System.currentTimeMillis()) {
				System.out.println("发现过期session sessionId为:" + map.getKey() + "清理此session");
				sessionMap.remove(map.getKey());
			}
		}
	}
	System.out.println("Session清理完毕");
}

public static Map<String, Session> getSessionMap() {
	return sessionMap;
}

public static void setSessionMap(Map<String, Session> sessionMap) {
	SessionManager.sessionMap = sessionMap;
}


public long getDestroyTime() {
	return destroyTime;
}

public void setDestroyTime(long destroyTime) {
	this.destroyTime = destroyTime;
}

}
