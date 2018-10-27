package tailf

import (
	"sync"
	"github.com/gostudy03/logagent/common"
	"github.com/gostudy03/xlog"
)

type TailTaskMgr struct {

	tailTaskMap map[string]*TailTask
	collectLogList []*common.CollectConfig
	etcdCh <- chan []*common.CollectConfig
}

var (
	tailTaskMgr *TailTaskMgr
)

func Init(collectLogList []*common.CollectConfig, etcdCh <- chan []*common.CollectConfig) (err error) {
	tailTaskMgr = &TailTaskMgr {
		collectLogList: collectLogList,
		etcdCh: etcdCh,
		tailTaskMap:make(map[string]*TailTask, 16),
	}

	for _, conf := range collectLogList {

		if tailTaskMgr.exists(conf) {
			xlog.LogWarn("init tail Task failed, conf:%#v is duplicate", conf)
			continue
		}

		tailTask, err := NewTailTask(conf.Path, conf.ModuleName, conf.Topic) 
		if err != nil {
			xlog.LogError("init tail Task failed, conf:%#v, err:%v", conf, err)
			continue
		}

		go tailTask.Run()
		tailTaskMgr.tailTaskMap[tailTask.Key()] = tailTask
		//tailTaskMgr.tailTaskList = append(tailTaskMgr.tailTaskList, tailTask)
	}
	return
}

func (t *TailTaskMgr) listTask() {
	for key, task := range t.tailTaskMap {
		xlog.LogDebug("=============key:%s task:%s======== is running",
			 key, task.Topic)
	}
}

func (t *TailTaskMgr) run() {

	for {
		t.listTask()
		tmpCollectLogList := <- t.etcdCh
		xlog.LogDebug("你好，有新的配置变更, data:%#v", tmpCollectLogList)

		//判断是否有新增的日志收集配置
		for _, conf := range tmpCollectLogList {
			//如果对应的日志收集配置，已经存在。那么不需要做任何事情
			if t.exists(conf) {
				xlog.LogDebug("the task of conf:%#v is already run", conf)
				continue
			}

			xlog.LogDebug("new task of conf:%#v is running", conf)
			//如果不存在，说明是一个新的日志收集配置。那么起一个新的日志收集任务
			tailTask, err := NewTailTask(conf.Path, conf.ModuleName, conf.Topic) 
			if err != nil {
				xlog.LogError("init tail Task failed, conf:%#v, err:%v", conf, err)
				continue
			}
	
			go tailTask.Run()
			tailTaskMgr.tailTaskMap[tailTask.Key()] = tailTask
			//tailTaskMgr.tailTaskList = append(tailTaskMgr.tailTaskList, tailTask)
		}

		//从已经运行的任务里面判断是否存在最新的配置当中，如果不存在的话，说明这个任务的配置已经被删除
		//我们就需要把这个任务给停了
		for key, task := range t.tailTaskMap {
			found := false
			for _, conf := range tmpCollectLogList {
				if task.Path == conf.Path && 
				task.ModuleName == conf.ModuleName && 
				task.Topic == conf.Topic {
					found =  true
					break
				}
			}

			if found == false {
				//需要把这个任务取消，并且从tailTaskList删除
				task.Stop()
				delete(t.tailTaskMap,key)
/*
				prevTask := t.tailTaskList[0:index]
				nextTask := t.tailTaskList[index+1:]

				var taskList []*TailTask
				taskList = append(taskList, prevTask...)
				taskList = append(taskList, nextTask...)

				t.tailTaskList = taskList
				*/
			}
		}
	}
}

func (t *TailTaskMgr) exists(conf *common.CollectConfig) (bool) {
	
	for _, tailTask := range t.tailTaskMap {
		if tailTask.Path == conf.Path && 
		tailTask.ModuleName == conf.ModuleName && 
		tailTask.Topic == conf.Topic {
			return true
		}
	}

	return false
}



func Run(wg *sync.WaitGroup) {
	//tailTaskMgr主要做一些日志任务收集的管理工作，比如新增日志收集任务、删除日志收集任务
	tailTaskMgr.run()
	wg.Done()
}