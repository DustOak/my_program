package com.cui.util;

import org.springframework.util.DigestUtils;

import java.sql.Time;
import java.util.Random;

public class Encryption {
private final static String salt = "cuiwenbin1397.+";

public static String getMd5(String data) {
	String base = data + salt;
	return DigestUtils.md5DigestAsHex(base.getBytes());
}

public static String getRandomToken() {
	Random r = new Random(System.currentTimeMillis());
	String md5 = String.valueOf(r.nextInt());
	return DigestUtils.md5DigestAsHex(md5.getBytes());
}
}
