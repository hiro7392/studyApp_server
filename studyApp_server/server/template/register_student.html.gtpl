
<head>
  <meta charset="UTF-8">
  <title>講師用画面(生徒情報)</title>
  <link href="../css/register.css" rel="stylesheet">
</head>

<body>
    <div class="head_contents">以下の内容で登録します。</div>

<div class="check_field">
  <ul>
    <li class="contents"> 学年 {{.Grade}}
    </li>
    <li class="contents"> 担当講師 {{.Teacher_name}}
    </li>
    <li class="contents"> 中学校 {{.NowSchool}}
    </li>
    <li class="contents"> 志望校 {{.WantSchool}}
    </li>
  </ul>
</div>


<a class="a_button" href="teacher.html">登録</a>
<a class="a_button" href="register_student2.html">修正</a>
</body>
