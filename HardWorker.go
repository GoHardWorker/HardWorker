/*
	HardWorkerパッケージはHardWorkerフレームワーク本体です。

	メインの処理からHardWorker.Initializeを呼び出すことで、
	あとの処理をフレームワークに任せることができます。
 */
package HardWorker

import (
	"strconv"
	"net/http"
	"./Util"
)

// ApplicationはVizualizer本体のクラスになります。
// HTTPサーバーとして動作するものであるため、HTTPサーバーとして必要な処理を
// 記述することになります。
type Application struct {
	serverName string
	serverIp string
	serverPort int
}



// ApplicationはVizualizer本体のクラスを生成するためのファクトリ関数です。
func Application() *Application {
	return &Application{}
}

// ServeHTTPは受け取ったリクエストを処理するための汎用メソッドです。
func (this *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}

// Initializeはフレームワーク自体を初期化するための処理を行います。
func (this *Application) Initialize() {
	// アプリケーションインスタンスを生成
	application := Application()

	http.HandleFunc("/", application)
	http.ListenAndServe(Util.Concat(application.serverIp, ":", strconv.FormatInt(application.serverPort, 10)), nil)
}
