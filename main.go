package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func isZeroString(formstrs []string)bool{
		formstr := formstrs[0]
		fmt.Println("len: ",len(formstrs))
		if len(formstr)==0{
				return false
		}
		return true
}

func outErrorPage(w http.ResponseWriter){
		t, _ := template.ParseFiles("./views/error.gtpl")
		t.Execute(w,nil)
}

func sayYourName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("r.Form", r.Form)
	fmt.Println("r.Form[name]", r.Form["name"])
	var Name string
	for k, v := range r.Form {
		fmt.Println("key:", k)
		Name = strings.Join(v, ",")
	}
	fmt.Println(Name)
	fmt.Fprintf(w, Name)
}

func passwdCheck(username string, passwd string) {
    db, err := sql.Open("mysql","root:rootwolf@tcp(mysql)/vulnapp")
    if err != nil {
        log.Fatal(err)
        }

    res, err := db.Query("select id from user where passwd="+passwd)
    fmt.Println(res)
    defer db.Close()
}

func searchId(mail string)int{
	db, err := sql.Open("mysql","root:rootwolf@tcp(mysql)/vulnapp?parseTime=true")
	if err != nil {
		log.Fatal(err)
    }
	defer db.Close()

    sql := "select id from user where mail=?"
    res, err := db.Query(sql,mail)
    if err != nil {
        log.Fatal(err)
    }

    var id int

	for res.Next() {
		err := res.Scan(&id)
        if err != nil {
			log.Fatal(err)
        }
            log.Println("ID :",id)
    }

    fmt.Println(id)
	return id
}

func checkPasswd(id int, passwd string)string{
	db, err := sql.Open("mysql","root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	sql := "select name from user where id=? and passwd=?"
	res, err := db.Query(sql,id,passwd)
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		err := res.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
			log.Println(name)
	}

	return name
}

type Person struct {
    UserName string
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method ", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		t.Execute(w, nil)
	} else {

		r.ParseForm()
		if ( isZeroString(r.Form["mail"]) && isZeroString(r.Form["passwd"]) ){
				fmt.Println("passwd", r.Form["passwd"])
				fmt.Println("mail", r.Form["mail"])

				mail := r.Form["mail"]
				id := searchId(mail[0])

                var viewsFile string
                if id != 0 {
					passwd := r.Form["passwd"]
					name := checkPasswd(id,passwd[0])

					if name != "" {
                        fmt.Println(name)
						viewsFile = "./views/logined.gtpl"
                        t, _ := template.ParseFiles(viewsFile)
                        p := Person{UserName: name}
                        t.Execute(w,p)
					}else{
						fmt.Println(name)
						viewsFile = "./views/error.gtpl"
                        t, _ := template.ParseFiles(viewsFile)
                        t.Execute(w,nil)
					}
                }else{
                    viewsFile = "./views/error.gtpl"
                    t, _ := template.ParseFiles(viewsFile)
                    t.Execute(w,nil)
                }


		}else{
				fmt.Println("username or passwd are empty")
				outErrorPage(w)
		}
	}
}

func main() {
	http.HandleFunc("/", sayYourName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
