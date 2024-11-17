package main

import (
	"fmt"
	"net/http"
	"time"
)

func setDetailedCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Cookieを作成
	cookie1 := &http.Cookie{
		Name:     "detailed_cookie",              // Cookieの名前
		Value:    "detailed_value",               // Cookieの値
		Path:     "/",                            // Cookieの有効なパス（"/"は全体）
		Domain:   "localhost",                    // Cookieが有効なドメイン、通常はホスト名を指定する
		Expires:  time.Now().Add(24 * time.Hour), // Cookieの有効期限（24時間後）、ブラウザに保存されるCookieの場合に使用
		MaxAge:   3600,                           // Cookieの有効期間（秒単位、優先される）、０の場合、即時削除
		Secure:   true,                           // HTTPS通信でのみ送信、http通信の場合のみCookieが送信される
		HttpOnly: true,                           // JavaScriptからのアクセスを禁止、クライアントサイドのJavaScriptからは触れなくなる
		SameSite: http.SameSiteStrictMode,        // クロスサイトリクエストへの制限
	}

	// クライアントにCookieを送信
	http.SetCookie(w, cookie1)

	// 複数種類設定することも可能
	cookie2 := &http.Cookie{
		Name:     "cookie2",
		Value:    "cookie2",
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, cookie2)
	fmt.Fprintln(w, "Detailed Cookie has been set!")
}

func getCookieDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Cookieを取得
	cookie, err := r.Cookie("detailed_cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No cookie found", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Cookieの詳細情報を表示
	fmt.Fprintf(w, "Cookie Name: %s\n", cookie.Name)
	fmt.Fprintf(w, "Cookie Value: %s\n", cookie.Value)
	fmt.Fprintf(w, "Cookie Path: %s\n", cookie.Path)
	fmt.Fprintf(w, "Cookie Domain: %s\n", cookie.Domain)
	fmt.Fprintf(w, "Cookie Expires: %s\n", cookie.Expires)
	fmt.Fprintf(w, "Cookie MaxAge: %d\n", cookie.MaxAge)
	fmt.Fprintf(w, "Cookie Secure: %t\n", cookie.Secure)
	fmt.Fprintf(w, "Cookie HttpOnly: %t\n", cookie.HttpOnly)
	fmt.Fprintf(w, "Cookie SameSite: %v\n", cookie.SameSite)
}

func main() {
	// ハンドラーを設定
	http.HandleFunc("/set-detailed-cookie", setDetailedCookieHandler)
	http.HandleFunc("/get-cookie-details", getCookieDetailsHandler)

	// サーバーを起動
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
