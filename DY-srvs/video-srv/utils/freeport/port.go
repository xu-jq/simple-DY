/*
 * @Date: 2023-01-29 09:22:19
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-29 09:22:45
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/freeport/port.go
 * @Description: 查找指定数量的空闲端口
 */
package freeport

import "net"

func GetFreePorts(count int) ([]int, error) {
	var ports []int
	for i := 0; i < count; i++ {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			return nil, err
		}
		defer l.Close()
		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
	}
	return ports, nil
}
