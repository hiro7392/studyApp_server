<head>
  <meta charset="UTF-8">
  <title>講師用画面(生徒検索結果)</title>
  <link href="../css/style.css" rel="stylesheet">
</head>

<body>
  <header>
  <a class="a_button" href="teacher.html">ホーム</a>
  </header>

  <!-- 前のページで入力した条件に合う生徒を表示。 -->
  <!-- 生徒名をクリックして。生徒のページ(teacher_student.html)へ移動。その生徒の情報を次のページで表示したい -->
  <h2>生徒一覧</h2>
    <p>＊生徒名をクリックして各生徒の情報を表示することができます</p>
    <table>
      <colgroup span="5"></colgroup>
      <tr>
        <th>生徒Id</th>
        <th>氏名</th>
        <th>学年</th>
        <th>担当の講師</th>
        <th>志望校</th>
      </tr>
      {{range $nowdata:=.}}
        <tr>
            <td>{{$nowdata.Id}}</td>
            <td><a href="/teacher_student/{{$nowdata.Id}}">{{$nowdata.Name}}</td>
            <td>{{$nowdata.Grade}}</td>
            <td>{{$nowdata.NowSchool}}</td>
            <td>{{$nowdata.WantSchool}}</td>
        </tr>
        {{end}}
      <!-- example -->
      
    </table>
</body>
