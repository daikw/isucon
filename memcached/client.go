// see also: [GoのMemcacheパッケージ比較 - Qiita](https://qiita.com/masahikoofjoyto/items/a62a1c2b6c4affca772f)

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Memcache struct {
	conn     net.Conn
	buffered bufio.ReadWriter
}

func Mem(addr string) (conn *Memcache, err error) {
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Memcache{
		conn: nc,
		buffered: bufio.ReadWriter{
			Reader: bufio.NewReader(nc),
			Writer: bufio.NewWriter(nc),
		},
	}, err
}

func (mc *Memcache) get(key string) (result []byte, err error) {
	_, err = mc.buffered.WriteString("get " + key + "\n")
	if err == nil {
		err = mc.buffered.Flush()
		if err == nil {
			for {
				b, _, err := mc.buffered.ReadLine()
				l := string(b)
				if err == nil {
					if strings.HasPrefix(l, "END") {
						break
					}
					if strings.Contains(l, "ERROR") {
						panic("ERROR")
					}
					if !strings.HasPrefix(l, "VALUE") {
						result = append(result, l...)
						result = append(result, '\n')
					}
				} else {
					panic(err)
				}
			}
		} else {
			panic(err)
		}
	}
	return result, err
}

func (mc *Memcache) set(key string, value []byte) (err error) {
	_, err = mc.buffered.WriteString("set " + key + " 0 0 " + strconv.Itoa(len(value)) + "\r\n")
	if err == nil {
		v := append(value, "\r\n"...)
		_, err = mc.buffered.Write(v)
		if err != nil {
			panic(err)
		}
		err = mc.buffered.Flush()
		if err == nil {
			mc.buffered.ReadLine()
		}
	}

	return err
}

func main() {
	m, err := Mem("127.0.0.1:11211")
	if err == nil {
		err = m.set("foo", []byte("unko"))
		if err == nil {
			res, merr := m.get("foo")
			if merr == nil {
				fmt.Printf("%s", res)
			}
		}
	}
	defer m.conn.Close()
}
