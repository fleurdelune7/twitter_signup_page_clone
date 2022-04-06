{{template "base" .}} 



{{define "content"}}

{{$user := index .Data "user"}}

<p>New User Successful Created!!!</p>
<p>Hi {{$user.FirstName}}</p>
<p>Email: {{$user.Email}}</p>
{{end}}



