package cmd

//func (con *connection) Auth(password string) (bool, error) {
//	if con.p != nil {
//		c, _ := con.GetConnection()
//		defer c.Release()
//		return c.Auth(password)
//	} else {
//		res, err := rg.String(con.c.Do("AUTH", password))
//		return getBool(res), err
//	}
//}

//func (con *connection) Echo(message string) (string, error) {
//	if con.p != nil {
//		c, _ := con.GetConnection()
//		defer c.Release()
//		return c.Echo(message)
//	} else {
//		return rg.String(con.c.Do("ECHO", message))
//	}
//}
//
//func (con *connection) Ping() (string, error) {
//	if con.p != nil {
//		c, _ := con.GetConnection()
//		defer c.Release()
//		return c.Ping()
//	} else {
//		return rg.String(con.c.Do("PING"))
//	}
//}
//
//func (con *connection) Select(index int) (bool, error) {
//	if con.p != nil {
//		c, _ := con.GetConnection()
//		defer c.Release()
//		return c.Select(index)
//	} else {
//		res, err := con.c.Do("SELECT", index)
//		return rg.Bool(res, err)
//	}
//}
//
//func (con *connection) Quit() (string, error) {
//	return rg.String(con.c.Do("QUIT"))
//}