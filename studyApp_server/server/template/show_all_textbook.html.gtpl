<head>
  <meta charset="UTF-8">
  <title>講師用画面(生徒検索結果)</title>
  <link href="./client/css/style.css" rel="stylesheet">
</head>

<body>
  <header>
  <a class="a_button" href="teacher.html">ホーム</a>
  </header>

  <!-- 前のページで入力した条件に合う生徒を表示。 -->
  <!-- 生徒名をクリックして。生徒のページ(teacher_student.html)へ移動。その生徒の情報を次のページで表示したい -->
  <h2>教材一覧</h2>
   <table>
      <colgroup span="5"></colgroup>
      <tr>
        <th>テキストId</th>
        <th>科目</th>
        <th>テキスト名</th>
        <th>ページ数</th>
        
      </tr>
      {{range $nowdata:=.}}
        <tr>
            <td>{{$nowdata.Id}}</td>
            <td>{{$nowdata.Subject_name}}</td>
            <td>{{$nowdata.Name}}</td>
            <td>{{$nowdata.Page}}</td>
            <td>削除</td>
        </tr>
        {{end}}
      <!-- example -->
      
    </table>
</body>
