// Auto-generated by avdl-compiler v1.3.5 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_session.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type LoggedOutArg struct {
}

type LoggedInArg struct {
	Username string `codec:"username" json:"username"`
}

type ClientOutOfDateArg struct {
	UpgradeTo  string `codec:"upgradeTo" json:"upgradeTo"`
	UpgradeURI string `codec:"upgradeURI" json:"upgradeURI"`
	UpgradeMsg string `codec:"upgradeMsg" json:"upgradeMsg"`
}

type NotifySessionInterface interface {
	LoggedOut(context.Context) error
	LoggedIn(context.Context, string) error
	ClientOutOfDate(context.Context, ClientOutOfDateArg) error
}

func NotifySessionProtocol(i NotifySessionInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifySession",
		Methods: map[string]rpc.ServeHandlerDescription{
			"loggedOut": {
				MakeArg: func() interface{} {
					ret := make([]LoggedOutArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.LoggedOut(ctx)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"loggedIn": {
				MakeArg: func() interface{} {
					ret := make([]LoggedInArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LoggedInArg)
					if !ok {
						err = rpc.NewTypeError((*[]LoggedInArg)(nil), args)
						return
					}
					err = i.LoggedIn(ctx, (*typedArgs)[0].Username)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"clientOutOfDate": {
				MakeArg: func() interface{} {
					ret := make([]ClientOutOfDateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ClientOutOfDateArg)
					if !ok {
						err = rpc.NewTypeError((*[]ClientOutOfDateArg)(nil), args)
						return
					}
					err = i.ClientOutOfDate(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type NotifySessionClient struct {
	Cli rpc.GenericClient
}

func (c NotifySessionClient) LoggedOut(ctx context.Context) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifySession.loggedOut", []interface{}{LoggedOutArg{}})
	return
}

func (c NotifySessionClient) LoggedIn(ctx context.Context, username string) (err error) {
	__arg := LoggedInArg{Username: username}
	err = c.Cli.Call(ctx, "keybase.1.NotifySession.loggedIn", []interface{}{__arg}, nil)
	return
}

func (c NotifySessionClient) ClientOutOfDate(ctx context.Context, __arg ClientOutOfDateArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.NotifySession.clientOutOfDate", []interface{}{__arg}, nil)
	return
}
