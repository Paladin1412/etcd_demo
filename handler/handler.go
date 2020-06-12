package handler

import (
	"context"
	"code.byted.org/lark_data/etcd_demo/etcd_client"
	"code.byted.org/lark_data/etcd_demo/svc"
	"fmt"
)

/*
@Desc :
@Time : 2020/6/11 4:58 下午
@Author : Chang yg
@File : handler
*/


func AddTask(groupId, taskId string) error {
	_, err := etcd_client.EtcdClient.Put(context.Background(), svc.All_TASK_ROOT + groupId + "/" + taskId, fmt.Sprintf("%s-%s", groupId, taskId))
	return err
}

func DelTask(groupId, taskId string) error {
	_, err := etcd_client.EtcdClient.Delete(context.Background(), svc.All_TASK_ROOT + groupId + "/" + taskId)
	return err
}

