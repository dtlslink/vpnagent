package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/tlslink/simplejson"
    "golang.org/x/crypto/ssh/terminal"
    "os"
    "strings"
    "vpnagent/rpc"
)

var (
    host     string
    username string
    password string
    group    string
)

var connect = &cobra.Command{
    Use:   "connect",
    Short: "Connect to the VPN server",
    // Args:  cobra.MinimumNArgs(1), // 至少1个非选项参数
    Run: func(cmd *cobra.Command, args []string) {
        if host == "" || username == "" {
            cmd.Help()
        } else {
            if password == "" {
                fmt.Print("Enter your password:")
                bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
                if err != nil {
                    fmt.Println("Error reading password:", err)
                    return
                }
                password = string(bytePassword)
                fmt.Println()
            }
        }
        params := make(map[string]string)
        params["log_level"] = logLevel
        params["log_path"] = logPath

        result := simplejson.New()
        err := rpcCall("config", params, result, rpc.CONFIG)
        if err != nil {
            after, _ := strings.CutPrefix(err.Error(), "jsonrpc2: code 1 message: ")
            fmt.Println(after)
        } else {
            // pretty, _ := result.EncodePretty()
            // fmt.Println(string(pretty))

            // fmt.Println(host, username, password, group)
            params := make(map[string]string)
            params["host"] = host
            params["username"] = username
            params["password"] = password
            params["group"] = group

            result := simplejson.New()
            err := rpcCall("connect", params, result, rpc.CONNECT)
            if err != nil {
                after, _ := strings.CutPrefix(err.Error(), "jsonrpc2: code 1 message: ")
                fmt.Println(after)
            } else {
                pretty, _ := result.EncodePretty()
                fmt.Println(string(pretty))
            }
        }
    },
}

func init() {
    // 子命令自己被编译、添加到主命令当中
    rootCmd.AddCommand(connect)

    // 将 Flag 解析到全局变量
    connect.Flags().StringVar(&host, "host", "", "The hostname of the VPN server")
    connect.Flags().StringVarP(&username, "username", "u", "", "User name")
    connect.Flags().StringVarP(&password, "password", "p", "", "User password")
    connect.Flags().StringVarP(&group, "group", "g", "", "User group")

    connect.Flags().StringVarP(&logLevel, "log_level", "l", "info", "Set the log level")
    connect.Flags().StringVarP(&logPath, "log_path", "d", os.TempDir(), "Set the log directory")
}
