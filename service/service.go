package service

import (
	"apiserver/model"
	"apiserver/util"
	"fmt"
	"sync"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {

	userInfos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)

	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}

	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}

	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}
	channelError := make(chan error, 1)
	isChannelFinished := make(chan bool, 1)

	// Improve efficiency using parallel query
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortId, err := util.GenShortId()

			if err != nil {
				channelError <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()

			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2000-01-02 12:30:05"),
				UpdatedAt: u.UpdatedAt.Format("2000-01-02 12:30:05"),
			}

		}(u)
	}

	go func() {
		wg.Wait()
		close(isChannelFinished)
	}()

	select {
	case <-isChannelFinished:
	case err := <-channelError:
		return nil, count, err
	}

	for _, id := range ids {
		userInfos = append(userInfos, userList.IdMap[id])
	}

	return userInfos, count, nil
}
