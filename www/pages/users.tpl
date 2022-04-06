{{template "base" .}} 



{{define "content"}}

{{$users := index .Data "users"}}

<table border="1">
    <tr>
      <th>First Name</th>
      <th>Email</th>
    </tr>
    {{range $users}}
        <tr>
            <td>{{.FirstName}}</td>
            <td>{{.Email}}</td>
        </tr>
    {{end}}

</table>


{{end}}