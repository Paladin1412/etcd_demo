package http_svc

import (
	"code.byted.org/lark_data/etcd_demo/handler"
	"fmt"
	"log"
	"net/http"
)

/*
@Desc :
@Time : 2020/3/4 6:28 下午
@Author : Chang yg
@File : http_svc
*/

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ping ...")
	fmt.Fprintln(w, "pong")
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	groupId := r.FormValue("groupId")
	taskId := r.FormValue("taskId")
	err := handler.AddTask(groupId, taskId)
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("err : %v", err))
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("add : %s-%s", groupId, taskId))
}

func DelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	taskId := r.FormValue("taskId")
	groupId := r.FormValue("groupId")
	err := handler.DelTask(groupId, taskId)
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("err : %v", err))
		return
	}
	fmt.Fprintln(w, fmt.Sprintf("delete : %s-%s", groupId, taskId))
}