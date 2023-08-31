
package server

import "net/http"

type Route struct {
	path    string
	methods map[string]func(http.ResponseWriter, *http.Request)
}


type Server struct {
	route []Route
	
}


func NewServer()*Server{
	return &Server{
		route: make([]Route,0),
	}
	

}


func (s Server) addMethod(path string,method string,handleFunc func(http.ResponseWriter, *http.Request)){
	route := Route{
		path: path,
		methods: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
	route.methods[method] = handleFunc
	http.HandleFunc(path,func(w http.ResponseWriter,r *http.Request){
		if(r.Method == method){
			handleFunc(w,r)
		}else{
			http.Error(w,"http method not allowed Please Write GET",http.StatusMethodNotAllowed)
		}
	})

	s.route = append(s.route, route)

}

func (s *Server) Get(path string, handleFunc func(http.ResponseWriter, *http.Request)) {
	s.addMethod(path, http.MethodGet, handleFunc)
}

func (s *Server) Post(path string, handleFunc func(http.ResponseWriter, *http.Request)) {
	s.addMethod(path, http.MethodPost, handleFunc)
}


func (s *Server) Put(path string, handleFunc func(http.ResponseWriter, *http.Request)) {
	s.addMethod(path, http.MethodPut, handleFunc)
}


func (s *Server) Delete(path string, handleFunc func(http.ResponseWriter, *http.Request)) {
	s.addMethod(path, http.MethodDelete, handleFunc)
}



func (s Server)RunServer(port string,handler http.Handler){
	http.ListenAndServe(port,handler)
}







