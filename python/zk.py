#! /usr/bin/python
# -*- coding:utf-8 -*-
# @zhuchen    : 2020/4/8 20:44

from kazoo.client import KazooClient

client = KazooClient()
client.start()


class ZkHosts:

    go_host = []
    python_host = []


zk_host = ZkHosts()


@client.ChildrenWatch('/zhuchen/golang')
def golang_watch(*args):
    print('golang update')
    hosts = args[0] if args else []
    new_hosts = []
    for host_name in hosts:
        d, _ = client.get(f'/zhuchen/golang/{host_name}')
        new_hosts.append(d.decode())
    zk_host.go_host = new_hosts


@client.DataWatch('/zhuchen/python')
def python_watch(*args):
    print("python update")
    if not args:
        return
    zk_host.python_host = [args[0].decode()]


