package app

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
)

// Chat структура чата
type Chat struct {
	config  *Config
	logger  *logrus.Logger
	connMap *sync.Map
}

// New конструктор чата
func New(config *Config, logger *logrus.Logger) *Chat {
	return &Chat{
		config:  config,
		logger:  logger,
		connMap: &sync.Map{},
	}
}

func (c *Chat) Start() {
	// инициализируем слушателя подключений по TCP
	listener, err := net.Listen(c.config.Protocol, fmt.Sprintf("%s:%s", c.config.Host, c.config.Port))
	if err != nil {
		c.logger.WithError(err)
		return
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)

	c.listenForConnections(listener)
}

func (c *Chat) listenForConnections(listener net.Listener) {
	for {
		// бесконечно слушаем новые подключения
		conn, err := listener.Accept()
		if err != nil {
			c.logger.WithError(err)
			return
		}

		connId := uuid.New().String() // генерируем уникальный строковый идентирфикатор подключения
		c.connMap.Store(connId, conn) // сохраняем новое подключение в карту

		go c.handleUserConnection(conn) // обрабатываем каждое новое подключение в отдельной горутине
	}
}

// обрабатывает каждое отдельное подключение
func (c *Chat) handleUserConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	for {
		// бесконечно ожидаем новых данных
		input, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}

		// когда данные получены, проходимся по всем сохраненным подключениям чтобы в них писать
		c.connMap.Range(func(key, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				if _, err := conn.Write([]byte(input)); err != nil {
					c.logger.Error("error on writing to connection")
				}
			}

			return true
		})
	}
}
