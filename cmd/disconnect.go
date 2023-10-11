package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/tlslink/simplejson"
    "strings"
    "vpnagent/rpc"
)

var disconnect = &cobra.Command{
    Use:   "disconnect",
    Short: "Disconnect from the VPN server",
    Run: func(cmd *cobra.Command, args []string) {
        result := simplejson.New()
        err := rpcCall("disconnect", nil, result, rpc.DISCONNECT)
        if err != nil {
            after, _ := strings.CutPrefix(err.Error(), "jsonrpc2: code 1 message: ")
            fmt.Println(after)
        } else {
            pretty, _ := result.EncodePretty()
            fmt.Println(string(pretty))
        }
    },
}

func init() {
    rootCmd.AddCommand(disconnect)
}
