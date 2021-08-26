<head>
  <meta charset="UTF-8">
  <title>講師用画面(生徒情報)</title>
  <link href="../client/css/register_style.css" rel="stylesheet">
</head>

<body>
    <div class="head_contents">以下の内容で登録します。</div> 

<div class="check_field">
  <ul>
    <li class="contents"> タイトル {{.Name}}
    </li>
    <li class="contents"> 科目 {{.Subject_name}}
    </li>
    <li class="contents"> ページ数 {{.Page}}
  </ul>
</div>


<a class="a_button" href="teacher.html">登録</a>
<a class="a_button" href="register_textbook2.html">修正</a>
</body>
