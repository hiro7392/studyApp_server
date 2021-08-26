<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="UTF-8">
  <title>student info</title>
</head>
<body>
  <h1>student info</h1>
  <table>
    <thead>
        <tr>
            <th>名前</th><th>タスクの内容</th><th>期限</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>{{.Name}}</td>
            <td>{{.Content}}</td>
            <td>{{.Deadline}}</td>
        </tr>
    </tbody>
</table>

  
</body>
</html>