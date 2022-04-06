{{template "base" .}} 




{{define "content"}}

{{$user := index .Data "user"}}

<table>
    <tr>
        <td>Name</td>
        <td>{{$user.FirstName}}</td>
    </tr>
    <tr>
        <td>Email</td>
        <td>{{$user.Email}}</td>
    </tr>
</table>


{{end}}