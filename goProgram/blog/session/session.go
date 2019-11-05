package session

import (
	"blog/logger"
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

//session 对象
type session struct {
	sessionId  string
	createTime int64
}

//session 管理
type manager interface {
	//创建session 返回sessionId和value
	CreateSession() string
	//验证是否过期
	CheckLifeTime(sessionId string) bool
	//检查是否存在
	CheckIsExist(sessionId string) bool
	//销毁session
	DestroySession(sessionId string)
	//定时清理
	TimingCleanSession(afterTime time.Duration)
	//遍历清除过期session
	destroyExpireSession()
	//返回cookie值
	CookieIsExist(r *http.Request, name string) string
}

type SessionManager struct {
	lifeTime int64
	lock     sync.Mutex
	sessions map[string]session
}

//判断cookie是否存在
func (s *SessionManager) CookieIsExist(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	} else {
		return cookie.Value
	}

}

//创建一个session管理器 lifeTime单位为秒
func NewSessionManager(lifeTime int64) *SessionManager {
	return &SessionManager{
		lifeTime: lifeTime,
		sessions: make(map[string]session),
	}
}

//新建session 并放入内存中管理 返回sessionid 和value
func (s *SessionManager) CreateSession() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	session := session{
		sessionId:  fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s", id)))),
		createTime: time.Now().Unix(),
	}
	s.sessions[session.sessionId] = session
	return session.sessionId
}

//检查是否超时
func (s *SessionManager) CheckLifeTime(sessionId string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if (time.Now().Unix() - s.sessions[sessionId].createTime) >= s.lifeTime {
		return true
	}
	return false
}

//销毁session
func (s *SessionManager) DestroySession(sessionId string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.sessions, sessionId)
}

//遍历清除过期session
func (s *SessionManager) destroyExpireSession() {
	s.lock.Lock()
	defer s.lock.Lock()
	for _, v := range s.sessions {
		if (time.Now().Unix() - v.createTime) >= s.lifeTime {
			logger.InfoLog.Println("删除过期的session id：", v.sessionId)
			delete(s.sessions, v.sessionId)
		}
	}
	logger.InfoLog.Println("清理完毕")
	return

}

//定时清理session afterTime为定时时间
func (s *SessionManager) TimingCleanSession(afterTime time.Duration) {
	s.lock.Lock()
	defer s.lock.Unlock()
	time.AfterFunc(afterTime, func() {
		logger.InfoLog.Println("定时清理session")
		s.destroyExpireSession()
	})
}

//检查是否存在
func (s *SessionManager) CheckIsExist(sessionId string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.sessions[sessionId]
	return ok
}
