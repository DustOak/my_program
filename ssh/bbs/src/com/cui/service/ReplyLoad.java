package com.cui.service;

import com.cui.po.Reply;

import java.util.List;

public interface ReplyLoad {
boolean saveOrUpdate(Reply reply);

void deleteReply(Reply reply);

List<Reply> getReplies(Integer id);
}
