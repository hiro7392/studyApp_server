<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="UTF-8">
  <title>student info</title>
</head>
<body>
  <h1>task info student {{.Id}}</h1>
  <table>
    <thead>
        <tr>
            <th>タスク名</th> <th>タスクの内容</th>　<th>期限</th>
        </tr>
    </thead>
    <tbody>
        {{range $nowdata:=.D}}
        <tr>
            <td>{{$nowdata.Name}}</td>
            <td>{{$nowdata.Content}}</td>
            <td>{{$nowdata.Deadline}}</td>
        </tr>
        {{end}}
    </tbody>
</table>


</body>
</html>