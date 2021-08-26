<head>
  <meta charset="UTF-8">
  <title>講師用画面(生徒検索結果)</title>
  <link href="../client/css/search_student.css" rel="stylesheet">
</head>

<body>
  <header>
    <a class="header-logo" href="/teacher">アプリ名</a>
  </header>

  <div class="container">
    <!-- 前のページで入力した条件に合う生徒を表示。 -->
    <!-- 生徒名をクリックして。生徒のページ(teacher_student.html)へ移動。その生徒の情報を次のページで表示したい -->
    <h2>検索結果</h2>
    <p>＊ 生徒名をクリックして各生徒の情報を表示することができます</p>
    <table class="type06">
      <colgroup span="5"></colgroup>
      <tr>
        <th>生徒ID</th>
        <th>氏名</th>
        <th>学年</th>
        <th>担当の講師</th>
        <th>志望校</th>
        <th>削除</th> 
      </tr>
      <tr>
        {{range $nowdata:=.}}
        <tr>
            <td>{{$nowdata.Id}}</td>
            <td><a href="/teacher_student/{{$nowdata.Id}}">{{$nowdata.Name}}</td>
            <td>{{$nowdata.Grade}}</td>
            <td>{{$nowdata.NowSchool}}</td>
            <td>{{$nowdata.WantSchool}}</td>
            <td>
              <a class="a_button" href="#delete-request">　　</a>
          </td>
        </tr>
        {{end}}
      </tr>
      <!-- example -->
    </table>
  </div>
</body>

